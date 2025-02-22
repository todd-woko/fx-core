package server

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb/util"
	sm "github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/store"
	dbm "github.com/tendermint/tm-db"
)

const (
	flagHeight    = "height"
	flagPruning   = "enable_pruning"
	flagDBBackend = "db_backend"

	BlockDBName = "blockstore"
	StateDBName = "state"
	AppDBName   = "application"
)

func DataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data",
		Short: "modify data or query data in database",
	}

	cmd.AddCommand(
		dataQueryBlockCmd(),
		dataPruningCmd(),
	)

	cmd.PersistentFlags().String(flagDBBackend, "goleveldb", "Database backend: goleveldb")
	return cmd
}

func dataQueryBlockCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "block",
		Short: "Query blocks heights in database",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)
			backendType := dbm.BackendType(ctx.Viper.GetString(flagDBBackend))

			blockStoreDB, err := openDB(BlockDBName, backendType, ctx.Config.RootDir)
			if err != nil {
				return err
			}
			blockStore := store.NewBlockStore(blockStoreDB)
			fmt.Printf("[%d ~ %d]\n", blockStore.Base(), blockStore.Height())

			return nil
		},
	}
}

func dataPruningCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prune-compact",
		Short: "Prune and Compact blocks and application states",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			backend, err := cmd.Flags().GetString(flagDBBackend)
			if err != nil {
				return err
			}
			if backend != string(dbm.GoLevelDBBackend) {
				return fmt.Errorf("nonsupport db_backend %s, expected <%s>", backend, dbm.GoLevelDBBackend)
			}
			return nil
		},
	}

	cmd.AddCommand(
		pruneAllCmd(),
		pruneAppCmd(),
		pruneBlockCmd(),
	)

	cmd.PersistentFlags().Int64P(flagHeight, "r", 0, "Removes block or state up to (but not including) a height")
	cmd.PersistentFlags().BoolP(flagPruning, "p", false, "Enable pruning")
	return cmd
}

func pruneAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "Compact both application states and blocks",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)

			blockStoreDB := GetDB(ctx.Config.RootDir, BlockDBName)
			stateDB := GetDB(ctx.Config.RootDir, StateDBName)
			appDB := GetDB(ctx.Config.RootDir, AppDBName)

			if viper.GetBool(flagPruning) {
				baseHeight, retainHeight := getPruneBlockParams(blockStoreDB)
				var wg sync.WaitGroup
				log.Println("--------- pruning start... ---------")
				wg.Add(2)
				go pruneBlocks(blockStoreDB, baseHeight, retainHeight, &wg)
				go pruneStates(stateDB, baseHeight, retainHeight, &wg)
				wg.Wait()
				log.Println("--------- pruning end!!!   ---------")
			}
			log.Println("--------- compact start... ---------")
			var wg sync.WaitGroup
			wg.Add(3)
			go compactDB(blockStoreDB, BlockDBName, &wg)
			go compactDB(stateDB, StateDBName, &wg)
			go compactDB(appDB, AppDBName, &wg)
			wg.Wait()
			log.Println("--------- compact end!!!   ---------")

			return nil
		},
	}

	return cmd
}

func pruneAppCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: "Compact application state",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)

			appDB := GetDB(ctx.Config.RootDir, AppDBName)
			log.Println("--------- compact start ---------")
			var wg sync.WaitGroup
			wg.Add(1)
			compactDB(appDB, AppDBName, &wg)
			log.Println("--------- compact end ---------")

			return nil
		},
	}

	return cmd
}

func pruneBlockCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block",
		Short: "Compact while pruning blocks and states",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := server.GetServerContextFromCmd(cmd)

			blockStoreDB := GetDB(ctx.Config.RootDir, BlockDBName)
			stateDB := GetDB(ctx.Config.RootDir, StateDBName)

			if viper.GetBool(flagPruning) {
				baseHeight, retainHeight := getPruneBlockParams(blockStoreDB)
				//
				log.Println("--------- pruning start... ---------")
				var wg sync.WaitGroup
				wg.Add(2)
				go pruneBlocks(blockStoreDB, baseHeight, retainHeight, &wg)
				go pruneStates(stateDB, baseHeight, retainHeight, &wg)
				wg.Wait()
				log.Println("--------- pruning end!!!   ---------")
			}

			log.Println("--------- compact start... ---------")
			var wg sync.WaitGroup
			wg.Add(2)
			go compactDB(blockStoreDB, BlockDBName, &wg)
			go compactDB(stateDB, StateDBName, &wg)
			wg.Wait()
			log.Println("--------- compact end!!!   ---------")

			return nil
		},
	}

	return cmd
}

func getPruneBlockParams(blockStoreDB dbm.DB) (baseHeight, retainHeight int64) {
	baseHeight, size := getBlockInfo(blockStoreDB)

	retainHeight = viper.GetInt64(flagHeight)
	if retainHeight >= baseHeight+size-1 || retainHeight <= baseHeight {
		retainHeight = baseHeight + size - 2
	}

	return
}

// pruneBlocks deletes blocks between the given heights (including from, excluding to).
func pruneBlocks(blockStoreDB dbm.DB, baseHeight, retainHeight int64, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Prune blocks [%d,%d)...", baseHeight, retainHeight)
	if retainHeight <= baseHeight {
		return
	}

	baseHeightBefore, sizeBefore := getBlockInfo(blockStoreDB)
	start := time.Now()
	_, err := store.NewBlockStore(blockStoreDB).PruneBlocks(retainHeight)
	if err != nil {
		panic(fmt.Errorf("failed to prune block store: %w", err))
	}

	baseHeightAfter, sizeAfter := getBlockInfo(blockStoreDB)
	log.Printf("Block db info [baseHeight,size]: [%d,%d] --> [%d,%d]\n", baseHeightBefore, sizeBefore, baseHeightAfter, sizeAfter)
	log.Printf("Prune blocks done in %v \n", time.Since(start))
}

// pruneStates deletes states between the given heights (including from, excluding to).
func pruneStates(stateDB dbm.DB, from, to int64, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Prune states [%d,%d)...", from, to)
	if to <= from {
		return
	}

	start := time.Now()
	stateStore := sm.NewStore(stateDB, sm.StoreOptions{
		DiscardABCIResponses: false,
	})
	if err := stateStore.PruneStates(from, to); err != nil {
		panic(fmt.Errorf("failed to prune state database: %w", err))
	}

	log.Printf("Prune states done in %v \n", time.Since(start))
}

func compactDB(db dbm.DB, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Compact %s... \n", name)
	start := time.Now()

	if ldb, ok := db.(*dbm.GoLevelDB); ok {
		if err := ldb.DB().CompactRange(util.Range{}); err != nil {
			panic(err)
		}
	}

	log.Printf("Compact %s done in %v \n", name, time.Since(start))
}

func getBlockInfo(blockStoreDB dbm.DB) (baseHeight, size int64) {
	blockStore := store.NewBlockStore(blockStoreDB)
	baseHeight = blockStore.Base()
	size = blockStore.Size()
	return
}

func GetDB(rootDir, dbName string) dbm.DB {
	if dbName != BlockDBName && dbName != StateDBName && dbName != AppDBName {
		panic(fmt.Sprintf("unknow db name: %s", dbName))
	}
	db, err := openDB(BlockDBName, dbm.GoLevelDBBackend, rootDir)
	if err != nil {
		panic(err)
	}
	return db
}

func openDB(name string, backendType dbm.BackendType, rootDir string) (dbm.DB, error) {
	dataDir := filepath.Join(rootDir, "data")
	return dbm.NewDB(name, backendType, dataDir)
}
