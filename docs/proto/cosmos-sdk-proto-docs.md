<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [cosmos/app/v1alpha1/module.proto](#cosmos/app/v1alpha1/module.proto)
    - [MigrateFromInfo](#cosmos.app.v1alpha1.MigrateFromInfo)
    - [ModuleDescriptor](#cosmos.app.v1alpha1.ModuleDescriptor)
    - [PackageReference](#cosmos.app.v1alpha1.PackageReference)
  
    - [File-level Extensions](#cosmos/app/v1alpha1/module.proto-extensions)
  
- [cosmos/app/module/v1alpha1/module.proto](#cosmos/app/module/v1alpha1/module.proto)
    - [Module](#cosmos.app.module.v1alpha1.Module)
  
- [cosmos/app/v1alpha1/config.proto](#cosmos/app/v1alpha1/config.proto)
    - [Config](#cosmos.app.v1alpha1.Config)
    - [ModuleConfig](#cosmos.app.v1alpha1.ModuleConfig)
  
- [cosmos/app/v1alpha1/query.proto](#cosmos/app/v1alpha1/query.proto)
    - [QueryConfigRequest](#cosmos.app.v1alpha1.QueryConfigRequest)
    - [QueryConfigResponse](#cosmos.app.v1alpha1.QueryConfigResponse)
  
    - [Query](#cosmos.app.v1alpha1.Query)
  
- [cosmos/auth/v1beta1/auth.proto](#cosmos/auth/v1beta1/auth.proto)
    - [BaseAccount](#cosmos.auth.v1beta1.BaseAccount)
    - [ModuleAccount](#cosmos.auth.v1beta1.ModuleAccount)
    - [Params](#cosmos.auth.v1beta1.Params)
  
- [cosmos/auth/v1beta1/genesis.proto](#cosmos/auth/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.auth.v1beta1.GenesisState)
  
- [cosmos/base/query/v1beta1/pagination.proto](#cosmos/base/query/v1beta1/pagination.proto)
    - [PageRequest](#cosmos.base.query.v1beta1.PageRequest)
    - [PageResponse](#cosmos.base.query.v1beta1.PageResponse)
  
- [cosmos/auth/v1beta1/query.proto](#cosmos/auth/v1beta1/query.proto)
    - [AddressBytesToStringRequest](#cosmos.auth.v1beta1.AddressBytesToStringRequest)
    - [AddressBytesToStringResponse](#cosmos.auth.v1beta1.AddressBytesToStringResponse)
    - [AddressStringToBytesRequest](#cosmos.auth.v1beta1.AddressStringToBytesRequest)
    - [AddressStringToBytesResponse](#cosmos.auth.v1beta1.AddressStringToBytesResponse)
    - [Bech32PrefixRequest](#cosmos.auth.v1beta1.Bech32PrefixRequest)
    - [Bech32PrefixResponse](#cosmos.auth.v1beta1.Bech32PrefixResponse)
    - [QueryAccountAddressByIDRequest](#cosmos.auth.v1beta1.QueryAccountAddressByIDRequest)
    - [QueryAccountAddressByIDResponse](#cosmos.auth.v1beta1.QueryAccountAddressByIDResponse)
    - [QueryAccountRequest](#cosmos.auth.v1beta1.QueryAccountRequest)
    - [QueryAccountResponse](#cosmos.auth.v1beta1.QueryAccountResponse)
    - [QueryAccountsRequest](#cosmos.auth.v1beta1.QueryAccountsRequest)
    - [QueryAccountsResponse](#cosmos.auth.v1beta1.QueryAccountsResponse)
    - [QueryModuleAccountByNameRequest](#cosmos.auth.v1beta1.QueryModuleAccountByNameRequest)
    - [QueryModuleAccountByNameResponse](#cosmos.auth.v1beta1.QueryModuleAccountByNameResponse)
    - [QueryModuleAccountsRequest](#cosmos.auth.v1beta1.QueryModuleAccountsRequest)
    - [QueryModuleAccountsResponse](#cosmos.auth.v1beta1.QueryModuleAccountsResponse)
    - [QueryParamsRequest](#cosmos.auth.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.auth.v1beta1.QueryParamsResponse)
  
    - [Query](#cosmos.auth.v1beta1.Query)
  
- [cosmos/authz/v1beta1/authz.proto](#cosmos/authz/v1beta1/authz.proto)
    - [GenericAuthorization](#cosmos.authz.v1beta1.GenericAuthorization)
    - [Grant](#cosmos.authz.v1beta1.Grant)
    - [GrantAuthorization](#cosmos.authz.v1beta1.GrantAuthorization)
    - [GrantQueueItem](#cosmos.authz.v1beta1.GrantQueueItem)
  
- [cosmos/authz/v1beta1/event.proto](#cosmos/authz/v1beta1/event.proto)
    - [EventGrant](#cosmos.authz.v1beta1.EventGrant)
    - [EventRevoke](#cosmos.authz.v1beta1.EventRevoke)
  
- [cosmos/authz/v1beta1/genesis.proto](#cosmos/authz/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.authz.v1beta1.GenesisState)
  
- [cosmos/authz/v1beta1/query.proto](#cosmos/authz/v1beta1/query.proto)
    - [QueryGranteeGrantsRequest](#cosmos.authz.v1beta1.QueryGranteeGrantsRequest)
    - [QueryGranteeGrantsResponse](#cosmos.authz.v1beta1.QueryGranteeGrantsResponse)
    - [QueryGranterGrantsRequest](#cosmos.authz.v1beta1.QueryGranterGrantsRequest)
    - [QueryGranterGrantsResponse](#cosmos.authz.v1beta1.QueryGranterGrantsResponse)
    - [QueryGrantsRequest](#cosmos.authz.v1beta1.QueryGrantsRequest)
    - [QueryGrantsResponse](#cosmos.authz.v1beta1.QueryGrantsResponse)
  
    - [Query](#cosmos.authz.v1beta1.Query)
  
- [cosmos/msg/v1/msg.proto](#cosmos/msg/v1/msg.proto)
    - [File-level Extensions](#cosmos/msg/v1/msg.proto-extensions)
  
- [cosmos/authz/v1beta1/tx.proto](#cosmos/authz/v1beta1/tx.proto)
    - [MsgExec](#cosmos.authz.v1beta1.MsgExec)
    - [MsgExecResponse](#cosmos.authz.v1beta1.MsgExecResponse)
    - [MsgGrant](#cosmos.authz.v1beta1.MsgGrant)
    - [MsgGrantResponse](#cosmos.authz.v1beta1.MsgGrantResponse)
    - [MsgRevoke](#cosmos.authz.v1beta1.MsgRevoke)
    - [MsgRevokeResponse](#cosmos.authz.v1beta1.MsgRevokeResponse)
  
    - [Msg](#cosmos.authz.v1beta1.Msg)
  
- [cosmos/base/v1beta1/coin.proto](#cosmos/base/v1beta1/coin.proto)
    - [Coin](#cosmos.base.v1beta1.Coin)
    - [DecCoin](#cosmos.base.v1beta1.DecCoin)
    - [DecProto](#cosmos.base.v1beta1.DecProto)
    - [IntProto](#cosmos.base.v1beta1.IntProto)
  
- [cosmos/bank/v1beta1/authz.proto](#cosmos/bank/v1beta1/authz.proto)
    - [SendAuthorization](#cosmos.bank.v1beta1.SendAuthorization)
  
- [cosmos/bank/v1beta1/bank.proto](#cosmos/bank/v1beta1/bank.proto)
    - [DenomUnit](#cosmos.bank.v1beta1.DenomUnit)
    - [Input](#cosmos.bank.v1beta1.Input)
    - [Metadata](#cosmos.bank.v1beta1.Metadata)
    - [Output](#cosmos.bank.v1beta1.Output)
    - [Params](#cosmos.bank.v1beta1.Params)
    - [SendEnabled](#cosmos.bank.v1beta1.SendEnabled)
    - [Supply](#cosmos.bank.v1beta1.Supply)
  
- [cosmos/bank/v1beta1/genesis.proto](#cosmos/bank/v1beta1/genesis.proto)
    - [Balance](#cosmos.bank.v1beta1.Balance)
    - [GenesisState](#cosmos.bank.v1beta1.GenesisState)
  
- [cosmos/bank/v1beta1/query.proto](#cosmos/bank/v1beta1/query.proto)
    - [DenomOwner](#cosmos.bank.v1beta1.DenomOwner)
    - [QueryAllBalancesRequest](#cosmos.bank.v1beta1.QueryAllBalancesRequest)
    - [QueryAllBalancesResponse](#cosmos.bank.v1beta1.QueryAllBalancesResponse)
    - [QueryBalanceRequest](#cosmos.bank.v1beta1.QueryBalanceRequest)
    - [QueryBalanceResponse](#cosmos.bank.v1beta1.QueryBalanceResponse)
    - [QueryDenomMetadataRequest](#cosmos.bank.v1beta1.QueryDenomMetadataRequest)
    - [QueryDenomMetadataResponse](#cosmos.bank.v1beta1.QueryDenomMetadataResponse)
    - [QueryDenomOwnersRequest](#cosmos.bank.v1beta1.QueryDenomOwnersRequest)
    - [QueryDenomOwnersResponse](#cosmos.bank.v1beta1.QueryDenomOwnersResponse)
    - [QueryDenomsMetadataRequest](#cosmos.bank.v1beta1.QueryDenomsMetadataRequest)
    - [QueryDenomsMetadataResponse](#cosmos.bank.v1beta1.QueryDenomsMetadataResponse)
    - [QueryParamsRequest](#cosmos.bank.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.bank.v1beta1.QueryParamsResponse)
    - [QuerySpendableBalancesRequest](#cosmos.bank.v1beta1.QuerySpendableBalancesRequest)
    - [QuerySpendableBalancesResponse](#cosmos.bank.v1beta1.QuerySpendableBalancesResponse)
    - [QuerySupplyOfRequest](#cosmos.bank.v1beta1.QuerySupplyOfRequest)
    - [QuerySupplyOfResponse](#cosmos.bank.v1beta1.QuerySupplyOfResponse)
    - [QueryTotalSupplyRequest](#cosmos.bank.v1beta1.QueryTotalSupplyRequest)
    - [QueryTotalSupplyResponse](#cosmos.bank.v1beta1.QueryTotalSupplyResponse)
  
    - [Query](#cosmos.bank.v1beta1.Query)
  
- [cosmos/bank/v1beta1/tx.proto](#cosmos/bank/v1beta1/tx.proto)
    - [MsgMultiSend](#cosmos.bank.v1beta1.MsgMultiSend)
    - [MsgMultiSendResponse](#cosmos.bank.v1beta1.MsgMultiSendResponse)
    - [MsgSend](#cosmos.bank.v1beta1.MsgSend)
    - [MsgSendResponse](#cosmos.bank.v1beta1.MsgSendResponse)
  
    - [Msg](#cosmos.bank.v1beta1.Msg)
  
- [tendermint/crypto/proof.proto](#tendermint/crypto/proof.proto)
    - [DominoOp](#tendermint.crypto.DominoOp)
    - [Proof](#tendermint.crypto.Proof)
    - [ProofOp](#tendermint.crypto.ProofOp)
    - [ProofOps](#tendermint.crypto.ProofOps)
    - [ValueOp](#tendermint.crypto.ValueOp)
  
- [tendermint/version/types.proto](#tendermint/version/types.proto)
    - [App](#tendermint.version.App)
    - [Consensus](#tendermint.version.Consensus)
  
- [tendermint/crypto/keys.proto](#tendermint/crypto/keys.proto)
    - [PublicKey](#tendermint.crypto.PublicKey)
  
- [tendermint/types/validator.proto](#tendermint/types/validator.proto)
    - [SimpleValidator](#tendermint.types.SimpleValidator)
    - [Validator](#tendermint.types.Validator)
    - [ValidatorSet](#tendermint.types.ValidatorSet)
  
- [tendermint/types/types.proto](#tendermint/types/types.proto)
    - [BlockID](#tendermint.types.BlockID)
    - [BlockMeta](#tendermint.types.BlockMeta)
    - [Commit](#tendermint.types.Commit)
    - [CommitSig](#tendermint.types.CommitSig)
    - [Data](#tendermint.types.Data)
    - [Header](#tendermint.types.Header)
    - [LightBlock](#tendermint.types.LightBlock)
    - [Part](#tendermint.types.Part)
    - [PartSetHeader](#tendermint.types.PartSetHeader)
    - [Proposal](#tendermint.types.Proposal)
    - [SignedHeader](#tendermint.types.SignedHeader)
    - [TxProof](#tendermint.types.TxProof)
    - [Vote](#tendermint.types.Vote)
  
    - [BlockIDFlag](#tendermint.types.BlockIDFlag)
    - [SignedMsgType](#tendermint.types.SignedMsgType)
  
- [tendermint/types/params.proto](#tendermint/types/params.proto)
    - [BlockParams](#tendermint.types.BlockParams)
    - [ConsensusParams](#tendermint.types.ConsensusParams)
    - [EvidenceParams](#tendermint.types.EvidenceParams)
    - [HashedParams](#tendermint.types.HashedParams)
    - [ValidatorParams](#tendermint.types.ValidatorParams)
    - [VersionParams](#tendermint.types.VersionParams)
  
- [tendermint/abci/types.proto](#tendermint/abci/types.proto)
    - [BlockParams](#tendermint.abci.BlockParams)
    - [ConsensusParams](#tendermint.abci.ConsensusParams)
    - [Event](#tendermint.abci.Event)
    - [EventAttribute](#tendermint.abci.EventAttribute)
    - [Evidence](#tendermint.abci.Evidence)
    - [LastCommitInfo](#tendermint.abci.LastCommitInfo)
    - [Request](#tendermint.abci.Request)
    - [RequestApplySnapshotChunk](#tendermint.abci.RequestApplySnapshotChunk)
    - [RequestBeginBlock](#tendermint.abci.RequestBeginBlock)
    - [RequestCheckTx](#tendermint.abci.RequestCheckTx)
    - [RequestCommit](#tendermint.abci.RequestCommit)
    - [RequestDeliverTx](#tendermint.abci.RequestDeliverTx)
    - [RequestEcho](#tendermint.abci.RequestEcho)
    - [RequestEndBlock](#tendermint.abci.RequestEndBlock)
    - [RequestFlush](#tendermint.abci.RequestFlush)
    - [RequestInfo](#tendermint.abci.RequestInfo)
    - [RequestInitChain](#tendermint.abci.RequestInitChain)
    - [RequestListSnapshots](#tendermint.abci.RequestListSnapshots)
    - [RequestLoadSnapshotChunk](#tendermint.abci.RequestLoadSnapshotChunk)
    - [RequestOfferSnapshot](#tendermint.abci.RequestOfferSnapshot)
    - [RequestQuery](#tendermint.abci.RequestQuery)
    - [RequestSetOption](#tendermint.abci.RequestSetOption)
    - [Response](#tendermint.abci.Response)
    - [ResponseApplySnapshotChunk](#tendermint.abci.ResponseApplySnapshotChunk)
    - [ResponseBeginBlock](#tendermint.abci.ResponseBeginBlock)
    - [ResponseCheckTx](#tendermint.abci.ResponseCheckTx)
    - [ResponseCommit](#tendermint.abci.ResponseCommit)
    - [ResponseDeliverTx](#tendermint.abci.ResponseDeliverTx)
    - [ResponseEcho](#tendermint.abci.ResponseEcho)
    - [ResponseEndBlock](#tendermint.abci.ResponseEndBlock)
    - [ResponseException](#tendermint.abci.ResponseException)
    - [ResponseFlush](#tendermint.abci.ResponseFlush)
    - [ResponseInfo](#tendermint.abci.ResponseInfo)
    - [ResponseInitChain](#tendermint.abci.ResponseInitChain)
    - [ResponseListSnapshots](#tendermint.abci.ResponseListSnapshots)
    - [ResponseLoadSnapshotChunk](#tendermint.abci.ResponseLoadSnapshotChunk)
    - [ResponseOfferSnapshot](#tendermint.abci.ResponseOfferSnapshot)
    - [ResponseQuery](#tendermint.abci.ResponseQuery)
    - [ResponseSetOption](#tendermint.abci.ResponseSetOption)
    - [Snapshot](#tendermint.abci.Snapshot)
    - [TxResult](#tendermint.abci.TxResult)
    - [Validator](#tendermint.abci.Validator)
    - [ValidatorUpdate](#tendermint.abci.ValidatorUpdate)
    - [VoteInfo](#tendermint.abci.VoteInfo)
  
    - [CheckTxType](#tendermint.abci.CheckTxType)
    - [EvidenceType](#tendermint.abci.EvidenceType)
    - [ResponseApplySnapshotChunk.Result](#tendermint.abci.ResponseApplySnapshotChunk.Result)
    - [ResponseOfferSnapshot.Result](#tendermint.abci.ResponseOfferSnapshot.Result)
  
    - [ABCIApplication](#tendermint.abci.ABCIApplication)
  
- [cosmos/base/abci/v1beta1/abci.proto](#cosmos/base/abci/v1beta1/abci.proto)
    - [ABCIMessageLog](#cosmos.base.abci.v1beta1.ABCIMessageLog)
    - [Attribute](#cosmos.base.abci.v1beta1.Attribute)
    - [GasInfo](#cosmos.base.abci.v1beta1.GasInfo)
    - [MsgData](#cosmos.base.abci.v1beta1.MsgData)
    - [Result](#cosmos.base.abci.v1beta1.Result)
    - [SearchTxsResult](#cosmos.base.abci.v1beta1.SearchTxsResult)
    - [SimulationResponse](#cosmos.base.abci.v1beta1.SimulationResponse)
    - [StringEvent](#cosmos.base.abci.v1beta1.StringEvent)
    - [TxMsgData](#cosmos.base.abci.v1beta1.TxMsgData)
    - [TxResponse](#cosmos.base.abci.v1beta1.TxResponse)
  
- [cosmos/base/kv/v1beta1/kv.proto](#cosmos/base/kv/v1beta1/kv.proto)
    - [Pair](#cosmos.base.kv.v1beta1.Pair)
    - [Pairs](#cosmos.base.kv.v1beta1.Pairs)
  
- [cosmos/base/node/v1beta1/query.proto](#cosmos/base/node/v1beta1/query.proto)
    - [ConfigRequest](#cosmos.base.node.v1beta1.ConfigRequest)
    - [ConfigResponse](#cosmos.base.node.v1beta1.ConfigResponse)
  
    - [Service](#cosmos.base.node.v1beta1.Service)
  
- [cosmos/base/reflection/v1beta1/reflection.proto](#cosmos/base/reflection/v1beta1/reflection.proto)
    - [ListAllInterfacesRequest](#cosmos.base.reflection.v1beta1.ListAllInterfacesRequest)
    - [ListAllInterfacesResponse](#cosmos.base.reflection.v1beta1.ListAllInterfacesResponse)
    - [ListImplementationsRequest](#cosmos.base.reflection.v1beta1.ListImplementationsRequest)
    - [ListImplementationsResponse](#cosmos.base.reflection.v1beta1.ListImplementationsResponse)
  
    - [ReflectionService](#cosmos.base.reflection.v1beta1.ReflectionService)
  
- [cosmos/base/reflection/v2alpha1/reflection.proto](#cosmos/base/reflection/v2alpha1/reflection.proto)
    - [AppDescriptor](#cosmos.base.reflection.v2alpha1.AppDescriptor)
    - [AuthnDescriptor](#cosmos.base.reflection.v2alpha1.AuthnDescriptor)
    - [ChainDescriptor](#cosmos.base.reflection.v2alpha1.ChainDescriptor)
    - [CodecDescriptor](#cosmos.base.reflection.v2alpha1.CodecDescriptor)
    - [ConfigurationDescriptor](#cosmos.base.reflection.v2alpha1.ConfigurationDescriptor)
    - [GetAuthnDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest)
    - [GetAuthnDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse)
    - [GetChainDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest)
    - [GetChainDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse)
    - [GetCodecDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest)
    - [GetCodecDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse)
    - [GetConfigurationDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest)
    - [GetConfigurationDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse)
    - [GetQueryServicesDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest)
    - [GetQueryServicesDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse)
    - [GetTxDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest)
    - [GetTxDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse)
    - [InterfaceAcceptingMessageDescriptor](#cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor)
    - [InterfaceDescriptor](#cosmos.base.reflection.v2alpha1.InterfaceDescriptor)
    - [InterfaceImplementerDescriptor](#cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor)
    - [MsgDescriptor](#cosmos.base.reflection.v2alpha1.MsgDescriptor)
    - [QueryMethodDescriptor](#cosmos.base.reflection.v2alpha1.QueryMethodDescriptor)
    - [QueryServiceDescriptor](#cosmos.base.reflection.v2alpha1.QueryServiceDescriptor)
    - [QueryServicesDescriptor](#cosmos.base.reflection.v2alpha1.QueryServicesDescriptor)
    - [SigningModeDescriptor](#cosmos.base.reflection.v2alpha1.SigningModeDescriptor)
    - [TxDescriptor](#cosmos.base.reflection.v2alpha1.TxDescriptor)
  
    - [ReflectionService](#cosmos.base.reflection.v2alpha1.ReflectionService)
  
- [cosmos/base/snapshots/v1beta1/snapshot.proto](#cosmos/base/snapshots/v1beta1/snapshot.proto)
    - [Metadata](#cosmos.base.snapshots.v1beta1.Metadata)
    - [Snapshot](#cosmos.base.snapshots.v1beta1.Snapshot)
    - [SnapshotExtensionMeta](#cosmos.base.snapshots.v1beta1.SnapshotExtensionMeta)
    - [SnapshotExtensionPayload](#cosmos.base.snapshots.v1beta1.SnapshotExtensionPayload)
    - [SnapshotIAVLItem](#cosmos.base.snapshots.v1beta1.SnapshotIAVLItem)
    - [SnapshotItem](#cosmos.base.snapshots.v1beta1.SnapshotItem)
    - [SnapshotKVItem](#cosmos.base.snapshots.v1beta1.SnapshotKVItem)
    - [SnapshotSchema](#cosmos.base.snapshots.v1beta1.SnapshotSchema)
    - [SnapshotStoreItem](#cosmos.base.snapshots.v1beta1.SnapshotStoreItem)
  
- [cosmos/base/store/v1beta1/commit_info.proto](#cosmos/base/store/v1beta1/commit_info.proto)
    - [CommitID](#cosmos.base.store.v1beta1.CommitID)
    - [CommitInfo](#cosmos.base.store.v1beta1.CommitInfo)
    - [StoreInfo](#cosmos.base.store.v1beta1.StoreInfo)
  
- [cosmos/base/store/v1beta1/listening.proto](#cosmos/base/store/v1beta1/listening.proto)
    - [BlockMetadata](#cosmos.base.store.v1beta1.BlockMetadata)
    - [BlockMetadata.DeliverTx](#cosmos.base.store.v1beta1.BlockMetadata.DeliverTx)
    - [StoreKVPair](#cosmos.base.store.v1beta1.StoreKVPair)
  
- [tendermint/p2p/types.proto](#tendermint/p2p/types.proto)
    - [DefaultNodeInfo](#tendermint.p2p.DefaultNodeInfo)
    - [DefaultNodeInfoOther](#tendermint.p2p.DefaultNodeInfoOther)
    - [NetAddress](#tendermint.p2p.NetAddress)
    - [ProtocolVersion](#tendermint.p2p.ProtocolVersion)
  
- [tendermint/types/evidence.proto](#tendermint/types/evidence.proto)
    - [DuplicateVoteEvidence](#tendermint.types.DuplicateVoteEvidence)
    - [Evidence](#tendermint.types.Evidence)
    - [EvidenceList](#tendermint.types.EvidenceList)
    - [LightClientAttackEvidence](#tendermint.types.LightClientAttackEvidence)
  
- [cosmos/base/tendermint/v1beta1/types.proto](#cosmos/base/tendermint/v1beta1/types.proto)
    - [Block](#cosmos.base.tendermint.v1beta1.Block)
    - [Header](#cosmos.base.tendermint.v1beta1.Header)
  
- [tendermint/types/block.proto](#tendermint/types/block.proto)
    - [Block](#tendermint.types.Block)
  
- [cosmos/base/tendermint/v1beta1/query.proto](#cosmos/base/tendermint/v1beta1/query.proto)
    - [ABCIQueryRequest](#cosmos.base.tendermint.v1beta1.ABCIQueryRequest)
    - [ABCIQueryResponse](#cosmos.base.tendermint.v1beta1.ABCIQueryResponse)
    - [GetBlockByHeightRequest](#cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest)
    - [GetBlockByHeightResponse](#cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse)
    - [GetLatestBlockRequest](#cosmos.base.tendermint.v1beta1.GetLatestBlockRequest)
    - [GetLatestBlockResponse](#cosmos.base.tendermint.v1beta1.GetLatestBlockResponse)
    - [GetLatestValidatorSetRequest](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest)
    - [GetLatestValidatorSetResponse](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse)
    - [GetNodeInfoRequest](#cosmos.base.tendermint.v1beta1.GetNodeInfoRequest)
    - [GetNodeInfoResponse](#cosmos.base.tendermint.v1beta1.GetNodeInfoResponse)
    - [GetSyncingRequest](#cosmos.base.tendermint.v1beta1.GetSyncingRequest)
    - [GetSyncingResponse](#cosmos.base.tendermint.v1beta1.GetSyncingResponse)
    - [GetValidatorSetByHeightRequest](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest)
    - [GetValidatorSetByHeightResponse](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse)
    - [Module](#cosmos.base.tendermint.v1beta1.Module)
    - [ProofOp](#cosmos.base.tendermint.v1beta1.ProofOp)
    - [ProofOps](#cosmos.base.tendermint.v1beta1.ProofOps)
    - [Validator](#cosmos.base.tendermint.v1beta1.Validator)
    - [VersionInfo](#cosmos.base.tendermint.v1beta1.VersionInfo)
  
    - [Service](#cosmos.base.tendermint.v1beta1.Service)
  
- [cosmos/capability/v1beta1/capability.proto](#cosmos/capability/v1beta1/capability.proto)
    - [Capability](#cosmos.capability.v1beta1.Capability)
    - [CapabilityOwners](#cosmos.capability.v1beta1.CapabilityOwners)
    - [Owner](#cosmos.capability.v1beta1.Owner)
  
- [cosmos/capability/v1beta1/genesis.proto](#cosmos/capability/v1beta1/genesis.proto)
    - [GenesisOwners](#cosmos.capability.v1beta1.GenesisOwners)
    - [GenesisState](#cosmos.capability.v1beta1.GenesisState)
  
- [cosmos/crisis/v1beta1/genesis.proto](#cosmos/crisis/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.crisis.v1beta1.GenesisState)
  
- [cosmos/crisis/v1beta1/tx.proto](#cosmos/crisis/v1beta1/tx.proto)
    - [MsgVerifyInvariant](#cosmos.crisis.v1beta1.MsgVerifyInvariant)
    - [MsgVerifyInvariantResponse](#cosmos.crisis.v1beta1.MsgVerifyInvariantResponse)
  
    - [Msg](#cosmos.crisis.v1beta1.Msg)
  
- [cosmos/crypto/ed25519/keys.proto](#cosmos/crypto/ed25519/keys.proto)
    - [PrivKey](#cosmos.crypto.ed25519.PrivKey)
    - [PubKey](#cosmos.crypto.ed25519.PubKey)
  
- [cosmos/crypto/hd/v1/hd.proto](#cosmos/crypto/hd/v1/hd.proto)
    - [BIP44Params](#cosmos.crypto.hd.v1.BIP44Params)
  
- [cosmos/crypto/keyring/v1/record.proto](#cosmos/crypto/keyring/v1/record.proto)
    - [Record](#cosmos.crypto.keyring.v1.Record)
    - [Record.Ledger](#cosmos.crypto.keyring.v1.Record.Ledger)
    - [Record.Local](#cosmos.crypto.keyring.v1.Record.Local)
    - [Record.Multi](#cosmos.crypto.keyring.v1.Record.Multi)
    - [Record.Offline](#cosmos.crypto.keyring.v1.Record.Offline)
  
- [cosmos/crypto/multisig/keys.proto](#cosmos/crypto/multisig/keys.proto)
    - [LegacyAminoPubKey](#cosmos.crypto.multisig.LegacyAminoPubKey)
  
- [cosmos/crypto/multisig/v1beta1/multisig.proto](#cosmos/crypto/multisig/v1beta1/multisig.proto)
    - [CompactBitArray](#cosmos.crypto.multisig.v1beta1.CompactBitArray)
    - [MultiSignature](#cosmos.crypto.multisig.v1beta1.MultiSignature)
  
- [cosmos/crypto/secp256k1/keys.proto](#cosmos/crypto/secp256k1/keys.proto)
    - [PrivKey](#cosmos.crypto.secp256k1.PrivKey)
    - [PubKey](#cosmos.crypto.secp256k1.PubKey)
  
- [cosmos/crypto/secp256r1/keys.proto](#cosmos/crypto/secp256r1/keys.proto)
    - [PrivKey](#cosmos.crypto.secp256r1.PrivKey)
    - [PubKey](#cosmos.crypto.secp256r1.PubKey)
  
- [cosmos/distribution/v1beta1/distribution.proto](#cosmos/distribution/v1beta1/distribution.proto)
    - [CommunityPoolSpendProposal](#cosmos.distribution.v1beta1.CommunityPoolSpendProposal)
    - [CommunityPoolSpendProposalWithDeposit](#cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit)
    - [DelegationDelegatorReward](#cosmos.distribution.v1beta1.DelegationDelegatorReward)
    - [DelegatorStartingInfo](#cosmos.distribution.v1beta1.DelegatorStartingInfo)
    - [FeePool](#cosmos.distribution.v1beta1.FeePool)
    - [Params](#cosmos.distribution.v1beta1.Params)
    - [ValidatorAccumulatedCommission](#cosmos.distribution.v1beta1.ValidatorAccumulatedCommission)
    - [ValidatorCurrentRewards](#cosmos.distribution.v1beta1.ValidatorCurrentRewards)
    - [ValidatorHistoricalRewards](#cosmos.distribution.v1beta1.ValidatorHistoricalRewards)
    - [ValidatorOutstandingRewards](#cosmos.distribution.v1beta1.ValidatorOutstandingRewards)
    - [ValidatorSlashEvent](#cosmos.distribution.v1beta1.ValidatorSlashEvent)
    - [ValidatorSlashEvents](#cosmos.distribution.v1beta1.ValidatorSlashEvents)
  
- [cosmos/distribution/v1beta1/genesis.proto](#cosmos/distribution/v1beta1/genesis.proto)
    - [DelegatorStartingInfoRecord](#cosmos.distribution.v1beta1.DelegatorStartingInfoRecord)
    - [DelegatorWithdrawInfo](#cosmos.distribution.v1beta1.DelegatorWithdrawInfo)
    - [GenesisState](#cosmos.distribution.v1beta1.GenesisState)
    - [ValidatorAccumulatedCommissionRecord](#cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord)
    - [ValidatorCurrentRewardsRecord](#cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord)
    - [ValidatorHistoricalRewardsRecord](#cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord)
    - [ValidatorOutstandingRewardsRecord](#cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord)
    - [ValidatorSlashEventRecord](#cosmos.distribution.v1beta1.ValidatorSlashEventRecord)
  
- [cosmos/distribution/v1beta1/query.proto](#cosmos/distribution/v1beta1/query.proto)
    - [QueryCommunityPoolRequest](#cosmos.distribution.v1beta1.QueryCommunityPoolRequest)
    - [QueryCommunityPoolResponse](#cosmos.distribution.v1beta1.QueryCommunityPoolResponse)
    - [QueryDelegationRewardsRequest](#cosmos.distribution.v1beta1.QueryDelegationRewardsRequest)
    - [QueryDelegationRewardsResponse](#cosmos.distribution.v1beta1.QueryDelegationRewardsResponse)
    - [QueryDelegationTotalRewardsRequest](#cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest)
    - [QueryDelegationTotalRewardsResponse](#cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse)
    - [QueryDelegatorValidatorsRequest](#cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest)
    - [QueryDelegatorValidatorsResponse](#cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse)
    - [QueryDelegatorWithdrawAddressRequest](#cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest)
    - [QueryDelegatorWithdrawAddressResponse](#cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse)
    - [QueryParamsRequest](#cosmos.distribution.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.distribution.v1beta1.QueryParamsResponse)
    - [QueryValidatorCommissionRequest](#cosmos.distribution.v1beta1.QueryValidatorCommissionRequest)
    - [QueryValidatorCommissionResponse](#cosmos.distribution.v1beta1.QueryValidatorCommissionResponse)
    - [QueryValidatorOutstandingRewardsRequest](#cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest)
    - [QueryValidatorOutstandingRewardsResponse](#cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse)
    - [QueryValidatorSlashesRequest](#cosmos.distribution.v1beta1.QueryValidatorSlashesRequest)
    - [QueryValidatorSlashesResponse](#cosmos.distribution.v1beta1.QueryValidatorSlashesResponse)
  
    - [Query](#cosmos.distribution.v1beta1.Query)
  
- [cosmos/distribution/v1beta1/tx.proto](#cosmos/distribution/v1beta1/tx.proto)
    - [MsgFundCommunityPool](#cosmos.distribution.v1beta1.MsgFundCommunityPool)
    - [MsgFundCommunityPoolResponse](#cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse)
    - [MsgSetWithdrawAddress](#cosmos.distribution.v1beta1.MsgSetWithdrawAddress)
    - [MsgSetWithdrawAddressResponse](#cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse)
    - [MsgWithdrawDelegatorReward](#cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward)
    - [MsgWithdrawDelegatorRewardResponse](#cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse)
    - [MsgWithdrawValidatorCommission](#cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission)
    - [MsgWithdrawValidatorCommissionResponse](#cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse)
  
    - [Msg](#cosmos.distribution.v1beta1.Msg)
  
- [cosmos/evidence/v1beta1/evidence.proto](#cosmos/evidence/v1beta1/evidence.proto)
    - [Equivocation](#cosmos.evidence.v1beta1.Equivocation)
  
- [cosmos/evidence/v1beta1/genesis.proto](#cosmos/evidence/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.evidence.v1beta1.GenesisState)
  
- [cosmos/evidence/v1beta1/query.proto](#cosmos/evidence/v1beta1/query.proto)
    - [QueryAllEvidenceRequest](#cosmos.evidence.v1beta1.QueryAllEvidenceRequest)
    - [QueryAllEvidenceResponse](#cosmos.evidence.v1beta1.QueryAllEvidenceResponse)
    - [QueryEvidenceRequest](#cosmos.evidence.v1beta1.QueryEvidenceRequest)
    - [QueryEvidenceResponse](#cosmos.evidence.v1beta1.QueryEvidenceResponse)
  
    - [Query](#cosmos.evidence.v1beta1.Query)
  
- [cosmos/evidence/v1beta1/tx.proto](#cosmos/evidence/v1beta1/tx.proto)
    - [MsgSubmitEvidence](#cosmos.evidence.v1beta1.MsgSubmitEvidence)
    - [MsgSubmitEvidenceResponse](#cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse)
  
    - [Msg](#cosmos.evidence.v1beta1.Msg)
  
- [cosmos/feegrant/v1beta1/feegrant.proto](#cosmos/feegrant/v1beta1/feegrant.proto)
    - [AllowedMsgAllowance](#cosmos.feegrant.v1beta1.AllowedMsgAllowance)
    - [BasicAllowance](#cosmos.feegrant.v1beta1.BasicAllowance)
    - [Grant](#cosmos.feegrant.v1beta1.Grant)
    - [PeriodicAllowance](#cosmos.feegrant.v1beta1.PeriodicAllowance)
  
- [cosmos/feegrant/v1beta1/genesis.proto](#cosmos/feegrant/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.feegrant.v1beta1.GenesisState)
  
- [cosmos/feegrant/v1beta1/query.proto](#cosmos/feegrant/v1beta1/query.proto)
    - [QueryAllowanceRequest](#cosmos.feegrant.v1beta1.QueryAllowanceRequest)
    - [QueryAllowanceResponse](#cosmos.feegrant.v1beta1.QueryAllowanceResponse)
    - [QueryAllowancesByGranterRequest](#cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest)
    - [QueryAllowancesByGranterResponse](#cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse)
    - [QueryAllowancesRequest](#cosmos.feegrant.v1beta1.QueryAllowancesRequest)
    - [QueryAllowancesResponse](#cosmos.feegrant.v1beta1.QueryAllowancesResponse)
  
    - [Query](#cosmos.feegrant.v1beta1.Query)
  
- [cosmos/feegrant/v1beta1/tx.proto](#cosmos/feegrant/v1beta1/tx.proto)
    - [MsgGrantAllowance](#cosmos.feegrant.v1beta1.MsgGrantAllowance)
    - [MsgGrantAllowanceResponse](#cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse)
    - [MsgRevokeAllowance](#cosmos.feegrant.v1beta1.MsgRevokeAllowance)
    - [MsgRevokeAllowanceResponse](#cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse)
  
    - [Msg](#cosmos.feegrant.v1beta1.Msg)
  
- [cosmos/genutil/v1beta1/genesis.proto](#cosmos/genutil/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.genutil.v1beta1.GenesisState)
  
- [cosmos/gov/v1/gov.proto](#cosmos/gov/v1/gov.proto)
    - [Deposit](#cosmos.gov.v1.Deposit)
    - [DepositParams](#cosmos.gov.v1.DepositParams)
    - [Proposal](#cosmos.gov.v1.Proposal)
    - [TallyParams](#cosmos.gov.v1.TallyParams)
    - [TallyResult](#cosmos.gov.v1.TallyResult)
    - [Vote](#cosmos.gov.v1.Vote)
    - [VotingParams](#cosmos.gov.v1.VotingParams)
    - [WeightedVoteOption](#cosmos.gov.v1.WeightedVoteOption)
  
    - [ProposalStatus](#cosmos.gov.v1.ProposalStatus)
    - [VoteOption](#cosmos.gov.v1.VoteOption)
  
- [cosmos/gov/v1/genesis.proto](#cosmos/gov/v1/genesis.proto)
    - [GenesisState](#cosmos.gov.v1.GenesisState)
  
- [cosmos/gov/v1/query.proto](#cosmos/gov/v1/query.proto)
    - [QueryDepositRequest](#cosmos.gov.v1.QueryDepositRequest)
    - [QueryDepositResponse](#cosmos.gov.v1.QueryDepositResponse)
    - [QueryDepositsRequest](#cosmos.gov.v1.QueryDepositsRequest)
    - [QueryDepositsResponse](#cosmos.gov.v1.QueryDepositsResponse)
    - [QueryParamsRequest](#cosmos.gov.v1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.gov.v1.QueryParamsResponse)
    - [QueryProposalRequest](#cosmos.gov.v1.QueryProposalRequest)
    - [QueryProposalResponse](#cosmos.gov.v1.QueryProposalResponse)
    - [QueryProposalsRequest](#cosmos.gov.v1.QueryProposalsRequest)
    - [QueryProposalsResponse](#cosmos.gov.v1.QueryProposalsResponse)
    - [QueryTallyResultRequest](#cosmos.gov.v1.QueryTallyResultRequest)
    - [QueryTallyResultResponse](#cosmos.gov.v1.QueryTallyResultResponse)
    - [QueryVoteRequest](#cosmos.gov.v1.QueryVoteRequest)
    - [QueryVoteResponse](#cosmos.gov.v1.QueryVoteResponse)
    - [QueryVotesRequest](#cosmos.gov.v1.QueryVotesRequest)
    - [QueryVotesResponse](#cosmos.gov.v1.QueryVotesResponse)
  
    - [Query](#cosmos.gov.v1.Query)
  
- [cosmos/gov/v1/tx.proto](#cosmos/gov/v1/tx.proto)
    - [MsgDeposit](#cosmos.gov.v1.MsgDeposit)
    - [MsgDepositResponse](#cosmos.gov.v1.MsgDepositResponse)
    - [MsgExecLegacyContent](#cosmos.gov.v1.MsgExecLegacyContent)
    - [MsgExecLegacyContentResponse](#cosmos.gov.v1.MsgExecLegacyContentResponse)
    - [MsgSubmitProposal](#cosmos.gov.v1.MsgSubmitProposal)
    - [MsgSubmitProposalResponse](#cosmos.gov.v1.MsgSubmitProposalResponse)
    - [MsgVote](#cosmos.gov.v1.MsgVote)
    - [MsgVoteResponse](#cosmos.gov.v1.MsgVoteResponse)
    - [MsgVoteWeighted](#cosmos.gov.v1.MsgVoteWeighted)
    - [MsgVoteWeightedResponse](#cosmos.gov.v1.MsgVoteWeightedResponse)
  
    - [Msg](#cosmos.gov.v1.Msg)
  
- [cosmos/gov/v1beta1/gov.proto](#cosmos/gov/v1beta1/gov.proto)
    - [Deposit](#cosmos.gov.v1beta1.Deposit)
    - [DepositParams](#cosmos.gov.v1beta1.DepositParams)
    - [Proposal](#cosmos.gov.v1beta1.Proposal)
    - [TallyParams](#cosmos.gov.v1beta1.TallyParams)
    - [TallyResult](#cosmos.gov.v1beta1.TallyResult)
    - [TextProposal](#cosmos.gov.v1beta1.TextProposal)
    - [Vote](#cosmos.gov.v1beta1.Vote)
    - [VotingParams](#cosmos.gov.v1beta1.VotingParams)
    - [WeightedVoteOption](#cosmos.gov.v1beta1.WeightedVoteOption)
  
    - [ProposalStatus](#cosmos.gov.v1beta1.ProposalStatus)
    - [VoteOption](#cosmos.gov.v1beta1.VoteOption)
  
- [cosmos/gov/v1beta1/genesis.proto](#cosmos/gov/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.gov.v1beta1.GenesisState)
  
- [cosmos/gov/v1beta1/query.proto](#cosmos/gov/v1beta1/query.proto)
    - [QueryDepositRequest](#cosmos.gov.v1beta1.QueryDepositRequest)
    - [QueryDepositResponse](#cosmos.gov.v1beta1.QueryDepositResponse)
    - [QueryDepositsRequest](#cosmos.gov.v1beta1.QueryDepositsRequest)
    - [QueryDepositsResponse](#cosmos.gov.v1beta1.QueryDepositsResponse)
    - [QueryParamsRequest](#cosmos.gov.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.gov.v1beta1.QueryParamsResponse)
    - [QueryProposalRequest](#cosmos.gov.v1beta1.QueryProposalRequest)
    - [QueryProposalResponse](#cosmos.gov.v1beta1.QueryProposalResponse)
    - [QueryProposalsRequest](#cosmos.gov.v1beta1.QueryProposalsRequest)
    - [QueryProposalsResponse](#cosmos.gov.v1beta1.QueryProposalsResponse)
    - [QueryTallyResultRequest](#cosmos.gov.v1beta1.QueryTallyResultRequest)
    - [QueryTallyResultResponse](#cosmos.gov.v1beta1.QueryTallyResultResponse)
    - [QueryVoteRequest](#cosmos.gov.v1beta1.QueryVoteRequest)
    - [QueryVoteResponse](#cosmos.gov.v1beta1.QueryVoteResponse)
    - [QueryVotesRequest](#cosmos.gov.v1beta1.QueryVotesRequest)
    - [QueryVotesResponse](#cosmos.gov.v1beta1.QueryVotesResponse)
  
    - [Query](#cosmos.gov.v1beta1.Query)
  
- [cosmos/gov/v1beta1/tx.proto](#cosmos/gov/v1beta1/tx.proto)
    - [MsgDeposit](#cosmos.gov.v1beta1.MsgDeposit)
    - [MsgDepositResponse](#cosmos.gov.v1beta1.MsgDepositResponse)
    - [MsgSubmitProposal](#cosmos.gov.v1beta1.MsgSubmitProposal)
    - [MsgSubmitProposalResponse](#cosmos.gov.v1beta1.MsgSubmitProposalResponse)
    - [MsgVote](#cosmos.gov.v1beta1.MsgVote)
    - [MsgVoteResponse](#cosmos.gov.v1beta1.MsgVoteResponse)
    - [MsgVoteWeighted](#cosmos.gov.v1beta1.MsgVoteWeighted)
    - [MsgVoteWeightedResponse](#cosmos.gov.v1beta1.MsgVoteWeightedResponse)
  
    - [Msg](#cosmos.gov.v1beta1.Msg)
  
- [cosmos/group/v1/types.proto](#cosmos/group/v1/types.proto)
    - [DecisionPolicyWindows](#cosmos.group.v1.DecisionPolicyWindows)
    - [GroupInfo](#cosmos.group.v1.GroupInfo)
    - [GroupMember](#cosmos.group.v1.GroupMember)
    - [GroupPolicyInfo](#cosmos.group.v1.GroupPolicyInfo)
    - [Member](#cosmos.group.v1.Member)
    - [MemberRequest](#cosmos.group.v1.MemberRequest)
    - [PercentageDecisionPolicy](#cosmos.group.v1.PercentageDecisionPolicy)
    - [Proposal](#cosmos.group.v1.Proposal)
    - [TallyResult](#cosmos.group.v1.TallyResult)
    - [ThresholdDecisionPolicy](#cosmos.group.v1.ThresholdDecisionPolicy)
    - [Vote](#cosmos.group.v1.Vote)
  
    - [ProposalExecutorResult](#cosmos.group.v1.ProposalExecutorResult)
    - [ProposalStatus](#cosmos.group.v1.ProposalStatus)
    - [VoteOption](#cosmos.group.v1.VoteOption)
  
- [cosmos/group/v1/events.proto](#cosmos/group/v1/events.proto)
    - [EventCreateGroup](#cosmos.group.v1.EventCreateGroup)
    - [EventCreateGroupPolicy](#cosmos.group.v1.EventCreateGroupPolicy)
    - [EventExec](#cosmos.group.v1.EventExec)
    - [EventLeaveGroup](#cosmos.group.v1.EventLeaveGroup)
    - [EventSubmitProposal](#cosmos.group.v1.EventSubmitProposal)
    - [EventUpdateGroup](#cosmos.group.v1.EventUpdateGroup)
    - [EventUpdateGroupPolicy](#cosmos.group.v1.EventUpdateGroupPolicy)
    - [EventVote](#cosmos.group.v1.EventVote)
    - [EventWithdrawProposal](#cosmos.group.v1.EventWithdrawProposal)
  
- [cosmos/group/v1/genesis.proto](#cosmos/group/v1/genesis.proto)
    - [GenesisState](#cosmos.group.v1.GenesisState)
  
- [cosmos/group/v1/query.proto](#cosmos/group/v1/query.proto)
    - [QueryGroupInfoRequest](#cosmos.group.v1.QueryGroupInfoRequest)
    - [QueryGroupInfoResponse](#cosmos.group.v1.QueryGroupInfoResponse)
    - [QueryGroupMembersRequest](#cosmos.group.v1.QueryGroupMembersRequest)
    - [QueryGroupMembersResponse](#cosmos.group.v1.QueryGroupMembersResponse)
    - [QueryGroupPoliciesByAdminRequest](#cosmos.group.v1.QueryGroupPoliciesByAdminRequest)
    - [QueryGroupPoliciesByAdminResponse](#cosmos.group.v1.QueryGroupPoliciesByAdminResponse)
    - [QueryGroupPoliciesByGroupRequest](#cosmos.group.v1.QueryGroupPoliciesByGroupRequest)
    - [QueryGroupPoliciesByGroupResponse](#cosmos.group.v1.QueryGroupPoliciesByGroupResponse)
    - [QueryGroupPolicyInfoRequest](#cosmos.group.v1.QueryGroupPolicyInfoRequest)
    - [QueryGroupPolicyInfoResponse](#cosmos.group.v1.QueryGroupPolicyInfoResponse)
    - [QueryGroupsByAdminRequest](#cosmos.group.v1.QueryGroupsByAdminRequest)
    - [QueryGroupsByAdminResponse](#cosmos.group.v1.QueryGroupsByAdminResponse)
    - [QueryGroupsByMemberRequest](#cosmos.group.v1.QueryGroupsByMemberRequest)
    - [QueryGroupsByMemberResponse](#cosmos.group.v1.QueryGroupsByMemberResponse)
    - [QueryProposalRequest](#cosmos.group.v1.QueryProposalRequest)
    - [QueryProposalResponse](#cosmos.group.v1.QueryProposalResponse)
    - [QueryProposalsByGroupPolicyRequest](#cosmos.group.v1.QueryProposalsByGroupPolicyRequest)
    - [QueryProposalsByGroupPolicyResponse](#cosmos.group.v1.QueryProposalsByGroupPolicyResponse)
    - [QueryTallyResultRequest](#cosmos.group.v1.QueryTallyResultRequest)
    - [QueryTallyResultResponse](#cosmos.group.v1.QueryTallyResultResponse)
    - [QueryVoteByProposalVoterRequest](#cosmos.group.v1.QueryVoteByProposalVoterRequest)
    - [QueryVoteByProposalVoterResponse](#cosmos.group.v1.QueryVoteByProposalVoterResponse)
    - [QueryVotesByProposalRequest](#cosmos.group.v1.QueryVotesByProposalRequest)
    - [QueryVotesByProposalResponse](#cosmos.group.v1.QueryVotesByProposalResponse)
    - [QueryVotesByVoterRequest](#cosmos.group.v1.QueryVotesByVoterRequest)
    - [QueryVotesByVoterResponse](#cosmos.group.v1.QueryVotesByVoterResponse)
  
    - [Query](#cosmos.group.v1.Query)
  
- [cosmos/group/v1/tx.proto](#cosmos/group/v1/tx.proto)
    - [MsgCreateGroup](#cosmos.group.v1.MsgCreateGroup)
    - [MsgCreateGroupPolicy](#cosmos.group.v1.MsgCreateGroupPolicy)
    - [MsgCreateGroupPolicyResponse](#cosmos.group.v1.MsgCreateGroupPolicyResponse)
    - [MsgCreateGroupResponse](#cosmos.group.v1.MsgCreateGroupResponse)
    - [MsgCreateGroupWithPolicy](#cosmos.group.v1.MsgCreateGroupWithPolicy)
    - [MsgCreateGroupWithPolicyResponse](#cosmos.group.v1.MsgCreateGroupWithPolicyResponse)
    - [MsgExec](#cosmos.group.v1.MsgExec)
    - [MsgExecResponse](#cosmos.group.v1.MsgExecResponse)
    - [MsgLeaveGroup](#cosmos.group.v1.MsgLeaveGroup)
    - [MsgLeaveGroupResponse](#cosmos.group.v1.MsgLeaveGroupResponse)
    - [MsgSubmitProposal](#cosmos.group.v1.MsgSubmitProposal)
    - [MsgSubmitProposalResponse](#cosmos.group.v1.MsgSubmitProposalResponse)
    - [MsgUpdateGroupAdmin](#cosmos.group.v1.MsgUpdateGroupAdmin)
    - [MsgUpdateGroupAdminResponse](#cosmos.group.v1.MsgUpdateGroupAdminResponse)
    - [MsgUpdateGroupMembers](#cosmos.group.v1.MsgUpdateGroupMembers)
    - [MsgUpdateGroupMembersResponse](#cosmos.group.v1.MsgUpdateGroupMembersResponse)
    - [MsgUpdateGroupMetadata](#cosmos.group.v1.MsgUpdateGroupMetadata)
    - [MsgUpdateGroupMetadataResponse](#cosmos.group.v1.MsgUpdateGroupMetadataResponse)
    - [MsgUpdateGroupPolicyAdmin](#cosmos.group.v1.MsgUpdateGroupPolicyAdmin)
    - [MsgUpdateGroupPolicyAdminResponse](#cosmos.group.v1.MsgUpdateGroupPolicyAdminResponse)
    - [MsgUpdateGroupPolicyDecisionPolicy](#cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicy)
    - [MsgUpdateGroupPolicyDecisionPolicyResponse](#cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicyResponse)
    - [MsgUpdateGroupPolicyMetadata](#cosmos.group.v1.MsgUpdateGroupPolicyMetadata)
    - [MsgUpdateGroupPolicyMetadataResponse](#cosmos.group.v1.MsgUpdateGroupPolicyMetadataResponse)
    - [MsgVote](#cosmos.group.v1.MsgVote)
    - [MsgVoteResponse](#cosmos.group.v1.MsgVoteResponse)
    - [MsgWithdrawProposal](#cosmos.group.v1.MsgWithdrawProposal)
    - [MsgWithdrawProposalResponse](#cosmos.group.v1.MsgWithdrawProposalResponse)
  
    - [Exec](#cosmos.group.v1.Exec)
  
    - [Msg](#cosmos.group.v1.Msg)
  
- [cosmos/mint/v1beta1/mint.proto](#cosmos/mint/v1beta1/mint.proto)
    - [Minter](#cosmos.mint.v1beta1.Minter)
    - [Params](#cosmos.mint.v1beta1.Params)
  
- [cosmos/mint/v1beta1/genesis.proto](#cosmos/mint/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.mint.v1beta1.GenesisState)
  
- [cosmos/mint/v1beta1/query.proto](#cosmos/mint/v1beta1/query.proto)
    - [QueryAnnualProvisionsRequest](#cosmos.mint.v1beta1.QueryAnnualProvisionsRequest)
    - [QueryAnnualProvisionsResponse](#cosmos.mint.v1beta1.QueryAnnualProvisionsResponse)
    - [QueryInflationRequest](#cosmos.mint.v1beta1.QueryInflationRequest)
    - [QueryInflationResponse](#cosmos.mint.v1beta1.QueryInflationResponse)
    - [QueryParamsRequest](#cosmos.mint.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.mint.v1beta1.QueryParamsResponse)
  
    - [Query](#cosmos.mint.v1beta1.Query)
  
- [cosmos/nft/v1beta1/event.proto](#cosmos/nft/v1beta1/event.proto)
    - [EventBurn](#cosmos.nft.v1beta1.EventBurn)
    - [EventMint](#cosmos.nft.v1beta1.EventMint)
    - [EventSend](#cosmos.nft.v1beta1.EventSend)
  
- [cosmos/nft/v1beta1/nft.proto](#cosmos/nft/v1beta1/nft.proto)
    - [Class](#cosmos.nft.v1beta1.Class)
    - [NFT](#cosmos.nft.v1beta1.NFT)
  
- [cosmos/nft/v1beta1/genesis.proto](#cosmos/nft/v1beta1/genesis.proto)
    - [Entry](#cosmos.nft.v1beta1.Entry)
    - [GenesisState](#cosmos.nft.v1beta1.GenesisState)
  
- [cosmos/nft/v1beta1/query.proto](#cosmos/nft/v1beta1/query.proto)
    - [QueryBalanceRequest](#cosmos.nft.v1beta1.QueryBalanceRequest)
    - [QueryBalanceResponse](#cosmos.nft.v1beta1.QueryBalanceResponse)
    - [QueryClassRequest](#cosmos.nft.v1beta1.QueryClassRequest)
    - [QueryClassResponse](#cosmos.nft.v1beta1.QueryClassResponse)
    - [QueryClassesRequest](#cosmos.nft.v1beta1.QueryClassesRequest)
    - [QueryClassesResponse](#cosmos.nft.v1beta1.QueryClassesResponse)
    - [QueryNFTRequest](#cosmos.nft.v1beta1.QueryNFTRequest)
    - [QueryNFTResponse](#cosmos.nft.v1beta1.QueryNFTResponse)
    - [QueryNFTsRequest](#cosmos.nft.v1beta1.QueryNFTsRequest)
    - [QueryNFTsResponse](#cosmos.nft.v1beta1.QueryNFTsResponse)
    - [QueryOwnerRequest](#cosmos.nft.v1beta1.QueryOwnerRequest)
    - [QueryOwnerResponse](#cosmos.nft.v1beta1.QueryOwnerResponse)
    - [QuerySupplyRequest](#cosmos.nft.v1beta1.QuerySupplyRequest)
    - [QuerySupplyResponse](#cosmos.nft.v1beta1.QuerySupplyResponse)
  
    - [Query](#cosmos.nft.v1beta1.Query)
  
- [cosmos/nft/v1beta1/tx.proto](#cosmos/nft/v1beta1/tx.proto)
    - [MsgSend](#cosmos.nft.v1beta1.MsgSend)
    - [MsgSendResponse](#cosmos.nft.v1beta1.MsgSendResponse)
  
    - [Msg](#cosmos.nft.v1beta1.Msg)
  
- [cosmos/orm/module/v1alpha1/module.proto](#cosmos/orm/module/v1alpha1/module.proto)
    - [Module](#cosmos.orm.module.v1alpha1.Module)
  
- [cosmos/orm/v1/orm.proto](#cosmos/orm/v1/orm.proto)
    - [PrimaryKeyDescriptor](#cosmos.orm.v1.PrimaryKeyDescriptor)
    - [SecondaryIndexDescriptor](#cosmos.orm.v1.SecondaryIndexDescriptor)
    - [SingletonDescriptor](#cosmos.orm.v1.SingletonDescriptor)
    - [TableDescriptor](#cosmos.orm.v1.TableDescriptor)
  
    - [File-level Extensions](#cosmos/orm/v1/orm.proto-extensions)
    - [File-level Extensions](#cosmos/orm/v1/orm.proto-extensions)
  
- [cosmos/orm/v1alpha1/schema.proto](#cosmos/orm/v1alpha1/schema.proto)
    - [ModuleSchemaDescriptor](#cosmos.orm.v1alpha1.ModuleSchemaDescriptor)
    - [ModuleSchemaDescriptor.FileEntry](#cosmos.orm.v1alpha1.ModuleSchemaDescriptor.FileEntry)
  
    - [StorageType](#cosmos.orm.v1alpha1.StorageType)
  
    - [File-level Extensions](#cosmos/orm/v1alpha1/schema.proto-extensions)
  
- [cosmos/params/v1beta1/params.proto](#cosmos/params/v1beta1/params.proto)
    - [ParamChange](#cosmos.params.v1beta1.ParamChange)
    - [ParameterChangeProposal](#cosmos.params.v1beta1.ParameterChangeProposal)
  
- [cosmos/params/v1beta1/query.proto](#cosmos/params/v1beta1/query.proto)
    - [QueryParamsRequest](#cosmos.params.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.params.v1beta1.QueryParamsResponse)
    - [QuerySubspacesRequest](#cosmos.params.v1beta1.QuerySubspacesRequest)
    - [QuerySubspacesResponse](#cosmos.params.v1beta1.QuerySubspacesResponse)
    - [Subspace](#cosmos.params.v1beta1.Subspace)
  
    - [Query](#cosmos.params.v1beta1.Query)
  
- [cosmos/slashing/v1beta1/slashing.proto](#cosmos/slashing/v1beta1/slashing.proto)
    - [Params](#cosmos.slashing.v1beta1.Params)
    - [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo)
  
- [cosmos/slashing/v1beta1/genesis.proto](#cosmos/slashing/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.slashing.v1beta1.GenesisState)
    - [MissedBlock](#cosmos.slashing.v1beta1.MissedBlock)
    - [SigningInfo](#cosmos.slashing.v1beta1.SigningInfo)
    - [ValidatorMissedBlocks](#cosmos.slashing.v1beta1.ValidatorMissedBlocks)
  
- [cosmos/slashing/v1beta1/query.proto](#cosmos/slashing/v1beta1/query.proto)
    - [QueryParamsRequest](#cosmos.slashing.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.slashing.v1beta1.QueryParamsResponse)
    - [QuerySigningInfoRequest](#cosmos.slashing.v1beta1.QuerySigningInfoRequest)
    - [QuerySigningInfoResponse](#cosmos.slashing.v1beta1.QuerySigningInfoResponse)
    - [QuerySigningInfosRequest](#cosmos.slashing.v1beta1.QuerySigningInfosRequest)
    - [QuerySigningInfosResponse](#cosmos.slashing.v1beta1.QuerySigningInfosResponse)
  
    - [Query](#cosmos.slashing.v1beta1.Query)
  
- [cosmos/slashing/v1beta1/tx.proto](#cosmos/slashing/v1beta1/tx.proto)
    - [MsgUnjail](#cosmos.slashing.v1beta1.MsgUnjail)
    - [MsgUnjailResponse](#cosmos.slashing.v1beta1.MsgUnjailResponse)
  
    - [Msg](#cosmos.slashing.v1beta1.Msg)
  
- [cosmos/staking/v1beta1/authz.proto](#cosmos/staking/v1beta1/authz.proto)
    - [StakeAuthorization](#cosmos.staking.v1beta1.StakeAuthorization)
    - [StakeAuthorization.Validators](#cosmos.staking.v1beta1.StakeAuthorization.Validators)
  
    - [AuthorizationType](#cosmos.staking.v1beta1.AuthorizationType)
  
- [cosmos/staking/v1beta1/staking.proto](#cosmos/staking/v1beta1/staking.proto)
    - [Commission](#cosmos.staking.v1beta1.Commission)
    - [CommissionRates](#cosmos.staking.v1beta1.CommissionRates)
    - [DVPair](#cosmos.staking.v1beta1.DVPair)
    - [DVPairs](#cosmos.staking.v1beta1.DVPairs)
    - [DVVTriplet](#cosmos.staking.v1beta1.DVVTriplet)
    - [DVVTriplets](#cosmos.staking.v1beta1.DVVTriplets)
    - [Delegation](#cosmos.staking.v1beta1.Delegation)
    - [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse)
    - [Description](#cosmos.staking.v1beta1.Description)
    - [HistoricalInfo](#cosmos.staking.v1beta1.HistoricalInfo)
    - [Params](#cosmos.staking.v1beta1.Params)
    - [Pool](#cosmos.staking.v1beta1.Pool)
    - [Redelegation](#cosmos.staking.v1beta1.Redelegation)
    - [RedelegationEntry](#cosmos.staking.v1beta1.RedelegationEntry)
    - [RedelegationEntryResponse](#cosmos.staking.v1beta1.RedelegationEntryResponse)
    - [RedelegationResponse](#cosmos.staking.v1beta1.RedelegationResponse)
    - [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation)
    - [UnbondingDelegationEntry](#cosmos.staking.v1beta1.UnbondingDelegationEntry)
    - [ValAddresses](#cosmos.staking.v1beta1.ValAddresses)
    - [Validator](#cosmos.staking.v1beta1.Validator)
  
    - [BondStatus](#cosmos.staking.v1beta1.BondStatus)
  
- [cosmos/staking/v1beta1/genesis.proto](#cosmos/staking/v1beta1/genesis.proto)
    - [GenesisState](#cosmos.staking.v1beta1.GenesisState)
    - [LastValidatorPower](#cosmos.staking.v1beta1.LastValidatorPower)
  
- [cosmos/staking/v1beta1/query.proto](#cosmos/staking/v1beta1/query.proto)
    - [QueryDelegationRequest](#cosmos.staking.v1beta1.QueryDelegationRequest)
    - [QueryDelegationResponse](#cosmos.staking.v1beta1.QueryDelegationResponse)
    - [QueryDelegatorDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest)
    - [QueryDelegatorDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse)
    - [QueryDelegatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest)
    - [QueryDelegatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse)
    - [QueryDelegatorValidatorRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorRequest)
    - [QueryDelegatorValidatorResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorResponse)
    - [QueryDelegatorValidatorsRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest)
    - [QueryDelegatorValidatorsResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse)
    - [QueryHistoricalInfoRequest](#cosmos.staking.v1beta1.QueryHistoricalInfoRequest)
    - [QueryHistoricalInfoResponse](#cosmos.staking.v1beta1.QueryHistoricalInfoResponse)
    - [QueryParamsRequest](#cosmos.staking.v1beta1.QueryParamsRequest)
    - [QueryParamsResponse](#cosmos.staking.v1beta1.QueryParamsResponse)
    - [QueryPoolRequest](#cosmos.staking.v1beta1.QueryPoolRequest)
    - [QueryPoolResponse](#cosmos.staking.v1beta1.QueryPoolResponse)
    - [QueryRedelegationsRequest](#cosmos.staking.v1beta1.QueryRedelegationsRequest)
    - [QueryRedelegationsResponse](#cosmos.staking.v1beta1.QueryRedelegationsResponse)
    - [QueryUnbondingDelegationRequest](#cosmos.staking.v1beta1.QueryUnbondingDelegationRequest)
    - [QueryUnbondingDelegationResponse](#cosmos.staking.v1beta1.QueryUnbondingDelegationResponse)
    - [QueryValidatorDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorDelegationsRequest)
    - [QueryValidatorDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorDelegationsResponse)
    - [QueryValidatorRequest](#cosmos.staking.v1beta1.QueryValidatorRequest)
    - [QueryValidatorResponse](#cosmos.staking.v1beta1.QueryValidatorResponse)
    - [QueryValidatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest)
    - [QueryValidatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse)
    - [QueryValidatorsRequest](#cosmos.staking.v1beta1.QueryValidatorsRequest)
    - [QueryValidatorsResponse](#cosmos.staking.v1beta1.QueryValidatorsResponse)
  
    - [Query](#cosmos.staking.v1beta1.Query)
  
- [cosmos/staking/v1beta1/tx.proto](#cosmos/staking/v1beta1/tx.proto)
    - [MsgBeginRedelegate](#cosmos.staking.v1beta1.MsgBeginRedelegate)
    - [MsgBeginRedelegateResponse](#cosmos.staking.v1beta1.MsgBeginRedelegateResponse)
    - [MsgCancelUnbondingDelegation](#cosmos.staking.v1beta1.MsgCancelUnbondingDelegation)
    - [MsgCancelUnbondingDelegationResponse](#cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse)
    - [MsgCreateValidator](#cosmos.staking.v1beta1.MsgCreateValidator)
    - [MsgCreateValidatorResponse](#cosmos.staking.v1beta1.MsgCreateValidatorResponse)
    - [MsgDelegate](#cosmos.staking.v1beta1.MsgDelegate)
    - [MsgDelegateResponse](#cosmos.staking.v1beta1.MsgDelegateResponse)
    - [MsgEditValidator](#cosmos.staking.v1beta1.MsgEditValidator)
    - [MsgEditValidatorResponse](#cosmos.staking.v1beta1.MsgEditValidatorResponse)
    - [MsgUndelegate](#cosmos.staking.v1beta1.MsgUndelegate)
    - [MsgUndelegateResponse](#cosmos.staking.v1beta1.MsgUndelegateResponse)
  
    - [Msg](#cosmos.staking.v1beta1.Msg)
  
- [cosmos/tx/signing/v1beta1/signing.proto](#cosmos/tx/signing/v1beta1/signing.proto)
    - [SignatureDescriptor](#cosmos.tx.signing.v1beta1.SignatureDescriptor)
    - [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data)
    - [SignatureDescriptor.Data.Multi](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi)
    - [SignatureDescriptor.Data.Single](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single)
    - [SignatureDescriptors](#cosmos.tx.signing.v1beta1.SignatureDescriptors)
  
    - [SignMode](#cosmos.tx.signing.v1beta1.SignMode)
  
- [cosmos/tx/v1beta1/tx.proto](#cosmos/tx/v1beta1/tx.proto)
    - [AuthInfo](#cosmos.tx.v1beta1.AuthInfo)
    - [AuxSignerData](#cosmos.tx.v1beta1.AuxSignerData)
    - [Fee](#cosmos.tx.v1beta1.Fee)
    - [ModeInfo](#cosmos.tx.v1beta1.ModeInfo)
    - [ModeInfo.Multi](#cosmos.tx.v1beta1.ModeInfo.Multi)
    - [ModeInfo.Single](#cosmos.tx.v1beta1.ModeInfo.Single)
    - [SignDoc](#cosmos.tx.v1beta1.SignDoc)
    - [SignDocDirectAux](#cosmos.tx.v1beta1.SignDocDirectAux)
    - [SignerInfo](#cosmos.tx.v1beta1.SignerInfo)
    - [Tip](#cosmos.tx.v1beta1.Tip)
    - [Tx](#cosmos.tx.v1beta1.Tx)
    - [TxBody](#cosmos.tx.v1beta1.TxBody)
    - [TxRaw](#cosmos.tx.v1beta1.TxRaw)
  
- [cosmos/tx/v1beta1/service.proto](#cosmos/tx/v1beta1/service.proto)
    - [BroadcastTxRequest](#cosmos.tx.v1beta1.BroadcastTxRequest)
    - [BroadcastTxResponse](#cosmos.tx.v1beta1.BroadcastTxResponse)
    - [GetBlockWithTxsRequest](#cosmos.tx.v1beta1.GetBlockWithTxsRequest)
    - [GetBlockWithTxsResponse](#cosmos.tx.v1beta1.GetBlockWithTxsResponse)
    - [GetTxRequest](#cosmos.tx.v1beta1.GetTxRequest)
    - [GetTxResponse](#cosmos.tx.v1beta1.GetTxResponse)
    - [GetTxsEventRequest](#cosmos.tx.v1beta1.GetTxsEventRequest)
    - [GetTxsEventResponse](#cosmos.tx.v1beta1.GetTxsEventResponse)
    - [SimulateRequest](#cosmos.tx.v1beta1.SimulateRequest)
    - [SimulateResponse](#cosmos.tx.v1beta1.SimulateResponse)
  
    - [BroadcastMode](#cosmos.tx.v1beta1.BroadcastMode)
    - [OrderBy](#cosmos.tx.v1beta1.OrderBy)
  
    - [Service](#cosmos.tx.v1beta1.Service)
  
- [cosmos/upgrade/v1beta1/upgrade.proto](#cosmos/upgrade/v1beta1/upgrade.proto)
    - [CancelSoftwareUpgradeProposal](#cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal)
    - [ModuleVersion](#cosmos.upgrade.v1beta1.ModuleVersion)
    - [Plan](#cosmos.upgrade.v1beta1.Plan)
    - [SoftwareUpgradeProposal](#cosmos.upgrade.v1beta1.SoftwareUpgradeProposal)
  
- [cosmos/upgrade/v1beta1/query.proto](#cosmos/upgrade/v1beta1/query.proto)
    - [QueryAppliedPlanRequest](#cosmos.upgrade.v1beta1.QueryAppliedPlanRequest)
    - [QueryAppliedPlanResponse](#cosmos.upgrade.v1beta1.QueryAppliedPlanResponse)
    - [QueryAuthorityRequest](#cosmos.upgrade.v1beta1.QueryAuthorityRequest)
    - [QueryAuthorityResponse](#cosmos.upgrade.v1beta1.QueryAuthorityResponse)
    - [QueryCurrentPlanRequest](#cosmos.upgrade.v1beta1.QueryCurrentPlanRequest)
    - [QueryCurrentPlanResponse](#cosmos.upgrade.v1beta1.QueryCurrentPlanResponse)
    - [QueryModuleVersionsRequest](#cosmos.upgrade.v1beta1.QueryModuleVersionsRequest)
    - [QueryModuleVersionsResponse](#cosmos.upgrade.v1beta1.QueryModuleVersionsResponse)
    - [QueryUpgradedConsensusStateRequest](#cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest)
    - [QueryUpgradedConsensusStateResponse](#cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse)
  
    - [Query](#cosmos.upgrade.v1beta1.Query)
  
- [cosmos/upgrade/v1beta1/tx.proto](#cosmos/upgrade/v1beta1/tx.proto)
    - [MsgCancelUpgrade](#cosmos.upgrade.v1beta1.MsgCancelUpgrade)
    - [MsgCancelUpgradeResponse](#cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse)
    - [MsgSoftwareUpgrade](#cosmos.upgrade.v1beta1.MsgSoftwareUpgrade)
    - [MsgSoftwareUpgradeResponse](#cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse)
  
    - [Msg](#cosmos.upgrade.v1beta1.Msg)
  
- [cosmos/vesting/v1beta1/vesting.proto](#cosmos/vesting/v1beta1/vesting.proto)
    - [BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount)
    - [ContinuousVestingAccount](#cosmos.vesting.v1beta1.ContinuousVestingAccount)
    - [DelayedVestingAccount](#cosmos.vesting.v1beta1.DelayedVestingAccount)
    - [Period](#cosmos.vesting.v1beta1.Period)
    - [PeriodicVestingAccount](#cosmos.vesting.v1beta1.PeriodicVestingAccount)
    - [PermanentLockedAccount](#cosmos.vesting.v1beta1.PermanentLockedAccount)
  
- [cosmos/vesting/v1beta1/tx.proto](#cosmos/vesting/v1beta1/tx.proto)
    - [MsgCreatePeriodicVestingAccount](#cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount)
    - [MsgCreatePeriodicVestingAccountResponse](#cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse)
    - [MsgCreatePermanentLockedAccount](#cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccount)
    - [MsgCreatePermanentLockedAccountResponse](#cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccountResponse)
    - [MsgCreateVestingAccount](#cosmos.vesting.v1beta1.MsgCreateVestingAccount)
    - [MsgCreateVestingAccountResponse](#cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse)
  
    - [Msg](#cosmos.vesting.v1beta1.Msg)
  
- [tendermint/libs/bits/types.proto](#tendermint/libs/bits/types.proto)
    - [BitArray](#tendermint.libs.bits.BitArray)
  
- [Scalar Value Types](#scalar-value-types)



<a name="cosmos/app/v1alpha1/module.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/app/v1alpha1/module.proto



<a name="cosmos.app.v1alpha1.MigrateFromInfo"></a>

### MigrateFromInfo
MigrateFromInfo is information on a module version that a newer module
can migrate from.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  | module is the fully-qualified protobuf name of the module config object for the previous module version, ex: "cosmos.group.module.v1.Module". |






<a name="cosmos.app.v1alpha1.ModuleDescriptor"></a>

### ModuleDescriptor
ModuleDescriptor describes an app module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `go_import` | [string](#string) |  | go_import names the package that should be imported by an app to load the module in the runtime module registry. It is required to make debugging of configuration errors easier for users. |
| `use_package` | [PackageReference](#cosmos.app.v1alpha1.PackageReference) | repeated | use_package refers to a protobuf package that this module uses and exposes to the world. In an app, only one module should "use" or own a single protobuf package. It is assumed that the module uses all of the .proto files in a single package. |
| `can_migrate_from` | [MigrateFromInfo](#cosmos.app.v1alpha1.MigrateFromInfo) | repeated | can_migrate_from defines which module versions this module can migrate state from. The framework will check that one module version is able to migrate from a previous module version before attempting to update its config. It is assumed that modules can transitively migrate from earlier versions. For instance if v3 declares it can migrate from v2, and v2 declares it can migrate from v1, the framework knows how to migrate from v1 to v3, assuming all 3 module versions are registered at runtime. |






<a name="cosmos.app.v1alpha1.PackageReference"></a>

### PackageReference
PackageReference is a reference to a protobuf package used by a module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name is the fully-qualified name of the package. |
| `revision` | [uint32](#uint32) |  | revision is the optional revision of the package that is being used. Protobuf packages used in Cosmos should generally have a major version as the last part of the package name, ex. foo.bar.baz.v1. The revision of a package can be thought of as the minor version of a package which has additional backwards compatible definitions that weren't present in a previous version.

A package should indicate its revision with a source code comment above the package declaration in one of its files containing the text "Revision N" where N is an integer revision. All packages start at revision 0 the first time they are released in a module.

When a new version of a module is released and items are added to existing .proto files, these definitions should contain comments of the form "Since Revision N" where N is an integer revision.

When the module runtime starts up, it will check the pinned proto image and panic if there are runtime protobuf definitions that are not in the pinned descriptor which do not have a "Since Revision N" comment or have a "Since Revision N" comment where N is <= to the revision specified here. This indicates that the protobuf files have been updated, but the pinned file descriptor hasn't.

If there are items in the pinned file descriptor with a revision greater than the value indicated here, this will also cause a panic as it may mean that the pinned descriptor for a legacy module has been improperly updated or that there is some other versioning discrepancy. Runtime protobuf definitions will also be checked for compatibility with pinned file descriptors to make sure there are no incompatible changes.

This behavior ensures that: * pinned proto images are up-to-date * protobuf files are carefully annotated with revision comments which are important good client UX * protobuf files are changed in backwards and forwards compatible ways |





 <!-- end messages -->

 <!-- end enums -->


<a name="cosmos/app/v1alpha1/module.proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| `module` | ModuleDescriptor | .google.protobuf.MessageOptions | 57193479 | module indicates that this proto type is a config object for an app module and optionally provides other descriptive information about the module. It is recommended that a new module config object and go module is versioned for every state machine breaking version of a module. The recommended pattern for doing this is to put module config objects in a separate proto package from the API they expose. Ex: the cosmos.group.v1 API would be exposed by module configs cosmos.group.module.v1, cosmos.group.module.v2, etc. |

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/app/module/v1alpha1/module.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/app/module/v1alpha1/module.proto



<a name="cosmos.app.module.v1alpha1.Module"></a>

### Module
Module is the module config object for the cosmos.app v1 app module.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/app/v1alpha1/config.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/app/v1alpha1/config.proto



<a name="cosmos.app.v1alpha1.Config"></a>

### Config
Config represents the configuration for a Cosmos SDK ABCI app.
It is intended that all state machine logic including the version of
baseapp and tx handlers (and possibly even Tendermint) that an app needs
can be described in a config object. For compatibility, the framework should
allow a mixture of declarative and imperative app wiring, however, apps
that strive for the maximum ease of maintainability should be able to describe
their state machine with a config object alone.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `modules` | [ModuleConfig](#cosmos.app.v1alpha1.ModuleConfig) | repeated | modules are the module configurations for the app. |






<a name="cosmos.app.v1alpha1.ModuleConfig"></a>

### ModuleConfig
ModuleConfig is a module configuration for an app.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name is the unique name of the module within the app. It should be a name that persists between different versions of a module so that modules can be smoothly upgraded to new versions.

For example, for the module cosmos.bank.module.v1.Module, we may chose to simply name the module "bank" in the app. When we upgrade to cosmos.bank.module.v2.Module, the app-specific name "bank" stays the same and the framework knows that the v2 module should receive all the same state that the v1 module had. Note: modules should provide info on which versions they can migrate from in the ModuleDescriptor.can_migration_from field. |
| `config` | [google.protobuf.Any](#google.protobuf.Any) |  | config is the config object for the module. Module config messages should define a ModuleDescriptor using the cosmos.app.v1alpha1.is_module extension. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/app/v1alpha1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/app/v1alpha1/query.proto



<a name="cosmos.app.v1alpha1.QueryConfigRequest"></a>

### QueryConfigRequest
QueryConfigRequest is the Query/Config request type.






<a name="cosmos.app.v1alpha1.QueryConfigResponse"></a>

### QueryConfigResponse
QueryConfigRequest is the Query/Config response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `config` | [Config](#cosmos.app.v1alpha1.Config) |  | config is the current app config. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.app.v1alpha1.Query"></a>

### Query
Query is the app module query service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Config` | [QueryConfigRequest](#cosmos.app.v1alpha1.QueryConfigRequest) | [QueryConfigResponse](#cosmos.app.v1alpha1.QueryConfigResponse) | Config returns the current app config. | |

 <!-- end services -->



<a name="cosmos/auth/v1beta1/auth.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/auth/v1beta1/auth.proto



<a name="cosmos.auth.v1beta1.BaseAccount"></a>

### BaseAccount
BaseAccount defines a base account type. It contains all the necessary fields
for basic account functionality. Any custom account type should extend this
type for additional functionality (e.g. vesting).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `pub_key` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `account_number` | [uint64](#uint64) |  |  |
| `sequence` | [uint64](#uint64) |  |  |






<a name="cosmos.auth.v1beta1.ModuleAccount"></a>

### ModuleAccount
ModuleAccount defines an account for modules that holds coins on a pool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_account` | [BaseAccount](#cosmos.auth.v1beta1.BaseAccount) |  |  |
| `name` | [string](#string) |  |  |
| `permissions` | [string](#string) | repeated |  |






<a name="cosmos.auth.v1beta1.Params"></a>

### Params
Params defines the parameters for the auth module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_memo_characters` | [uint64](#uint64) |  |  |
| `tx_sig_limit` | [uint64](#uint64) |  |  |
| `tx_size_cost_per_byte` | [uint64](#uint64) |  |  |
| `sig_verify_cost_ed25519` | [uint64](#uint64) |  |  |
| `sig_verify_cost_secp256k1` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/auth/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/auth/v1beta1/genesis.proto



<a name="cosmos.auth.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the auth module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.auth.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated | accounts are the accounts present at genesis. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/query/v1beta1/pagination.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/query/v1beta1/pagination.proto



<a name="cosmos.base.query.v1beta1.PageRequest"></a>

### PageRequest
PageRequest is to be embedded in gRPC request messages for efficient
pagination. Ex:

 message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  | key is a value returned in PageResponse.next_key to begin querying the next page most efficiently. Only one of offset or key should be set. |
| `offset` | [uint64](#uint64) |  | offset is a numeric offset that can be used when key is unavailable. It is less efficient than using key. Only one of offset or key should be set. |
| `limit` | [uint64](#uint64) |  | limit is the total number of results to be returned in the result page. If left empty it will default to a value to be set by each app. |
| `count_total` | [bool](#bool) |  | count_total is set to true to indicate that the result set should include a count of the total number of items available for pagination in UIs. count_total is only respected when offset is used. It is ignored when key is set. |
| `reverse` | [bool](#bool) |  | reverse is set to true if results are to be returned in the descending order.

Since: cosmos-sdk 0.43 |






<a name="cosmos.base.query.v1beta1.PageResponse"></a>

### PageResponse
PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `next_key` | [bytes](#bytes) |  | next_key is the key to be passed to PageRequest.key to query the next page most efficiently. It will be empty if there are no more results. |
| `total` | [uint64](#uint64) |  | total is total number of results available if PageRequest.count_total was set, its value is undefined otherwise |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/auth/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/auth/v1beta1/query.proto



<a name="cosmos.auth.v1beta1.AddressBytesToStringRequest"></a>

### AddressBytesToStringRequest
AddressBytesToStringRequest is the request type for AddressString rpc method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address_bytes` | [bytes](#bytes) |  |  |






<a name="cosmos.auth.v1beta1.AddressBytesToStringResponse"></a>

### AddressBytesToStringResponse
AddressBytesToStringResponse is the response type for AddressString rpc method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address_string` | [string](#string) |  |  |






<a name="cosmos.auth.v1beta1.AddressStringToBytesRequest"></a>

### AddressStringToBytesRequest
AddressStringToBytesRequest is the request type for AccountBytes rpc method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address_string` | [string](#string) |  |  |






<a name="cosmos.auth.v1beta1.AddressStringToBytesResponse"></a>

### AddressStringToBytesResponse
AddressStringToBytesResponse is the response type for AddressBytes rpc method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address_bytes` | [bytes](#bytes) |  |  |






<a name="cosmos.auth.v1beta1.Bech32PrefixRequest"></a>

### Bech32PrefixRequest
Bech32PrefixRequest is the request type for Bech32Prefix rpc method.

Since: cosmos-sdk 0.46






<a name="cosmos.auth.v1beta1.Bech32PrefixResponse"></a>

### Bech32PrefixResponse
Bech32PrefixResponse is the response type for Bech32Prefix rpc method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bech32_prefix` | [string](#string) |  |  |






<a name="cosmos.auth.v1beta1.QueryAccountAddressByIDRequest"></a>

### QueryAccountAddressByIDRequest
QueryAccountAddressByIDRequest is the request type for AccountAddressByID rpc method

Since: cosmos-sdk 0.46.2


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [int64](#int64) |  | id is the account number of the address to be queried. This field should have been an uint64 (like all account numbers), and will be updated to uint64 in a future version of the auth query. |






<a name="cosmos.auth.v1beta1.QueryAccountAddressByIDResponse"></a>

### QueryAccountAddressByIDResponse
QueryAccountAddressByIDResponse is the response type for AccountAddressByID rpc method

Since: cosmos-sdk 0.46.2


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account_address` | [string](#string) |  |  |






<a name="cosmos.auth.v1beta1.QueryAccountRequest"></a>

### QueryAccountRequest
QueryAccountRequest is the request type for the Query/Account RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address defines the address to query for. |






<a name="cosmos.auth.v1beta1.QueryAccountResponse"></a>

### QueryAccountResponse
QueryAccountResponse is the response type for the Query/Account RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account` | [google.protobuf.Any](#google.protobuf.Any) |  | account defines the account of the corresponding address. |






<a name="cosmos.auth.v1beta1.QueryAccountsRequest"></a>

### QueryAccountsRequest
QueryAccountsRequest is the request type for the Query/Accounts RPC method.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.auth.v1beta1.QueryAccountsResponse"></a>

### QueryAccountsResponse
QueryAccountsResponse is the response type for the Query/Accounts RPC method.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated | accounts are the existing accounts |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.auth.v1beta1.QueryModuleAccountByNameRequest"></a>

### QueryModuleAccountByNameRequest
QueryModuleAccountByNameRequest is the request type for the Query/ModuleAccountByName RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |






<a name="cosmos.auth.v1beta1.QueryModuleAccountByNameResponse"></a>

### QueryModuleAccountByNameResponse
QueryModuleAccountByNameResponse is the response type for the Query/ModuleAccountByName RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `account` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="cosmos.auth.v1beta1.QueryModuleAccountsRequest"></a>

### QueryModuleAccountsRequest
QueryModuleAccountsRequest is the request type for the Query/ModuleAccounts RPC method.

Since: cosmos-sdk 0.46






<a name="cosmos.auth.v1beta1.QueryModuleAccountsResponse"></a>

### QueryModuleAccountsResponse
QueryModuleAccountsResponse is the response type for the Query/ModuleAccounts RPC method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `accounts` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |






<a name="cosmos.auth.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="cosmos.auth.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.auth.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.auth.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Accounts` | [QueryAccountsRequest](#cosmos.auth.v1beta1.QueryAccountsRequest) | [QueryAccountsResponse](#cosmos.auth.v1beta1.QueryAccountsResponse) | Accounts returns all the existing accounts

Since: cosmos-sdk 0.43 | GET|/cosmos/auth/v1beta1/accounts|
| `Account` | [QueryAccountRequest](#cosmos.auth.v1beta1.QueryAccountRequest) | [QueryAccountResponse](#cosmos.auth.v1beta1.QueryAccountResponse) | Account returns account details based on address. | GET|/cosmos/auth/v1beta1/accounts/{address}|
| `AccountAddressByID` | [QueryAccountAddressByIDRequest](#cosmos.auth.v1beta1.QueryAccountAddressByIDRequest) | [QueryAccountAddressByIDResponse](#cosmos.auth.v1beta1.QueryAccountAddressByIDResponse) | AccountAddressByID returns account address based on account number.

Since: cosmos-sdk 0.46.2 | GET|/cosmos/auth/v1beta1/address_by_id/{id}|
| `Params` | [QueryParamsRequest](#cosmos.auth.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.auth.v1beta1.QueryParamsResponse) | Params queries all parameters. | GET|/cosmos/auth/v1beta1/params|
| `ModuleAccounts` | [QueryModuleAccountsRequest](#cosmos.auth.v1beta1.QueryModuleAccountsRequest) | [QueryModuleAccountsResponse](#cosmos.auth.v1beta1.QueryModuleAccountsResponse) | ModuleAccounts returns all the existing module accounts.

Since: cosmos-sdk 0.46 | GET|/cosmos/auth/v1beta1/module_accounts|
| `ModuleAccountByName` | [QueryModuleAccountByNameRequest](#cosmos.auth.v1beta1.QueryModuleAccountByNameRequest) | [QueryModuleAccountByNameResponse](#cosmos.auth.v1beta1.QueryModuleAccountByNameResponse) | ModuleAccountByName returns the module account info by module name | GET|/cosmos/auth/v1beta1/module_accounts/{name}|
| `Bech32Prefix` | [Bech32PrefixRequest](#cosmos.auth.v1beta1.Bech32PrefixRequest) | [Bech32PrefixResponse](#cosmos.auth.v1beta1.Bech32PrefixResponse) | Bech32Prefix queries bech32Prefix

Since: cosmos-sdk 0.46 | GET|/cosmos/auth/v1beta1/bech32|
| `AddressBytesToString` | [AddressBytesToStringRequest](#cosmos.auth.v1beta1.AddressBytesToStringRequest) | [AddressBytesToStringResponse](#cosmos.auth.v1beta1.AddressBytesToStringResponse) | AddressBytesToString converts Account Address bytes to string

Since: cosmos-sdk 0.46 | GET|/cosmos/auth/v1beta1/bech32/{address_bytes}|
| `AddressStringToBytes` | [AddressStringToBytesRequest](#cosmos.auth.v1beta1.AddressStringToBytesRequest) | [AddressStringToBytesResponse](#cosmos.auth.v1beta1.AddressStringToBytesResponse) | AddressStringToBytes converts Address string to bytes

Since: cosmos-sdk 0.46 | GET|/cosmos/auth/v1beta1/bech32/{address_string}|

 <!-- end services -->



<a name="cosmos/authz/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/authz/v1beta1/authz.proto
Since: cosmos-sdk 0.43


<a name="cosmos.authz.v1beta1.GenericAuthorization"></a>

### GenericAuthorization
GenericAuthorization gives the grantee unrestricted permissions to execute
the provided method on behalf of the granter's account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg` | [string](#string) |  | Msg, identified by it's type URL, to grant unrestricted permissions to execute |






<a name="cosmos.authz.v1beta1.Grant"></a>

### Grant
Grant gives permissions to execute
the provide method with expiration time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authorization` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `expiration` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | time when the grant will expire and will be pruned. If null, then the grant doesn't have a time expiration (other conditions in `authorization` may apply to invalidate the grant) |






<a name="cosmos.authz.v1beta1.GrantAuthorization"></a>

### GrantAuthorization
GrantAuthorization extends a grant with both the addresses of the grantee and granter.
It is used in genesis.proto and query.proto


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  |  |
| `grantee` | [string](#string) |  |  |
| `authorization` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `expiration` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="cosmos.authz.v1beta1.GrantQueueItem"></a>

### GrantQueueItem
GrantQueueItem contains the list of TypeURL of a sdk.Msg.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_type_urls` | [string](#string) | repeated | msg_type_urls contains the list of TypeURL of a sdk.Msg. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/authz/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/authz/v1beta1/event.proto
Since: cosmos-sdk 0.43


<a name="cosmos.authz.v1beta1.EventGrant"></a>

### EventGrant
EventGrant is emitted on Msg/Grant


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_type_url` | [string](#string) |  | Msg type URL for which an autorization is granted |
| `granter` | [string](#string) |  | Granter account address |
| `grantee` | [string](#string) |  | Grantee account address |






<a name="cosmos.authz.v1beta1.EventRevoke"></a>

### EventRevoke
EventRevoke is emitted on Msg/Revoke


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_type_url` | [string](#string) |  | Msg type URL for which an autorization is revoked |
| `granter` | [string](#string) |  | Granter account address |
| `grantee` | [string](#string) |  | Grantee account address |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/authz/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/authz/v1beta1/genesis.proto
Since: cosmos-sdk 0.43


<a name="cosmos.authz.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the authz module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authorization` | [GrantAuthorization](#cosmos.authz.v1beta1.GrantAuthorization) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/authz/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/authz/v1beta1/query.proto
Since: cosmos-sdk 0.43


<a name="cosmos.authz.v1beta1.QueryGranteeGrantsRequest"></a>

### QueryGranteeGrantsRequest
QueryGranteeGrantsRequest is the request type for the Query/IssuedGrants RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grantee` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.authz.v1beta1.QueryGranteeGrantsResponse"></a>

### QueryGranteeGrantsResponse
QueryGranteeGrantsResponse is the response type for the Query/GranteeGrants RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grants` | [GrantAuthorization](#cosmos.authz.v1beta1.GrantAuthorization) | repeated | grants is a list of grants granted to the grantee. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.authz.v1beta1.QueryGranterGrantsRequest"></a>

### QueryGranterGrantsRequest
QueryGranterGrantsRequest is the request type for the Query/GranterGrants RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.authz.v1beta1.QueryGranterGrantsResponse"></a>

### QueryGranterGrantsResponse
QueryGranterGrantsResponse is the response type for the Query/GranterGrants RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grants` | [GrantAuthorization](#cosmos.authz.v1beta1.GrantAuthorization) | repeated | grants is a list of grants granted by the granter. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.authz.v1beta1.QueryGrantsRequest"></a>

### QueryGrantsRequest
QueryGrantsRequest is the request type for the Query/Grants RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  |  |
| `grantee` | [string](#string) |  |  |
| `msg_type_url` | [string](#string) |  | Optional, msg_type_url, when set, will query only grants matching given msg type. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.authz.v1beta1.QueryGrantsResponse"></a>

### QueryGrantsResponse
QueryGrantsResponse is the response type for the Query/Authorizations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grants` | [Grant](#cosmos.authz.v1beta1.Grant) | repeated | authorizations is a list of grants granted for grantee by granter. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.authz.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Grants` | [QueryGrantsRequest](#cosmos.authz.v1beta1.QueryGrantsRequest) | [QueryGrantsResponse](#cosmos.authz.v1beta1.QueryGrantsResponse) | Returns list of `Authorization`, granted to the grantee by the granter. | GET|/cosmos/authz/v1beta1/grants|
| `GranterGrants` | [QueryGranterGrantsRequest](#cosmos.authz.v1beta1.QueryGranterGrantsRequest) | [QueryGranterGrantsResponse](#cosmos.authz.v1beta1.QueryGranterGrantsResponse) | GranterGrants returns list of `GrantAuthorization`, granted by granter.

Since: cosmos-sdk 0.46 | GET|/cosmos/authz/v1beta1/grants/granter/{granter}|
| `GranteeGrants` | [QueryGranteeGrantsRequest](#cosmos.authz.v1beta1.QueryGranteeGrantsRequest) | [QueryGranteeGrantsResponse](#cosmos.authz.v1beta1.QueryGranteeGrantsResponse) | GranteeGrants returns a list of `GrantAuthorization` by grantee.

Since: cosmos-sdk 0.46 | GET|/cosmos/authz/v1beta1/grants/grantee/{grantee}|

 <!-- end services -->



<a name="cosmos/msg/v1/msg.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/msg/v1/msg.proto


 <!-- end messages -->

 <!-- end enums -->


<a name="cosmos/msg/v1/msg.proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| `signer` | string | .google.protobuf.MessageOptions | 11110000 | signer must be used in cosmos messages in order to signal to external clients which fields in a given cosmos message must be filled with signer information (address). The field must be the protobuf name of the message field extended with this MessageOption. The field must either be of string kind, or of message kind in case the signer information is contained within a message inside the cosmos message. |

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/authz/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/authz/v1beta1/tx.proto
Since: cosmos-sdk 0.43


<a name="cosmos.authz.v1beta1.MsgExec"></a>

### MsgExec
MsgExec attempts to execute the provided messages using
authorizations granted to the grantee. Each message should have only
one signer corresponding to the granter of the authorization.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grantee` | [string](#string) |  |  |
| `msgs` | [google.protobuf.Any](#google.protobuf.Any) | repeated | Authorization Msg requests to execute. Each msg must implement Authorization interface The x/authz will try to find a grant matching (msg.signers[0], grantee, MsgTypeURL(msg)) triple and validate it. |






<a name="cosmos.authz.v1beta1.MsgExecResponse"></a>

### MsgExecResponse
MsgExecResponse defines the Msg/MsgExecResponse response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `results` | [bytes](#bytes) | repeated |  |






<a name="cosmos.authz.v1beta1.MsgGrant"></a>

### MsgGrant
MsgGrant is a request type for Grant method. It declares authorization to the grantee
on behalf of the granter with the provided expiration time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  |  |
| `grantee` | [string](#string) |  |  |
| `grant` | [Grant](#cosmos.authz.v1beta1.Grant) |  |  |






<a name="cosmos.authz.v1beta1.MsgGrantResponse"></a>

### MsgGrantResponse
MsgGrantResponse defines the Msg/MsgGrant response type.






<a name="cosmos.authz.v1beta1.MsgRevoke"></a>

### MsgRevoke
MsgRevoke revokes any authorization with the provided sdk.Msg type on the
granter's account with that has been granted to the grantee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  |  |
| `grantee` | [string](#string) |  |  |
| `msg_type_url` | [string](#string) |  |  |






<a name="cosmos.authz.v1beta1.MsgRevokeResponse"></a>

### MsgRevokeResponse
MsgRevokeResponse defines the Msg/MsgRevokeResponse response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.authz.v1beta1.Msg"></a>

### Msg
Msg defines the authz Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Grant` | [MsgGrant](#cosmos.authz.v1beta1.MsgGrant) | [MsgGrantResponse](#cosmos.authz.v1beta1.MsgGrantResponse) | Grant grants the provided authorization to the grantee on the granter's account with the provided expiration time. If there is already a grant for the given (granter, grantee, Authorization) triple, then the grant will be overwritten. | |
| `Exec` | [MsgExec](#cosmos.authz.v1beta1.MsgExec) | [MsgExecResponse](#cosmos.authz.v1beta1.MsgExecResponse) | Exec attempts to execute the provided messages using authorizations granted to the grantee. Each message should have only one signer corresponding to the granter of the authorization. | |
| `Revoke` | [MsgRevoke](#cosmos.authz.v1beta1.MsgRevoke) | [MsgRevokeResponse](#cosmos.authz.v1beta1.MsgRevokeResponse) | Revoke revokes any authorization corresponding to the provided method name on the granter's account that has been granted to the grantee. | |

 <!-- end services -->



<a name="cosmos/base/v1beta1/coin.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/v1beta1/coin.proto



<a name="cosmos.base.v1beta1.Coin"></a>

### Coin
Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.DecCoin"></a>

### DecCoin
DecCoin defines a token with a denomination and a decimal amount.

NOTE: The amount field is an Dec which implements the custom method
signatures required by gogoproto.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.DecProto"></a>

### DecProto
DecProto defines a Protobuf wrapper around a Dec object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `dec` | [string](#string) |  |  |






<a name="cosmos.base.v1beta1.IntProto"></a>

### IntProto
IntProto defines a Protobuf wrapper around an Int object.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `int` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/authz.proto



<a name="cosmos.bank.v1beta1.SendAuthorization"></a>

### SendAuthorization
SendAuthorization allows the grantee to spend up to spend_limit coins from
the granter's account.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `spend_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/bank.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/bank.proto



<a name="cosmos.bank.v1beta1.DenomUnit"></a>

### DenomUnit
DenomUnit represents a struct that describes a given
denomination unit of the basic token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom represents the string name of the given denom unit (e.g uatom). |
| `exponent` | [uint32](#uint32) |  | exponent represents power of 10 exponent that one must raise the base_denom to in order to equal the given DenomUnit's denom 1 denom = 10^exponent base_denom (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with exponent = 6, thus: 1 atom = 10^6 uatom). |
| `aliases` | [string](#string) | repeated | aliases is a list of string aliases for the given denom |






<a name="cosmos.bank.v1beta1.Input"></a>

### Input
Input models transaction input.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.bank.v1beta1.Metadata"></a>

### Metadata
Metadata represents a struct that describes
a basic token.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [string](#string) |  |  |
| `denom_units` | [DenomUnit](#cosmos.bank.v1beta1.DenomUnit) | repeated | denom_units represents the list of DenomUnit's for a given coin |
| `base` | [string](#string) |  | base represents the base denom (should be the DenomUnit with exponent = 0). |
| `display` | [string](#string) |  | display indicates the suggested denom that should be displayed in clients. |
| `name` | [string](#string) |  | name defines the name of the token (eg: Cosmos Atom)

Since: cosmos-sdk 0.43 |
| `symbol` | [string](#string) |  | symbol is the token symbol usually shown on exchanges (eg: ATOM). This can be the same as the display.

Since: cosmos-sdk 0.43 |
| `uri` | [string](#string) |  | URI to a document (on or off-chain) that contains additional information. Optional.

Since: cosmos-sdk 0.46 |
| `uri_hash` | [string](#string) |  | URIHash is a sha256 hash of a document pointed by URI. It's used to verify that the document didn't change. Optional.

Since: cosmos-sdk 0.46 |






<a name="cosmos.bank.v1beta1.Output"></a>

### Output
Output models transaction outputs.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.bank.v1beta1.Params"></a>

### Params
Params defines the parameters for the bank module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `send_enabled` | [SendEnabled](#cosmos.bank.v1beta1.SendEnabled) | repeated |  |
| `default_send_enabled` | [bool](#bool) |  |  |






<a name="cosmos.bank.v1beta1.SendEnabled"></a>

### SendEnabled
SendEnabled maps coin denom to a send_enabled status (whether a denom is
sendable).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `enabled` | [bool](#bool) |  |  |






<a name="cosmos.bank.v1beta1.Supply"></a>

### Supply
Supply represents a struct that passively keeps track of the total supply
amounts in the network.
This message is deprecated now that supply is indexed by denom.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/genesis.proto



<a name="cosmos.bank.v1beta1.Balance"></a>

### Balance
Balance defines an account address and balance pair used in the bank module's
genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the balance holder. |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | coins defines the different coins this balance holds. |






<a name="cosmos.bank.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the bank module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.bank.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `balances` | [Balance](#cosmos.bank.v1beta1.Balance) | repeated | balances is an array containing the balances of all the accounts. |
| `supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | supply represents the total supply. If it is left empty, then supply will be calculated based on the provided balances. Otherwise, it will be used to validate that the sum of the balances equals this amount. |
| `denom_metadata` | [Metadata](#cosmos.bank.v1beta1.Metadata) | repeated | denom_metadata defines the metadata of the differents coins. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/bank/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/query.proto



<a name="cosmos.bank.v1beta1.DenomOwner"></a>

### DenomOwner
DenomOwner defines structure representing an account that owns or holds a
particular denominated token. It contains the account address and account
balance of the denominated token.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address defines the address that owns a particular denomination. |
| `balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | balance is the balance of the denominated coin for an account. |






<a name="cosmos.bank.v1beta1.QueryAllBalancesRequest"></a>

### QueryAllBalancesRequest
QueryBalanceRequest is the request type for the Query/AllBalances RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address to query balances for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.bank.v1beta1.QueryAllBalancesResponse"></a>

### QueryAllBalancesResponse
QueryAllBalancesResponse is the response type for the Query/AllBalances RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | balances is the balances of all the coins. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.bank.v1beta1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the request type for the Query/Balance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address to query balances for. |
| `denom` | [string](#string) |  | denom is the coin denom to query balances for. |






<a name="cosmos.bank.v1beta1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the response type for the Query/Balance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | balance is the balance of the coin. |






<a name="cosmos.bank.v1beta1.QueryDenomMetadataRequest"></a>

### QueryDenomMetadataRequest
QueryDenomMetadataRequest is the request type for the Query/DenomMetadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom is the coin denom to query the metadata for. |






<a name="cosmos.bank.v1beta1.QueryDenomMetadataResponse"></a>

### QueryDenomMetadataResponse
QueryDenomMetadataResponse is the response type for the Query/DenomMetadata RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [Metadata](#cosmos.bank.v1beta1.Metadata) |  | metadata describes and provides all the client information for the requested token. |






<a name="cosmos.bank.v1beta1.QueryDenomOwnersRequest"></a>

### QueryDenomOwnersRequest
QueryDenomOwnersRequest defines the request type for the DenomOwners RPC query,
which queries for a paginated set of all account holders of a particular
denomination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom defines the coin denomination to query all account holders for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.bank.v1beta1.QueryDenomOwnersResponse"></a>

### QueryDenomOwnersResponse
QueryDenomOwnersResponse defines the RPC response of a DenomOwners RPC query.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom_owners` | [DenomOwner](#cosmos.bank.v1beta1.DenomOwner) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.bank.v1beta1.QueryDenomsMetadataRequest"></a>

### QueryDenomsMetadataRequest
QueryDenomsMetadataRequest is the request type for the Query/DenomsMetadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.bank.v1beta1.QueryDenomsMetadataResponse"></a>

### QueryDenomsMetadataResponse
QueryDenomsMetadataResponse is the response type for the Query/DenomsMetadata RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadatas` | [Metadata](#cosmos.bank.v1beta1.Metadata) | repeated | metadata provides the client information for all the registered tokens. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.bank.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest defines the request type for querying x/bank parameters.






<a name="cosmos.bank.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse defines the response type for querying x/bank parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.bank.v1beta1.Params) |  |  |






<a name="cosmos.bank.v1beta1.QuerySpendableBalancesRequest"></a>

### QuerySpendableBalancesRequest
QuerySpendableBalancesRequest defines the gRPC request structure for querying
an account's spendable balances.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address to query spendable balances for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.bank.v1beta1.QuerySpendableBalancesResponse"></a>

### QuerySpendableBalancesResponse
QuerySpendableBalancesResponse defines the gRPC response structure for querying
an account's spendable balances.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balances` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | balances is the spendable balances of all the coins. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.bank.v1beta1.QuerySupplyOfRequest"></a>

### QuerySupplyOfRequest
QuerySupplyOfRequest is the request type for the Query/SupplyOf RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  | denom is the coin denom to query balances for. |






<a name="cosmos.bank.v1beta1.QuerySupplyOfResponse"></a>

### QuerySupplyOfResponse
QuerySupplyOfResponse is the response type for the Query/SupplyOf RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is the supply of the coin. |






<a name="cosmos.bank.v1beta1.QueryTotalSupplyRequest"></a>

### QueryTotalSupplyRequest
QueryTotalSupplyRequest is the request type for the Query/TotalSupply RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request.

Since: cosmos-sdk 0.43 |






<a name="cosmos.bank.v1beta1.QueryTotalSupplyResponse"></a>

### QueryTotalSupplyResponse
QueryTotalSupplyResponse is the response type for the Query/TotalSupply RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `supply` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | supply is the supply of the coins |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response.

Since: cosmos-sdk 0.43 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.bank.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Balance` | [QueryBalanceRequest](#cosmos.bank.v1beta1.QueryBalanceRequest) | [QueryBalanceResponse](#cosmos.bank.v1beta1.QueryBalanceResponse) | Balance queries the balance of a single coin for a single account. | GET|/cosmos/bank/v1beta1/balances/{address}/by_denom|
| `AllBalances` | [QueryAllBalancesRequest](#cosmos.bank.v1beta1.QueryAllBalancesRequest) | [QueryAllBalancesResponse](#cosmos.bank.v1beta1.QueryAllBalancesResponse) | AllBalances queries the balance of all coins for a single account. | GET|/cosmos/bank/v1beta1/balances/{address}|
| `SpendableBalances` | [QuerySpendableBalancesRequest](#cosmos.bank.v1beta1.QuerySpendableBalancesRequest) | [QuerySpendableBalancesResponse](#cosmos.bank.v1beta1.QuerySpendableBalancesResponse) | SpendableBalances queries the spenable balance of all coins for a single account.

Since: cosmos-sdk 0.46 | GET|/cosmos/bank/v1beta1/spendable_balances/{address}|
| `TotalSupply` | [QueryTotalSupplyRequest](#cosmos.bank.v1beta1.QueryTotalSupplyRequest) | [QueryTotalSupplyResponse](#cosmos.bank.v1beta1.QueryTotalSupplyResponse) | TotalSupply queries the total supply of all coins. | GET|/cosmos/bank/v1beta1/supply|
| `SupplyOf` | [QuerySupplyOfRequest](#cosmos.bank.v1beta1.QuerySupplyOfRequest) | [QuerySupplyOfResponse](#cosmos.bank.v1beta1.QuerySupplyOfResponse) | SupplyOf queries the supply of a single coin. | GET|/cosmos/bank/v1beta1/supply/by_denom|
| `Params` | [QueryParamsRequest](#cosmos.bank.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.bank.v1beta1.QueryParamsResponse) | Params queries the parameters of x/bank module. | GET|/cosmos/bank/v1beta1/params|
| `DenomMetadata` | [QueryDenomMetadataRequest](#cosmos.bank.v1beta1.QueryDenomMetadataRequest) | [QueryDenomMetadataResponse](#cosmos.bank.v1beta1.QueryDenomMetadataResponse) | DenomsMetadata queries the client metadata of a given coin denomination. | GET|/cosmos/bank/v1beta1/denoms_metadata/{denom}|
| `DenomsMetadata` | [QueryDenomsMetadataRequest](#cosmos.bank.v1beta1.QueryDenomsMetadataRequest) | [QueryDenomsMetadataResponse](#cosmos.bank.v1beta1.QueryDenomsMetadataResponse) | DenomsMetadata queries the client metadata for all registered coin denominations. | GET|/cosmos/bank/v1beta1/denoms_metadata|
| `DenomOwners` | [QueryDenomOwnersRequest](#cosmos.bank.v1beta1.QueryDenomOwnersRequest) | [QueryDenomOwnersResponse](#cosmos.bank.v1beta1.QueryDenomOwnersResponse) | DenomOwners queries for all account addresses that own a particular token denomination.

Since: cosmos-sdk 0.46 | GET|/cosmos/bank/v1beta1/denom_owners/{denom}|

 <!-- end services -->



<a name="cosmos/bank/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/bank/v1beta1/tx.proto



<a name="cosmos.bank.v1beta1.MsgMultiSend"></a>

### MsgMultiSend
MsgMultiSend represents an arbitrary multi-in, multi-out send message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inputs` | [Input](#cosmos.bank.v1beta1.Input) | repeated |  |
| `outputs` | [Output](#cosmos.bank.v1beta1.Output) | repeated |  |






<a name="cosmos.bank.v1beta1.MsgMultiSendResponse"></a>

### MsgMultiSendResponse
MsgMultiSendResponse defines the Msg/MultiSend response type.






<a name="cosmos.bank.v1beta1.MsgSend"></a>

### MsgSend
MsgSend represents a message to send coins from one account to another.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.bank.v1beta1.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse defines the Msg/Send response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.bank.v1beta1.Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Send` | [MsgSend](#cosmos.bank.v1beta1.MsgSend) | [MsgSendResponse](#cosmos.bank.v1beta1.MsgSendResponse) | Send defines a method for sending coins from one account to another account. | |
| `MultiSend` | [MsgMultiSend](#cosmos.bank.v1beta1.MsgMultiSend) | [MsgMultiSendResponse](#cosmos.bank.v1beta1.MsgMultiSendResponse) | MultiSend defines a method for sending coins from some accounts to other accounts. | |

 <!-- end services -->



<a name="tendermint/crypto/proof.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/crypto/proof.proto



<a name="tendermint.crypto.DominoOp"></a>

### DominoOp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `input` | [string](#string) |  |  |
| `output` | [string](#string) |  |  |






<a name="tendermint.crypto.Proof"></a>

### Proof



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total` | [int64](#int64) |  |  |
| `index` | [int64](#int64) |  |  |
| `leaf_hash` | [bytes](#bytes) |  |  |
| `aunts` | [bytes](#bytes) | repeated |  |






<a name="tendermint.crypto.ProofOp"></a>

### ProofOp
ProofOp defines an operation used for calculating Merkle root
The data could be arbitrary format, providing nessecary data
for example neighbouring node hash


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `key` | [bytes](#bytes) |  |  |
| `data` | [bytes](#bytes) |  |  |






<a name="tendermint.crypto.ProofOps"></a>

### ProofOps
ProofOps is Merkle proof defined by the list of ProofOps


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ops` | [ProofOp](#tendermint.crypto.ProofOp) | repeated |  |






<a name="tendermint.crypto.ValueOp"></a>

### ValueOp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  | Encoded in ProofOp.Key. |
| `proof` | [Proof](#tendermint.crypto.Proof) |  | To encode in ProofOp.Data |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/version/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/version/types.proto



<a name="tendermint.version.App"></a>

### App
App includes the protocol and software version for the application.
This information is included in ResponseInfo. The App.Protocol can be
updated in ResponseEndBlock.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol` | [uint64](#uint64) |  |  |
| `software` | [string](#string) |  |  |






<a name="tendermint.version.Consensus"></a>

### Consensus
Consensus captures the consensus rules for processing a block in the blockchain,
including all blockchain data structures and the rules of the application's
state transition machine.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [uint64](#uint64) |  |  |
| `app` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/crypto/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/crypto/keys.proto



<a name="tendermint.crypto.PublicKey"></a>

### PublicKey
PublicKey defines the keys available for use with Tendermint Validators


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ed25519` | [bytes](#bytes) |  |  |
| `secp256k1` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/types/validator.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/types/validator.proto



<a name="tendermint.types.SimpleValidator"></a>

### SimpleValidator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key` | [tendermint.crypto.PublicKey](#tendermint.crypto.PublicKey) |  |  |
| `voting_power` | [int64](#int64) |  |  |






<a name="tendermint.types.Validator"></a>

### Validator



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `pub_key` | [tendermint.crypto.PublicKey](#tendermint.crypto.PublicKey) |  |  |
| `voting_power` | [int64](#int64) |  |  |
| `proposer_priority` | [int64](#int64) |  |  |






<a name="tendermint.types.ValidatorSet"></a>

### ValidatorSet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [Validator](#tendermint.types.Validator) | repeated |  |
| `proposer` | [Validator](#tendermint.types.Validator) |  |  |
| `total_voting_power` | [int64](#int64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/types/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/types/types.proto



<a name="tendermint.types.BlockID"></a>

### BlockID
BlockID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [bytes](#bytes) |  |  |
| `part_set_header` | [PartSetHeader](#tendermint.types.PartSetHeader) |  |  |






<a name="tendermint.types.BlockMeta"></a>

### BlockMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_id` | [BlockID](#tendermint.types.BlockID) |  |  |
| `block_size` | [int64](#int64) |  |  |
| `header` | [Header](#tendermint.types.Header) |  |  |
| `num_txs` | [int64](#int64) |  |  |






<a name="tendermint.types.Commit"></a>

### Commit
Commit contains the evidence that a block was committed by a set of validators.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  |  |
| `round` | [int32](#int32) |  |  |
| `block_id` | [BlockID](#tendermint.types.BlockID) |  |  |
| `signatures` | [CommitSig](#tendermint.types.CommitSig) | repeated |  |






<a name="tendermint.types.CommitSig"></a>

### CommitSig
CommitSig is a part of the Vote included in a Commit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_id_flag` | [BlockIDFlag](#tendermint.types.BlockIDFlag) |  |  |
| `validator_address` | [bytes](#bytes) |  |  |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `signature` | [bytes](#bytes) |  |  |






<a name="tendermint.types.Data"></a>

### Data
Data contains the set of transactions included in the block


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `txs` | [bytes](#bytes) | repeated | Txs that will be applied by state @ block.Height+1. NOTE: not all txs here are valid. We're just agreeing on the order first. This means that block.AppHash does not include these txs. |






<a name="tendermint.types.Header"></a>

### Header
Header defines the structure of a Tendermint block header.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [tendermint.version.Consensus](#tendermint.version.Consensus) |  | basic block info |
| `chain_id` | [string](#string) |  |  |
| `height` | [int64](#int64) |  |  |
| `time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `last_block_id` | [BlockID](#tendermint.types.BlockID) |  | prev block info |
| `last_commit_hash` | [bytes](#bytes) |  | hashes of block data

commit from validators from the last block |
| `data_hash` | [bytes](#bytes) |  | transactions |
| `validators_hash` | [bytes](#bytes) |  | hashes from the app output from the prev block

validators for the current block |
| `next_validators_hash` | [bytes](#bytes) |  | validators for the next block |
| `consensus_hash` | [bytes](#bytes) |  | consensus params for current block |
| `app_hash` | [bytes](#bytes) |  | state after txs from the previous block |
| `last_results_hash` | [bytes](#bytes) |  | root hash of all results from the txs from the previous block |
| `evidence_hash` | [bytes](#bytes) |  | consensus info

evidence included in the block |
| `proposer_address` | [bytes](#bytes) |  | original proposer of the block |






<a name="tendermint.types.LightBlock"></a>

### LightBlock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signed_header` | [SignedHeader](#tendermint.types.SignedHeader) |  |  |
| `validator_set` | [ValidatorSet](#tendermint.types.ValidatorSet) |  |  |






<a name="tendermint.types.Part"></a>

### Part



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint32](#uint32) |  |  |
| `bytes` | [bytes](#bytes) |  |  |
| `proof` | [tendermint.crypto.Proof](#tendermint.crypto.Proof) |  |  |






<a name="tendermint.types.PartSetHeader"></a>

### PartSetHeader
PartsetHeader


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total` | [uint32](#uint32) |  |  |
| `hash` | [bytes](#bytes) |  |  |






<a name="tendermint.types.Proposal"></a>

### Proposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [SignedMsgType](#tendermint.types.SignedMsgType) |  |  |
| `height` | [int64](#int64) |  |  |
| `round` | [int32](#int32) |  |  |
| `pol_round` | [int32](#int32) |  |  |
| `block_id` | [BlockID](#tendermint.types.BlockID) |  |  |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `signature` | [bytes](#bytes) |  |  |






<a name="tendermint.types.SignedHeader"></a>

### SignedHeader



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header` | [Header](#tendermint.types.Header) |  |  |
| `commit` | [Commit](#tendermint.types.Commit) |  |  |






<a name="tendermint.types.TxProof"></a>

### TxProof
TxProof represents a Merkle proof of the presence of a transaction in the Merkle tree.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `root_hash` | [bytes](#bytes) |  |  |
| `data` | [bytes](#bytes) |  |  |
| `proof` | [tendermint.crypto.Proof](#tendermint.crypto.Proof) |  |  |






<a name="tendermint.types.Vote"></a>

### Vote
Vote represents a prevote, precommit, or commit vote from validators for
consensus.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [SignedMsgType](#tendermint.types.SignedMsgType) |  |  |
| `height` | [int64](#int64) |  |  |
| `round` | [int32](#int32) |  |  |
| `block_id` | [BlockID](#tendermint.types.BlockID) |  | zero if vote is nil. |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `validator_address` | [bytes](#bytes) |  |  |
| `validator_index` | [int32](#int32) |  |  |
| `signature` | [bytes](#bytes) |  |  |





 <!-- end messages -->


<a name="tendermint.types.BlockIDFlag"></a>

### BlockIDFlag
BlockIdFlag indicates which BlcokID the signature is for

| Name | Number | Description |
| ---- | ------ | ----------- |
| BLOCK_ID_FLAG_UNKNOWN | 0 |  |
| BLOCK_ID_FLAG_ABSENT | 1 |  |
| BLOCK_ID_FLAG_COMMIT | 2 |  |
| BLOCK_ID_FLAG_NIL | 3 |  |



<a name="tendermint.types.SignedMsgType"></a>

### SignedMsgType
SignedMsgType is a type of signed message in the consensus.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SIGNED_MSG_TYPE_UNKNOWN | 0 |  |
| SIGNED_MSG_TYPE_PREVOTE | 1 | Votes |
| SIGNED_MSG_TYPE_PRECOMMIT | 2 |  |
| SIGNED_MSG_TYPE_PROPOSAL | 32 | Proposals |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/types/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/types/params.proto



<a name="tendermint.types.BlockParams"></a>

### BlockParams
BlockParams contains limits on the block size.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_bytes` | [int64](#int64) |  | Max block size, in bytes. Note: must be greater than 0 |
| `max_gas` | [int64](#int64) |  | Max gas per block. Note: must be greater or equal to -1 |
| `time_iota_ms` | [int64](#int64) |  | Minimum time increment between consecutive blocks (in milliseconds) If the block header timestamp is ahead of the system clock, decrease this value.

Not exposed to the application. |






<a name="tendermint.types.ConsensusParams"></a>

### ConsensusParams
ConsensusParams contains consensus critical parameters that determine the
validity of blocks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [BlockParams](#tendermint.types.BlockParams) |  |  |
| `evidence` | [EvidenceParams](#tendermint.types.EvidenceParams) |  |  |
| `validator` | [ValidatorParams](#tendermint.types.ValidatorParams) |  |  |
| `version` | [VersionParams](#tendermint.types.VersionParams) |  |  |






<a name="tendermint.types.EvidenceParams"></a>

### EvidenceParams
EvidenceParams determine how we handle evidence of malfeasance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_age_num_blocks` | [int64](#int64) |  | Max age of evidence, in blocks.

The basic formula for calculating this is: MaxAgeDuration / {average block time}. |
| `max_age_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Max age of evidence, in time.

It should correspond with an app's "unbonding period" or other similar mechanism for handling [Nothing-At-Stake attacks](https://github.com/ethereum/wiki/wiki/Proof-of-Stake-FAQ#what-is-the-nothing-at-stake-problem-and-how-can-it-be-fixed). |
| `max_bytes` | [int64](#int64) |  | This sets the maximum size of total evidence in bytes that can be committed in a single block. and should fall comfortably under the max block bytes. Default is 1048576 or 1MB |






<a name="tendermint.types.HashedParams"></a>

### HashedParams
HashedParams is a subset of ConsensusParams.

It is hashed into the Header.ConsensusHash.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_max_bytes` | [int64](#int64) |  |  |
| `block_max_gas` | [int64](#int64) |  |  |






<a name="tendermint.types.ValidatorParams"></a>

### ValidatorParams
ValidatorParams restrict the public key types validators can use.
NOTE: uses ABCI pubkey naming, not Amino names.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key_types` | [string](#string) | repeated |  |






<a name="tendermint.types.VersionParams"></a>

### VersionParams
VersionParams contains the ABCI application version.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `app_version` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/abci/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/abci/types.proto



<a name="tendermint.abci.BlockParams"></a>

### BlockParams
BlockParams contains limits on the block size.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_bytes` | [int64](#int64) |  | Note: must be greater than 0 |
| `max_gas` | [int64](#int64) |  | Note: must be greater or equal to -1 |






<a name="tendermint.abci.ConsensusParams"></a>

### ConsensusParams
ConsensusParams contains all consensus-relevant parameters
that can be adjusted by the abci app


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block` | [BlockParams](#tendermint.abci.BlockParams) |  |  |
| `evidence` | [tendermint.types.EvidenceParams](#tendermint.types.EvidenceParams) |  |  |
| `validator` | [tendermint.types.ValidatorParams](#tendermint.types.ValidatorParams) |  |  |
| `version` | [tendermint.types.VersionParams](#tendermint.types.VersionParams) |  |  |






<a name="tendermint.abci.Event"></a>

### Event
Event allows application developers to attach additional information to
ResponseBeginBlock, ResponseEndBlock, ResponseCheckTx and ResponseDeliverTx.
Later, transactions may be queried using these events.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `attributes` | [EventAttribute](#tendermint.abci.EventAttribute) | repeated |  |






<a name="tendermint.abci.EventAttribute"></a>

### EventAttribute
EventAttribute is a single key-value pair, associated with an event.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |
| `index` | [bool](#bool) |  | nondeterministic |






<a name="tendermint.abci.Evidence"></a>

### Evidence



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [EvidenceType](#tendermint.abci.EvidenceType) |  |  |
| `validator` | [Validator](#tendermint.abci.Validator) |  | The offending validator |
| `height` | [int64](#int64) |  | The height when the offense occurred |
| `time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The corresponding time where the offense occurred |
| `total_voting_power` | [int64](#int64) |  | Total voting power of the validator set in case the ABCI application does not store historical validators. https://github.com/tendermint/tendermint/issues/4581 |






<a name="tendermint.abci.LastCommitInfo"></a>

### LastCommitInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `round` | [int32](#int32) |  |  |
| `votes` | [VoteInfo](#tendermint.abci.VoteInfo) | repeated |  |






<a name="tendermint.abci.Request"></a>

### Request



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `echo` | [RequestEcho](#tendermint.abci.RequestEcho) |  |  |
| `flush` | [RequestFlush](#tendermint.abci.RequestFlush) |  |  |
| `info` | [RequestInfo](#tendermint.abci.RequestInfo) |  |  |
| `set_option` | [RequestSetOption](#tendermint.abci.RequestSetOption) |  |  |
| `init_chain` | [RequestInitChain](#tendermint.abci.RequestInitChain) |  |  |
| `query` | [RequestQuery](#tendermint.abci.RequestQuery) |  |  |
| `begin_block` | [RequestBeginBlock](#tendermint.abci.RequestBeginBlock) |  |  |
| `check_tx` | [RequestCheckTx](#tendermint.abci.RequestCheckTx) |  |  |
| `deliver_tx` | [RequestDeliverTx](#tendermint.abci.RequestDeliverTx) |  |  |
| `end_block` | [RequestEndBlock](#tendermint.abci.RequestEndBlock) |  |  |
| `commit` | [RequestCommit](#tendermint.abci.RequestCommit) |  |  |
| `list_snapshots` | [RequestListSnapshots](#tendermint.abci.RequestListSnapshots) |  |  |
| `offer_snapshot` | [RequestOfferSnapshot](#tendermint.abci.RequestOfferSnapshot) |  |  |
| `load_snapshot_chunk` | [RequestLoadSnapshotChunk](#tendermint.abci.RequestLoadSnapshotChunk) |  |  |
| `apply_snapshot_chunk` | [RequestApplySnapshotChunk](#tendermint.abci.RequestApplySnapshotChunk) |  |  |






<a name="tendermint.abci.RequestApplySnapshotChunk"></a>

### RequestApplySnapshotChunk
Applies a snapshot chunk


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint32](#uint32) |  |  |
| `chunk` | [bytes](#bytes) |  |  |
| `sender` | [string](#string) |  |  |






<a name="tendermint.abci.RequestBeginBlock"></a>

### RequestBeginBlock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [bytes](#bytes) |  |  |
| `header` | [tendermint.types.Header](#tendermint.types.Header) |  |  |
| `last_commit_info` | [LastCommitInfo](#tendermint.abci.LastCommitInfo) |  |  |
| `byzantine_validators` | [Evidence](#tendermint.abci.Evidence) | repeated |  |






<a name="tendermint.abci.RequestCheckTx"></a>

### RequestCheckTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [bytes](#bytes) |  |  |
| `type` | [CheckTxType](#tendermint.abci.CheckTxType) |  |  |






<a name="tendermint.abci.RequestCommit"></a>

### RequestCommit







<a name="tendermint.abci.RequestDeliverTx"></a>

### RequestDeliverTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [bytes](#bytes) |  |  |






<a name="tendermint.abci.RequestEcho"></a>

### RequestEcho



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message` | [string](#string) |  |  |






<a name="tendermint.abci.RequestEndBlock"></a>

### RequestEndBlock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  |  |






<a name="tendermint.abci.RequestFlush"></a>

### RequestFlush







<a name="tendermint.abci.RequestInfo"></a>

### RequestInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [string](#string) |  |  |
| `block_version` | [uint64](#uint64) |  |  |
| `p2p_version` | [uint64](#uint64) |  |  |






<a name="tendermint.abci.RequestInitChain"></a>

### RequestInitChain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `chain_id` | [string](#string) |  |  |
| `consensus_params` | [ConsensusParams](#tendermint.abci.ConsensusParams) |  |  |
| `validators` | [ValidatorUpdate](#tendermint.abci.ValidatorUpdate) | repeated |  |
| `app_state_bytes` | [bytes](#bytes) |  |  |
| `initial_height` | [int64](#int64) |  |  |






<a name="tendermint.abci.RequestListSnapshots"></a>

### RequestListSnapshots
lists available snapshots






<a name="tendermint.abci.RequestLoadSnapshotChunk"></a>

### RequestLoadSnapshotChunk
loads a snapshot chunk


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [uint64](#uint64) |  |  |
| `format` | [uint32](#uint32) |  |  |
| `chunk` | [uint32](#uint32) |  |  |






<a name="tendermint.abci.RequestOfferSnapshot"></a>

### RequestOfferSnapshot
offers a snapshot to the application


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `snapshot` | [Snapshot](#tendermint.abci.Snapshot) |  | snapshot offered by peers |
| `app_hash` | [bytes](#bytes) |  | light client-verified app hash for snapshot height |






<a name="tendermint.abci.RequestQuery"></a>

### RequestQuery



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  |  |
| `path` | [string](#string) |  |  |
| `height` | [int64](#int64) |  |  |
| `prove` | [bool](#bool) |  |  |






<a name="tendermint.abci.RequestSetOption"></a>

### RequestSetOption
nondeterministic


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="tendermint.abci.Response"></a>

### Response



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `exception` | [ResponseException](#tendermint.abci.ResponseException) |  |  |
| `echo` | [ResponseEcho](#tendermint.abci.ResponseEcho) |  |  |
| `flush` | [ResponseFlush](#tendermint.abci.ResponseFlush) |  |  |
| `info` | [ResponseInfo](#tendermint.abci.ResponseInfo) |  |  |
| `set_option` | [ResponseSetOption](#tendermint.abci.ResponseSetOption) |  |  |
| `init_chain` | [ResponseInitChain](#tendermint.abci.ResponseInitChain) |  |  |
| `query` | [ResponseQuery](#tendermint.abci.ResponseQuery) |  |  |
| `begin_block` | [ResponseBeginBlock](#tendermint.abci.ResponseBeginBlock) |  |  |
| `check_tx` | [ResponseCheckTx](#tendermint.abci.ResponseCheckTx) |  |  |
| `deliver_tx` | [ResponseDeliverTx](#tendermint.abci.ResponseDeliverTx) |  |  |
| `end_block` | [ResponseEndBlock](#tendermint.abci.ResponseEndBlock) |  |  |
| `commit` | [ResponseCommit](#tendermint.abci.ResponseCommit) |  |  |
| `list_snapshots` | [ResponseListSnapshots](#tendermint.abci.ResponseListSnapshots) |  |  |
| `offer_snapshot` | [ResponseOfferSnapshot](#tendermint.abci.ResponseOfferSnapshot) |  |  |
| `load_snapshot_chunk` | [ResponseLoadSnapshotChunk](#tendermint.abci.ResponseLoadSnapshotChunk) |  |  |
| `apply_snapshot_chunk` | [ResponseApplySnapshotChunk](#tendermint.abci.ResponseApplySnapshotChunk) |  |  |






<a name="tendermint.abci.ResponseApplySnapshotChunk"></a>

### ResponseApplySnapshotChunk



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `result` | [ResponseApplySnapshotChunk.Result](#tendermint.abci.ResponseApplySnapshotChunk.Result) |  |  |
| `refetch_chunks` | [uint32](#uint32) | repeated | Chunks to refetch and reapply |
| `reject_senders` | [string](#string) | repeated | Chunk senders to reject and ban |






<a name="tendermint.abci.ResponseBeginBlock"></a>

### ResponseBeginBlock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `events` | [Event](#tendermint.abci.Event) | repeated |  |






<a name="tendermint.abci.ResponseCheckTx"></a>

### ResponseCheckTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [uint32](#uint32) |  |  |
| `data` | [bytes](#bytes) |  |  |
| `log` | [string](#string) |  | nondeterministic |
| `info` | [string](#string) |  | nondeterministic |
| `gas_wanted` | [int64](#int64) |  |  |
| `gas_used` | [int64](#int64) |  |  |
| `events` | [Event](#tendermint.abci.Event) | repeated |  |
| `codespace` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `priority` | [int64](#int64) |  |  |
| `mempool_error` | [string](#string) |  | mempool_error is set by Tendermint. ABCI applictions creating a ResponseCheckTX should not set mempool_error. |






<a name="tendermint.abci.ResponseCommit"></a>

### ResponseCommit



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | reserve 1 |
| `retain_height` | [int64](#int64) |  |  |






<a name="tendermint.abci.ResponseDeliverTx"></a>

### ResponseDeliverTx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [uint32](#uint32) |  |  |
| `data` | [bytes](#bytes) |  |  |
| `log` | [string](#string) |  | nondeterministic |
| `info` | [string](#string) |  | nondeterministic |
| `gas_wanted` | [int64](#int64) |  |  |
| `gas_used` | [int64](#int64) |  |  |
| `events` | [Event](#tendermint.abci.Event) | repeated | nondeterministic |
| `codespace` | [string](#string) |  |  |






<a name="tendermint.abci.ResponseEcho"></a>

### ResponseEcho



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `message` | [string](#string) |  |  |






<a name="tendermint.abci.ResponseEndBlock"></a>

### ResponseEndBlock



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_updates` | [ValidatorUpdate](#tendermint.abci.ValidatorUpdate) | repeated |  |
| `consensus_param_updates` | [ConsensusParams](#tendermint.abci.ConsensusParams) |  |  |
| `events` | [Event](#tendermint.abci.Event) | repeated |  |






<a name="tendermint.abci.ResponseException"></a>

### ResponseException
nondeterministic


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `error` | [string](#string) |  |  |






<a name="tendermint.abci.ResponseFlush"></a>

### ResponseFlush







<a name="tendermint.abci.ResponseInfo"></a>

### ResponseInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [string](#string) |  |  |
| `version` | [string](#string) |  |  |
| `app_version` | [uint64](#uint64) |  |  |
| `last_block_height` | [int64](#int64) |  |  |
| `last_block_app_hash` | [bytes](#bytes) |  |  |






<a name="tendermint.abci.ResponseInitChain"></a>

### ResponseInitChain



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `consensus_params` | [ConsensusParams](#tendermint.abci.ConsensusParams) |  |  |
| `validators` | [ValidatorUpdate](#tendermint.abci.ValidatorUpdate) | repeated |  |
| `app_hash` | [bytes](#bytes) |  |  |






<a name="tendermint.abci.ResponseListSnapshots"></a>

### ResponseListSnapshots



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `snapshots` | [Snapshot](#tendermint.abci.Snapshot) | repeated |  |






<a name="tendermint.abci.ResponseLoadSnapshotChunk"></a>

### ResponseLoadSnapshotChunk



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chunk` | [bytes](#bytes) |  |  |






<a name="tendermint.abci.ResponseOfferSnapshot"></a>

### ResponseOfferSnapshot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `result` | [ResponseOfferSnapshot.Result](#tendermint.abci.ResponseOfferSnapshot.Result) |  |  |






<a name="tendermint.abci.ResponseQuery"></a>

### ResponseQuery



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [uint32](#uint32) |  |  |
| `log` | [string](#string) |  | bytes data = 2; // use "value" instead.

nondeterministic |
| `info` | [string](#string) |  | nondeterministic |
| `index` | [int64](#int64) |  |  |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |
| `proof_ops` | [tendermint.crypto.ProofOps](#tendermint.crypto.ProofOps) |  |  |
| `height` | [int64](#int64) |  |  |
| `codespace` | [string](#string) |  |  |






<a name="tendermint.abci.ResponseSetOption"></a>

### ResponseSetOption
nondeterministic


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [uint32](#uint32) |  |  |
| `log` | [string](#string) |  | bytes data = 2; |
| `info` | [string](#string) |  |  |






<a name="tendermint.abci.Snapshot"></a>

### Snapshot



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [uint64](#uint64) |  | The height at which the snapshot was taken |
| `format` | [uint32](#uint32) |  | The application-specific snapshot format |
| `chunks` | [uint32](#uint32) |  | Number of chunks in the snapshot |
| `hash` | [bytes](#bytes) |  | Arbitrary snapshot hash, equal only if identical |
| `metadata` | [bytes](#bytes) |  | Arbitrary application metadata |






<a name="tendermint.abci.TxResult"></a>

### TxResult
TxResult contains results of executing the transaction.

One usage is indexing transaction results.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  |  |
| `index` | [uint32](#uint32) |  |  |
| `tx` | [bytes](#bytes) |  |  |
| `result` | [ResponseDeliverTx](#tendermint.abci.ResponseDeliverTx) |  |  |






<a name="tendermint.abci.Validator"></a>

### Validator
Validator


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | The first 20 bytes of SHA256(public key) |
| `power` | [int64](#int64) |  | PubKey pub_key = 2 [(gogoproto.nullable)=false];

The voting power |






<a name="tendermint.abci.ValidatorUpdate"></a>

### ValidatorUpdate
ValidatorUpdate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pub_key` | [tendermint.crypto.PublicKey](#tendermint.crypto.PublicKey) |  |  |
| `power` | [int64](#int64) |  |  |






<a name="tendermint.abci.VoteInfo"></a>

### VoteInfo
VoteInfo


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [Validator](#tendermint.abci.Validator) |  |  |
| `signed_last_block` | [bool](#bool) |  |  |





 <!-- end messages -->


<a name="tendermint.abci.CheckTxType"></a>

### CheckTxType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NEW | 0 |  |
| RECHECK | 1 |  |



<a name="tendermint.abci.EvidenceType"></a>

### EvidenceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 |  |
| DUPLICATE_VOTE | 1 |  |
| LIGHT_CLIENT_ATTACK | 2 |  |



<a name="tendermint.abci.ResponseApplySnapshotChunk.Result"></a>

### ResponseApplySnapshotChunk.Result


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | Unknown result, abort all snapshot restoration |
| ACCEPT | 1 | Chunk successfully accepted |
| ABORT | 2 | Abort all snapshot restoration |
| RETRY | 3 | Retry chunk (combine with refetch and reject) |
| RETRY_SNAPSHOT | 4 | Retry snapshot (combine with refetch and reject) |
| REJECT_SNAPSHOT | 5 | Reject this snapshot, try others |



<a name="tendermint.abci.ResponseOfferSnapshot.Result"></a>

### ResponseOfferSnapshot.Result


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | Unknown result, abort all snapshot restoration |
| ACCEPT | 1 | Snapshot accepted, apply chunks |
| ABORT | 2 | Abort all snapshot restoration |
| REJECT | 3 | Reject this specific snapshot, try others |
| REJECT_FORMAT | 4 | Reject all snapshots of this format, try others |
| REJECT_SENDER | 5 | Reject all snapshots from the sender(s), try others |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="tendermint.abci.ABCIApplication"></a>

### ABCIApplication


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Echo` | [RequestEcho](#tendermint.abci.RequestEcho) | [ResponseEcho](#tendermint.abci.ResponseEcho) |  | |
| `Flush` | [RequestFlush](#tendermint.abci.RequestFlush) | [ResponseFlush](#tendermint.abci.ResponseFlush) |  | |
| `Info` | [RequestInfo](#tendermint.abci.RequestInfo) | [ResponseInfo](#tendermint.abci.ResponseInfo) |  | |
| `SetOption` | [RequestSetOption](#tendermint.abci.RequestSetOption) | [ResponseSetOption](#tendermint.abci.ResponseSetOption) |  | |
| `DeliverTx` | [RequestDeliverTx](#tendermint.abci.RequestDeliverTx) | [ResponseDeliverTx](#tendermint.abci.ResponseDeliverTx) |  | |
| `CheckTx` | [RequestCheckTx](#tendermint.abci.RequestCheckTx) | [ResponseCheckTx](#tendermint.abci.ResponseCheckTx) |  | |
| `Query` | [RequestQuery](#tendermint.abci.RequestQuery) | [ResponseQuery](#tendermint.abci.ResponseQuery) |  | |
| `Commit` | [RequestCommit](#tendermint.abci.RequestCommit) | [ResponseCommit](#tendermint.abci.ResponseCommit) |  | |
| `InitChain` | [RequestInitChain](#tendermint.abci.RequestInitChain) | [ResponseInitChain](#tendermint.abci.ResponseInitChain) |  | |
| `BeginBlock` | [RequestBeginBlock](#tendermint.abci.RequestBeginBlock) | [ResponseBeginBlock](#tendermint.abci.ResponseBeginBlock) |  | |
| `EndBlock` | [RequestEndBlock](#tendermint.abci.RequestEndBlock) | [ResponseEndBlock](#tendermint.abci.ResponseEndBlock) |  | |
| `ListSnapshots` | [RequestListSnapshots](#tendermint.abci.RequestListSnapshots) | [ResponseListSnapshots](#tendermint.abci.ResponseListSnapshots) |  | |
| `OfferSnapshot` | [RequestOfferSnapshot](#tendermint.abci.RequestOfferSnapshot) | [ResponseOfferSnapshot](#tendermint.abci.ResponseOfferSnapshot) |  | |
| `LoadSnapshotChunk` | [RequestLoadSnapshotChunk](#tendermint.abci.RequestLoadSnapshotChunk) | [ResponseLoadSnapshotChunk](#tendermint.abci.ResponseLoadSnapshotChunk) |  | |
| `ApplySnapshotChunk` | [RequestApplySnapshotChunk](#tendermint.abci.RequestApplySnapshotChunk) | [ResponseApplySnapshotChunk](#tendermint.abci.ResponseApplySnapshotChunk) |  | |

 <!-- end services -->



<a name="cosmos/base/abci/v1beta1/abci.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/abci/v1beta1/abci.proto



<a name="cosmos.base.abci.v1beta1.ABCIMessageLog"></a>

### ABCIMessageLog
ABCIMessageLog defines a structure containing an indexed tx ABCI message log.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_index` | [uint32](#uint32) |  |  |
| `log` | [string](#string) |  |  |
| `events` | [StringEvent](#cosmos.base.abci.v1beta1.StringEvent) | repeated | Events contains a slice of Event objects that were emitted during some execution. |






<a name="cosmos.base.abci.v1beta1.Attribute"></a>

### Attribute
Attribute defines an attribute wrapper where the key and value are
strings instead of raw bytes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="cosmos.base.abci.v1beta1.GasInfo"></a>

### GasInfo
GasInfo defines tx execution gas context.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_wanted` | [uint64](#uint64) |  | GasWanted is the maximum units of work we allow this tx to perform. |
| `gas_used` | [uint64](#uint64) |  | GasUsed is the amount of gas actually consumed. |






<a name="cosmos.base.abci.v1beta1.MsgData"></a>

### MsgData
MsgData defines the data returned in a Result object during message
execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_type` | [string](#string) |  |  |
| `data` | [bytes](#bytes) |  |  |






<a name="cosmos.base.abci.v1beta1.Result"></a>

### Result
Result is the union of ResponseFormat and ResponseCheckTx.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  | **Deprecated.** Data is any data returned from message or handler execution. It MUST be length prefixed in order to separate data from multiple message executions. Deprecated. This field is still populated, but prefer msg_response instead because it also contains the Msg response typeURL. |
| `log` | [string](#string) |  | Log contains the log information from message or handler execution. |
| `events` | [tendermint.abci.Event](#tendermint.abci.Event) | repeated | Events contains a slice of Event objects that were emitted during message or handler execution. |
| `msg_responses` | [google.protobuf.Any](#google.protobuf.Any) | repeated | msg_responses contains the Msg handler responses type packed in Anys.

Since: cosmos-sdk 0.46 |






<a name="cosmos.base.abci.v1beta1.SearchTxsResult"></a>

### SearchTxsResult
SearchTxsResult defines a structure for querying txs pageable


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `total_count` | [uint64](#uint64) |  | Count of all txs |
| `count` | [uint64](#uint64) |  | Count of txs in current page |
| `page_number` | [uint64](#uint64) |  | Index of current page, start from 1 |
| `page_total` | [uint64](#uint64) |  | Count of total pages |
| `limit` | [uint64](#uint64) |  | Max count txs per page |
| `txs` | [TxResponse](#cosmos.base.abci.v1beta1.TxResponse) | repeated | List of txs in current page |






<a name="cosmos.base.abci.v1beta1.SimulationResponse"></a>

### SimulationResponse
SimulationResponse defines the response generated when a transaction is
successfully simulated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_info` | [GasInfo](#cosmos.base.abci.v1beta1.GasInfo) |  |  |
| `result` | [Result](#cosmos.base.abci.v1beta1.Result) |  |  |






<a name="cosmos.base.abci.v1beta1.StringEvent"></a>

### StringEvent
StringEvent defines en Event object wrapper where all the attributes
contain key/value pairs that are strings instead of raw bytes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `attributes` | [Attribute](#cosmos.base.abci.v1beta1.Attribute) | repeated |  |






<a name="cosmos.base.abci.v1beta1.TxMsgData"></a>

### TxMsgData
TxMsgData defines a list of MsgData. A transaction will have a MsgData object
for each message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [MsgData](#cosmos.base.abci.v1beta1.MsgData) | repeated | **Deprecated.** data field is deprecated and not populated. |
| `msg_responses` | [google.protobuf.Any](#google.protobuf.Any) | repeated | msg_responses contains the Msg handler responses packed into Anys.

Since: cosmos-sdk 0.46 |






<a name="cosmos.base.abci.v1beta1.TxResponse"></a>

### TxResponse
TxResponse defines a structure containing relevant tx data and metadata. The
tags are stringified and the log is JSON decoded.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | The block height |
| `txhash` | [string](#string) |  | The transaction hash. |
| `codespace` | [string](#string) |  | Namespace for the Code |
| `code` | [uint32](#uint32) |  | Response code. |
| `data` | [string](#string) |  | Result bytes, if any. |
| `raw_log` | [string](#string) |  | The output of the application's logger (raw string). May be non-deterministic. |
| `logs` | [ABCIMessageLog](#cosmos.base.abci.v1beta1.ABCIMessageLog) | repeated | The output of the application's logger (typed). May be non-deterministic. |
| `info` | [string](#string) |  | Additional information. May be non-deterministic. |
| `gas_wanted` | [int64](#int64) |  | Amount of gas requested for transaction. |
| `gas_used` | [int64](#int64) |  | Amount of gas consumed by transaction. |
| `tx` | [google.protobuf.Any](#google.protobuf.Any) |  | The request transaction bytes. |
| `timestamp` | [string](#string) |  | Time of the previous block. For heights > 1, it's the weighted median of the timestamps of the valid votes in the block.LastCommit. For height == 1, it's genesis time. |
| `events` | [tendermint.abci.Event](#tendermint.abci.Event) | repeated | Events defines all the events emitted by processing a transaction. Note, these events include those emitted by processing all the messages and those emitted from the ante. Whereas Logs contains the events, with additional metadata, emitted only by processing the messages.

Since: cosmos-sdk 0.42.11, 0.44.5, 0.45 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/kv/v1beta1/kv.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/kv/v1beta1/kv.proto



<a name="cosmos.base.kv.v1beta1.Pair"></a>

### Pair
Pair defines a key/value bytes tuple.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="cosmos.base.kv.v1beta1.Pairs"></a>

### Pairs
Pairs defines a repeated slice of Pair objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pairs` | [Pair](#cosmos.base.kv.v1beta1.Pair) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/node/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/node/v1beta1/query.proto



<a name="cosmos.base.node.v1beta1.ConfigRequest"></a>

### ConfigRequest
ConfigRequest defines the request structure for the Config gRPC query.






<a name="cosmos.base.node.v1beta1.ConfigResponse"></a>

### ConfigResponse
ConfigResponse defines the response structure for the Config gRPC query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minimum_gas_price` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.base.node.v1beta1.Service"></a>

### Service
Service defines the gRPC querier service for node related queries.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Config` | [ConfigRequest](#cosmos.base.node.v1beta1.ConfigRequest) | [ConfigResponse](#cosmos.base.node.v1beta1.ConfigResponse) | Config queries for the operator configuration. | GET|/cosmos/base/node/v1beta1/config|

 <!-- end services -->



<a name="cosmos/base/reflection/v1beta1/reflection.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/reflection/v1beta1/reflection.proto



<a name="cosmos.base.reflection.v1beta1.ListAllInterfacesRequest"></a>

### ListAllInterfacesRequest
ListAllInterfacesRequest is the request type of the ListAllInterfaces RPC.






<a name="cosmos.base.reflection.v1beta1.ListAllInterfacesResponse"></a>

### ListAllInterfacesResponse
ListAllInterfacesResponse is the response type of the ListAllInterfaces RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interface_names` | [string](#string) | repeated | interface_names is an array of all the registered interfaces. |






<a name="cosmos.base.reflection.v1beta1.ListImplementationsRequest"></a>

### ListImplementationsRequest
ListImplementationsRequest is the request type of the ListImplementations
RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interface_name` | [string](#string) |  | interface_name defines the interface to query the implementations for. |






<a name="cosmos.base.reflection.v1beta1.ListImplementationsResponse"></a>

### ListImplementationsResponse
ListImplementationsResponse is the response type of the ListImplementations
RPC.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `implementation_message_names` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.base.reflection.v1beta1.ReflectionService"></a>

### ReflectionService
ReflectionService defines a service for interface reflection.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ListAllInterfaces` | [ListAllInterfacesRequest](#cosmos.base.reflection.v1beta1.ListAllInterfacesRequest) | [ListAllInterfacesResponse](#cosmos.base.reflection.v1beta1.ListAllInterfacesResponse) | ListAllInterfaces lists all the interfaces registered in the interface registry. | GET|/cosmos/base/reflection/v1beta1/interfaces|
| `ListImplementations` | [ListImplementationsRequest](#cosmos.base.reflection.v1beta1.ListImplementationsRequest) | [ListImplementationsResponse](#cosmos.base.reflection.v1beta1.ListImplementationsResponse) | ListImplementations list all the concrete types that implement a given interface. | GET|/cosmos/base/reflection/v1beta1/interfaces/{interface_name}/implementations|

 <!-- end services -->



<a name="cosmos/base/reflection/v2alpha1/reflection.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/reflection/v2alpha1/reflection.proto
Since: cosmos-sdk 0.43


<a name="cosmos.base.reflection.v2alpha1.AppDescriptor"></a>

### AppDescriptor
AppDescriptor describes a cosmos-sdk based application


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authn` | [AuthnDescriptor](#cosmos.base.reflection.v2alpha1.AuthnDescriptor) |  | AuthnDescriptor provides information on how to authenticate transactions on the application NOTE: experimental and subject to change in future releases. |
| `chain` | [ChainDescriptor](#cosmos.base.reflection.v2alpha1.ChainDescriptor) |  | chain provides the chain descriptor |
| `codec` | [CodecDescriptor](#cosmos.base.reflection.v2alpha1.CodecDescriptor) |  | codec provides metadata information regarding codec related types |
| `configuration` | [ConfigurationDescriptor](#cosmos.base.reflection.v2alpha1.ConfigurationDescriptor) |  | configuration provides metadata information regarding the sdk.Config type |
| `query_services` | [QueryServicesDescriptor](#cosmos.base.reflection.v2alpha1.QueryServicesDescriptor) |  | query_services provides metadata information regarding the available queriable endpoints |
| `tx` | [TxDescriptor](#cosmos.base.reflection.v2alpha1.TxDescriptor) |  | tx provides metadata information regarding how to send transactions to the given application |






<a name="cosmos.base.reflection.v2alpha1.AuthnDescriptor"></a>

### AuthnDescriptor
AuthnDescriptor provides information on how to sign transactions without relying
on the online RPCs GetTxMetadata and CombineUnsignedTxAndSignatures


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sign_modes` | [SigningModeDescriptor](#cosmos.base.reflection.v2alpha1.SigningModeDescriptor) | repeated | sign_modes defines the supported signature algorithm |






<a name="cosmos.base.reflection.v2alpha1.ChainDescriptor"></a>

### ChainDescriptor
ChainDescriptor describes chain information of the application


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  | id is the chain id |






<a name="cosmos.base.reflection.v2alpha1.CodecDescriptor"></a>

### CodecDescriptor
CodecDescriptor describes the registered interfaces and provides metadata information on the types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `interfaces` | [InterfaceDescriptor](#cosmos.base.reflection.v2alpha1.InterfaceDescriptor) | repeated | interfaces is a list of the registerted interfaces descriptors |






<a name="cosmos.base.reflection.v2alpha1.ConfigurationDescriptor"></a>

### ConfigurationDescriptor
ConfigurationDescriptor contains metadata information on the sdk.Config


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bech32_account_address_prefix` | [string](#string) |  | bech32_account_address_prefix is the account address prefix |






<a name="cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest"></a>

### GetAuthnDescriptorRequest
GetAuthnDescriptorRequest is the request used for the GetAuthnDescriptor RPC






<a name="cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse"></a>

### GetAuthnDescriptorResponse
GetAuthnDescriptorResponse is the response returned by the GetAuthnDescriptor RPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authn` | [AuthnDescriptor](#cosmos.base.reflection.v2alpha1.AuthnDescriptor) |  | authn describes how to authenticate to the application when sending transactions |






<a name="cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest"></a>

### GetChainDescriptorRequest
GetChainDescriptorRequest is the request used for the GetChainDescriptor RPC






<a name="cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse"></a>

### GetChainDescriptorResponse
GetChainDescriptorResponse is the response returned by the GetChainDescriptor RPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chain` | [ChainDescriptor](#cosmos.base.reflection.v2alpha1.ChainDescriptor) |  | chain describes application chain information |






<a name="cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest"></a>

### GetCodecDescriptorRequest
GetCodecDescriptorRequest is the request used for the GetCodecDescriptor RPC






<a name="cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse"></a>

### GetCodecDescriptorResponse
GetCodecDescriptorResponse is the response returned by the GetCodecDescriptor RPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `codec` | [CodecDescriptor](#cosmos.base.reflection.v2alpha1.CodecDescriptor) |  | codec describes the application codec such as registered interfaces and implementations |






<a name="cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest"></a>

### GetConfigurationDescriptorRequest
GetConfigurationDescriptorRequest is the request used for the GetConfigurationDescriptor RPC






<a name="cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse"></a>

### GetConfigurationDescriptorResponse
GetConfigurationDescriptorResponse is the response returned by the GetConfigurationDescriptor RPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `config` | [ConfigurationDescriptor](#cosmos.base.reflection.v2alpha1.ConfigurationDescriptor) |  | config describes the application's sdk.Config |






<a name="cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest"></a>

### GetQueryServicesDescriptorRequest
GetQueryServicesDescriptorRequest is the request used for the GetQueryServicesDescriptor RPC






<a name="cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse"></a>

### GetQueryServicesDescriptorResponse
GetQueryServicesDescriptorResponse is the response returned by the GetQueryServicesDescriptor RPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `queries` | [QueryServicesDescriptor](#cosmos.base.reflection.v2alpha1.QueryServicesDescriptor) |  | queries provides information on the available queryable services |






<a name="cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest"></a>

### GetTxDescriptorRequest
GetTxDescriptorRequest is the request used for the GetTxDescriptor RPC






<a name="cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse"></a>

### GetTxDescriptorResponse
GetTxDescriptorResponse is the response returned by the GetTxDescriptor RPC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [TxDescriptor](#cosmos.base.reflection.v2alpha1.TxDescriptor) |  | tx provides information on msgs that can be forwarded to the application alongside the accepted transaction protobuf type |






<a name="cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor"></a>

### InterfaceAcceptingMessageDescriptor
InterfaceAcceptingMessageDescriptor describes a protobuf message which contains
an interface represented as a google.protobuf.Any


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fullname` | [string](#string) |  | fullname is the protobuf fullname of the type containing the interface |
| `field_descriptor_names` | [string](#string) | repeated | field_descriptor_names is a list of the protobuf name (not fullname) of the field which contains the interface as google.protobuf.Any (the interface is the same, but it can be in multiple fields of the same proto message) |






<a name="cosmos.base.reflection.v2alpha1.InterfaceDescriptor"></a>

### InterfaceDescriptor
InterfaceDescriptor describes the implementation of an interface


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fullname` | [string](#string) |  | fullname is the name of the interface |
| `interface_accepting_messages` | [InterfaceAcceptingMessageDescriptor](#cosmos.base.reflection.v2alpha1.InterfaceAcceptingMessageDescriptor) | repeated | interface_accepting_messages contains information regarding the proto messages which contain the interface as google.protobuf.Any field |
| `interface_implementers` | [InterfaceImplementerDescriptor](#cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor) | repeated | interface_implementers is a list of the descriptors of the interface implementers |






<a name="cosmos.base.reflection.v2alpha1.InterfaceImplementerDescriptor"></a>

### InterfaceImplementerDescriptor
InterfaceImplementerDescriptor describes an interface implementer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fullname` | [string](#string) |  | fullname is the protobuf queryable name of the interface implementer |
| `type_url` | [string](#string) |  | type_url defines the type URL used when marshalling the type as any this is required so we can provide type safe google.protobuf.Any marshalling and unmarshalling, making sure that we don't accept just 'any' type in our interface fields |






<a name="cosmos.base.reflection.v2alpha1.MsgDescriptor"></a>

### MsgDescriptor
MsgDescriptor describes a cosmos-sdk message that can be delivered with a transaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `msg_type_url` | [string](#string) |  | msg_type_url contains the TypeURL of a sdk.Msg. |






<a name="cosmos.base.reflection.v2alpha1.QueryMethodDescriptor"></a>

### QueryMethodDescriptor
QueryMethodDescriptor describes a queryable method of a query service
no other info is provided beside method name and tendermint queryable path
because it would be redundant with the grpc reflection service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name is the protobuf name (not fullname) of the method |
| `full_query_path` | [string](#string) |  | full_query_path is the path that can be used to query this method via tendermint abci.Query |






<a name="cosmos.base.reflection.v2alpha1.QueryServiceDescriptor"></a>

### QueryServiceDescriptor
QueryServiceDescriptor describes a cosmos-sdk queryable service


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fullname` | [string](#string) |  | fullname is the protobuf fullname of the service descriptor |
| `is_module` | [bool](#bool) |  | is_module describes if this service is actually exposed by an application's module |
| `methods` | [QueryMethodDescriptor](#cosmos.base.reflection.v2alpha1.QueryMethodDescriptor) | repeated | methods provides a list of query service methods |






<a name="cosmos.base.reflection.v2alpha1.QueryServicesDescriptor"></a>

### QueryServicesDescriptor
QueryServicesDescriptor contains the list of cosmos-sdk queriable services


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `query_services` | [QueryServiceDescriptor](#cosmos.base.reflection.v2alpha1.QueryServiceDescriptor) | repeated | query_services is a list of cosmos-sdk QueryServiceDescriptor |






<a name="cosmos.base.reflection.v2alpha1.SigningModeDescriptor"></a>

### SigningModeDescriptor
SigningModeDescriptor provides information on a signing flow of the application
NOTE(fdymylja): here we could go as far as providing an entire flow on how
to sign a message given a SigningModeDescriptor, but it's better to think about
this another time


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name defines the unique name of the signing mode |
| `number` | [int32](#int32) |  | number is the unique int32 identifier for the sign_mode enum |
| `authn_info_provider_method_fullname` | [string](#string) |  | authn_info_provider_method_fullname defines the fullname of the method to call to get the metadata required to authenticate using the provided sign_modes |






<a name="cosmos.base.reflection.v2alpha1.TxDescriptor"></a>

### TxDescriptor
TxDescriptor describes the accepted transaction type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fullname` | [string](#string) |  | fullname is the protobuf fullname of the raw transaction type (for instance the tx.Tx type) it is not meant to support polymorphism of transaction types, it is supposed to be used by reflection clients to understand if they can handle a specific transaction type in an application. |
| `msgs` | [MsgDescriptor](#cosmos.base.reflection.v2alpha1.MsgDescriptor) | repeated | msgs lists the accepted application messages (sdk.Msg) |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.base.reflection.v2alpha1.ReflectionService"></a>

### ReflectionService
ReflectionService defines a service for application reflection.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetAuthnDescriptor` | [GetAuthnDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetAuthnDescriptorRequest) | [GetAuthnDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetAuthnDescriptorResponse) | GetAuthnDescriptor returns information on how to authenticate transactions in the application NOTE: this RPC is still experimental and might be subject to breaking changes or removal in future releases of the cosmos-sdk. | GET|/cosmos/base/reflection/v1beta1/app_descriptor/authn|
| `GetChainDescriptor` | [GetChainDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetChainDescriptorRequest) | [GetChainDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetChainDescriptorResponse) | GetChainDescriptor returns the description of the chain | GET|/cosmos/base/reflection/v1beta1/app_descriptor/chain|
| `GetCodecDescriptor` | [GetCodecDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetCodecDescriptorRequest) | [GetCodecDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetCodecDescriptorResponse) | GetCodecDescriptor returns the descriptor of the codec of the application | GET|/cosmos/base/reflection/v1beta1/app_descriptor/codec|
| `GetConfigurationDescriptor` | [GetConfigurationDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorRequest) | [GetConfigurationDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetConfigurationDescriptorResponse) | GetConfigurationDescriptor returns the descriptor for the sdk.Config of the application | GET|/cosmos/base/reflection/v1beta1/app_descriptor/configuration|
| `GetQueryServicesDescriptor` | [GetQueryServicesDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorRequest) | [GetQueryServicesDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetQueryServicesDescriptorResponse) | GetQueryServicesDescriptor returns the available gRPC queryable services of the application | GET|/cosmos/base/reflection/v1beta1/app_descriptor/query_services|
| `GetTxDescriptor` | [GetTxDescriptorRequest](#cosmos.base.reflection.v2alpha1.GetTxDescriptorRequest) | [GetTxDescriptorResponse](#cosmos.base.reflection.v2alpha1.GetTxDescriptorResponse) | GetTxDescriptor returns information on the used transaction object and available msgs that can be used | GET|/cosmos/base/reflection/v1beta1/app_descriptor/tx_descriptor|

 <!-- end services -->



<a name="cosmos/base/snapshots/v1beta1/snapshot.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/snapshots/v1beta1/snapshot.proto



<a name="cosmos.base.snapshots.v1beta1.Metadata"></a>

### Metadata
Metadata contains SDK-specific snapshot metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `chunk_hashes` | [bytes](#bytes) | repeated | SHA-256 chunk hashes |






<a name="cosmos.base.snapshots.v1beta1.Snapshot"></a>

### Snapshot
Snapshot contains Tendermint state sync snapshot info.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [uint64](#uint64) |  |  |
| `format` | [uint32](#uint32) |  |  |
| `chunks` | [uint32](#uint32) |  |  |
| `hash` | [bytes](#bytes) |  |  |
| `metadata` | [Metadata](#cosmos.base.snapshots.v1beta1.Metadata) |  |  |






<a name="cosmos.base.snapshots.v1beta1.SnapshotExtensionMeta"></a>

### SnapshotExtensionMeta
SnapshotExtensionMeta contains metadata about an external snapshotter.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `format` | [uint32](#uint32) |  |  |






<a name="cosmos.base.snapshots.v1beta1.SnapshotExtensionPayload"></a>

### SnapshotExtensionPayload
SnapshotExtensionPayload contains payloads of an external snapshotter.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `payload` | [bytes](#bytes) |  |  |






<a name="cosmos.base.snapshots.v1beta1.SnapshotIAVLItem"></a>

### SnapshotIAVLItem
SnapshotIAVLItem is an exported IAVL node.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |
| `version` | [int64](#int64) |  | version is block height |
| `height` | [int32](#int32) |  | height is depth of the tree. |






<a name="cosmos.base.snapshots.v1beta1.SnapshotItem"></a>

### SnapshotItem
SnapshotItem is an item contained in a rootmulti.Store snapshot.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `store` | [SnapshotStoreItem](#cosmos.base.snapshots.v1beta1.SnapshotStoreItem) |  |  |
| `iavl` | [SnapshotIAVLItem](#cosmos.base.snapshots.v1beta1.SnapshotIAVLItem) |  |  |
| `extension` | [SnapshotExtensionMeta](#cosmos.base.snapshots.v1beta1.SnapshotExtensionMeta) |  |  |
| `extension_payload` | [SnapshotExtensionPayload](#cosmos.base.snapshots.v1beta1.SnapshotExtensionPayload) |  |  |
| `kv` | [SnapshotKVItem](#cosmos.base.snapshots.v1beta1.SnapshotKVItem) |  |  |
| `schema` | [SnapshotSchema](#cosmos.base.snapshots.v1beta1.SnapshotSchema) |  |  |






<a name="cosmos.base.snapshots.v1beta1.SnapshotKVItem"></a>

### SnapshotKVItem
SnapshotKVItem is an exported Key/Value Pair

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="cosmos.base.snapshots.v1beta1.SnapshotSchema"></a>

### SnapshotSchema
SnapshotSchema is an exported schema of smt store

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `keys` | [bytes](#bytes) | repeated |  |






<a name="cosmos.base.snapshots.v1beta1.SnapshotStoreItem"></a>

### SnapshotStoreItem
SnapshotStoreItem contains metadata about a snapshotted store.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/store/v1beta1/commit_info.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/store/v1beta1/commit_info.proto



<a name="cosmos.base.store.v1beta1.CommitID"></a>

### CommitID
CommitID defines the committment information when a specific store is
committed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [int64](#int64) |  |  |
| `hash` | [bytes](#bytes) |  |  |






<a name="cosmos.base.store.v1beta1.CommitInfo"></a>

### CommitInfo
CommitInfo defines commit information used by the multi-store when committing
a version/height.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [int64](#int64) |  |  |
| `store_infos` | [StoreInfo](#cosmos.base.store.v1beta1.StoreInfo) | repeated |  |






<a name="cosmos.base.store.v1beta1.StoreInfo"></a>

### StoreInfo
StoreInfo defines store-specific commit information. It contains a reference
between a store name and the commit ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `commit_id` | [CommitID](#cosmos.base.store.v1beta1.CommitID) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/store/v1beta1/listening.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/store/v1beta1/listening.proto



<a name="cosmos.base.store.v1beta1.BlockMetadata"></a>

### BlockMetadata
BlockMetadata contains all the abci event data of a block
the file streamer dump them into files together with the state changes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `request_begin_block` | [tendermint.abci.RequestBeginBlock](#tendermint.abci.RequestBeginBlock) |  |  |
| `response_begin_block` | [tendermint.abci.ResponseBeginBlock](#tendermint.abci.ResponseBeginBlock) |  |  |
| `deliver_txs` | [BlockMetadata.DeliverTx](#cosmos.base.store.v1beta1.BlockMetadata.DeliverTx) | repeated |  |
| `request_end_block` | [tendermint.abci.RequestEndBlock](#tendermint.abci.RequestEndBlock) |  |  |
| `response_end_block` | [tendermint.abci.ResponseEndBlock](#tendermint.abci.ResponseEndBlock) |  |  |
| `response_commit` | [tendermint.abci.ResponseCommit](#tendermint.abci.ResponseCommit) |  |  |






<a name="cosmos.base.store.v1beta1.BlockMetadata.DeliverTx"></a>

### BlockMetadata.DeliverTx
DeliverTx encapulate deliver tx request and response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `request` | [tendermint.abci.RequestDeliverTx](#tendermint.abci.RequestDeliverTx) |  |  |
| `response` | [tendermint.abci.ResponseDeliverTx](#tendermint.abci.ResponseDeliverTx) |  |  |






<a name="cosmos.base.store.v1beta1.StoreKVPair"></a>

### StoreKVPair
StoreKVPair is a KVStore KVPair used for listening to state changes (Sets and Deletes)
It optionally includes the StoreKey for the originating KVStore and a Boolean flag to distinguish between Sets and
Deletes

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `store_key` | [string](#string) |  | the store key for the KVStore this pair originates from |
| `delete` | [bool](#bool) |  | true indicates a delete operation, false indicates a set operation |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/p2p/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/p2p/types.proto



<a name="tendermint.p2p.DefaultNodeInfo"></a>

### DefaultNodeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `protocol_version` | [ProtocolVersion](#tendermint.p2p.ProtocolVersion) |  |  |
| `default_node_id` | [string](#string) |  |  |
| `listen_addr` | [string](#string) |  |  |
| `network` | [string](#string) |  |  |
| `version` | [string](#string) |  |  |
| `channels` | [bytes](#bytes) |  |  |
| `moniker` | [string](#string) |  |  |
| `other` | [DefaultNodeInfoOther](#tendermint.p2p.DefaultNodeInfoOther) |  |  |






<a name="tendermint.p2p.DefaultNodeInfoOther"></a>

### DefaultNodeInfoOther



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_index` | [string](#string) |  |  |
| `rpc_address` | [string](#string) |  |  |






<a name="tendermint.p2p.NetAddress"></a>

### NetAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  |  |
| `ip` | [string](#string) |  |  |
| `port` | [uint32](#uint32) |  |  |






<a name="tendermint.p2p.ProtocolVersion"></a>

### ProtocolVersion



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `p2p` | [uint64](#uint64) |  |  |
| `block` | [uint64](#uint64) |  |  |
| `app` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/types/evidence.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/types/evidence.proto



<a name="tendermint.types.DuplicateVoteEvidence"></a>

### DuplicateVoteEvidence
DuplicateVoteEvidence contains evidence of a validator signed two conflicting votes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote_a` | [Vote](#tendermint.types.Vote) |  |  |
| `vote_b` | [Vote](#tendermint.types.Vote) |  |  |
| `total_voting_power` | [int64](#int64) |  |  |
| `validator_power` | [int64](#int64) |  |  |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="tendermint.types.Evidence"></a>

### Evidence



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `duplicate_vote_evidence` | [DuplicateVoteEvidence](#tendermint.types.DuplicateVoteEvidence) |  |  |
| `light_client_attack_evidence` | [LightClientAttackEvidence](#tendermint.types.LightClientAttackEvidence) |  |  |






<a name="tendermint.types.EvidenceList"></a>

### EvidenceList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evidence` | [Evidence](#tendermint.types.Evidence) | repeated |  |






<a name="tendermint.types.LightClientAttackEvidence"></a>

### LightClientAttackEvidence
LightClientAttackEvidence contains evidence of a set of validators attempting to mislead a light client.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `conflicting_block` | [LightBlock](#tendermint.types.LightBlock) |  |  |
| `common_height` | [int64](#int64) |  |  |
| `byzantine_validators` | [Validator](#tendermint.types.Validator) | repeated |  |
| `total_voting_power` | [int64](#int64) |  |  |
| `timestamp` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/tendermint/v1beta1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/tendermint/v1beta1/types.proto



<a name="cosmos.base.tendermint.v1beta1.Block"></a>

### Block
Block is tendermint type Block, with the Header proposer address
field converted to bech32 string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header` | [Header](#cosmos.base.tendermint.v1beta1.Header) |  |  |
| `data` | [tendermint.types.Data](#tendermint.types.Data) |  |  |
| `evidence` | [tendermint.types.EvidenceList](#tendermint.types.EvidenceList) |  |  |
| `last_commit` | [tendermint.types.Commit](#tendermint.types.Commit) |  |  |






<a name="cosmos.base.tendermint.v1beta1.Header"></a>

### Header
Header defines the structure of a Tendermint block header.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `version` | [tendermint.version.Consensus](#tendermint.version.Consensus) |  | basic block info |
| `chain_id` | [string](#string) |  |  |
| `height` | [int64](#int64) |  |  |
| `time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `last_block_id` | [tendermint.types.BlockID](#tendermint.types.BlockID) |  | prev block info |
| `last_commit_hash` | [bytes](#bytes) |  | hashes of block data

commit from validators from the last block |
| `data_hash` | [bytes](#bytes) |  | transactions |
| `validators_hash` | [bytes](#bytes) |  | hashes from the app output from the prev block

validators for the current block |
| `next_validators_hash` | [bytes](#bytes) |  | validators for the next block |
| `consensus_hash` | [bytes](#bytes) |  | consensus params for current block |
| `app_hash` | [bytes](#bytes) |  | state after txs from the previous block |
| `last_results_hash` | [bytes](#bytes) |  | root hash of all results from the txs from the previous block |
| `evidence_hash` | [bytes](#bytes) |  | consensus info

evidence included in the block |
| `proposer_address` | [string](#string) |  | proposer_address is the original block proposer address, formatted as a Bech32 string. In Tendermint, this type is `bytes`, but in the SDK, we convert it to a Bech32 string for better UX.

original proposer of the block |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="tendermint/types/block.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/types/block.proto



<a name="tendermint.types.Block"></a>

### Block



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header` | [Header](#tendermint.types.Header) |  |  |
| `data` | [Data](#tendermint.types.Data) |  |  |
| `evidence` | [EvidenceList](#tendermint.types.EvidenceList) |  |  |
| `last_commit` | [Commit](#tendermint.types.Commit) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/base/tendermint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/base/tendermint/v1beta1/query.proto



<a name="cosmos.base.tendermint.v1beta1.ABCIQueryRequest"></a>

### ABCIQueryRequest
ABCIQueryRequest defines the request structure for the ABCIQuery gRPC query.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `data` | [bytes](#bytes) |  |  |
| `path` | [string](#string) |  |  |
| `height` | [int64](#int64) |  |  |
| `prove` | [bool](#bool) |  |  |






<a name="cosmos.base.tendermint.v1beta1.ABCIQueryResponse"></a>

### ABCIQueryResponse
ABCIQueryResponse defines the response structure for the ABCIQuery gRPC
query.

Note: This type is a duplicate of the ResponseQuery proto type defined in
Tendermint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [uint32](#uint32) |  |  |
| `log` | [string](#string) |  | nondeterministic |
| `info` | [string](#string) |  | nondeterministic |
| `index` | [int64](#int64) |  |  |
| `key` | [bytes](#bytes) |  |  |
| `value` | [bytes](#bytes) |  |  |
| `proof_ops` | [ProofOps](#cosmos.base.tendermint.v1beta1.ProofOps) |  |  |
| `height` | [int64](#int64) |  |  |
| `codespace` | [string](#string) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest"></a>

### GetBlockByHeightRequest
GetBlockByHeightRequest is the request type for the Query/GetBlockByHeight
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse"></a>

### GetBlockByHeightResponse
GetBlockByHeightResponse is the response type for the Query/GetBlockByHeight
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_id` | [tendermint.types.BlockID](#tendermint.types.BlockID) |  |  |
| `block` | [tendermint.types.Block](#tendermint.types.Block) |  | Deprecated: please use `sdk_block` instead |
| `sdk_block` | [Block](#cosmos.base.tendermint.v1beta1.Block) |  | Since: cosmos-sdk 0.47 |






<a name="cosmos.base.tendermint.v1beta1.GetLatestBlockRequest"></a>

### GetLatestBlockRequest
GetLatestBlockRequest is the request type for the Query/GetLatestBlock RPC
method.






<a name="cosmos.base.tendermint.v1beta1.GetLatestBlockResponse"></a>

### GetLatestBlockResponse
GetLatestBlockResponse is the response type for the Query/GetLatestBlock RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_id` | [tendermint.types.BlockID](#tendermint.types.BlockID) |  |  |
| `block` | [tendermint.types.Block](#tendermint.types.Block) |  | Deprecated: please use `sdk_block` instead |
| `sdk_block` | [Block](#cosmos.base.tendermint.v1beta1.Block) |  | Since: cosmos-sdk 0.47 |






<a name="cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest"></a>

### GetLatestValidatorSetRequest
GetLatestValidatorSetRequest is the request type for the
Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse"></a>

### GetLatestValidatorSetResponse
GetLatestValidatorSetResponse is the response type for the
Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_height` | [int64](#int64) |  |  |
| `validators` | [Validator](#cosmos.base.tendermint.v1beta1.Validator) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.base.tendermint.v1beta1.GetNodeInfoRequest"></a>

### GetNodeInfoRequest
GetNodeInfoRequest is the request type for the Query/GetNodeInfo RPC method.






<a name="cosmos.base.tendermint.v1beta1.GetNodeInfoResponse"></a>

### GetNodeInfoResponse
GetNodeInfoResponse is the response type for the Query/GetNodeInfo RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `default_node_info` | [tendermint.p2p.DefaultNodeInfo](#tendermint.p2p.DefaultNodeInfo) |  |  |
| `application_version` | [VersionInfo](#cosmos.base.tendermint.v1beta1.VersionInfo) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetSyncingRequest"></a>

### GetSyncingRequest
GetSyncingRequest is the request type for the Query/GetSyncing RPC method.






<a name="cosmos.base.tendermint.v1beta1.GetSyncingResponse"></a>

### GetSyncingResponse
GetSyncingResponse is the response type for the Query/GetSyncing RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `syncing` | [bool](#bool) |  |  |






<a name="cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest"></a>

### GetValidatorSetByHeightRequest
GetValidatorSetByHeightRequest is the request type for the
Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse"></a>

### GetValidatorSetByHeightResponse
GetValidatorSetByHeightResponse is the response type for the
Query/GetValidatorSetByHeight RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `block_height` | [int64](#int64) |  |  |
| `validators` | [Validator](#cosmos.base.tendermint.v1beta1.Validator) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.base.tendermint.v1beta1.Module"></a>

### Module
Module is the type for VersionInfo


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `path` | [string](#string) |  | module path |
| `version` | [string](#string) |  | module version |
| `sum` | [string](#string) |  | checksum |






<a name="cosmos.base.tendermint.v1beta1.ProofOp"></a>

### ProofOp
ProofOp defines an operation used for calculating Merkle root. The data could
be arbitrary format, providing nessecary data for example neighbouring node
hash.

Note: This type is a duplicate of the ProofOp proto type defined in
Tendermint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [string](#string) |  |  |
| `key` | [bytes](#bytes) |  |  |
| `data` | [bytes](#bytes) |  |  |






<a name="cosmos.base.tendermint.v1beta1.ProofOps"></a>

### ProofOps
ProofOps is Merkle proof defined by the list of ProofOps.

Note: This type is a duplicate of the ProofOps proto type defined in
Tendermint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ops` | [ProofOp](#cosmos.base.tendermint.v1beta1.ProofOp) | repeated |  |






<a name="cosmos.base.tendermint.v1beta1.Validator"></a>

### Validator
Validator is the type for the validator-set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `pub_key` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `voting_power` | [int64](#int64) |  |  |
| `proposer_priority` | [int64](#int64) |  |  |






<a name="cosmos.base.tendermint.v1beta1.VersionInfo"></a>

### VersionInfo
VersionInfo is the type for the GetNodeInfoResponse message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `app_name` | [string](#string) |  |  |
| `version` | [string](#string) |  |  |
| `git_commit` | [string](#string) |  |  |
| `build_tags` | [string](#string) |  |  |
| `go_version` | [string](#string) |  |  |
| `build_deps` | [Module](#cosmos.base.tendermint.v1beta1.Module) | repeated |  |
| `cosmos_sdk_version` | [string](#string) |  | Since: cosmos-sdk 0.43 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.base.tendermint.v1beta1.Service"></a>

### Service
Service defines the gRPC querier service for tendermint queries.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetNodeInfo` | [GetNodeInfoRequest](#cosmos.base.tendermint.v1beta1.GetNodeInfoRequest) | [GetNodeInfoResponse](#cosmos.base.tendermint.v1beta1.GetNodeInfoResponse) | GetNodeInfo queries the current node info. | GET|/cosmos/base/tendermint/v1beta1/node_info|
| `GetSyncing` | [GetSyncingRequest](#cosmos.base.tendermint.v1beta1.GetSyncingRequest) | [GetSyncingResponse](#cosmos.base.tendermint.v1beta1.GetSyncingResponse) | GetSyncing queries node syncing. | GET|/cosmos/base/tendermint/v1beta1/syncing|
| `GetLatestBlock` | [GetLatestBlockRequest](#cosmos.base.tendermint.v1beta1.GetLatestBlockRequest) | [GetLatestBlockResponse](#cosmos.base.tendermint.v1beta1.GetLatestBlockResponse) | GetLatestBlock returns the latest block. | GET|/cosmos/base/tendermint/v1beta1/blocks/latest|
| `GetBlockByHeight` | [GetBlockByHeightRequest](#cosmos.base.tendermint.v1beta1.GetBlockByHeightRequest) | [GetBlockByHeightResponse](#cosmos.base.tendermint.v1beta1.GetBlockByHeightResponse) | GetBlockByHeight queries block for given height. | GET|/cosmos/base/tendermint/v1beta1/blocks/{height}|
| `GetLatestValidatorSet` | [GetLatestValidatorSetRequest](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetRequest) | [GetLatestValidatorSetResponse](#cosmos.base.tendermint.v1beta1.GetLatestValidatorSetResponse) | GetLatestValidatorSet queries latest validator-set. | GET|/cosmos/base/tendermint/v1beta1/validatorsets/latest|
| `GetValidatorSetByHeight` | [GetValidatorSetByHeightRequest](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightRequest) | [GetValidatorSetByHeightResponse](#cosmos.base.tendermint.v1beta1.GetValidatorSetByHeightResponse) | GetValidatorSetByHeight queries validator-set at a given height. | GET|/cosmos/base/tendermint/v1beta1/validatorsets/{height}|
| `ABCIQuery` | [ABCIQueryRequest](#cosmos.base.tendermint.v1beta1.ABCIQueryRequest) | [ABCIQueryResponse](#cosmos.base.tendermint.v1beta1.ABCIQueryResponse) | ABCIQuery defines a query handler that supports ABCI queries directly to the application, bypassing Tendermint completely. The ABCI query must contain a valid and supported path, including app, custom, p2p, and store.

Since: cosmos-sdk 0.46 | GET|/cosmos/base/tendermint/v1beta1/abci_query|

 <!-- end services -->



<a name="cosmos/capability/v1beta1/capability.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/capability/v1beta1/capability.proto



<a name="cosmos.capability.v1beta1.Capability"></a>

### Capability
Capability defines an implementation of an object capability. The index
provided to a Capability must be globally unique.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint64](#uint64) |  |  |






<a name="cosmos.capability.v1beta1.CapabilityOwners"></a>

### CapabilityOwners
CapabilityOwners defines a set of owners of a single Capability. The set of
owners must be unique.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owners` | [Owner](#cosmos.capability.v1beta1.Owner) | repeated |  |






<a name="cosmos.capability.v1beta1.Owner"></a>

### Owner
Owner defines a single capability owner. An owner is defined by the name of
capability and the module name.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/capability/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/capability/v1beta1/genesis.proto



<a name="cosmos.capability.v1beta1.GenesisOwners"></a>

### GenesisOwners
GenesisOwners defines the capability owners with their corresponding index.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint64](#uint64) |  | index is the index of the capability owner. |
| `index_owners` | [CapabilityOwners](#cosmos.capability.v1beta1.CapabilityOwners) |  | index_owners are the owners at the given index. |






<a name="cosmos.capability.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the capability module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [uint64](#uint64) |  | index is the capability global index. |
| `owners` | [GenesisOwners](#cosmos.capability.v1beta1.GenesisOwners) | repeated | owners represents a map from index to owners of the capability index index key is string to allow amino marshalling. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crisis/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crisis/v1beta1/genesis.proto



<a name="cosmos.crisis.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the crisis module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `constant_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | constant_fee is the fee used to verify the invariant in the crisis module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crisis/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crisis/v1beta1/tx.proto



<a name="cosmos.crisis.v1beta1.MsgVerifyInvariant"></a>

### MsgVerifyInvariant
MsgVerifyInvariant represents a message to verify a particular invariance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `invariant_module_name` | [string](#string) |  |  |
| `invariant_route` | [string](#string) |  |  |






<a name="cosmos.crisis.v1beta1.MsgVerifyInvariantResponse"></a>

### MsgVerifyInvariantResponse
MsgVerifyInvariantResponse defines the Msg/VerifyInvariant response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.crisis.v1beta1.Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `VerifyInvariant` | [MsgVerifyInvariant](#cosmos.crisis.v1beta1.MsgVerifyInvariant) | [MsgVerifyInvariantResponse](#cosmos.crisis.v1beta1.MsgVerifyInvariantResponse) | VerifyInvariant defines a method to verify a particular invariance. | |

 <!-- end services -->



<a name="cosmos/crypto/ed25519/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/ed25519/keys.proto



<a name="cosmos.crypto.ed25519.PrivKey"></a>

### PrivKey
Deprecated: PrivKey defines a ed25519 private key.
NOTE: ed25519 keys must not be used in SDK apps except in a tendermint validator context.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.ed25519.PubKey"></a>

### PubKey
PubKey is an ed25519 public key for handling Tendermint keys in SDK.
It's needed for Any serialization and SDK compatibility.
It must not be used in a non Tendermint key context because it doesn't implement
ADR-28. Nevertheless, you will like to use ed25519 in app user level
then you must create a new proto message and follow ADR-28 for Address construction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/hd/v1/hd.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/hd/v1/hd.proto
Since: cosmos-sdk 0.46


<a name="cosmos.crypto.hd.v1.BIP44Params"></a>

### BIP44Params
BIP44Params is used as path field in ledger item in Record.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `purpose` | [uint32](#uint32) |  | purpose is a constant set to 44' (or 0x8000002C) following the BIP43 recommendation |
| `coin_type` | [uint32](#uint32) |  | coin_type is a constant that improves privacy |
| `account` | [uint32](#uint32) |  | account splits the key space into independent user identities |
| `change` | [bool](#bool) |  | change is a constant used for public derivation. Constant 0 is used for external chain and constant 1 for internal chain. |
| `address_index` | [uint32](#uint32) |  | address_index is used as child index in BIP32 derivation |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/keyring/v1/record.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/keyring/v1/record.proto
Since: cosmos-sdk 0.46


<a name="cosmos.crypto.keyring.v1.Record"></a>

### Record
Record is used for representing a key in the keyring.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name represents a name of Record |
| `pub_key` | [google.protobuf.Any](#google.protobuf.Any) |  | pub_key represents a public key in any format |
| `local` | [Record.Local](#cosmos.crypto.keyring.v1.Record.Local) |  | local stores the private key locally. |
| `ledger` | [Record.Ledger](#cosmos.crypto.keyring.v1.Record.Ledger) |  | ledger stores the information about a Ledger key. |
| `multi` | [Record.Multi](#cosmos.crypto.keyring.v1.Record.Multi) |  | Multi does not store any other information. |
| `offline` | [Record.Offline](#cosmos.crypto.keyring.v1.Record.Offline) |  | Offline does not store any other information. |






<a name="cosmos.crypto.keyring.v1.Record.Ledger"></a>

### Record.Ledger
Ledger item


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `path` | [cosmos.crypto.hd.v1.BIP44Params](#cosmos.crypto.hd.v1.BIP44Params) |  |  |






<a name="cosmos.crypto.keyring.v1.Record.Local"></a>

### Record.Local
Item is a keyring item stored in a keyring backend.
Local item


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `priv_key` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="cosmos.crypto.keyring.v1.Record.Multi"></a>

### Record.Multi
Multi item






<a name="cosmos.crypto.keyring.v1.Record.Offline"></a>

### Record.Offline
Offline item





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/multisig/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/multisig/keys.proto



<a name="cosmos.crypto.multisig.LegacyAminoPubKey"></a>

### LegacyAminoPubKey
LegacyAminoPubKey specifies a public key type
which nests multiple public keys and a threshold,
it uses legacy amino address rules.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `threshold` | [uint32](#uint32) |  |  |
| `public_keys` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/multisig/v1beta1/multisig.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/multisig/v1beta1/multisig.proto



<a name="cosmos.crypto.multisig.v1beta1.CompactBitArray"></a>

### CompactBitArray
CompactBitArray is an implementation of a space efficient bit array.
This is used to ensure that the encoded data takes up a minimal amount of
space after proto encoding.
This is not thread safe, and is not intended for concurrent usage.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `extra_bits_stored` | [uint32](#uint32) |  |  |
| `elems` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.multisig.v1beta1.MultiSignature"></a>

### MultiSignature
MultiSignature wraps the signatures from a multisig.LegacyAminoPubKey.
See cosmos.tx.v1betata1.ModeInfo.Multi for how to specify which signers
signed and with which modes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signatures` | [bytes](#bytes) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/secp256k1/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/secp256k1/keys.proto



<a name="cosmos.crypto.secp256k1.PrivKey"></a>

### PrivKey
PrivKey defines a secp256k1 private key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |






<a name="cosmos.crypto.secp256k1.PubKey"></a>

### PubKey
PubKey defines a secp256k1 public key
Key is the compressed form of the pubkey. The first byte depends is a 0x02 byte
if the y-coordinate is the lexicographically largest of the two associated with
the x-coordinate. Otherwise the first byte is a 0x03.
This prefix is followed with the x-coordinate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/crypto/secp256r1/keys.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/crypto/secp256r1/keys.proto
Since: cosmos-sdk 0.43


<a name="cosmos.crypto.secp256r1.PrivKey"></a>

### PrivKey
PrivKey defines a secp256r1 ECDSA private key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `secret` | [bytes](#bytes) |  | secret number serialized using big-endian encoding |






<a name="cosmos.crypto.secp256r1.PubKey"></a>

### PubKey
PubKey defines a secp256r1 ECDSA public key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `key` | [bytes](#bytes) |  | Point on secp256r1 curve in a compressed representation as specified in section 4.3.6 of ANSI X9.62: https://webstore.ansi.org/standards/ascx9/ansix9621998 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/distribution/v1beta1/distribution.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/distribution/v1beta1/distribution.proto



<a name="cosmos.distribution.v1beta1.CommunityPoolSpendProposal"></a>

### CommunityPoolSpendProposal
CommunityPoolSpendProposal details a proposal for use of community funds,
together with how many coins are proposed to be spent, and to which
recipient account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.distribution.v1beta1.CommunityPoolSpendProposalWithDeposit"></a>

### CommunityPoolSpendProposalWithDeposit
CommunityPoolSpendProposalWithDeposit defines a CommunityPoolSpendProposal
with a deposit


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `recipient` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |
| `deposit` | [string](#string) |  |  |






<a name="cosmos.distribution.v1beta1.DelegationDelegatorReward"></a>

### DelegationDelegatorReward
DelegationDelegatorReward represents the properties
of a delegator's delegation reward.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  |  |
| `reward` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="cosmos.distribution.v1beta1.DelegatorStartingInfo"></a>

### DelegatorStartingInfo
DelegatorStartingInfo represents the starting info for a delegator reward
period. It tracks the previous validator period, the delegation's amount of
staking token, and the creation height (to check later on if any slashes have
occurred). NOTE: Even though validators are slashed to whole staking tokens,
the delegators within the validator may be left with less than a full token,
thus sdk.Dec is used.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `previous_period` | [uint64](#uint64) |  |  |
| `stake` | [string](#string) |  |  |
| `height` | [uint64](#uint64) |  |  |






<a name="cosmos.distribution.v1beta1.FeePool"></a>

### FeePool
FeePool is the global fee pool for distribution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `community_pool` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="cosmos.distribution.v1beta1.Params"></a>

### Params
Params defines the set of params for the distribution module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `community_tax` | [string](#string) |  |  |
| `base_proposer_reward` | [string](#string) |  |  |
| `bonus_proposer_reward` | [string](#string) |  |  |
| `withdraw_addr_enabled` | [bool](#bool) |  |  |






<a name="cosmos.distribution.v1beta1.ValidatorAccumulatedCommission"></a>

### ValidatorAccumulatedCommission
ValidatorAccumulatedCommission represents accumulated commission
for a validator kept as a running counter, can be withdrawn at any time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commission` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="cosmos.distribution.v1beta1.ValidatorCurrentRewards"></a>

### ValidatorCurrentRewards
ValidatorCurrentRewards represents current rewards and current
period for a validator kept as a running counter and incremented
each block as long as the validator's tokens remain constant.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| `period` | [uint64](#uint64) |  |  |






<a name="cosmos.distribution.v1beta1.ValidatorHistoricalRewards"></a>

### ValidatorHistoricalRewards
ValidatorHistoricalRewards represents historical rewards for a validator.
Height is implicit within the store key.
Cumulative reward ratio is the sum from the zeroeth period
until this period of rewards / tokens, per the spec.
The reference count indicates the number of objects
which might need to reference this historical entry at any point.
ReferenceCount =
   number of outstanding delegations which ended the associated period (and
   might need to read that record)
 + number of slashes which ended the associated period (and might need to
 read that record)
 + one per validator for the zeroeth period, set on initialization


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cumulative_reward_ratio` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |
| `reference_count` | [uint32](#uint32) |  |  |






<a name="cosmos.distribution.v1beta1.ValidatorOutstandingRewards"></a>

### ValidatorOutstandingRewards
ValidatorOutstandingRewards represents outstanding (un-withdrawn) rewards
for a validator inexpensive to track, allows simple sanity checks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated |  |






<a name="cosmos.distribution.v1beta1.ValidatorSlashEvent"></a>

### ValidatorSlashEvent
ValidatorSlashEvent represents a validator slash event.
Height is implicit within the store key.
This is needed to calculate appropriate amount of staking tokens
for delegations which are withdrawn after a slash has occurred.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_period` | [uint64](#uint64) |  |  |
| `fraction` | [string](#string) |  |  |






<a name="cosmos.distribution.v1beta1.ValidatorSlashEvents"></a>

### ValidatorSlashEvents
ValidatorSlashEvents is a collection of ValidatorSlashEvent messages.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_slash_events` | [ValidatorSlashEvent](#cosmos.distribution.v1beta1.ValidatorSlashEvent) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/distribution/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/distribution/v1beta1/genesis.proto



<a name="cosmos.distribution.v1beta1.DelegatorStartingInfoRecord"></a>

### DelegatorStartingInfoRecord
DelegatorStartingInfoRecord used for import / export via genesis json.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the address of the delegator. |
| `validator_address` | [string](#string) |  | validator_address is the address of the validator. |
| `starting_info` | [DelegatorStartingInfo](#cosmos.distribution.v1beta1.DelegatorStartingInfo) |  | starting_info defines the starting info of a delegator. |






<a name="cosmos.distribution.v1beta1.DelegatorWithdrawInfo"></a>

### DelegatorWithdrawInfo
DelegatorWithdrawInfo is the address for where distributions rewards are
withdrawn to by default this struct is only used at genesis to feed in
default withdraw addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the address of the delegator. |
| `withdraw_address` | [string](#string) |  | withdraw_address is the address to withdraw the delegation rewards to. |






<a name="cosmos.distribution.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the distribution module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.distribution.v1beta1.Params) |  | params defines all the paramaters of the module. |
| `fee_pool` | [FeePool](#cosmos.distribution.v1beta1.FeePool) |  | fee_pool defines the fee pool at genesis. |
| `delegator_withdraw_infos` | [DelegatorWithdrawInfo](#cosmos.distribution.v1beta1.DelegatorWithdrawInfo) | repeated | fee_pool defines the delegator withdraw infos at genesis. |
| `previous_proposer` | [string](#string) |  | fee_pool defines the previous proposer at genesis. |
| `outstanding_rewards` | [ValidatorOutstandingRewardsRecord](#cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord) | repeated | fee_pool defines the outstanding rewards of all validators at genesis. |
| `validator_accumulated_commissions` | [ValidatorAccumulatedCommissionRecord](#cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord) | repeated | fee_pool defines the accumulated commisions of all validators at genesis. |
| `validator_historical_rewards` | [ValidatorHistoricalRewardsRecord](#cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord) | repeated | fee_pool defines the historical rewards of all validators at genesis. |
| `validator_current_rewards` | [ValidatorCurrentRewardsRecord](#cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord) | repeated | fee_pool defines the current rewards of all validators at genesis. |
| `delegator_starting_infos` | [DelegatorStartingInfoRecord](#cosmos.distribution.v1beta1.DelegatorStartingInfoRecord) | repeated | fee_pool defines the delegator starting infos at genesis. |
| `validator_slash_events` | [ValidatorSlashEventRecord](#cosmos.distribution.v1beta1.ValidatorSlashEventRecord) | repeated | fee_pool defines the validator slash events at genesis. |






<a name="cosmos.distribution.v1beta1.ValidatorAccumulatedCommissionRecord"></a>

### ValidatorAccumulatedCommissionRecord
ValidatorAccumulatedCommissionRecord is used for import / export via genesis
json.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address is the address of the validator. |
| `accumulated` | [ValidatorAccumulatedCommission](#cosmos.distribution.v1beta1.ValidatorAccumulatedCommission) |  | accumulated is the accumulated commission of a validator. |






<a name="cosmos.distribution.v1beta1.ValidatorCurrentRewardsRecord"></a>

### ValidatorCurrentRewardsRecord
ValidatorCurrentRewardsRecord is used for import / export via genesis json.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address is the address of the validator. |
| `rewards` | [ValidatorCurrentRewards](#cosmos.distribution.v1beta1.ValidatorCurrentRewards) |  | rewards defines the current rewards of a validator. |






<a name="cosmos.distribution.v1beta1.ValidatorHistoricalRewardsRecord"></a>

### ValidatorHistoricalRewardsRecord
ValidatorHistoricalRewardsRecord is used for import / export via genesis
json.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address is the address of the validator. |
| `period` | [uint64](#uint64) |  | period defines the period the historical rewards apply to. |
| `rewards` | [ValidatorHistoricalRewards](#cosmos.distribution.v1beta1.ValidatorHistoricalRewards) |  | rewards defines the historical rewards of a validator. |






<a name="cosmos.distribution.v1beta1.ValidatorOutstandingRewardsRecord"></a>

### ValidatorOutstandingRewardsRecord
ValidatorOutstandingRewardsRecord is used for import/export via genesis json.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address is the address of the validator. |
| `outstanding_rewards` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | outstanding_rewards represents the oustanding rewards of a validator. |






<a name="cosmos.distribution.v1beta1.ValidatorSlashEventRecord"></a>

### ValidatorSlashEventRecord
ValidatorSlashEventRecord is used for import / export via genesis json.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address is the address of the validator. |
| `height` | [uint64](#uint64) |  | height defines the block height at which the slash event occured. |
| `period` | [uint64](#uint64) |  | period is the period of the slash event. |
| `validator_slash_event` | [ValidatorSlashEvent](#cosmos.distribution.v1beta1.ValidatorSlashEvent) |  | validator_slash_event describes the slash event. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/distribution/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/distribution/v1beta1/query.proto



<a name="cosmos.distribution.v1beta1.QueryCommunityPoolRequest"></a>

### QueryCommunityPoolRequest
QueryCommunityPoolRequest is the request type for the Query/CommunityPool RPC
method.






<a name="cosmos.distribution.v1beta1.QueryCommunityPoolResponse"></a>

### QueryCommunityPoolResponse
QueryCommunityPoolResponse is the response type for the Query/CommunityPool
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | pool defines community pool's coins. |






<a name="cosmos.distribution.v1beta1.QueryDelegationRewardsRequest"></a>

### QueryDelegationRewardsRequest
QueryDelegationRewardsRequest is the request type for the
Query/DelegationRewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address defines the delegator address to query for. |
| `validator_address` | [string](#string) |  | validator_address defines the validator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryDelegationRewardsResponse"></a>

### QueryDelegationRewardsResponse
QueryDelegationRewardsResponse is the response type for the
Query/DelegationRewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | rewards defines the rewards accrued by a delegation. |






<a name="cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest"></a>

### QueryDelegationTotalRewardsRequest
QueryDelegationTotalRewardsRequest is the request type for the
Query/DelegationTotalRewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address defines the delegator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse"></a>

### QueryDelegationTotalRewardsResponse
QueryDelegationTotalRewardsResponse is the response type for the
Query/DelegationTotalRewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [DelegationDelegatorReward](#cosmos.distribution.v1beta1.DelegationDelegatorReward) | repeated | rewards defines all the rewards accrued by a delegator. |
| `total` | [cosmos.base.v1beta1.DecCoin](#cosmos.base.v1beta1.DecCoin) | repeated | total defines the sum of all the rewards. |






<a name="cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest"></a>

### QueryDelegatorValidatorsRequest
QueryDelegatorValidatorsRequest is the request type for the
Query/DelegatorValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address defines the delegator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse"></a>

### QueryDelegatorValidatorsResponse
QueryDelegatorValidatorsResponse is the response type for the
Query/DelegatorValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [string](#string) | repeated | validators defines the validators a delegator is delegating for. |






<a name="cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest"></a>

### QueryDelegatorWithdrawAddressRequest
QueryDelegatorWithdrawAddressRequest is the request type for the
Query/DelegatorWithdrawAddress RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address defines the delegator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse"></a>

### QueryDelegatorWithdrawAddressResponse
QueryDelegatorWithdrawAddressResponse is the response type for the
Query/DelegatorWithdrawAddress RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `withdraw_address` | [string](#string) |  | withdraw_address defines the delegator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="cosmos.distribution.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.distribution.v1beta1.Params) |  | params defines the parameters of the module. |






<a name="cosmos.distribution.v1beta1.QueryValidatorCommissionRequest"></a>

### QueryValidatorCommissionRequest
QueryValidatorCommissionRequest is the request type for the
Query/ValidatorCommission RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address defines the validator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryValidatorCommissionResponse"></a>

### QueryValidatorCommissionResponse
QueryValidatorCommissionResponse is the response type for the
Query/ValidatorCommission RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commission` | [ValidatorAccumulatedCommission](#cosmos.distribution.v1beta1.ValidatorAccumulatedCommission) |  | commission defines the commision the validator received. |






<a name="cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest"></a>

### QueryValidatorOutstandingRewardsRequest
QueryValidatorOutstandingRewardsRequest is the request type for the
Query/ValidatorOutstandingRewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address defines the validator address to query for. |






<a name="cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse"></a>

### QueryValidatorOutstandingRewardsResponse
QueryValidatorOutstandingRewardsResponse is the response type for the
Query/ValidatorOutstandingRewards RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rewards` | [ValidatorOutstandingRewards](#cosmos.distribution.v1beta1.ValidatorOutstandingRewards) |  |  |






<a name="cosmos.distribution.v1beta1.QueryValidatorSlashesRequest"></a>

### QueryValidatorSlashesRequest
QueryValidatorSlashesRequest is the request type for the
Query/ValidatorSlashes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  | validator_address defines the validator address to query for. |
| `starting_height` | [uint64](#uint64) |  | starting_height defines the optional starting height to query the slashes. |
| `ending_height` | [uint64](#uint64) |  | starting_height defines the optional ending height to query the slashes. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.distribution.v1beta1.QueryValidatorSlashesResponse"></a>

### QueryValidatorSlashesResponse
QueryValidatorSlashesResponse is the response type for the
Query/ValidatorSlashes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `slashes` | [ValidatorSlashEvent](#cosmos.distribution.v1beta1.ValidatorSlashEvent) | repeated | slashes defines the slashes the validator received. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.distribution.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for distribution module.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#cosmos.distribution.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.distribution.v1beta1.QueryParamsResponse) | Params queries params of the distribution module. | GET|/cosmos/distribution/v1beta1/params|
| `ValidatorOutstandingRewards` | [QueryValidatorOutstandingRewardsRequest](#cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsRequest) | [QueryValidatorOutstandingRewardsResponse](#cosmos.distribution.v1beta1.QueryValidatorOutstandingRewardsResponse) | ValidatorOutstandingRewards queries rewards of a validator address. | GET|/cosmos/distribution/v1beta1/validators/{validator_address}/outstanding_rewards|
| `ValidatorCommission` | [QueryValidatorCommissionRequest](#cosmos.distribution.v1beta1.QueryValidatorCommissionRequest) | [QueryValidatorCommissionResponse](#cosmos.distribution.v1beta1.QueryValidatorCommissionResponse) | ValidatorCommission queries accumulated commission for a validator. | GET|/cosmos/distribution/v1beta1/validators/{validator_address}/commission|
| `ValidatorSlashes` | [QueryValidatorSlashesRequest](#cosmos.distribution.v1beta1.QueryValidatorSlashesRequest) | [QueryValidatorSlashesResponse](#cosmos.distribution.v1beta1.QueryValidatorSlashesResponse) | ValidatorSlashes queries slash events of a validator. | GET|/cosmos/distribution/v1beta1/validators/{validator_address}/slashes|
| `DelegationRewards` | [QueryDelegationRewardsRequest](#cosmos.distribution.v1beta1.QueryDelegationRewardsRequest) | [QueryDelegationRewardsResponse](#cosmos.distribution.v1beta1.QueryDelegationRewardsResponse) | DelegationRewards queries the total rewards accrued by a delegation. | GET|/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards/{validator_address}|
| `DelegationTotalRewards` | [QueryDelegationTotalRewardsRequest](#cosmos.distribution.v1beta1.QueryDelegationTotalRewardsRequest) | [QueryDelegationTotalRewardsResponse](#cosmos.distribution.v1beta1.QueryDelegationTotalRewardsResponse) | DelegationTotalRewards queries the total rewards accrued by a each validator. | GET|/cosmos/distribution/v1beta1/delegators/{delegator_address}/rewards|
| `DelegatorValidators` | [QueryDelegatorValidatorsRequest](#cosmos.distribution.v1beta1.QueryDelegatorValidatorsRequest) | [QueryDelegatorValidatorsResponse](#cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse) | DelegatorValidators queries the validators of a delegator. | GET|/cosmos/distribution/v1beta1/delegators/{delegator_address}/validators|
| `DelegatorWithdrawAddress` | [QueryDelegatorWithdrawAddressRequest](#cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressRequest) | [QueryDelegatorWithdrawAddressResponse](#cosmos.distribution.v1beta1.QueryDelegatorWithdrawAddressResponse) | DelegatorWithdrawAddress queries withdraw address of a delegator. | GET|/cosmos/distribution/v1beta1/delegators/{delegator_address}/withdraw_address|
| `CommunityPool` | [QueryCommunityPoolRequest](#cosmos.distribution.v1beta1.QueryCommunityPoolRequest) | [QueryCommunityPoolResponse](#cosmos.distribution.v1beta1.QueryCommunityPoolResponse) | CommunityPool queries the community pool coins. | GET|/cosmos/distribution/v1beta1/community_pool|

 <!-- end services -->



<a name="cosmos/distribution/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/distribution/v1beta1/tx.proto



<a name="cosmos.distribution.v1beta1.MsgFundCommunityPool"></a>

### MsgFundCommunityPool
MsgFundCommunityPool allows an account to directly
fund the community pool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `depositor` | [string](#string) |  |  |






<a name="cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse"></a>

### MsgFundCommunityPoolResponse
MsgFundCommunityPoolResponse defines the Msg/FundCommunityPool response type.






<a name="cosmos.distribution.v1beta1.MsgSetWithdrawAddress"></a>

### MsgSetWithdrawAddress
MsgSetWithdrawAddress sets the withdraw address for
a delegator (or validator self-delegation).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `withdraw_address` | [string](#string) |  |  |






<a name="cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse"></a>

### MsgSetWithdrawAddressResponse
MsgSetWithdrawAddressResponse defines the Msg/SetWithdrawAddress response type.






<a name="cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"></a>

### MsgWithdrawDelegatorReward
MsgWithdrawDelegatorReward represents delegation withdrawal to a delegator
from a single validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |






<a name="cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse"></a>

### MsgWithdrawDelegatorRewardResponse
MsgWithdrawDelegatorRewardResponse defines the Msg/WithdrawDelegatorReward response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Since: cosmos-sdk 0.46 |






<a name="cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission"></a>

### MsgWithdrawValidatorCommission
MsgWithdrawValidatorCommission withdraws the full commission to the validator
address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  |  |






<a name="cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse"></a>

### MsgWithdrawValidatorCommissionResponse
MsgWithdrawValidatorCommissionResponse defines the Msg/WithdrawValidatorCommission response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Since: cosmos-sdk 0.46 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.distribution.v1beta1.Msg"></a>

### Msg
Msg defines the distribution Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetWithdrawAddress` | [MsgSetWithdrawAddress](#cosmos.distribution.v1beta1.MsgSetWithdrawAddress) | [MsgSetWithdrawAddressResponse](#cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse) | SetWithdrawAddress defines a method to change the withdraw address for a delegator (or validator self-delegation). | |
| `WithdrawDelegatorReward` | [MsgWithdrawDelegatorReward](#cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward) | [MsgWithdrawDelegatorRewardResponse](#cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse) | WithdrawDelegatorReward defines a method to withdraw rewards of delegator from a single validator. | |
| `WithdrawValidatorCommission` | [MsgWithdrawValidatorCommission](#cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission) | [MsgWithdrawValidatorCommissionResponse](#cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse) | WithdrawValidatorCommission defines a method to withdraw the full commission to the validator address. | |
| `FundCommunityPool` | [MsgFundCommunityPool](#cosmos.distribution.v1beta1.MsgFundCommunityPool) | [MsgFundCommunityPoolResponse](#cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse) | FundCommunityPool defines a method to allow an account to directly fund the community pool. | |

 <!-- end services -->



<a name="cosmos/evidence/v1beta1/evidence.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/evidence/v1beta1/evidence.proto



<a name="cosmos.evidence.v1beta1.Equivocation"></a>

### Equivocation
Equivocation implements the Evidence interface and defines evidence of double
signing misbehavior.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  |  |
| `time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `power` | [int64](#int64) |  |  |
| `consensus_address` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/evidence/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/evidence/v1beta1/genesis.proto



<a name="cosmos.evidence.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the evidence module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evidence` | [google.protobuf.Any](#google.protobuf.Any) | repeated | evidence defines all the evidence at genesis. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/evidence/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/evidence/v1beta1/query.proto



<a name="cosmos.evidence.v1beta1.QueryAllEvidenceRequest"></a>

### QueryAllEvidenceRequest
QueryEvidenceRequest is the request type for the Query/AllEvidence RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.evidence.v1beta1.QueryAllEvidenceResponse"></a>

### QueryAllEvidenceResponse
QueryAllEvidenceResponse is the response type for the Query/AllEvidence RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evidence` | [google.protobuf.Any](#google.protobuf.Any) | repeated | evidence returns all evidences. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.evidence.v1beta1.QueryEvidenceRequest"></a>

### QueryEvidenceRequest
QueryEvidenceRequest is the request type for the Query/Evidence RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evidence_hash` | [bytes](#bytes) |  | evidence_hash defines the hash of the requested evidence. |






<a name="cosmos.evidence.v1beta1.QueryEvidenceResponse"></a>

### QueryEvidenceResponse
QueryEvidenceResponse is the response type for the Query/Evidence RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `evidence` | [google.protobuf.Any](#google.protobuf.Any) |  | evidence returns the requested evidence. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.evidence.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Evidence` | [QueryEvidenceRequest](#cosmos.evidence.v1beta1.QueryEvidenceRequest) | [QueryEvidenceResponse](#cosmos.evidence.v1beta1.QueryEvidenceResponse) | Evidence queries evidence based on evidence hash. | GET|/cosmos/evidence/v1beta1/evidence/{evidence_hash}|
| `AllEvidence` | [QueryAllEvidenceRequest](#cosmos.evidence.v1beta1.QueryAllEvidenceRequest) | [QueryAllEvidenceResponse](#cosmos.evidence.v1beta1.QueryAllEvidenceResponse) | AllEvidence queries all evidence. | GET|/cosmos/evidence/v1beta1/evidence|

 <!-- end services -->



<a name="cosmos/evidence/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/evidence/v1beta1/tx.proto



<a name="cosmos.evidence.v1beta1.MsgSubmitEvidence"></a>

### MsgSubmitEvidence
MsgSubmitEvidence represents a message that supports submitting arbitrary
Evidence of misbehavior such as equivocation or counterfactual signing.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `submitter` | [string](#string) |  |  |
| `evidence` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse"></a>

### MsgSubmitEvidenceResponse
MsgSubmitEvidenceResponse defines the Msg/SubmitEvidence response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [bytes](#bytes) |  | hash defines the hash of the evidence. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.evidence.v1beta1.Msg"></a>

### Msg
Msg defines the evidence Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SubmitEvidence` | [MsgSubmitEvidence](#cosmos.evidence.v1beta1.MsgSubmitEvidence) | [MsgSubmitEvidenceResponse](#cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse) | SubmitEvidence submits an arbitrary Evidence of misbehavior such as equivocation or counterfactual signing. | |

 <!-- end services -->



<a name="cosmos/feegrant/v1beta1/feegrant.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/feegrant/v1beta1/feegrant.proto
Since: cosmos-sdk 0.43


<a name="cosmos.feegrant.v1beta1.AllowedMsgAllowance"></a>

### AllowedMsgAllowance
AllowedMsgAllowance creates allowance only for specified message types.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowance` | [google.protobuf.Any](#google.protobuf.Any) |  | allowance can be any of basic and periodic fee allowance. |
| `allowed_messages` | [string](#string) | repeated | allowed_messages are the messages for which the grantee has the access. |






<a name="cosmos.feegrant.v1beta1.BasicAllowance"></a>

### BasicAllowance
BasicAllowance implements Allowance with a one-time grant of coins
that optionally expires. The grantee can use up to SpendLimit to cover fees.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `spend_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | spend_limit specifies the maximum amount of coins that can be spent by this allowance and will be updated as coins are spent. If it is empty, there is no spend limit and any amount of coins can be spent. |
| `expiration` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | expiration specifies an optional time when this allowance expires |






<a name="cosmos.feegrant.v1beta1.Grant"></a>

### Grant
Grant is stored in the KVStore to record a grant with full context


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  | granter is the address of the user granting an allowance of their funds. |
| `grantee` | [string](#string) |  | grantee is the address of the user being granted an allowance of another user's funds. |
| `allowance` | [google.protobuf.Any](#google.protobuf.Any) |  | allowance can be any of basic, periodic, allowed fee allowance. |






<a name="cosmos.feegrant.v1beta1.PeriodicAllowance"></a>

### PeriodicAllowance
PeriodicAllowance extends Allowance to allow for both a maximum cap,
as well as a limit per time period.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `basic` | [BasicAllowance](#cosmos.feegrant.v1beta1.BasicAllowance) |  | basic specifies a struct of `BasicAllowance` |
| `period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | period specifies the time duration in which period_spend_limit coins can be spent before that allowance is reset |
| `period_spend_limit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | period_spend_limit specifies the maximum number of coins that can be spent in the period |
| `period_can_spend` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | period_can_spend is the number of coins left to be spent before the period_reset time |
| `period_reset` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | period_reset is the time at which this period resets and a new one begins, it is calculated from the start time of the first transaction after the last period ended |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/feegrant/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/feegrant/v1beta1/genesis.proto
Since: cosmos-sdk 0.43


<a name="cosmos.feegrant.v1beta1.GenesisState"></a>

### GenesisState
GenesisState contains a set of fee allowances, persisted from the store


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowances` | [Grant](#cosmos.feegrant.v1beta1.Grant) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/feegrant/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/feegrant/v1beta1/query.proto
Since: cosmos-sdk 0.43


<a name="cosmos.feegrant.v1beta1.QueryAllowanceRequest"></a>

### QueryAllowanceRequest
QueryAllowanceRequest is the request type for the Query/Allowance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  | granter is the address of the user granting an allowance of their funds. |
| `grantee` | [string](#string) |  | grantee is the address of the user being granted an allowance of another user's funds. |






<a name="cosmos.feegrant.v1beta1.QueryAllowanceResponse"></a>

### QueryAllowanceResponse
QueryAllowanceResponse is the response type for the Query/Allowance RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowance` | [Grant](#cosmos.feegrant.v1beta1.Grant) |  | allowance is a allowance granted for grantee by granter. |






<a name="cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest"></a>

### QueryAllowancesByGranterRequest
QueryAllowancesByGranterRequest is the request type for the Query/AllowancesByGranter RPC method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse"></a>

### QueryAllowancesByGranterResponse
QueryAllowancesByGranterResponse is the response type for the Query/AllowancesByGranter RPC method.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowances` | [Grant](#cosmos.feegrant.v1beta1.Grant) | repeated | allowances that have been issued by the granter. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |






<a name="cosmos.feegrant.v1beta1.QueryAllowancesRequest"></a>

### QueryAllowancesRequest
QueryAllowancesRequest is the request type for the Query/Allowances RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `grantee` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an pagination for the request. |






<a name="cosmos.feegrant.v1beta1.QueryAllowancesResponse"></a>

### QueryAllowancesResponse
QueryAllowancesResponse is the response type for the Query/Allowances RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `allowances` | [Grant](#cosmos.feegrant.v1beta1.Grant) | repeated | allowances are allowance's granted for grantee by granter. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines an pagination for the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.feegrant.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Allowance` | [QueryAllowanceRequest](#cosmos.feegrant.v1beta1.QueryAllowanceRequest) | [QueryAllowanceResponse](#cosmos.feegrant.v1beta1.QueryAllowanceResponse) | Allowance returns fee granted to the grantee by the granter. | GET|/cosmos/feegrant/v1beta1/allowance/{granter}/{grantee}|
| `Allowances` | [QueryAllowancesRequest](#cosmos.feegrant.v1beta1.QueryAllowancesRequest) | [QueryAllowancesResponse](#cosmos.feegrant.v1beta1.QueryAllowancesResponse) | Allowances returns all the grants for address. | GET|/cosmos/feegrant/v1beta1/allowances/{grantee}|
| `AllowancesByGranter` | [QueryAllowancesByGranterRequest](#cosmos.feegrant.v1beta1.QueryAllowancesByGranterRequest) | [QueryAllowancesByGranterResponse](#cosmos.feegrant.v1beta1.QueryAllowancesByGranterResponse) | AllowancesByGranter returns all the grants given by an address

Since: cosmos-sdk 0.46 | GET|/cosmos/feegrant/v1beta1/issued/{granter}|

 <!-- end services -->



<a name="cosmos/feegrant/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/feegrant/v1beta1/tx.proto
Since: cosmos-sdk 0.43


<a name="cosmos.feegrant.v1beta1.MsgGrantAllowance"></a>

### MsgGrantAllowance
MsgGrantAllowance adds permission for Grantee to spend up to Allowance
of fees from the account of Granter.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  | granter is the address of the user granting an allowance of their funds. |
| `grantee` | [string](#string) |  | grantee is the address of the user being granted an allowance of another user's funds. |
| `allowance` | [google.protobuf.Any](#google.protobuf.Any) |  | allowance can be any of basic, periodic, allowed fee allowance. |






<a name="cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse"></a>

### MsgGrantAllowanceResponse
MsgGrantAllowanceResponse defines the Msg/GrantAllowanceResponse response type.






<a name="cosmos.feegrant.v1beta1.MsgRevokeAllowance"></a>

### MsgRevokeAllowance
MsgRevokeAllowance removes any existing Allowance from Granter to Grantee.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `granter` | [string](#string) |  | granter is the address of the user granting an allowance of their funds. |
| `grantee` | [string](#string) |  | grantee is the address of the user being granted an allowance of another user's funds. |






<a name="cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse"></a>

### MsgRevokeAllowanceResponse
MsgRevokeAllowanceResponse defines the Msg/RevokeAllowanceResponse response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.feegrant.v1beta1.Msg"></a>

### Msg
Msg defines the feegrant msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GrantAllowance` | [MsgGrantAllowance](#cosmos.feegrant.v1beta1.MsgGrantAllowance) | [MsgGrantAllowanceResponse](#cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse) | GrantAllowance grants fee allowance to the grantee on the granter's account with the provided expiration time. | |
| `RevokeAllowance` | [MsgRevokeAllowance](#cosmos.feegrant.v1beta1.MsgRevokeAllowance) | [MsgRevokeAllowanceResponse](#cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse) | RevokeAllowance revokes any fee allowance of granter's account that has been granted to the grantee. | |

 <!-- end services -->



<a name="cosmos/genutil/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/genutil/v1beta1/genesis.proto



<a name="cosmos.genutil.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the raw genesis transaction in JSON.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gen_txs` | [bytes](#bytes) | repeated | gen_txs defines the genesis transactions. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1/gov.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1/gov.proto
Since: cosmos-sdk 0.46


<a name="cosmos.gov.v1.Deposit"></a>

### Deposit
Deposit defines an amount deposited by an account address to an active
proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.gov.v1.DepositParams"></a>

### DepositParams
DepositParams defines the params for deposits on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Minimum deposit for a proposal to enter voting period. |
| `max_deposit_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months. |






<a name="cosmos.gov.v1.Proposal"></a>

### Proposal
Proposal defines the core field members of a governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `status` | [ProposalStatus](#cosmos.gov.v1.ProposalStatus) |  |  |
| `final_tally_result` | [TallyResult](#cosmos.gov.v1.TallyResult) |  | final_tally_result is the final tally result of the proposal. When querying a proposal via gRPC, this field is not populated until the proposal's voting period has ended. |
| `submit_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `deposit_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `total_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `voting_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `voting_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata attached to the proposal. |






<a name="cosmos.gov.v1.TallyParams"></a>

### TallyParams
TallyParams defines the params for tallying votes on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `quorum` | [string](#string) |  | Minimum percentage of total stake needed to vote for a result to be considered valid. |
| `threshold` | [string](#string) |  | Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. |
| `veto_threshold` | [string](#string) |  | Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Default value: 1/3. |






<a name="cosmos.gov.v1.TallyResult"></a>

### TallyResult
TallyResult defines a standard tally for a governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `yes_count` | [string](#string) |  |  |
| `abstain_count` | [string](#string) |  |  |
| `no_count` | [string](#string) |  |  |
| `no_with_veto_count` | [string](#string) |  |  |






<a name="cosmos.gov.v1.Vote"></a>

### Vote
Vote defines a vote on a governance proposal.
A Vote consists of a proposal ID, the voter, and the vote option.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `options` | [WeightedVoteOption](#cosmos.gov.v1.WeightedVoteOption) | repeated |  |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the vote. |






<a name="cosmos.gov.v1.VotingParams"></a>

### VotingParams
VotingParams defines the params for voting on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Length of the voting period. |






<a name="cosmos.gov.v1.WeightedVoteOption"></a>

### WeightedVoteOption
WeightedVoteOption defines a unit of vote for vote split.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `option` | [VoteOption](#cosmos.gov.v1.VoteOption) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->


<a name="cosmos.gov.v1.ProposalStatus"></a>

### ProposalStatus
ProposalStatus enumerates the valid statuses of a proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROPOSAL_STATUS_UNSPECIFIED | 0 | PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status. |
| PROPOSAL_STATUS_DEPOSIT_PERIOD | 1 | PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit period. |
| PROPOSAL_STATUS_VOTING_PERIOD | 2 | PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting period. |
| PROPOSAL_STATUS_PASSED | 3 | PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has passed. |
| PROPOSAL_STATUS_REJECTED | 4 | PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has been rejected. |
| PROPOSAL_STATUS_FAILED | 5 | PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has failed. |



<a name="cosmos.gov.v1.VoteOption"></a>

### VoteOption
VoteOption enumerates the valid vote options for a given governance proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOTE_OPTION_UNSPECIFIED | 0 | VOTE_OPTION_UNSPECIFIED defines a no-op vote option. |
| VOTE_OPTION_YES | 1 | VOTE_OPTION_YES defines a yes vote option. |
| VOTE_OPTION_ABSTAIN | 2 | VOTE_OPTION_ABSTAIN defines an abstain vote option. |
| VOTE_OPTION_NO | 3 | VOTE_OPTION_NO defines a no vote option. |
| VOTE_OPTION_NO_WITH_VETO | 4 | VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1/genesis.proto
Since: cosmos-sdk 0.46


<a name="cosmos.gov.v1.GenesisState"></a>

### GenesisState
GenesisState defines the gov module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `starting_proposal_id` | [uint64](#uint64) |  | starting_proposal_id is the ID of the starting proposal. |
| `deposits` | [Deposit](#cosmos.gov.v1.Deposit) | repeated | deposits defines all the deposits present at genesis. |
| `votes` | [Vote](#cosmos.gov.v1.Vote) | repeated | votes defines all the votes present at genesis. |
| `proposals` | [Proposal](#cosmos.gov.v1.Proposal) | repeated | proposals defines all the proposals present at genesis. |
| `deposit_params` | [DepositParams](#cosmos.gov.v1.DepositParams) |  | params defines all the paramaters of related to deposit. |
| `voting_params` | [VotingParams](#cosmos.gov.v1.VotingParams) |  | params defines all the paramaters of related to voting. |
| `tally_params` | [TallyParams](#cosmos.gov.v1.TallyParams) |  | params defines all the paramaters of related to tally. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1/query.proto
Since: cosmos-sdk 0.46


<a name="cosmos.gov.v1.QueryDepositRequest"></a>

### QueryDepositRequest
QueryDepositRequest is the request type for the Query/Deposit RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `depositor` | [string](#string) |  | depositor defines the deposit addresses from the proposals. |






<a name="cosmos.gov.v1.QueryDepositResponse"></a>

### QueryDepositResponse
QueryDepositResponse is the response type for the Query/Deposit RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit` | [Deposit](#cosmos.gov.v1.Deposit) |  | deposit defines the requested deposit. |






<a name="cosmos.gov.v1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#cosmos.gov.v1.Deposit) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.gov.v1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params_type` | [string](#string) |  | params_type defines which parameters to query for, can be one of "voting", "tallying" or "deposit". |






<a name="cosmos.gov.v1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_params` | [VotingParams](#cosmos.gov.v1.VotingParams) |  | voting_params defines the parameters related to voting. |
| `deposit_params` | [DepositParams](#cosmos.gov.v1.DepositParams) |  | deposit_params defines the parameters related to deposit. |
| `tally_params` | [TallyParams](#cosmos.gov.v1.TallyParams) |  | tally_params defines the parameters related to tally. |






<a name="cosmos.gov.v1.QueryProposalRequest"></a>

### QueryProposalRequest
QueryProposalRequest is the request type for the Query/Proposal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |






<a name="cosmos.gov.v1.QueryProposalResponse"></a>

### QueryProposalResponse
QueryProposalResponse is the response type for the Query/Proposal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal` | [Proposal](#cosmos.gov.v1.Proposal) |  |  |






<a name="cosmos.gov.v1.QueryProposalsRequest"></a>

### QueryProposalsRequest
QueryProposalsRequest is the request type for the Query/Proposals RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_status` | [ProposalStatus](#cosmos.gov.v1.ProposalStatus) |  | proposal_status defines the status of the proposals. |
| `voter` | [string](#string) |  | voter defines the voter address for the proposals. |
| `depositor` | [string](#string) |  | depositor defines the deposit addresses from the proposals. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1.QueryProposalsResponse"></a>

### QueryProposalsResponse
QueryProposalsResponse is the response type for the Query/Proposals RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposals` | [Proposal](#cosmos.gov.v1.Proposal) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.gov.v1.QueryTallyResultRequest"></a>

### QueryTallyResultRequest
QueryTallyResultRequest is the request type for the Query/Tally RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |






<a name="cosmos.gov.v1.QueryTallyResultResponse"></a>

### QueryTallyResultResponse
QueryTallyResultResponse is the response type for the Query/Tally RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tally` | [TallyResult](#cosmos.gov.v1.TallyResult) |  | tally defines the requested tally. |






<a name="cosmos.gov.v1.QueryVoteRequest"></a>

### QueryVoteRequest
QueryVoteRequest is the request type for the Query/Vote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `voter` | [string](#string) |  | voter defines the voter address for the proposals. |






<a name="cosmos.gov.v1.QueryVoteResponse"></a>

### QueryVoteResponse
QueryVoteResponse is the response type for the Query/Vote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote` | [Vote](#cosmos.gov.v1.Vote) |  | vote defined the queried vote. |






<a name="cosmos.gov.v1.QueryVotesRequest"></a>

### QueryVotesRequest
QueryVotesRequest is the request type for the Query/Votes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1.QueryVotesResponse"></a>

### QueryVotesResponse
QueryVotesResponse is the response type for the Query/Votes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [Vote](#cosmos.gov.v1.Vote) | repeated | votes defined the queried votes. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.gov.v1.Query"></a>

### Query
Query defines the gRPC querier service for gov module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Proposal` | [QueryProposalRequest](#cosmos.gov.v1.QueryProposalRequest) | [QueryProposalResponse](#cosmos.gov.v1.QueryProposalResponse) | Proposal queries proposal details based on ProposalID. | GET|/cosmos/gov/v1/proposals/{proposal_id}|
| `Proposals` | [QueryProposalsRequest](#cosmos.gov.v1.QueryProposalsRequest) | [QueryProposalsResponse](#cosmos.gov.v1.QueryProposalsResponse) | Proposals queries all proposals based on given status. | GET|/cosmos/gov/v1/proposals|
| `Vote` | [QueryVoteRequest](#cosmos.gov.v1.QueryVoteRequest) | [QueryVoteResponse](#cosmos.gov.v1.QueryVoteResponse) | Vote queries voted information based on proposalID, voterAddr. | GET|/cosmos/gov/v1/proposals/{proposal_id}/votes/{voter}|
| `Votes` | [QueryVotesRequest](#cosmos.gov.v1.QueryVotesRequest) | [QueryVotesResponse](#cosmos.gov.v1.QueryVotesResponse) | Votes queries votes of a given proposal. | GET|/cosmos/gov/v1/proposals/{proposal_id}/votes|
| `Params` | [QueryParamsRequest](#cosmos.gov.v1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.gov.v1.QueryParamsResponse) | Params queries all parameters of the gov module. | GET|/cosmos/gov/v1/params/{params_type}|
| `Deposit` | [QueryDepositRequest](#cosmos.gov.v1.QueryDepositRequest) | [QueryDepositResponse](#cosmos.gov.v1.QueryDepositResponse) | Deposit queries single deposit information based proposalID, depositAddr. | GET|/cosmos/gov/v1/proposals/{proposal_id}/deposits/{depositor}|
| `Deposits` | [QueryDepositsRequest](#cosmos.gov.v1.QueryDepositsRequest) | [QueryDepositsResponse](#cosmos.gov.v1.QueryDepositsResponse) | Deposits queries all deposits of a single proposal. | GET|/cosmos/gov/v1/proposals/{proposal_id}/deposits|
| `TallyResult` | [QueryTallyResultRequest](#cosmos.gov.v1.QueryTallyResultRequest) | [QueryTallyResultResponse](#cosmos.gov.v1.QueryTallyResultResponse) | TallyResult queries the tally of a proposal vote. | GET|/cosmos/gov/v1/proposals/{proposal_id}/tally|

 <!-- end services -->



<a name="cosmos/gov/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1/tx.proto
Since: cosmos-sdk 0.46


<a name="cosmos.gov.v1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines a message to submit a deposit to an existing proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.gov.v1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="cosmos.gov.v1.MsgExecLegacyContent"></a>

### MsgExecLegacyContent
MsgExecLegacyContent is used to wrap the legacy content field into a message.
This ensures backwards compatibility with v1beta1.MsgSubmitProposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  | content is the proposal's content. |
| `authority` | [string](#string) |  | authority must be the gov module address. |






<a name="cosmos.gov.v1.MsgExecLegacyContentResponse"></a>

### MsgExecLegacyContentResponse
MsgExecLegacyContentResponse defines the Msg/ExecLegacyContent response type.






<a name="cosmos.gov.v1.MsgSubmitProposal"></a>

### MsgSubmitProposal
MsgSubmitProposal defines an sdk.Msg type that supports submitting arbitrary
proposal Content.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated |  |
| `initial_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `proposer` | [string](#string) |  |  |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata attached to the proposal. |






<a name="cosmos.gov.v1.MsgSubmitProposalResponse"></a>

### MsgSubmitProposalResponse
MsgSubmitProposalResponse defines the Msg/SubmitProposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="cosmos.gov.v1.MsgVote"></a>

### MsgVote
MsgVote defines a message to cast a vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `option` | [VoteOption](#cosmos.gov.v1.VoteOption) |  |  |
| `metadata` | [string](#string) |  |  |






<a name="cosmos.gov.v1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse defines the Msg/Vote response type.






<a name="cosmos.gov.v1.MsgVoteWeighted"></a>

### MsgVoteWeighted
MsgVoteWeighted defines a message to cast a vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `options` | [WeightedVoteOption](#cosmos.gov.v1.WeightedVoteOption) | repeated |  |
| `metadata` | [string](#string) |  |  |






<a name="cosmos.gov.v1.MsgVoteWeightedResponse"></a>

### MsgVoteWeightedResponse
MsgVoteWeightedResponse defines the Msg/VoteWeighted response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.gov.v1.Msg"></a>

### Msg
Msg defines the gov Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SubmitProposal` | [MsgSubmitProposal](#cosmos.gov.v1.MsgSubmitProposal) | [MsgSubmitProposalResponse](#cosmos.gov.v1.MsgSubmitProposalResponse) | SubmitProposal defines a method to create new proposal given a content. | |
| `ExecLegacyContent` | [MsgExecLegacyContent](#cosmos.gov.v1.MsgExecLegacyContent) | [MsgExecLegacyContentResponse](#cosmos.gov.v1.MsgExecLegacyContentResponse) | ExecLegacyContent defines a Msg to be in included in a MsgSubmitProposal to execute a legacy content-based proposal. | |
| `Vote` | [MsgVote](#cosmos.gov.v1.MsgVote) | [MsgVoteResponse](#cosmos.gov.v1.MsgVoteResponse) | Vote defines a method to add a vote on a specific proposal. | |
| `VoteWeighted` | [MsgVoteWeighted](#cosmos.gov.v1.MsgVoteWeighted) | [MsgVoteWeightedResponse](#cosmos.gov.v1.MsgVoteWeightedResponse) | VoteWeighted defines a method to add a weighted vote on a specific proposal. | |
| `Deposit` | [MsgDeposit](#cosmos.gov.v1.MsgDeposit) | [MsgDepositResponse](#cosmos.gov.v1.MsgDepositResponse) | Deposit defines a method to add deposit on a specific proposal. | |

 <!-- end services -->



<a name="cosmos/gov/v1beta1/gov.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/gov.proto



<a name="cosmos.gov.v1beta1.Deposit"></a>

### Deposit
Deposit defines an amount deposited by an account address to an active
proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.gov.v1beta1.DepositParams"></a>

### DepositParams
DepositParams defines the params for deposits on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `min_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | Minimum deposit for a proposal to enter voting period. |
| `max_deposit_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Maximum period for Atom holders to deposit on a proposal. Initial value: 2 months. |






<a name="cosmos.gov.v1beta1.Proposal"></a>

### Proposal
Proposal defines the core field members of a governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `status` | [ProposalStatus](#cosmos.gov.v1beta1.ProposalStatus) |  |  |
| `final_tally_result` | [TallyResult](#cosmos.gov.v1beta1.TallyResult) |  | final_tally_result is the final tally result of the proposal. When querying a proposal via gRPC, this field is not populated until the proposal's voting period has ended. |
| `submit_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `deposit_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `total_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `voting_start_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |
| `voting_end_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="cosmos.gov.v1beta1.TallyParams"></a>

### TallyParams
TallyParams defines the params for tallying votes on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `quorum` | [bytes](#bytes) |  | Minimum percentage of total stake needed to vote for a result to be considered valid. |
| `threshold` | [bytes](#bytes) |  | Minimum proportion of Yes votes for proposal to pass. Default value: 0.5. |
| `veto_threshold` | [bytes](#bytes) |  | Minimum value of Veto votes to Total votes ratio for proposal to be vetoed. Default value: 1/3. |






<a name="cosmos.gov.v1beta1.TallyResult"></a>

### TallyResult
TallyResult defines a standard tally for a governance proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `yes` | [string](#string) |  |  |
| `abstain` | [string](#string) |  |  |
| `no` | [string](#string) |  |  |
| `no_with_veto` | [string](#string) |  |  |






<a name="cosmos.gov.v1beta1.TextProposal"></a>

### TextProposal
TextProposal defines a standard text proposal whose changes need to be
manually updated in case of approval.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |






<a name="cosmos.gov.v1beta1.Vote"></a>

### Vote
Vote defines a vote on a governance proposal.
A Vote consists of a proposal ID, the voter, and the vote option.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `option` | [VoteOption](#cosmos.gov.v1beta1.VoteOption) |  | **Deprecated.** Deprecated: Prefer to use `options` instead. This field is set in queries if and only if `len(options) == 1` and that option has weight 1. In all other cases, this field will default to VOTE_OPTION_UNSPECIFIED. |
| `options` | [WeightedVoteOption](#cosmos.gov.v1beta1.WeightedVoteOption) | repeated | Since: cosmos-sdk 0.43 |






<a name="cosmos.gov.v1beta1.VotingParams"></a>

### VotingParams
VotingParams defines the params for voting on governance proposals.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | Length of the voting period. |






<a name="cosmos.gov.v1beta1.WeightedVoteOption"></a>

### WeightedVoteOption
WeightedVoteOption defines a unit of vote for vote split.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `option` | [VoteOption](#cosmos.gov.v1beta1.VoteOption) |  |  |
| `weight` | [string](#string) |  |  |





 <!-- end messages -->


<a name="cosmos.gov.v1beta1.ProposalStatus"></a>

### ProposalStatus
ProposalStatus enumerates the valid statuses of a proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROPOSAL_STATUS_UNSPECIFIED | 0 | PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status. |
| PROPOSAL_STATUS_DEPOSIT_PERIOD | 1 | PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit period. |
| PROPOSAL_STATUS_VOTING_PERIOD | 2 | PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting period. |
| PROPOSAL_STATUS_PASSED | 3 | PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has passed. |
| PROPOSAL_STATUS_REJECTED | 4 | PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has been rejected. |
| PROPOSAL_STATUS_FAILED | 5 | PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has failed. |



<a name="cosmos.gov.v1beta1.VoteOption"></a>

### VoteOption
VoteOption enumerates the valid vote options for a given governance proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOTE_OPTION_UNSPECIFIED | 0 | VOTE_OPTION_UNSPECIFIED defines a no-op vote option. |
| VOTE_OPTION_YES | 1 | VOTE_OPTION_YES defines a yes vote option. |
| VOTE_OPTION_ABSTAIN | 2 | VOTE_OPTION_ABSTAIN defines an abstain vote option. |
| VOTE_OPTION_NO | 3 | VOTE_OPTION_NO defines a no vote option. |
| VOTE_OPTION_NO_WITH_VETO | 4 | VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/genesis.proto



<a name="cosmos.gov.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the gov module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `starting_proposal_id` | [uint64](#uint64) |  | starting_proposal_id is the ID of the starting proposal. |
| `deposits` | [Deposit](#cosmos.gov.v1beta1.Deposit) | repeated | deposits defines all the deposits present at genesis. |
| `votes` | [Vote](#cosmos.gov.v1beta1.Vote) | repeated | votes defines all the votes present at genesis. |
| `proposals` | [Proposal](#cosmos.gov.v1beta1.Proposal) | repeated | proposals defines all the proposals present at genesis. |
| `deposit_params` | [DepositParams](#cosmos.gov.v1beta1.DepositParams) |  | params defines all the paramaters of related to deposit. |
| `voting_params` | [VotingParams](#cosmos.gov.v1beta1.VotingParams) |  | params defines all the paramaters of related to voting. |
| `tally_params` | [TallyParams](#cosmos.gov.v1beta1.TallyParams) |  | params defines all the paramaters of related to tally. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/gov/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/query.proto



<a name="cosmos.gov.v1beta1.QueryDepositRequest"></a>

### QueryDepositRequest
QueryDepositRequest is the request type for the Query/Deposit RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `depositor` | [string](#string) |  | depositor defines the deposit addresses from the proposals. |






<a name="cosmos.gov.v1beta1.QueryDepositResponse"></a>

### QueryDepositResponse
QueryDepositResponse is the response type for the Query/Deposit RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposit` | [Deposit](#cosmos.gov.v1beta1.Deposit) |  | deposit defines the requested deposit. |






<a name="cosmos.gov.v1beta1.QueryDepositsRequest"></a>

### QueryDepositsRequest
QueryDepositsRequest is the request type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1beta1.QueryDepositsResponse"></a>

### QueryDepositsResponse
QueryDepositsResponse is the response type for the Query/Deposits RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `deposits` | [Deposit](#cosmos.gov.v1beta1.Deposit) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.gov.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params_type` | [string](#string) |  | params_type defines which parameters to query for, can be one of "voting", "tallying" or "deposit". |






<a name="cosmos.gov.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_params` | [VotingParams](#cosmos.gov.v1beta1.VotingParams) |  | voting_params defines the parameters related to voting. |
| `deposit_params` | [DepositParams](#cosmos.gov.v1beta1.DepositParams) |  | deposit_params defines the parameters related to deposit. |
| `tally_params` | [TallyParams](#cosmos.gov.v1beta1.TallyParams) |  | tally_params defines the parameters related to tally. |






<a name="cosmos.gov.v1beta1.QueryProposalRequest"></a>

### QueryProposalRequest
QueryProposalRequest is the request type for the Query/Proposal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |






<a name="cosmos.gov.v1beta1.QueryProposalResponse"></a>

### QueryProposalResponse
QueryProposalResponse is the response type for the Query/Proposal RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal` | [Proposal](#cosmos.gov.v1beta1.Proposal) |  |  |






<a name="cosmos.gov.v1beta1.QueryProposalsRequest"></a>

### QueryProposalsRequest
QueryProposalsRequest is the request type for the Query/Proposals RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_status` | [ProposalStatus](#cosmos.gov.v1beta1.ProposalStatus) |  | proposal_status defines the status of the proposals. |
| `voter` | [string](#string) |  | voter defines the voter address for the proposals. |
| `depositor` | [string](#string) |  | depositor defines the deposit addresses from the proposals. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1beta1.QueryProposalsResponse"></a>

### QueryProposalsResponse
QueryProposalsResponse is the response type for the Query/Proposals RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposals` | [Proposal](#cosmos.gov.v1beta1.Proposal) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.gov.v1beta1.QueryTallyResultRequest"></a>

### QueryTallyResultRequest
QueryTallyResultRequest is the request type for the Query/Tally RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |






<a name="cosmos.gov.v1beta1.QueryTallyResultResponse"></a>

### QueryTallyResultResponse
QueryTallyResultResponse is the response type for the Query/Tally RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tally` | [TallyResult](#cosmos.gov.v1beta1.TallyResult) |  | tally defines the requested tally. |






<a name="cosmos.gov.v1beta1.QueryVoteRequest"></a>

### QueryVoteRequest
QueryVoteRequest is the request type for the Query/Vote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `voter` | [string](#string) |  | voter defines the voter address for the proposals. |






<a name="cosmos.gov.v1beta1.QueryVoteResponse"></a>

### QueryVoteResponse
QueryVoteResponse is the response type for the Query/Vote RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote` | [Vote](#cosmos.gov.v1beta1.Vote) |  | vote defined the queried vote. |






<a name="cosmos.gov.v1beta1.QueryVotesRequest"></a>

### QueryVotesRequest
QueryVotesRequest is the request type for the Query/Votes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id defines the unique id of the proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.gov.v1beta1.QueryVotesResponse"></a>

### QueryVotesResponse
QueryVotesResponse is the response type for the Query/Votes RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [Vote](#cosmos.gov.v1beta1.Vote) | repeated | votes defined the queried votes. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.gov.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service for gov module

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Proposal` | [QueryProposalRequest](#cosmos.gov.v1beta1.QueryProposalRequest) | [QueryProposalResponse](#cosmos.gov.v1beta1.QueryProposalResponse) | Proposal queries proposal details based on ProposalID. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}|
| `Proposals` | [QueryProposalsRequest](#cosmos.gov.v1beta1.QueryProposalsRequest) | [QueryProposalsResponse](#cosmos.gov.v1beta1.QueryProposalsResponse) | Proposals queries all proposals based on given status. | GET|/cosmos/gov/v1beta1/proposals|
| `Vote` | [QueryVoteRequest](#cosmos.gov.v1beta1.QueryVoteRequest) | [QueryVoteResponse](#cosmos.gov.v1beta1.QueryVoteResponse) | Vote queries voted information based on proposalID, voterAddr. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/votes/{voter}|
| `Votes` | [QueryVotesRequest](#cosmos.gov.v1beta1.QueryVotesRequest) | [QueryVotesResponse](#cosmos.gov.v1beta1.QueryVotesResponse) | Votes queries votes of a given proposal. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/votes|
| `Params` | [QueryParamsRequest](#cosmos.gov.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.gov.v1beta1.QueryParamsResponse) | Params queries all parameters of the gov module. | GET|/cosmos/gov/v1beta1/params/{params_type}|
| `Deposit` | [QueryDepositRequest](#cosmos.gov.v1beta1.QueryDepositRequest) | [QueryDepositResponse](#cosmos.gov.v1beta1.QueryDepositResponse) | Deposit queries single deposit information based proposalID, depositAddr. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits/{depositor}|
| `Deposits` | [QueryDepositsRequest](#cosmos.gov.v1beta1.QueryDepositsRequest) | [QueryDepositsResponse](#cosmos.gov.v1beta1.QueryDepositsResponse) | Deposits queries all deposits of a single proposal. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/deposits|
| `TallyResult` | [QueryTallyResultRequest](#cosmos.gov.v1beta1.QueryTallyResultRequest) | [QueryTallyResultResponse](#cosmos.gov.v1beta1.QueryTallyResultResponse) | TallyResult queries the tally of a proposal vote. | GET|/cosmos/gov/v1beta1/proposals/{proposal_id}/tally|

 <!-- end services -->



<a name="cosmos/gov/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/gov/v1beta1/tx.proto



<a name="cosmos.gov.v1beta1.MsgDeposit"></a>

### MsgDeposit
MsgDeposit defines a message to submit a deposit to an existing proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `depositor` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.gov.v1beta1.MsgDepositResponse"></a>

### MsgDepositResponse
MsgDepositResponse defines the Msg/Deposit response type.






<a name="cosmos.gov.v1beta1.MsgSubmitProposal"></a>

### MsgSubmitProposal
MsgSubmitProposal defines an sdk.Msg type that supports submitting arbitrary
proposal Content.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `initial_deposit` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `proposer` | [string](#string) |  |  |






<a name="cosmos.gov.v1beta1.MsgSubmitProposalResponse"></a>

### MsgSubmitProposalResponse
MsgSubmitProposalResponse defines the Msg/SubmitProposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |






<a name="cosmos.gov.v1beta1.MsgVote"></a>

### MsgVote
MsgVote defines a message to cast a vote.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `option` | [VoteOption](#cosmos.gov.v1beta1.VoteOption) |  |  |






<a name="cosmos.gov.v1beta1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse defines the Msg/Vote response type.






<a name="cosmos.gov.v1beta1.MsgVoteWeighted"></a>

### MsgVoteWeighted
MsgVoteWeighted defines a message to cast a vote.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  |  |
| `voter` | [string](#string) |  |  |
| `options` | [WeightedVoteOption](#cosmos.gov.v1beta1.WeightedVoteOption) | repeated |  |






<a name="cosmos.gov.v1beta1.MsgVoteWeightedResponse"></a>

### MsgVoteWeightedResponse
MsgVoteWeightedResponse defines the Msg/VoteWeighted response type.

Since: cosmos-sdk 0.43





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.gov.v1beta1.Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SubmitProposal` | [MsgSubmitProposal](#cosmos.gov.v1beta1.MsgSubmitProposal) | [MsgSubmitProposalResponse](#cosmos.gov.v1beta1.MsgSubmitProposalResponse) | SubmitProposal defines a method to create new proposal given a content. | |
| `Vote` | [MsgVote](#cosmos.gov.v1beta1.MsgVote) | [MsgVoteResponse](#cosmos.gov.v1beta1.MsgVoteResponse) | Vote defines a method to add a vote on a specific proposal. | |
| `VoteWeighted` | [MsgVoteWeighted](#cosmos.gov.v1beta1.MsgVoteWeighted) | [MsgVoteWeightedResponse](#cosmos.gov.v1beta1.MsgVoteWeightedResponse) | VoteWeighted defines a method to add a weighted vote on a specific proposal.

Since: cosmos-sdk 0.43 | |
| `Deposit` | [MsgDeposit](#cosmos.gov.v1beta1.MsgDeposit) | [MsgDepositResponse](#cosmos.gov.v1beta1.MsgDepositResponse) | Deposit defines a method to add deposit on a specific proposal. | |

 <!-- end services -->



<a name="cosmos/group/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/group/v1/types.proto
Since: cosmos-sdk 0.46


<a name="cosmos.group.v1.DecisionPolicyWindows"></a>

### DecisionPolicyWindows
DecisionPolicyWindows defines the different windows for voting and execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voting_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | voting_period is the duration from submission of a proposal to the end of voting period Within this times votes can be submitted with MsgVote. |
| `min_execution_period` | [google.protobuf.Duration](#google.protobuf.Duration) |  | min_execution_period is the minimum duration after the proposal submission where members can start sending MsgExec. This means that the window for sending a MsgExec transaction is: `[ submission + min_execution_period ; submission + voting_period + max_execution_period]` where max_execution_period is a app-specific config, defined in the keeper. If not set, min_execution_period will default to 0.

Please make sure to set a `min_execution_period` that is smaller than `voting_period + max_execution_period`, or else the above execution window is empty, meaning that all proposals created with this decision policy won't be able to be executed. |






<a name="cosmos.group.v1.GroupInfo"></a>

### GroupInfo
GroupInfo represents the high-level on-chain information for a group.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id is the unique ID of the group. |
| `admin` | [string](#string) |  | admin is the account address of the group's admin. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the group. |
| `version` | [uint64](#uint64) |  | version is used to track changes to a group's membership structure that would break existing proposals. Whenever any members weight is changed, or any member is added or removed this version is incremented and will cause proposals based on older versions of this group to fail |
| `total_weight` | [string](#string) |  | total_weight is the sum of the group members' weights. |
| `created_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | created_at is a timestamp specifying when a group was created. |






<a name="cosmos.group.v1.GroupMember"></a>

### GroupMember
GroupMember represents the relationship between a group and a member.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `member` | [Member](#cosmos.group.v1.Member) |  | member is the member data. |






<a name="cosmos.group.v1.GroupPolicyInfo"></a>

### GroupPolicyInfo
GroupPolicyInfo represents the high-level on-chain information for a group policy.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of group policy. |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the group policy. |
| `version` | [uint64](#uint64) |  | version is used to track changes to a group's GroupPolicyInfo structure that would create a different result on a running proposal. |
| `decision_policy` | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy specifies the group policy's decision policy. |
| `created_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | created_at is a timestamp specifying when a group policy was created. |






<a name="cosmos.group.v1.Member"></a>

### Member
Member represents a group member with an account address,
non-zero weight, metadata and added_at timestamp.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the member's account address. |
| `weight` | [string](#string) |  | weight is the member's voting weight that should be greater than 0. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata attached to the member. |
| `added_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | added_at is a timestamp specifying when a member was added. |






<a name="cosmos.group.v1.MemberRequest"></a>

### MemberRequest
MemberRequest represents a group member to be used in Msg server requests.
Contrary to `Member`, it doesn't have any `added_at` field
since this field cannot be set as part of requests.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the member's account address. |
| `weight` | [string](#string) |  | weight is the member's voting weight that should be greater than 0. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata attached to the member. |






<a name="cosmos.group.v1.PercentageDecisionPolicy"></a>

### PercentageDecisionPolicy
PercentageDecisionPolicy is a decision policy where a proposal passes when
it satisfies the two following conditions:
1. The percentage of all `YES` voters' weights out of the total group weight
   is greater or equal than the given `percentage`.
2. The voting and execution periods of the proposal respect the parameters
   given by `windows`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `percentage` | [string](#string) |  | percentage is the minimum percentage the weighted sum of `YES` votes must meet for a proposal to succeed. |
| `windows` | [DecisionPolicyWindows](#cosmos.group.v1.DecisionPolicyWindows) |  | windows defines the different windows for voting and execution. |






<a name="cosmos.group.v1.Proposal"></a>

### Proposal
Proposal defines a group proposal. Any member of a group can submit a proposal
for a group policy to decide upon.
A proposal consists of a set of `sdk.Msg`s that will be executed if the proposal
passes as well as some optional metadata associated with the proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  | id is the unique id of the proposal. |
| `group_policy_address` | [string](#string) |  | group_policy_address is the account address of group policy. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the proposal. |
| `proposers` | [string](#string) | repeated | proposers are the account addresses of the proposers. |
| `submit_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submit_time is a timestamp specifying when a proposal was submitted. |
| `group_version` | [uint64](#uint64) |  | group_version tracks the version of the group at proposal submission. This field is here for informational purposes only. |
| `group_policy_version` | [uint64](#uint64) |  | group_policy_version tracks the version of the group policy at proposal submission. When a decision policy is changed, existing proposals from previous policy versions will become invalid with the `ABORTED` status. This field is here for informational purposes only. |
| `status` | [ProposalStatus](#cosmos.group.v1.ProposalStatus) |  | status represents the high level position in the life cycle of the proposal. Initial value is Submitted. |
| `final_tally_result` | [TallyResult](#cosmos.group.v1.TallyResult) |  | final_tally_result contains the sums of all weighted votes for this proposal for each vote option. It is empty at submission, and only populated after tallying, at voting period end or at proposal execution, whichever happens first. |
| `voting_period_end` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | voting_period_end is the timestamp before which voting must be done. Unless a successfull MsgExec is called before (to execute a proposal whose tally is successful before the voting period ends), tallying will be done at this point, and the `final_tally_result`and `status` fields will be accordingly updated. |
| `executor_result` | [ProposalExecutorResult](#cosmos.group.v1.ProposalExecutorResult) |  | executor_result is the final result of the proposal execution. Initial value is NotRun. |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated | messages is a list of `sdk.Msg`s that will be executed if the proposal passes. |






<a name="cosmos.group.v1.TallyResult"></a>

### TallyResult
TallyResult represents the sum of weighted votes for each vote option.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `yes_count` | [string](#string) |  | yes_count is the weighted sum of yes votes. |
| `abstain_count` | [string](#string) |  | abstain_count is the weighted sum of abstainers. |
| `no_count` | [string](#string) |  | no_count is the weighted sum of no votes. |
| `no_with_veto_count` | [string](#string) |  | no_with_veto_count is the weighted sum of veto. |






<a name="cosmos.group.v1.ThresholdDecisionPolicy"></a>

### ThresholdDecisionPolicy
ThresholdDecisionPolicy is a decision policy where a proposal passes when it
satisfies the two following conditions:
1. The sum of all `YES` voters' weights is greater or equal than the defined
   `threshold`.
2. The voting and execution periods of the proposal respect the parameters
   given by `windows`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `threshold` | [string](#string) |  | threshold is the minimum weighted sum of `YES` votes that must be met or exceeded for a proposal to succeed. |
| `windows` | [DecisionPolicyWindows](#cosmos.group.v1.DecisionPolicyWindows) |  | windows defines the different windows for voting and execution. |






<a name="cosmos.group.v1.Vote"></a>

### Vote
Vote represents a vote for a proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| `voter` | [string](#string) |  | voter is the account address of the voter. |
| `option` | [VoteOption](#cosmos.group.v1.VoteOption) |  | option is the voter's choice on the proposal. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the vote. |
| `submit_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submit_time is the timestamp when the vote was submitted. |





 <!-- end messages -->


<a name="cosmos.group.v1.ProposalExecutorResult"></a>

### ProposalExecutorResult
ProposalExecutorResult defines types of proposal executor results.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROPOSAL_EXECUTOR_RESULT_UNSPECIFIED | 0 | An empty value is not allowed. |
| PROPOSAL_EXECUTOR_RESULT_NOT_RUN | 1 | We have not yet run the executor. |
| PROPOSAL_EXECUTOR_RESULT_SUCCESS | 2 | The executor was successful and proposed action updated state. |
| PROPOSAL_EXECUTOR_RESULT_FAILURE | 3 | The executor returned an error and proposed action didn't update state. |



<a name="cosmos.group.v1.ProposalStatus"></a>

### ProposalStatus
ProposalStatus defines proposal statuses.

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROPOSAL_STATUS_UNSPECIFIED | 0 | An empty value is invalid and not allowed. |
| PROPOSAL_STATUS_SUBMITTED | 1 | Initial status of a proposal when submitted. |
| PROPOSAL_STATUS_ACCEPTED | 2 | Final status of a proposal when the final tally is done and the outcome passes the group policy's decision policy. |
| PROPOSAL_STATUS_REJECTED | 3 | Final status of a proposal when the final tally is done and the outcome is rejected by the group policy's decision policy. |
| PROPOSAL_STATUS_ABORTED | 4 | Final status of a proposal when the group policy is modified before the final tally. |
| PROPOSAL_STATUS_WITHDRAWN | 5 | A proposal can be withdrawn before the voting start time by the owner. When this happens the final status is Withdrawn. |



<a name="cosmos.group.v1.VoteOption"></a>

### VoteOption
VoteOption enumerates the valid vote options for a given proposal.

| Name | Number | Description |
| ---- | ------ | ----------- |
| VOTE_OPTION_UNSPECIFIED | 0 | VOTE_OPTION_UNSPECIFIED defines an unspecified vote option which will return an error. |
| VOTE_OPTION_YES | 1 | VOTE_OPTION_YES defines a yes vote option. |
| VOTE_OPTION_ABSTAIN | 2 | VOTE_OPTION_ABSTAIN defines an abstain vote option. |
| VOTE_OPTION_NO | 3 | VOTE_OPTION_NO defines a no vote option. |
| VOTE_OPTION_NO_WITH_VETO | 4 | VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/group/v1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/group/v1/events.proto
Since: cosmos-sdk 0.46


<a name="cosmos.group.v1.EventCreateGroup"></a>

### EventCreateGroup
EventCreateGroup is an event emitted when a group is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |






<a name="cosmos.group.v1.EventCreateGroupPolicy"></a>

### EventCreateGroupPolicy
EventCreateGroupPolicy is an event emitted when a group policy is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of the group policy. |






<a name="cosmos.group.v1.EventExec"></a>

### EventExec
EventExec is an event emitted when a proposal is executed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of the proposal. |
| `result` | [ProposalExecutorResult](#cosmos.group.v1.ProposalExecutorResult) |  | result is the proposal execution result. |
| `logs` | [string](#string) |  | logs contains error logs in case the execution result is FAILURE. |






<a name="cosmos.group.v1.EventLeaveGroup"></a>

### EventLeaveGroup
EventLeaveGroup is an event emitted when group member leaves the group.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `address` | [string](#string) |  | address is the account address of the group member. |






<a name="cosmos.group.v1.EventSubmitProposal"></a>

### EventSubmitProposal
EventSubmitProposal is an event emitted when a proposal is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of the proposal. |






<a name="cosmos.group.v1.EventUpdateGroup"></a>

### EventUpdateGroup
EventUpdateGroup is an event emitted when a group is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |






<a name="cosmos.group.v1.EventUpdateGroupPolicy"></a>

### EventUpdateGroupPolicy
EventUpdateGroupPolicy is an event emitted when a group policy is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of the group policy. |






<a name="cosmos.group.v1.EventVote"></a>

### EventVote
EventVote is an event emitted when a voter votes on a proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of the proposal. |






<a name="cosmos.group.v1.EventWithdrawProposal"></a>

### EventWithdrawProposal
EventWithdrawProposal is an event emitted when a proposal is withdrawn.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of the proposal. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/group/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/group/v1/genesis.proto
Since: cosmos-sdk 0.46


<a name="cosmos.group.v1.GenesisState"></a>

### GenesisState
GenesisState defines the group module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_seq` | [uint64](#uint64) |  | group_seq is the group table orm.Sequence, it is used to get the next group ID. |
| `groups` | [GroupInfo](#cosmos.group.v1.GroupInfo) | repeated | groups is the list of groups info. |
| `group_members` | [GroupMember](#cosmos.group.v1.GroupMember) | repeated | group_members is the list of groups members. |
| `group_policy_seq` | [uint64](#uint64) |  | group_policy_seq is the group policy table orm.Sequence, it is used to generate the next group policy account address. |
| `group_policies` | [GroupPolicyInfo](#cosmos.group.v1.GroupPolicyInfo) | repeated | group_policies is the list of group policies info. |
| `proposal_seq` | [uint64](#uint64) |  | proposal_seq is the proposal table orm.Sequence, it is used to get the next proposal ID. |
| `proposals` | [Proposal](#cosmos.group.v1.Proposal) | repeated | proposals is the list of proposals. |
| `votes` | [Vote](#cosmos.group.v1.Vote) | repeated | votes is the list of votes. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/group/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/group/v1/query.proto
Since: cosmos-sdk 0.46


<a name="cosmos.group.v1.QueryGroupInfoRequest"></a>

### QueryGroupInfoRequest
QueryGroupInfoRequest is the Query/GroupInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |






<a name="cosmos.group.v1.QueryGroupInfoResponse"></a>

### QueryGroupInfoResponse
QueryGroupInfoResponse is the Query/GroupInfo response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [GroupInfo](#cosmos.group.v1.GroupInfo) |  | info is the GroupInfo for the group. |






<a name="cosmos.group.v1.QueryGroupMembersRequest"></a>

### QueryGroupMembersRequest
QueryGroupMembersRequest is the Query/GroupMembers request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryGroupMembersResponse"></a>

### QueryGroupMembersResponse
QueryGroupMembersResponse is the Query/GroupMembersResponse response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `members` | [GroupMember](#cosmos.group.v1.GroupMember) | repeated | members are the members of the group with given group_id. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryGroupPoliciesByAdminRequest"></a>

### QueryGroupPoliciesByAdminRequest
QueryGroupPoliciesByAdminRequest is the Query/GroupPoliciesByAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the admin address of the group policy. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryGroupPoliciesByAdminResponse"></a>

### QueryGroupPoliciesByAdminResponse
QueryGroupPoliciesByAdminResponse is the Query/GroupPoliciesByAdmin response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_policies` | [GroupPolicyInfo](#cosmos.group.v1.GroupPolicyInfo) | repeated | group_policies are the group policies info with provided admin. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryGroupPoliciesByGroupRequest"></a>

### QueryGroupPoliciesByGroupRequest
QueryGroupPoliciesByGroupRequest is the Query/GroupPoliciesByGroup request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group policy's group. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryGroupPoliciesByGroupResponse"></a>

### QueryGroupPoliciesByGroupResponse
QueryGroupPoliciesByGroupResponse is the Query/GroupPoliciesByGroup response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_policies` | [GroupPolicyInfo](#cosmos.group.v1.GroupPolicyInfo) | repeated | group_policies are the group policies info associated with the provided group. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryGroupPolicyInfoRequest"></a>

### QueryGroupPolicyInfoRequest
QueryGroupPolicyInfoRequest is the Query/GroupPolicyInfo request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of the group policy. |






<a name="cosmos.group.v1.QueryGroupPolicyInfoResponse"></a>

### QueryGroupPolicyInfoResponse
QueryGroupPolicyInfoResponse is the Query/GroupPolicyInfo response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [GroupPolicyInfo](#cosmos.group.v1.GroupPolicyInfo) |  | info is the GroupPolicyInfo for the group policy. |






<a name="cosmos.group.v1.QueryGroupsByAdminRequest"></a>

### QueryGroupsByAdminRequest
QueryGroupsByAdminRequest is the Query/GroupsByAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of a group's admin. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryGroupsByAdminResponse"></a>

### QueryGroupsByAdminResponse
QueryGroupsByAdminResponse is the Query/GroupsByAdminResponse response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `groups` | [GroupInfo](#cosmos.group.v1.GroupInfo) | repeated | groups are the groups info with the provided admin. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryGroupsByMemberRequest"></a>

### QueryGroupsByMemberRequest
QueryGroupsByMemberRequest is the Query/GroupsByMember request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the group member address. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryGroupsByMemberResponse"></a>

### QueryGroupsByMemberResponse
QueryGroupsByMemberResponse is the Query/GroupsByMember response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `groups` | [GroupInfo](#cosmos.group.v1.GroupInfo) | repeated | groups are the groups info with the provided group member. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryProposalRequest"></a>

### QueryProposalRequest
QueryProposalRequest is the Query/Proposal request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of a proposal. |






<a name="cosmos.group.v1.QueryProposalResponse"></a>

### QueryProposalResponse
QueryProposalResponse is the Query/Proposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal` | [Proposal](#cosmos.group.v1.Proposal) |  | proposal is the proposal info. |






<a name="cosmos.group.v1.QueryProposalsByGroupPolicyRequest"></a>

### QueryProposalsByGroupPolicyRequest
QueryProposalsByGroupPolicyRequest is the Query/ProposalByGroupPolicy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of the group policy related to proposals. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryProposalsByGroupPolicyResponse"></a>

### QueryProposalsByGroupPolicyResponse
QueryProposalsByGroupPolicyResponse is the Query/ProposalByGroupPolicy response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposals` | [Proposal](#cosmos.group.v1.Proposal) | repeated | proposals are the proposals with given group policy. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryTallyResultRequest"></a>

### QueryTallyResultRequest
QueryTallyResultRequest is the Query/TallyResult request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique id of a proposal. |






<a name="cosmos.group.v1.QueryTallyResultResponse"></a>

### QueryTallyResultResponse
QueryTallyResultResponse is the Query/TallyResult response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tally` | [TallyResult](#cosmos.group.v1.TallyResult) |  | tally defines the requested tally. |






<a name="cosmos.group.v1.QueryVoteByProposalVoterRequest"></a>

### QueryVoteByProposalVoterRequest
QueryVoteByProposalVoterRequest is the Query/VoteByProposalVoter request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of a proposal. |
| `voter` | [string](#string) |  | voter is a proposal voter account address. |






<a name="cosmos.group.v1.QueryVoteByProposalVoterResponse"></a>

### QueryVoteByProposalVoterResponse
QueryVoteByProposalVoterResponse is the Query/VoteByProposalVoter response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vote` | [Vote](#cosmos.group.v1.Vote) |  | vote is the vote with given proposal_id and voter. |






<a name="cosmos.group.v1.QueryVotesByProposalRequest"></a>

### QueryVotesByProposalRequest
QueryVotesByProposalRequest is the Query/VotesByProposal request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal_id is the unique ID of a proposal. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryVotesByProposalResponse"></a>

### QueryVotesByProposalResponse
QueryVotesByProposalResponse is the Query/VotesByProposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [Vote](#cosmos.group.v1.Vote) | repeated | votes are the list of votes for given proposal_id. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.group.v1.QueryVotesByVoterRequest"></a>

### QueryVotesByVoterRequest
QueryVotesByVoterRequest is the Query/VotesByVoter request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `voter` | [string](#string) |  | voter is a proposal voter account address. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.group.v1.QueryVotesByVoterResponse"></a>

### QueryVotesByVoterResponse
QueryVotesByVoterResponse is the Query/VotesByVoter response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `votes` | [Vote](#cosmos.group.v1.Vote) | repeated | votes are the list of votes by given voter. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.group.v1.Query"></a>

### Query
Query is the cosmos.group.v1 Query service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GroupInfo` | [QueryGroupInfoRequest](#cosmos.group.v1.QueryGroupInfoRequest) | [QueryGroupInfoResponse](#cosmos.group.v1.QueryGroupInfoResponse) | GroupInfo queries group info based on group id. | GET|/cosmos/group/v1/group_info/{group_id}|
| `GroupPolicyInfo` | [QueryGroupPolicyInfoRequest](#cosmos.group.v1.QueryGroupPolicyInfoRequest) | [QueryGroupPolicyInfoResponse](#cosmos.group.v1.QueryGroupPolicyInfoResponse) | GroupPolicyInfo queries group policy info based on account address of group policy. | GET|/cosmos/group/v1/group_policy_info/{address}|
| `GroupMembers` | [QueryGroupMembersRequest](#cosmos.group.v1.QueryGroupMembersRequest) | [QueryGroupMembersResponse](#cosmos.group.v1.QueryGroupMembersResponse) | GroupMembers queries members of a group | GET|/cosmos/group/v1/group_members/{group_id}|
| `GroupsByAdmin` | [QueryGroupsByAdminRequest](#cosmos.group.v1.QueryGroupsByAdminRequest) | [QueryGroupsByAdminResponse](#cosmos.group.v1.QueryGroupsByAdminResponse) | GroupsByAdmin queries groups by admin address. | GET|/cosmos/group/v1/groups_by_admin/{admin}|
| `GroupPoliciesByGroup` | [QueryGroupPoliciesByGroupRequest](#cosmos.group.v1.QueryGroupPoliciesByGroupRequest) | [QueryGroupPoliciesByGroupResponse](#cosmos.group.v1.QueryGroupPoliciesByGroupResponse) | GroupPoliciesByGroup queries group policies by group id. | GET|/cosmos/group/v1/group_policies_by_group/{group_id}|
| `GroupPoliciesByAdmin` | [QueryGroupPoliciesByAdminRequest](#cosmos.group.v1.QueryGroupPoliciesByAdminRequest) | [QueryGroupPoliciesByAdminResponse](#cosmos.group.v1.QueryGroupPoliciesByAdminResponse) | GroupsByAdmin queries group policies by admin address. | GET|/cosmos/group/v1/group_policies_by_admin/{admin}|
| `Proposal` | [QueryProposalRequest](#cosmos.group.v1.QueryProposalRequest) | [QueryProposalResponse](#cosmos.group.v1.QueryProposalResponse) | Proposal queries a proposal based on proposal id. | GET|/cosmos/group/v1/proposal/{proposal_id}|
| `ProposalsByGroupPolicy` | [QueryProposalsByGroupPolicyRequest](#cosmos.group.v1.QueryProposalsByGroupPolicyRequest) | [QueryProposalsByGroupPolicyResponse](#cosmos.group.v1.QueryProposalsByGroupPolicyResponse) | ProposalsByGroupPolicy queries proposals based on account address of group policy. | GET|/cosmos/group/v1/proposals_by_group_policy/{address}|
| `VoteByProposalVoter` | [QueryVoteByProposalVoterRequest](#cosmos.group.v1.QueryVoteByProposalVoterRequest) | [QueryVoteByProposalVoterResponse](#cosmos.group.v1.QueryVoteByProposalVoterResponse) | VoteByProposalVoter queries a vote by proposal id and voter. | GET|/cosmos/group/v1/vote_by_proposal_voter/{proposal_id}/{voter}|
| `VotesByProposal` | [QueryVotesByProposalRequest](#cosmos.group.v1.QueryVotesByProposalRequest) | [QueryVotesByProposalResponse](#cosmos.group.v1.QueryVotesByProposalResponse) | VotesByProposal queries a vote by proposal. | GET|/cosmos/group/v1/votes_by_proposal/{proposal_id}|
| `VotesByVoter` | [QueryVotesByVoterRequest](#cosmos.group.v1.QueryVotesByVoterRequest) | [QueryVotesByVoterResponse](#cosmos.group.v1.QueryVotesByVoterResponse) | VotesByVoter queries a vote by voter. | GET|/cosmos/group/v1/votes_by_voter/{voter}|
| `GroupsByMember` | [QueryGroupsByMemberRequest](#cosmos.group.v1.QueryGroupsByMemberRequest) | [QueryGroupsByMemberResponse](#cosmos.group.v1.QueryGroupsByMemberResponse) | GroupsByMember queries groups by member address. | GET|/cosmos/group/v1/groups_by_member/{address}|
| `TallyResult` | [QueryTallyResultRequest](#cosmos.group.v1.QueryTallyResultRequest) | [QueryTallyResultResponse](#cosmos.group.v1.QueryTallyResultResponse) | TallyResult returns the tally result of a proposal. If the proposal is still in voting period, then this query computes the current tally state, which might not be final. On the other hand, if the proposal is final, then it simply returns the `final_tally_result` state stored in the proposal itself. | GET|/cosmos/group/v1/proposals/{proposal_id}/tally|

 <!-- end services -->



<a name="cosmos/group/v1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/group/v1/tx.proto
Since: cosmos-sdk 0.46


<a name="cosmos.group.v1.MsgCreateGroup"></a>

### MsgCreateGroup
MsgCreateGroup is the Msg/CreateGroup request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `members` | [MemberRequest](#cosmos.group.v1.MemberRequest) | repeated | members defines the group members. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the group. |






<a name="cosmos.group.v1.MsgCreateGroupPolicy"></a>

### MsgCreateGroupPolicy
MsgCreateGroupPolicy is the Msg/CreateGroupPolicy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata attached to the group policy. |
| `decision_policy` | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy specifies the group policy's decision policy. |






<a name="cosmos.group.v1.MsgCreateGroupPolicyResponse"></a>

### MsgCreateGroupPolicyResponse
MsgCreateGroupPolicyResponse is the Msg/CreateGroupPolicy response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of the newly created group policy. |






<a name="cosmos.group.v1.MsgCreateGroupResponse"></a>

### MsgCreateGroupResponse
MsgCreateGroupResponse is the Msg/CreateGroup response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the newly created group. |






<a name="cosmos.group.v1.MsgCreateGroupWithPolicy"></a>

### MsgCreateGroupWithPolicy
MsgCreateGroupWithPolicy is the Msg/CreateGroupWithPolicy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group and group policy admin. |
| `members` | [MemberRequest](#cosmos.group.v1.MemberRequest) | repeated | members defines the group members. |
| `group_metadata` | [string](#string) |  | group_metadata is any arbitrary metadata attached to the group. |
| `group_policy_metadata` | [string](#string) |  | group_policy_metadata is any arbitrary metadata attached to the group policy. |
| `group_policy_as_admin` | [bool](#bool) |  | group_policy_as_admin is a boolean field, if set to true, the group policy account address will be used as group and group policy admin. |
| `decision_policy` | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy specifies the group policy's decision policy. |






<a name="cosmos.group.v1.MsgCreateGroupWithPolicyResponse"></a>

### MsgCreateGroupWithPolicyResponse
MsgCreateGroupWithPolicyResponse is the Msg/CreateGroupWithPolicy response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the newly created group with policy. |
| `group_policy_address` | [string](#string) |  | group_policy_address is the account address of the newly created group policy. |






<a name="cosmos.group.v1.MsgExec"></a>

### MsgExec
MsgExec is the Msg/Exec request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| `executor` | [string](#string) |  | executor is the account address used to execute the proposal. |






<a name="cosmos.group.v1.MsgExecResponse"></a>

### MsgExecResponse
MsgExecResponse is the Msg/Exec request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `result` | [ProposalExecutorResult](#cosmos.group.v1.ProposalExecutorResult) |  | result is the final result of the proposal execution. |






<a name="cosmos.group.v1.MsgLeaveGroup"></a>

### MsgLeaveGroup
MsgLeaveGroup is the Msg/LeaveGroup request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the account address of the group member. |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |






<a name="cosmos.group.v1.MsgLeaveGroupResponse"></a>

### MsgLeaveGroupResponse
MsgLeaveGroupResponse is the Msg/LeaveGroup response type.






<a name="cosmos.group.v1.MsgSubmitProposal"></a>

### MsgSubmitProposal
MsgSubmitProposal is the Msg/SubmitProposal request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `group_policy_address` | [string](#string) |  | group_policy_address is the account address of group policy. |
| `proposers` | [string](#string) | repeated | proposers are the account addresses of the proposers. Proposers signatures will be counted as yes votes. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the proposal. |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated | messages is a list of `sdk.Msg`s that will be executed if the proposal passes. |
| `exec` | [Exec](#cosmos.group.v1.Exec) |  | exec defines the mode of execution of the proposal, whether it should be executed immediately on creation or not. If so, proposers signatures are considered as Yes votes. |






<a name="cosmos.group.v1.MsgSubmitProposalResponse"></a>

### MsgSubmitProposalResponse
MsgSubmitProposalResponse is the Msg/SubmitProposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |






<a name="cosmos.group.v1.MsgUpdateGroupAdmin"></a>

### MsgUpdateGroupAdmin
MsgUpdateGroupAdmin is the Msg/UpdateGroupAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the current account address of the group admin. |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `new_admin` | [string](#string) |  | new_admin is the group new admin account address. |






<a name="cosmos.group.v1.MsgUpdateGroupAdminResponse"></a>

### MsgUpdateGroupAdminResponse
MsgUpdateGroupAdminResponse is the Msg/UpdateGroupAdmin response type.






<a name="cosmos.group.v1.MsgUpdateGroupMembers"></a>

### MsgUpdateGroupMembers
MsgUpdateGroupMembers is the Msg/UpdateGroupMembers request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `member_updates` | [MemberRequest](#cosmos.group.v1.MemberRequest) | repeated | member_updates is the list of members to update, set weight to 0 to remove a member. |






<a name="cosmos.group.v1.MsgUpdateGroupMembersResponse"></a>

### MsgUpdateGroupMembersResponse
MsgUpdateGroupMembersResponse is the Msg/UpdateGroupMembers response type.






<a name="cosmos.group.v1.MsgUpdateGroupMetadata"></a>

### MsgUpdateGroupMetadata
MsgUpdateGroupMetadata is the Msg/UpdateGroupMetadata request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `group_id` | [uint64](#uint64) |  | group_id is the unique ID of the group. |
| `metadata` | [string](#string) |  | metadata is the updated group's metadata. |






<a name="cosmos.group.v1.MsgUpdateGroupMetadataResponse"></a>

### MsgUpdateGroupMetadataResponse
MsgUpdateGroupMetadataResponse is the Msg/UpdateGroupMetadata response type.






<a name="cosmos.group.v1.MsgUpdateGroupPolicyAdmin"></a>

### MsgUpdateGroupPolicyAdmin
MsgUpdateGroupPolicyAdmin is the Msg/UpdateGroupPolicyAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `group_policy_address` | [string](#string) |  | group_policy_address is the account address of the group policy. |
| `new_admin` | [string](#string) |  | new_admin is the new group policy admin. |






<a name="cosmos.group.v1.MsgUpdateGroupPolicyAdminResponse"></a>

### MsgUpdateGroupPolicyAdminResponse
MsgUpdateGroupPolicyAdminResponse is the Msg/UpdateGroupPolicyAdmin response type.






<a name="cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicy"></a>

### MsgUpdateGroupPolicyDecisionPolicy
MsgUpdateGroupPolicyDecisionPolicy is the Msg/UpdateGroupPolicyDecisionPolicy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `group_policy_address` | [string](#string) |  | group_policy_address is the account address of group policy. |
| `decision_policy` | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy is the updated group policy's decision policy. |






<a name="cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicyResponse"></a>

### MsgUpdateGroupPolicyDecisionPolicyResponse
MsgUpdateGroupPolicyDecisionPolicyResponse is the Msg/UpdateGroupPolicyDecisionPolicy response type.






<a name="cosmos.group.v1.MsgUpdateGroupPolicyMetadata"></a>

### MsgUpdateGroupPolicyMetadata
MsgUpdateGroupPolicyMetadata is the Msg/UpdateGroupPolicyMetadata request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `admin` | [string](#string) |  | admin is the account address of the group admin. |
| `group_policy_address` | [string](#string) |  | group_policy_address is the account address of group policy. |
| `metadata` | [string](#string) |  | metadata is the updated group policy metadata. |






<a name="cosmos.group.v1.MsgUpdateGroupPolicyMetadataResponse"></a>

### MsgUpdateGroupPolicyMetadataResponse
MsgUpdateGroupPolicyMetadataResponse is the Msg/UpdateGroupPolicyMetadata response type.






<a name="cosmos.group.v1.MsgVote"></a>

### MsgVote
MsgVote is the Msg/Vote request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| `voter` | [string](#string) |  | voter is the voter account address. |
| `option` | [VoteOption](#cosmos.group.v1.VoteOption) |  | option is the voter's choice on the proposal. |
| `metadata` | [string](#string) |  | metadata is any arbitrary metadata to attached to the vote. |
| `exec` | [Exec](#cosmos.group.v1.Exec) |  | exec defines whether the proposal should be executed immediately after voting or not. |






<a name="cosmos.group.v1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse is the Msg/Vote response type.






<a name="cosmos.group.v1.MsgWithdrawProposal"></a>

### MsgWithdrawProposal
MsgWithdrawProposal is the Msg/WithdrawProposal request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `proposal_id` | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| `address` | [string](#string) |  | address is the admin of the group policy or one of the proposer of the proposal. |






<a name="cosmos.group.v1.MsgWithdrawProposalResponse"></a>

### MsgWithdrawProposalResponse
MsgWithdrawProposalResponse is the Msg/WithdrawProposal response type.





 <!-- end messages -->


<a name="cosmos.group.v1.Exec"></a>

### Exec
Exec defines modes of execution of a proposal on creation or on new vote.

| Name | Number | Description |
| ---- | ------ | ----------- |
| EXEC_UNSPECIFIED | 0 | An empty value means that there should be a separate MsgExec request for the proposal to execute. |
| EXEC_TRY | 1 | Try to execute the proposal immediately. If the proposal is not allowed per the DecisionPolicy, the proposal will still be open and could be executed at a later point. |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.group.v1.Msg"></a>

### Msg
Msg is the cosmos.group.v1 Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateGroup` | [MsgCreateGroup](#cosmos.group.v1.MsgCreateGroup) | [MsgCreateGroupResponse](#cosmos.group.v1.MsgCreateGroupResponse) | CreateGroup creates a new group with an admin account address, a list of members and some optional metadata. | |
| `UpdateGroupMembers` | [MsgUpdateGroupMembers](#cosmos.group.v1.MsgUpdateGroupMembers) | [MsgUpdateGroupMembersResponse](#cosmos.group.v1.MsgUpdateGroupMembersResponse) | UpdateGroupMembers updates the group members with given group id and admin address. | |
| `UpdateGroupAdmin` | [MsgUpdateGroupAdmin](#cosmos.group.v1.MsgUpdateGroupAdmin) | [MsgUpdateGroupAdminResponse](#cosmos.group.v1.MsgUpdateGroupAdminResponse) | UpdateGroupAdmin updates the group admin with given group id and previous admin address. | |
| `UpdateGroupMetadata` | [MsgUpdateGroupMetadata](#cosmos.group.v1.MsgUpdateGroupMetadata) | [MsgUpdateGroupMetadataResponse](#cosmos.group.v1.MsgUpdateGroupMetadataResponse) | UpdateGroupMetadata updates the group metadata with given group id and admin address. | |
| `CreateGroupPolicy` | [MsgCreateGroupPolicy](#cosmos.group.v1.MsgCreateGroupPolicy) | [MsgCreateGroupPolicyResponse](#cosmos.group.v1.MsgCreateGroupPolicyResponse) | CreateGroupPolicy creates a new group policy using given DecisionPolicy. | |
| `CreateGroupWithPolicy` | [MsgCreateGroupWithPolicy](#cosmos.group.v1.MsgCreateGroupWithPolicy) | [MsgCreateGroupWithPolicyResponse](#cosmos.group.v1.MsgCreateGroupWithPolicyResponse) | CreateGroupWithPolicy creates a new group with policy. | |
| `UpdateGroupPolicyAdmin` | [MsgUpdateGroupPolicyAdmin](#cosmos.group.v1.MsgUpdateGroupPolicyAdmin) | [MsgUpdateGroupPolicyAdminResponse](#cosmos.group.v1.MsgUpdateGroupPolicyAdminResponse) | UpdateGroupPolicyAdmin updates a group policy admin. | |
| `UpdateGroupPolicyDecisionPolicy` | [MsgUpdateGroupPolicyDecisionPolicy](#cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicy) | [MsgUpdateGroupPolicyDecisionPolicyResponse](#cosmos.group.v1.MsgUpdateGroupPolicyDecisionPolicyResponse) | UpdateGroupPolicyDecisionPolicy allows a group policy's decision policy to be updated. | |
| `UpdateGroupPolicyMetadata` | [MsgUpdateGroupPolicyMetadata](#cosmos.group.v1.MsgUpdateGroupPolicyMetadata) | [MsgUpdateGroupPolicyMetadataResponse](#cosmos.group.v1.MsgUpdateGroupPolicyMetadataResponse) | UpdateGroupPolicyMetadata updates a group policy metadata. | |
| `SubmitProposal` | [MsgSubmitProposal](#cosmos.group.v1.MsgSubmitProposal) | [MsgSubmitProposalResponse](#cosmos.group.v1.MsgSubmitProposalResponse) | SubmitProposal submits a new proposal. | |
| `WithdrawProposal` | [MsgWithdrawProposal](#cosmos.group.v1.MsgWithdrawProposal) | [MsgWithdrawProposalResponse](#cosmos.group.v1.MsgWithdrawProposalResponse) | WithdrawProposal withdraws a proposal. | |
| `Vote` | [MsgVote](#cosmos.group.v1.MsgVote) | [MsgVoteResponse](#cosmos.group.v1.MsgVoteResponse) | Vote allows a voter to vote on a proposal. | |
| `Exec` | [MsgExec](#cosmos.group.v1.MsgExec) | [MsgExecResponse](#cosmos.group.v1.MsgExecResponse) | Exec executes a proposal. | |
| `LeaveGroup` | [MsgLeaveGroup](#cosmos.group.v1.MsgLeaveGroup) | [MsgLeaveGroupResponse](#cosmos.group.v1.MsgLeaveGroupResponse) | LeaveGroup allows a group member to leave the group. | |

 <!-- end services -->



<a name="cosmos/mint/v1beta1/mint.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/mint/v1beta1/mint.proto



<a name="cosmos.mint.v1beta1.Minter"></a>

### Minter
Minter represents the minting state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation` | [string](#string) |  | current annual inflation rate |
| `annual_provisions` | [string](#string) |  | current annual expected provisions |






<a name="cosmos.mint.v1beta1.Params"></a>

### Params
Params holds parameters for the mint module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mint_denom` | [string](#string) |  | type of coin to mint |
| `inflation_rate_change` | [string](#string) |  | maximum annual change in inflation rate |
| `inflation_max` | [string](#string) |  | maximum inflation rate |
| `inflation_min` | [string](#string) |  | minimum inflation rate |
| `goal_bonded` | [string](#string) |  | goal of percent bonded atoms |
| `blocks_per_year` | [uint64](#uint64) |  | expected blocks per year |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/mint/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/mint/v1beta1/genesis.proto



<a name="cosmos.mint.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the mint module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `minter` | [Minter](#cosmos.mint.v1beta1.Minter) |  | minter is a space for holding current inflation information. |
| `params` | [Params](#cosmos.mint.v1beta1.Params) |  | params defines all the paramaters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/mint/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/mint/v1beta1/query.proto



<a name="cosmos.mint.v1beta1.QueryAnnualProvisionsRequest"></a>

### QueryAnnualProvisionsRequest
QueryAnnualProvisionsRequest is the request type for the
Query/AnnualProvisions RPC method.






<a name="cosmos.mint.v1beta1.QueryAnnualProvisionsResponse"></a>

### QueryAnnualProvisionsResponse
QueryAnnualProvisionsResponse is the response type for the
Query/AnnualProvisions RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `annual_provisions` | [bytes](#bytes) |  | annual_provisions is the current minting annual provisions value. |






<a name="cosmos.mint.v1beta1.QueryInflationRequest"></a>

### QueryInflationRequest
QueryInflationRequest is the request type for the Query/Inflation RPC method.






<a name="cosmos.mint.v1beta1.QueryInflationResponse"></a>

### QueryInflationResponse
QueryInflationResponse is the response type for the Query/Inflation RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `inflation` | [bytes](#bytes) |  | inflation is the current minting inflation value. |






<a name="cosmos.mint.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method.






<a name="cosmos.mint.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.mint.v1beta1.Params) |  | params defines the parameters of the module. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.mint.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#cosmos.mint.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.mint.v1beta1.QueryParamsResponse) | Params returns the total set of minting parameters. | GET|/cosmos/mint/v1beta1/params|
| `Inflation` | [QueryInflationRequest](#cosmos.mint.v1beta1.QueryInflationRequest) | [QueryInflationResponse](#cosmos.mint.v1beta1.QueryInflationResponse) | Inflation returns the current minting inflation value. | GET|/cosmos/mint/v1beta1/inflation|
| `AnnualProvisions` | [QueryAnnualProvisionsRequest](#cosmos.mint.v1beta1.QueryAnnualProvisionsRequest) | [QueryAnnualProvisionsResponse](#cosmos.mint.v1beta1.QueryAnnualProvisionsResponse) | AnnualProvisions current minting annual provisions value. | GET|/cosmos/mint/v1beta1/annual_provisions|

 <!-- end services -->



<a name="cosmos/nft/v1beta1/event.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/nft/v1beta1/event.proto



<a name="cosmos.nft.v1beta1.EventBurn"></a>

### EventBurn
EventBurn is emitted on Burn


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.EventMint"></a>

### EventMint
EventMint is emitted on Mint


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.EventSend"></a>

### EventSend
EventSend is emitted on Msg/Send


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |
| `sender` | [string](#string) |  |  |
| `receiver` | [string](#string) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/nft/v1beta1/nft.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/nft/v1beta1/nft.proto



<a name="cosmos.nft.v1beta1.Class"></a>

### Class
Class defines the class of the nft type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [string](#string) |  | id defines the unique identifier of the NFT classification, similar to the contract address of ERC721 |
| `name` | [string](#string) |  | name defines the human-readable name of the NFT classification. Optional |
| `symbol` | [string](#string) |  | symbol is an abbreviated name for nft classification. Optional |
| `description` | [string](#string) |  | description is a brief description of nft classification. Optional |
| `uri` | [string](#string) |  | uri for the class metadata stored off chain. It can define schema for Class and NFT `Data` attributes. Optional |
| `uri_hash` | [string](#string) |  | uri_hash is a hash of the document pointed by uri. Optional |
| `data` | [google.protobuf.Any](#google.protobuf.Any) |  | data is the app specific metadata of the NFT class. Optional |






<a name="cosmos.nft.v1beta1.NFT"></a>

### NFT
NFT defines the NFT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  | class_id associated with the NFT, similar to the contract address of ERC721 |
| `id` | [string](#string) |  | id is a unique identifier of the NFT |
| `uri` | [string](#string) |  | uri for the NFT metadata stored off chain |
| `uri_hash` | [string](#string) |  | uri_hash is a hash of the document pointed by uri |
| `data` | [google.protobuf.Any](#google.protobuf.Any) |  | data is an app specific data of the NFT. Optional |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/nft/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/nft/v1beta1/genesis.proto



<a name="cosmos.nft.v1beta1.Entry"></a>

### Entry
Entry Defines all nft owned by a person


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  | owner is the owner address of the following nft |
| `nfts` | [NFT](#cosmos.nft.v1beta1.NFT) | repeated | nfts is a group of nfts of the same owner |






<a name="cosmos.nft.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the nft module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `classes` | [Class](#cosmos.nft.v1beta1.Class) | repeated | class defines the class of the nft type. |
| `entries` | [Entry](#cosmos.nft.v1beta1.Entry) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/nft/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/nft/v1beta1/query.proto



<a name="cosmos.nft.v1beta1.QueryBalanceRequest"></a>

### QueryBalanceRequest
QueryBalanceRequest is the request type for the Query/Balance RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.QueryBalanceResponse"></a>

### QueryBalanceResponse
QueryBalanceResponse is the response type for the Query/Balance RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [uint64](#uint64) |  |  |






<a name="cosmos.nft.v1beta1.QueryClassRequest"></a>

### QueryClassRequest
QueryClassRequest is the request type for the Query/Class RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.QueryClassResponse"></a>

### QueryClassResponse
QueryClassResponse is the response type for the Query/Class RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class` | [Class](#cosmos.nft.v1beta1.Class) |  |  |






<a name="cosmos.nft.v1beta1.QueryClassesRequest"></a>

### QueryClassesRequest
QueryClassesRequest is the request type for the Query/Classes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.nft.v1beta1.QueryClassesResponse"></a>

### QueryClassesResponse
QueryClassesResponse is the response type for the Query/Classes RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `classes` | [Class](#cosmos.nft.v1beta1.Class) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="cosmos.nft.v1beta1.QueryNFTRequest"></a>

### QueryNFTRequest
QueryNFTRequest is the request type for the Query/NFT RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.QueryNFTResponse"></a>

### QueryNFTResponse
QueryNFTResponse is the response type for the Query/NFT RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nft` | [NFT](#cosmos.nft.v1beta1.NFT) |  |  |






<a name="cosmos.nft.v1beta1.QueryNFTsRequest"></a>

### QueryNFTsRequest
QueryNFTstRequest is the request type for the Query/NFTs RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `owner` | [string](#string) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="cosmos.nft.v1beta1.QueryNFTsResponse"></a>

### QueryNFTsResponse
QueryNFTsResponse is the response type for the Query/NFTs RPC methods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nfts` | [NFT](#cosmos.nft.v1beta1.NFT) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |






<a name="cosmos.nft.v1beta1.QueryOwnerRequest"></a>

### QueryOwnerRequest
QueryOwnerRequest is the request type for the Query/Owner RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |
| `id` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.QueryOwnerResponse"></a>

### QueryOwnerResponse
QueryOwnerResponse is the response type for the Query/Owner RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `owner` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.QuerySupplyRequest"></a>

### QuerySupplyRequest
QuerySupplyRequest is the request type for the Query/Supply RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  |  |






<a name="cosmos.nft.v1beta1.QuerySupplyResponse"></a>

### QuerySupplyResponse
QuerySupplyResponse is the response type for the Query/Supply RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.nft.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Balance` | [QueryBalanceRequest](#cosmos.nft.v1beta1.QueryBalanceRequest) | [QueryBalanceResponse](#cosmos.nft.v1beta1.QueryBalanceResponse) | Balance queries the number of NFTs of a given class owned by the owner, same as balanceOf in ERC721 | GET|/cosmos/nft/v1beta1/balance/{owner}/{class_id}|
| `Owner` | [QueryOwnerRequest](#cosmos.nft.v1beta1.QueryOwnerRequest) | [QueryOwnerResponse](#cosmos.nft.v1beta1.QueryOwnerResponse) | Owner queries the owner of the NFT based on its class and id, same as ownerOf in ERC721 | GET|/cosmos/nft/v1beta1/owner/{class_id}/{id}|
| `Supply` | [QuerySupplyRequest](#cosmos.nft.v1beta1.QuerySupplyRequest) | [QuerySupplyResponse](#cosmos.nft.v1beta1.QuerySupplyResponse) | Supply queries the number of NFTs from the given class, same as totalSupply of ERC721. | GET|/cosmos/nft/v1beta1/supply/{class_id}|
| `NFTs` | [QueryNFTsRequest](#cosmos.nft.v1beta1.QueryNFTsRequest) | [QueryNFTsResponse](#cosmos.nft.v1beta1.QueryNFTsResponse) | NFTs queries all NFTs of a given class or owner,choose at least one of the two, similar to tokenByIndex in ERC721Enumerable | GET|/cosmos/nft/v1beta1/nfts|
| `NFT` | [QueryNFTRequest](#cosmos.nft.v1beta1.QueryNFTRequest) | [QueryNFTResponse](#cosmos.nft.v1beta1.QueryNFTResponse) | NFT queries an NFT based on its class and id. | GET|/cosmos/nft/v1beta1/nfts/{class_id}/{id}|
| `Class` | [QueryClassRequest](#cosmos.nft.v1beta1.QueryClassRequest) | [QueryClassResponse](#cosmos.nft.v1beta1.QueryClassResponse) | Class queries an NFT class based on its id | GET|/cosmos/nft/v1beta1/classes/{class_id}|
| `Classes` | [QueryClassesRequest](#cosmos.nft.v1beta1.QueryClassesRequest) | [QueryClassesResponse](#cosmos.nft.v1beta1.QueryClassesResponse) | Classes queries all NFT classes | GET|/cosmos/nft/v1beta1/classes|

 <!-- end services -->



<a name="cosmos/nft/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/nft/v1beta1/tx.proto



<a name="cosmos.nft.v1beta1.MsgSend"></a>

### MsgSend
MsgSend represents a message to send a nft from one account to another account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `class_id` | [string](#string) |  | class_id defines the unique identifier of the nft classification, similar to the contract address of ERC721 |
| `id` | [string](#string) |  | id defines the unique identification of nft |
| `sender` | [string](#string) |  | sender is the address of the owner of nft |
| `receiver` | [string](#string) |  | receiver is the receiver address of nft |






<a name="cosmos.nft.v1beta1.MsgSendResponse"></a>

### MsgSendResponse
MsgSendResponse defines the Msg/Send response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.nft.v1beta1.Msg"></a>

### Msg
Msg defines the nft Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Send` | [MsgSend](#cosmos.nft.v1beta1.MsgSend) | [MsgSendResponse](#cosmos.nft.v1beta1.MsgSendResponse) | Send defines a method to send a nft from one account to another account. | |

 <!-- end services -->



<a name="cosmos/orm/module/v1alpha1/module.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/orm/module/v1alpha1/module.proto



<a name="cosmos.orm.module.v1alpha1.Module"></a>

### Module
Module defines the ORM module which adds providers to the app container for
module-scoped DB's. In the future it may provide gRPC services for interacting
with ORM data.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/orm/v1/orm.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/orm/v1/orm.proto



<a name="cosmos.orm.v1.PrimaryKeyDescriptor"></a>

### PrimaryKeyDescriptor
PrimaryKeyDescriptor describes a table primary key.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fields` | [string](#string) |  | fields is a comma-separated list of fields in the primary key. Spaces are not allowed. Supported field types, their encodings, and any applicable constraints are described below. - uint32 are encoded as 2,3,4 or 5 bytes using a compact encoding that is suitable for sorted iteration (not varint encoding). This type is well-suited for small integers. - uint64 are encoded as 2,4,6 or 9 bytes using a compact encoding that is suitable for sorted iteration (not varint encoding). This type is well-suited for small integers such as auto-incrementing sequences. - fixed32, fixed64 are encoded as big-endian fixed width bytes and support sorted iteration. These types are well-suited for encoding fixed with decimals as integers. - string's are encoded as raw bytes in terminal key segments and null-terminated in non-terminal segments. Null characters are thus forbidden in strings. string fields support sorted iteration. - bytes are encoded as raw bytes in terminal segments and length-prefixed with a 32-bit unsigned varint in non-terminal segments. - int32, sint32, int64, sint64, sfixed32, sfixed64 are encoded as fixed width bytes with an encoding that enables sorted iteration. - google.protobuf.Timestamp and google.protobuf.Duration are encoded as 12 bytes using an encoding that enables sorted iteration. - enum fields are encoded using varint encoding and do not support sorted iteration. - bool fields are encoded as a single byte 0 or 1.

All other fields types are unsupported in keys including repeated and oneof fields.

Primary keys are prefixed by the varint encoded table id and the byte 0x0 plus any additional prefix specified by the schema. |
| `auto_increment` | [bool](#bool) |  | auto_increment specifies that the primary key is generated by an auto-incrementing integer. If this is set to true fields must only contain one field of that is of type uint64. |






<a name="cosmos.orm.v1.SecondaryIndexDescriptor"></a>

### SecondaryIndexDescriptor
PrimaryKeyDescriptor describes a table secondary index.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `fields` | [string](#string) |  | fields is a comma-separated list of fields in the index. The supported field types are the same as those for PrimaryKeyDescriptor.fields. Index keys are prefixed by the varint encoded table id and the varint encoded index id plus any additional prefix specified by the schema.

In addition the field segments, non-unique index keys are suffixed with any additional primary key fields not present in the index fields so that the primary key can be reconstructed. Unique indexes instead of being suffixed store the remaining primary key fields in the value.. |
| `id` | [uint32](#uint32) |  | id is a non-zero integer ID that must be unique within the indexes for this table and less than 32768. It may be deprecated in the future when this can be auto-generated. |
| `unique` | [bool](#bool) |  | unique specifies that this an unique index. |






<a name="cosmos.orm.v1.SingletonDescriptor"></a>

### SingletonDescriptor
TableDescriptor describes an ORM singleton table which has at most one instance.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint32](#uint32) |  | id is a non-zero integer ID that must be unique within the tables and singletons in this file. It may be deprecated in the future when this can be auto-generated. |






<a name="cosmos.orm.v1.TableDescriptor"></a>

### TableDescriptor
TableDescriptor describes an ORM table.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `primary_key` | [PrimaryKeyDescriptor](#cosmos.orm.v1.PrimaryKeyDescriptor) |  | primary_key defines the primary key for the table. |
| `index` | [SecondaryIndexDescriptor](#cosmos.orm.v1.SecondaryIndexDescriptor) | repeated | index defines one or more secondary indexes. |
| `id` | [uint32](#uint32) |  | id is a non-zero integer ID that must be unique within the tables and singletons in this file. It may be deprecated in the future when this can be auto-generated. |





 <!-- end messages -->

 <!-- end enums -->


<a name="cosmos/orm/v1/orm.proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| `singleton` | SingletonDescriptor | .google.protobuf.MessageOptions | 104503791 | singleton specifies that this message will be used as an ORM singleton. It cannot be used together with the table option. |
| `table` | TableDescriptor | .google.protobuf.MessageOptions | 104503790 | table specifies that this message will be used as an ORM table. It cannot be used together with the singleton option. |

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/orm/v1alpha1/schema.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/orm/v1alpha1/schema.proto



<a name="cosmos.orm.v1alpha1.ModuleSchemaDescriptor"></a>

### ModuleSchemaDescriptor
ModuleSchemaDescriptor describe's a module's ORM schema.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `schema_file` | [ModuleSchemaDescriptor.FileEntry](#cosmos.orm.v1alpha1.ModuleSchemaDescriptor.FileEntry) | repeated |  |
| `prefix` | [bytes](#bytes) |  | prefix is an optional prefix that precedes all keys in this module's store. |






<a name="cosmos.orm.v1alpha1.ModuleSchemaDescriptor.FileEntry"></a>

### ModuleSchemaDescriptor.FileEntry
FileEntry describes an ORM file used in a module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint32](#uint32) |  | id is a prefix that will be varint encoded and prepended to all the table keys specified in the file's tables. |
| `proto_file_name` | [string](#string) |  | proto_file_name is the name of a file .proto in that contains table definitions. The .proto file must be in a package that the module has referenced using cosmos.app.v1.ModuleDescriptor.use_package. |
| `storage_type` | [StorageType](#cosmos.orm.v1alpha1.StorageType) |  | storage_type optionally indicates the type of storage this file's tables should used. If it is left unspecified, the default KV-storage of the app will be used. |





 <!-- end messages -->


<a name="cosmos.orm.v1alpha1.StorageType"></a>

### StorageType
StorageType

| Name | Number | Description |
| ---- | ------ | ----------- |
| STORAGE_TYPE_DEFAULT_UNSPECIFIED | 0 | STORAGE_TYPE_DEFAULT_UNSPECIFIED indicates the persistent KV-storage where primary key entries are stored in merkle-tree backed commitment storage and indexes and seqs are stored in fast index storage. Note that the Cosmos SDK before store/v2alpha1 does not support this. |
| STORAGE_TYPE_MEMORY | 1 | STORAGE_TYPE_MEMORY indicates in-memory storage that will be reloaded every time an app restarts. Tables with this type of storage will by default be ignored when importing and exporting a module's state from JSON. |
| STORAGE_TYPE_TRANSIENT | 2 | STORAGE_TYPE_TRANSIENT indicates transient storage that is reset at the end of every block. Tables with this type of storage will by default be ignored when importing and exporting a module's state from JSON. |
| STORAGE_TYPE_INDEX | 3 | STORAGE_TYPE_INDEX indicates persistent storage which is not backed by a merkle-tree and won't affect the app hash. Note that the Cosmos SDK before store/v2alpha1 does not support this. |
| STORAGE_TYPE_COMMITMENT | 4 | STORAGE_TYPE_INDEX indicates persistent storage which is backed by a merkle-tree. With this type of storage, both primary and index keys will affect the app hash and this is generally less efficient than using STORAGE_TYPE_DEFAULT_UNSPECIFIED which separates index keys into index storage. Note that modules built with the Cosmos SDK before store/v2alpha1 must specify STORAGE_TYPE_COMMITMENT instead of STORAGE_TYPE_DEFAULT_UNSPECIFIED or STORAGE_TYPE_INDEX because this is the only type of persistent storage available. |


 <!-- end enums -->


<a name="cosmos/orm/v1alpha1/schema.proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| `module_schema` | ModuleSchemaDescriptor | .google.protobuf.MessageOptions | 104503792 | module_schema is used to define the ORM schema for an app module. All module config messages that use module_schema must also declare themselves as app module config messages using the cosmos.app.v1.is_module option. |

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/params/v1beta1/params.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/params/v1beta1/params.proto



<a name="cosmos.params.v1beta1.ParamChange"></a>

### ParamChange
ParamChange defines an individual parameter change, for use in
ParameterChangeProposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  |  |
| `key` | [string](#string) |  |  |
| `value` | [string](#string) |  |  |






<a name="cosmos.params.v1beta1.ParameterChangeProposal"></a>

### ParameterChangeProposal
ParameterChangeProposal defines a proposal to change one or more parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `changes` | [ParamChange](#cosmos.params.v1beta1.ParamChange) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/params/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/params/v1beta1/query.proto



<a name="cosmos.params.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  | subspace defines the module to query the parameter for. |
| `key` | [string](#string) |  | key defines the key of the parameter in the subspace. |






<a name="cosmos.params.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `param` | [ParamChange](#cosmos.params.v1beta1.ParamChange) |  | param defines the queried parameter. |






<a name="cosmos.params.v1beta1.QuerySubspacesRequest"></a>

### QuerySubspacesRequest
QuerySubspacesRequest defines a request type for querying for all registered
subspaces and all keys for a subspace.

Since: cosmos-sdk 0.46






<a name="cosmos.params.v1beta1.QuerySubspacesResponse"></a>

### QuerySubspacesResponse
QuerySubspacesResponse defines the response types for querying for all
registered subspaces and all keys for a subspace.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspaces` | [Subspace](#cosmos.params.v1beta1.Subspace) | repeated |  |






<a name="cosmos.params.v1beta1.Subspace"></a>

### Subspace
Subspace defines a parameter subspace name and all the keys that exist for
the subspace.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `subspace` | [string](#string) |  |  |
| `keys` | [string](#string) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.params.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#cosmos.params.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.params.v1beta1.QueryParamsResponse) | Params queries a specific parameter of a module, given its subspace and key. | GET|/cosmos/params/v1beta1/params|
| `Subspaces` | [QuerySubspacesRequest](#cosmos.params.v1beta1.QuerySubspacesRequest) | [QuerySubspacesResponse](#cosmos.params.v1beta1.QuerySubspacesResponse) | Subspaces queries for all registered subspaces and all keys for a subspace.

Since: cosmos-sdk 0.46 | GET|/cosmos/params/v1beta1/subspaces|

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/slashing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/slashing.proto



<a name="cosmos.slashing.v1beta1.Params"></a>

### Params
Params represents the parameters used for by the slashing module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signed_blocks_window` | [int64](#int64) |  |  |
| `min_signed_per_window` | [bytes](#bytes) |  |  |
| `downtime_jail_duration` | [google.protobuf.Duration](#google.protobuf.Duration) |  |  |
| `slash_fraction_double_sign` | [bytes](#bytes) |  |  |
| `slash_fraction_downtime` | [bytes](#bytes) |  |  |






<a name="cosmos.slashing.v1beta1.ValidatorSigningInfo"></a>

### ValidatorSigningInfo
ValidatorSigningInfo defines a validator's signing info for monitoring their
liveness activity.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |
| `start_height` | [int64](#int64) |  | Height at which validator was first a candidate OR was unjailed |
| `index_offset` | [int64](#int64) |  | Index which is incremented each time the validator was a bonded in a block and may have signed a precommit or not. This in conjunction with the `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`. |
| `jailed_until` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Timestamp until which the validator is jailed due to liveness downtime. |
| `tombstoned` | [bool](#bool) |  | Whether or not a validator has been tombstoned (killed out of validator set). It is set once the validator commits an equivocation or for any other configured misbehiavor. |
| `missed_blocks_counter` | [int64](#int64) |  | A counter kept to avoid unnecessary array reads. Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/genesis.proto



<a name="cosmos.slashing.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the slashing module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.slashing.v1beta1.Params) |  | params defines all the paramaters of related to deposit. |
| `signing_infos` | [SigningInfo](#cosmos.slashing.v1beta1.SigningInfo) | repeated | signing_infos represents a map between validator addresses and their signing infos. |
| `missed_blocks` | [ValidatorMissedBlocks](#cosmos.slashing.v1beta1.ValidatorMissedBlocks) | repeated | missed_blocks represents a map between validator addresses and their missed blocks. |






<a name="cosmos.slashing.v1beta1.MissedBlock"></a>

### MissedBlock
MissedBlock contains height and missed status as boolean.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `index` | [int64](#int64) |  | index is the height at which the block was missed. |
| `missed` | [bool](#bool) |  | missed is the missed status. |






<a name="cosmos.slashing.v1beta1.SigningInfo"></a>

### SigningInfo
SigningInfo stores validator signing info of corresponding address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the validator address. |
| `validator_signing_info` | [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo) |  | validator_signing_info represents the signing info of this validator. |






<a name="cosmos.slashing.v1beta1.ValidatorMissedBlocks"></a>

### ValidatorMissedBlocks
ValidatorMissedBlocks contains array of missed blocks of corresponding
address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the validator address. |
| `missed_blocks` | [MissedBlock](#cosmos.slashing.v1beta1.MissedBlock) | repeated | missed_blocks is an array of missed blocks by the validator. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/query.proto



<a name="cosmos.slashing.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method






<a name="cosmos.slashing.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.slashing.v1beta1.Params) |  |  |






<a name="cosmos.slashing.v1beta1.QuerySigningInfoRequest"></a>

### QuerySigningInfoRequest
QuerySigningInfoRequest is the request type for the Query/SigningInfo RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cons_address` | [string](#string) |  | cons_address is the address to query signing info of |






<a name="cosmos.slashing.v1beta1.QuerySigningInfoResponse"></a>

### QuerySigningInfoResponse
QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `val_signing_info` | [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo) |  | val_signing_info is the signing info of requested val cons address |






<a name="cosmos.slashing.v1beta1.QuerySigningInfosRequest"></a>

### QuerySigningInfosRequest
QuerySigningInfosRequest is the request type for the Query/SigningInfos RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  |  |






<a name="cosmos.slashing.v1beta1.QuerySigningInfosResponse"></a>

### QuerySigningInfosResponse
QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [ValidatorSigningInfo](#cosmos.slashing.v1beta1.ValidatorSigningInfo) | repeated | info is the signing info of all validators |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.slashing.v1beta1.Query"></a>

### Query
Query provides defines the gRPC querier service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#cosmos.slashing.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.slashing.v1beta1.QueryParamsResponse) | Params queries the parameters of slashing module | GET|/cosmos/slashing/v1beta1/params|
| `SigningInfo` | [QuerySigningInfoRequest](#cosmos.slashing.v1beta1.QuerySigningInfoRequest) | [QuerySigningInfoResponse](#cosmos.slashing.v1beta1.QuerySigningInfoResponse) | SigningInfo queries the signing info of given cons address | GET|/cosmos/slashing/v1beta1/signing_infos/{cons_address}|
| `SigningInfos` | [QuerySigningInfosRequest](#cosmos.slashing.v1beta1.QuerySigningInfosRequest) | [QuerySigningInfosResponse](#cosmos.slashing.v1beta1.QuerySigningInfosResponse) | SigningInfos queries signing info of all validators | GET|/cosmos/slashing/v1beta1/signing_infos|

 <!-- end services -->



<a name="cosmos/slashing/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/slashing/v1beta1/tx.proto



<a name="cosmos.slashing.v1beta1.MsgUnjail"></a>

### MsgUnjail
MsgUnjail defines the Msg/Unjail request type


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  |  |






<a name="cosmos.slashing.v1beta1.MsgUnjailResponse"></a>

### MsgUnjailResponse
MsgUnjailResponse defines the Msg/Unjail response type





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.slashing.v1beta1.Msg"></a>

### Msg
Msg defines the slashing Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Unjail` | [MsgUnjail](#cosmos.slashing.v1beta1.MsgUnjail) | [MsgUnjailResponse](#cosmos.slashing.v1beta1.MsgUnjailResponse) | Unjail defines a method for unjailing a jailed validator, thus returning them into the bonded validator set, so they can begin receiving provisions and rewards again. | |

 <!-- end services -->



<a name="cosmos/staking/v1beta1/authz.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/authz.proto



<a name="cosmos.staking.v1beta1.StakeAuthorization"></a>

### StakeAuthorization
StakeAuthorization defines authorization for delegate/undelegate/redelegate.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `max_tokens` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | max_tokens specifies the maximum amount of tokens can be delegate to a validator. If it is empty, there is no spend limit and any amount of coins can be delegated. |
| `allow_list` | [StakeAuthorization.Validators](#cosmos.staking.v1beta1.StakeAuthorization.Validators) |  | allow_list specifies list of validator addresses to whom grantee can delegate tokens on behalf of granter's account. |
| `deny_list` | [StakeAuthorization.Validators](#cosmos.staking.v1beta1.StakeAuthorization.Validators) |  | deny_list specifies list of validator addresses to whom grantee can not delegate tokens. |
| `authorization_type` | [AuthorizationType](#cosmos.staking.v1beta1.AuthorizationType) |  | authorization_type defines one of AuthorizationType. |






<a name="cosmos.staking.v1beta1.StakeAuthorization.Validators"></a>

### StakeAuthorization.Validators
Validators defines list of validator addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) | repeated |  |





 <!-- end messages -->


<a name="cosmos.staking.v1beta1.AuthorizationType"></a>

### AuthorizationType
AuthorizationType defines the type of staking module authorization type

Since: cosmos-sdk 0.43

| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTHORIZATION_TYPE_UNSPECIFIED | 0 | AUTHORIZATION_TYPE_UNSPECIFIED specifies an unknown authorization type |
| AUTHORIZATION_TYPE_DELEGATE | 1 | AUTHORIZATION_TYPE_DELEGATE defines an authorization type for Msg/Delegate |
| AUTHORIZATION_TYPE_UNDELEGATE | 2 | AUTHORIZATION_TYPE_UNDELEGATE defines an authorization type for Msg/Undelegate |
| AUTHORIZATION_TYPE_REDELEGATE | 3 | AUTHORIZATION_TYPE_REDELEGATE defines an authorization type for Msg/BeginRedelegate |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/staking/v1beta1/staking.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/staking.proto



<a name="cosmos.staking.v1beta1.Commission"></a>

### Commission
Commission defines commission parameters for a given validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `commission_rates` | [CommissionRates](#cosmos.staking.v1beta1.CommissionRates) |  | commission_rates defines the initial commission rates to be used for creating a validator. |
| `update_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | update_time is the last time the commission rate was changed. |






<a name="cosmos.staking.v1beta1.CommissionRates"></a>

### CommissionRates
CommissionRates defines the initial commission rates to be used for creating
a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `rate` | [string](#string) |  | rate is the commission rate charged to delegators, as a fraction. |
| `max_rate` | [string](#string) |  | max_rate defines the maximum commission rate which validator can ever charge, as a fraction. |
| `max_change_rate` | [string](#string) |  | max_change_rate defines the maximum daily increase of the validator commission, as a fraction. |






<a name="cosmos.staking.v1beta1.DVPair"></a>

### DVPair
DVPair is struct that just has a delegator-validator pair with no other data.
It is intended to be used as a marshalable pointer. For example, a DVPair can
be used to construct the key to getting an UnbondingDelegation from state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.DVPairs"></a>

### DVPairs
DVPairs defines an array of DVPair objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pairs` | [DVPair](#cosmos.staking.v1beta1.DVPair) | repeated |  |






<a name="cosmos.staking.v1beta1.DVVTriplet"></a>

### DVVTriplet
DVVTriplet is struct that just has a delegator-validator-validator triplet
with no other data. It is intended to be used as a marshalable pointer. For
example, a DVVTriplet can be used to construct the key to getting a
Redelegation from state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_src_address` | [string](#string) |  |  |
| `validator_dst_address` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.DVVTriplets"></a>

### DVVTriplets
DVVTriplets defines an array of DVVTriplet objects.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `triplets` | [DVVTriplet](#cosmos.staking.v1beta1.DVVTriplet) | repeated |  |






<a name="cosmos.staking.v1beta1.Delegation"></a>

### Delegation
Delegation represents the bond with tokens held by an account. It is
owned by one delegator, and is associated with the voting power of one
validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the bech32-encoded address of the delegator. |
| `validator_address` | [string](#string) |  | validator_address is the bech32-encoded address of the validator. |
| `shares` | [string](#string) |  | shares define the delegation shares received. |






<a name="cosmos.staking.v1beta1.DelegationResponse"></a>

### DelegationResponse
DelegationResponse is equivalent to Delegation except that it contains a
balance in addition to shares which is more suitable for client responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation` | [Delegation](#cosmos.staking.v1beta1.Delegation) |  |  |
| `balance` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.Description"></a>

### Description
Description defines a validator description.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `moniker` | [string](#string) |  | moniker defines a human-readable name for the validator. |
| `identity` | [string](#string) |  | identity defines an optional identity signature (ex. UPort or Keybase). |
| `website` | [string](#string) |  | website defines an optional website link. |
| `security_contact` | [string](#string) |  | security_contact defines an optional email for security contact. |
| `details` | [string](#string) |  | details define other optional details. |






<a name="cosmos.staking.v1beta1.HistoricalInfo"></a>

### HistoricalInfo
HistoricalInfo contains header and validator information for a given block.
It is stored as part of staking module's state, which persists the `n` most
recent HistoricalInfo
(`n` is set by the staking module's `historical_entries` parameter).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `header` | [tendermint.types.Header](#tendermint.types.Header) |  |  |
| `valset` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated |  |






<a name="cosmos.staking.v1beta1.Params"></a>

### Params
Params defines the parameters for the staking module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbonding_time` | [google.protobuf.Duration](#google.protobuf.Duration) |  | unbonding_time is the time duration of unbonding. |
| `max_validators` | [uint32](#uint32) |  | max_validators is the maximum number of validators. |
| `max_entries` | [uint32](#uint32) |  | max_entries is the max entries for either unbonding delegation or redelegation (per pair/trio). |
| `historical_entries` | [uint32](#uint32) |  | historical_entries is the number of historical entries to persist. |
| `bond_denom` | [string](#string) |  | bond_denom defines the bondable coin denomination. |
| `min_commission_rate` | [string](#string) |  | min_commission_rate is the chain-wide minimum commission rate that a validator can charge their delegators |






<a name="cosmos.staking.v1beta1.Pool"></a>

### Pool
Pool is used for tracking bonded and not-bonded token supply of the bond
denomination.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `not_bonded_tokens` | [string](#string) |  |  |
| `bonded_tokens` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.Redelegation"></a>

### Redelegation
Redelegation contains the list of a particular delegator's redelegating bonds
from a particular source validator to a particular destination validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the bech32-encoded address of the delegator. |
| `validator_src_address` | [string](#string) |  | validator_src_address is the validator redelegation source operator address. |
| `validator_dst_address` | [string](#string) |  | validator_dst_address is the validator redelegation destination operator address. |
| `entries` | [RedelegationEntry](#cosmos.staking.v1beta1.RedelegationEntry) | repeated | entries are the redelegation entries.

redelegation entries |






<a name="cosmos.staking.v1beta1.RedelegationEntry"></a>

### RedelegationEntry
RedelegationEntry defines a redelegation object with relevant metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creation_height` | [int64](#int64) |  | creation_height defines the height which the redelegation took place. |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | completion_time defines the unix time for redelegation completion. |
| `initial_balance` | [string](#string) |  | initial_balance defines the initial balance when redelegation started. |
| `shares_dst` | [string](#string) |  | shares_dst is the amount of destination-validator shares created by redelegation. |






<a name="cosmos.staking.v1beta1.RedelegationEntryResponse"></a>

### RedelegationEntryResponse
RedelegationEntryResponse is equivalent to a RedelegationEntry except that it
contains a balance in addition to shares which is more suitable for client
responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redelegation_entry` | [RedelegationEntry](#cosmos.staking.v1beta1.RedelegationEntry) |  |  |
| `balance` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.RedelegationResponse"></a>

### RedelegationResponse
RedelegationResponse is equivalent to a Redelegation except that its entries
contain a balance in addition to shares which is more suitable for client
responses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redelegation` | [Redelegation](#cosmos.staking.v1beta1.Redelegation) |  |  |
| `entries` | [RedelegationEntryResponse](#cosmos.staking.v1beta1.RedelegationEntryResponse) | repeated |  |






<a name="cosmos.staking.v1beta1.UnbondingDelegation"></a>

### UnbondingDelegation
UnbondingDelegation stores all of a single delegator's unbonding bonds
for a single validator in an time-ordered list.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  | delegator_address is the bech32-encoded address of the delegator. |
| `validator_address` | [string](#string) |  | validator_address is the bech32-encoded address of the validator. |
| `entries` | [UnbondingDelegationEntry](#cosmos.staking.v1beta1.UnbondingDelegationEntry) | repeated | entries are the unbonding delegation entries.

unbonding delegation entries |






<a name="cosmos.staking.v1beta1.UnbondingDelegationEntry"></a>

### UnbondingDelegationEntry
UnbondingDelegationEntry defines an unbonding object with relevant metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `creation_height` | [int64](#int64) |  | creation_height is the height which the unbonding took place. |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | completion_time is the unix time for unbonding completion. |
| `initial_balance` | [string](#string) |  | initial_balance defines the tokens initially scheduled to receive at completion. |
| `balance` | [string](#string) |  | balance defines the tokens to receive at completion. |






<a name="cosmos.staking.v1beta1.ValAddresses"></a>

### ValAddresses
ValAddresses defines a repeated set of validator addresses.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `addresses` | [string](#string) | repeated |  |






<a name="cosmos.staking.v1beta1.Validator"></a>

### Validator
Validator defines a validator, together with the total amount of the
Validator's bond shares and their exchange rate to coins. Slashing results in
a decrease in the exchange rate, allowing correct calculation of future
undelegations without iterating over delegators. When coins are delegated to
this validator, the validator is credited with a delegation whose number of
bond shares is based on the amount of coins delegated divided by the current
exchange rate. Voting power can be calculated as total bonded shares
multiplied by exchange rate.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `operator_address` | [string](#string) |  | operator_address defines the address of the validator's operator; bech encoded in JSON. |
| `consensus_pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  | consensus_pubkey is the consensus public key of the validator, as a Protobuf Any. |
| `jailed` | [bool](#bool) |  | jailed defined whether the validator has been jailed from bonded status or not. |
| `status` | [BondStatus](#cosmos.staking.v1beta1.BondStatus) |  | status is the validator status (bonded/unbonding/unbonded). |
| `tokens` | [string](#string) |  | tokens define the delegated tokens (incl. self-delegation). |
| `delegator_shares` | [string](#string) |  | delegator_shares defines total shares issued to a validator's delegators. |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  | description defines the description terms for the validator. |
| `unbonding_height` | [int64](#int64) |  | unbonding_height defines, if unbonding, the height at which this validator has begun unbonding. |
| `unbonding_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | unbonding_time defines, if unbonding, the min time for the validator to complete unbonding. |
| `commission` | [Commission](#cosmos.staking.v1beta1.Commission) |  | commission defines the commission parameters. |
| `min_self_delegation` | [string](#string) |  | min_self_delegation is the validator's self declared minimum self delegation.

Since: cosmos-sdk 0.46 |





 <!-- end messages -->


<a name="cosmos.staking.v1beta1.BondStatus"></a>

### BondStatus
BondStatus is the status of a validator.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BOND_STATUS_UNSPECIFIED | 0 | UNSPECIFIED defines an invalid validator status. |
| BOND_STATUS_UNBONDED | 1 | UNBONDED defines a validator that is not bonded. |
| BOND_STATUS_UNBONDING | 2 | UNBONDING defines a validator that is unbonding. |
| BOND_STATUS_BONDED | 3 | BONDED defines a validator that is bonded. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/staking/v1beta1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/genesis.proto



<a name="cosmos.staking.v1beta1.GenesisState"></a>

### GenesisState
GenesisState defines the staking module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.staking.v1beta1.Params) |  | params defines all the paramaters of related to deposit. |
| `last_total_power` | [bytes](#bytes) |  | last_total_power tracks the total amounts of bonded tokens recorded during the previous end block. |
| `last_validator_powers` | [LastValidatorPower](#cosmos.staking.v1beta1.LastValidatorPower) | repeated | last_validator_powers is a special index that provides a historical list of the last-block's bonded validators. |
| `validators` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated | delegations defines the validator set at genesis. |
| `delegations` | [Delegation](#cosmos.staking.v1beta1.Delegation) | repeated | delegations defines the delegations active at genesis. |
| `unbonding_delegations` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) | repeated | unbonding_delegations defines the unbonding delegations active at genesis. |
| `redelegations` | [Redelegation](#cosmos.staking.v1beta1.Redelegation) | repeated | redelegations defines the redelegations active at genesis. |
| `exported` | [bool](#bool) |  |  |






<a name="cosmos.staking.v1beta1.LastValidatorPower"></a>

### LastValidatorPower
LastValidatorPower required for validator set update logic.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the address of the validator. |
| `power` | [int64](#int64) |  | power defines the power of the validator. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/staking/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/query.proto



<a name="cosmos.staking.v1beta1.QueryDelegationRequest"></a>

### QueryDelegationRequest
QueryDelegationRequest is request type for the Query/Delegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryDelegationResponse"></a>

### QueryDelegationResponse
QueryDelegationResponse is response type for the Query/Delegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation_response` | [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse) |  | delegation_responses defines the delegation info of a delegation. |






<a name="cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest"></a>

### QueryDelegatorDelegationsRequest
QueryDelegatorDelegationsRequest is request type for the
Query/DelegatorDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse"></a>

### QueryDelegatorDelegationsResponse
QueryDelegatorDelegationsResponse is response type for the
Query/DelegatorDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation_responses` | [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse) | repeated | delegation_responses defines all the delegations' info of a delegator. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest"></a>

### QueryDelegatorUnbondingDelegationsRequest
QueryDelegatorUnbondingDelegationsRequest is request type for the
Query/DelegatorUnbondingDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse"></a>

### QueryDelegatorUnbondingDelegationsResponse
QueryUnbondingDelegatorDelegationsResponse is response type for the
Query/UnbondingDelegatorDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbonding_responses` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorRequest"></a>

### QueryDelegatorValidatorRequest
QueryDelegatorValidatorRequest is request type for the
Query/DelegatorValidator RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorResponse"></a>

### QueryDelegatorValidatorResponse
QueryDelegatorValidatorResponse response type for the
Query/DelegatorValidator RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [Validator](#cosmos.staking.v1beta1.Validator) |  | validator defines the validator info. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest"></a>

### QueryDelegatorValidatorsRequest
QueryDelegatorValidatorsRequest is request type for the
Query/DelegatorValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse"></a>

### QueryDelegatorValidatorsResponse
QueryDelegatorValidatorsResponse is response type for the
Query/DelegatorValidators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated | validators defines the validators' info of a delegator. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryHistoricalInfoRequest"></a>

### QueryHistoricalInfoRequest
QueryHistoricalInfoRequest is request type for the Query/HistoricalInfo RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | height defines at which height to query the historical info. |






<a name="cosmos.staking.v1beta1.QueryHistoricalInfoResponse"></a>

### QueryHistoricalInfoResponse
QueryHistoricalInfoResponse is response type for the Query/HistoricalInfo RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hist` | [HistoricalInfo](#cosmos.staking.v1beta1.HistoricalInfo) |  | hist defines the historical info at the given height. |






<a name="cosmos.staking.v1beta1.QueryParamsRequest"></a>

### QueryParamsRequest
QueryParamsRequest is request type for the Query/Params RPC method.






<a name="cosmos.staking.v1beta1.QueryParamsResponse"></a>

### QueryParamsResponse
QueryParamsResponse is response type for the Query/Params RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#cosmos.staking.v1beta1.Params) |  | params holds all the parameters of this module. |






<a name="cosmos.staking.v1beta1.QueryPoolRequest"></a>

### QueryPoolRequest
QueryPoolRequest is request type for the Query/Pool RPC method.






<a name="cosmos.staking.v1beta1.QueryPoolResponse"></a>

### QueryPoolResponse
QueryPoolResponse is response type for the Query/Pool RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `pool` | [Pool](#cosmos.staking.v1beta1.Pool) |  | pool defines the pool info. |






<a name="cosmos.staking.v1beta1.QueryRedelegationsRequest"></a>

### QueryRedelegationsRequest
QueryRedelegationsRequest is request type for the Query/Redelegations RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `src_validator_addr` | [string](#string) |  | src_validator_addr defines the validator address to redelegate from. |
| `dst_validator_addr` | [string](#string) |  | dst_validator_addr defines the validator address to redelegate to. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryRedelegationsResponse"></a>

### QueryRedelegationsResponse
QueryRedelegationsResponse is response type for the Query/Redelegations RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `redelegation_responses` | [RedelegationResponse](#cosmos.staking.v1beta1.RedelegationResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryUnbondingDelegationRequest"></a>

### QueryUnbondingDelegationRequest
QueryUnbondingDelegationRequest is request type for the
Query/UnbondingDelegation RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_addr` | [string](#string) |  | delegator_addr defines the delegator address to query for. |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryUnbondingDelegationResponse"></a>

### QueryUnbondingDelegationResponse
QueryDelegationResponse is response type for the Query/UnbondingDelegation
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbond` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) |  | unbond defines the unbonding information of a delegation. |






<a name="cosmos.staking.v1beta1.QueryValidatorDelegationsRequest"></a>

### QueryValidatorDelegationsRequest
QueryValidatorDelegationsRequest is request type for the
Query/ValidatorDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryValidatorDelegationsResponse"></a>

### QueryValidatorDelegationsResponse
QueryValidatorDelegationsResponse is response type for the
Query/ValidatorDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegation_responses` | [DelegationResponse](#cosmos.staking.v1beta1.DelegationResponse) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryValidatorRequest"></a>

### QueryValidatorRequest
QueryValidatorRequest is response type for the Query/Validator RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |






<a name="cosmos.staking.v1beta1.QueryValidatorResponse"></a>

### QueryValidatorResponse
QueryValidatorResponse is response type for the Query/Validator RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [Validator](#cosmos.staking.v1beta1.Validator) |  | validator defines the validator info. |






<a name="cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest"></a>

### QueryValidatorUnbondingDelegationsRequest
QueryValidatorUnbondingDelegationsRequest is required type for the
Query/ValidatorUnbondingDelegations RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_addr` | [string](#string) |  | validator_addr defines the validator address to query for. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse"></a>

### QueryValidatorUnbondingDelegationsResponse
QueryValidatorUnbondingDelegationsResponse is response type for the
Query/ValidatorUnbondingDelegations RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `unbonding_responses` | [UnbondingDelegation](#cosmos.staking.v1beta1.UnbondingDelegation) | repeated |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |






<a name="cosmos.staking.v1beta1.QueryValidatorsRequest"></a>

### QueryValidatorsRequest
QueryValidatorsRequest is request type for Query/Validators RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [string](#string) |  | status enables to query for validators matching a given status. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines an optional pagination for the request. |






<a name="cosmos.staking.v1beta1.QueryValidatorsResponse"></a>

### QueryValidatorsResponse
QueryValidatorsResponse is response type for the Query/Validators RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validators` | [Validator](#cosmos.staking.v1beta1.Validator) | repeated | validators contains all the queried validators. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines the pagination in the response. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.staking.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Validators` | [QueryValidatorsRequest](#cosmos.staking.v1beta1.QueryValidatorsRequest) | [QueryValidatorsResponse](#cosmos.staking.v1beta1.QueryValidatorsResponse) | Validators queries all validators that match the given status. | GET|/cosmos/staking/v1beta1/validators|
| `Validator` | [QueryValidatorRequest](#cosmos.staking.v1beta1.QueryValidatorRequest) | [QueryValidatorResponse](#cosmos.staking.v1beta1.QueryValidatorResponse) | Validator queries validator info for given validator address. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}|
| `ValidatorDelegations` | [QueryValidatorDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorDelegationsRequest) | [QueryValidatorDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorDelegationsResponse) | ValidatorDelegations queries delegate info for given validator. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/delegations|
| `ValidatorUnbondingDelegations` | [QueryValidatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsRequest) | [QueryValidatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryValidatorUnbondingDelegationsResponse) | ValidatorUnbondingDelegations queries unbonding delegations of a validator. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/unbonding_delegations|
| `Delegation` | [QueryDelegationRequest](#cosmos.staking.v1beta1.QueryDelegationRequest) | [QueryDelegationResponse](#cosmos.staking.v1beta1.QueryDelegationResponse) | Delegation queries delegate info for given validator delegator pair. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}|
| `UnbondingDelegation` | [QueryUnbondingDelegationRequest](#cosmos.staking.v1beta1.QueryUnbondingDelegationRequest) | [QueryUnbondingDelegationResponse](#cosmos.staking.v1beta1.QueryUnbondingDelegationResponse) | UnbondingDelegation queries unbonding info for given validator delegator pair. | GET|/cosmos/staking/v1beta1/validators/{validator_addr}/delegations/{delegator_addr}/unbonding_delegation|
| `DelegatorDelegations` | [QueryDelegatorDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorDelegationsRequest) | [QueryDelegatorDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse) | DelegatorDelegations queries all delegations of a given delegator address. | GET|/cosmos/staking/v1beta1/delegations/{delegator_addr}|
| `DelegatorUnbondingDelegations` | [QueryDelegatorUnbondingDelegationsRequest](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsRequest) | [QueryDelegatorUnbondingDelegationsResponse](#cosmos.staking.v1beta1.QueryDelegatorUnbondingDelegationsResponse) | DelegatorUnbondingDelegations queries all unbonding delegations of a given delegator address. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/unbonding_delegations|
| `Redelegations` | [QueryRedelegationsRequest](#cosmos.staking.v1beta1.QueryRedelegationsRequest) | [QueryRedelegationsResponse](#cosmos.staking.v1beta1.QueryRedelegationsResponse) | Redelegations queries redelegations of given address. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/redelegations|
| `DelegatorValidators` | [QueryDelegatorValidatorsRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorsRequest) | [QueryDelegatorValidatorsResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorsResponse) | DelegatorValidators queries all validators info for given delegator address. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators|
| `DelegatorValidator` | [QueryDelegatorValidatorRequest](#cosmos.staking.v1beta1.QueryDelegatorValidatorRequest) | [QueryDelegatorValidatorResponse](#cosmos.staking.v1beta1.QueryDelegatorValidatorResponse) | DelegatorValidator queries validator info for given delegator validator pair. | GET|/cosmos/staking/v1beta1/delegators/{delegator_addr}/validators/{validator_addr}|
| `HistoricalInfo` | [QueryHistoricalInfoRequest](#cosmos.staking.v1beta1.QueryHistoricalInfoRequest) | [QueryHistoricalInfoResponse](#cosmos.staking.v1beta1.QueryHistoricalInfoResponse) | HistoricalInfo queries the historical info for given height. | GET|/cosmos/staking/v1beta1/historical_info/{height}|
| `Pool` | [QueryPoolRequest](#cosmos.staking.v1beta1.QueryPoolRequest) | [QueryPoolResponse](#cosmos.staking.v1beta1.QueryPoolResponse) | Pool queries the pool info. | GET|/cosmos/staking/v1beta1/pool|
| `Params` | [QueryParamsRequest](#cosmos.staking.v1beta1.QueryParamsRequest) | [QueryParamsResponse](#cosmos.staking.v1beta1.QueryParamsResponse) | Parameters queries the staking parameters. | GET|/cosmos/staking/v1beta1/params|

 <!-- end services -->



<a name="cosmos/staking/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/staking/v1beta1/tx.proto



<a name="cosmos.staking.v1beta1.MsgBeginRedelegate"></a>

### MsgBeginRedelegate
MsgBeginRedelegate defines a SDK message for performing a redelegation
of coins from a delegator and source validator to a destination validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_src_address` | [string](#string) |  |  |
| `validator_dst_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgBeginRedelegateResponse"></a>

### MsgBeginRedelegateResponse
MsgBeginRedelegateResponse defines the Msg/BeginRedelegate response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |






<a name="cosmos.staking.v1beta1.MsgCancelUnbondingDelegation"></a>

### MsgCancelUnbondingDelegation
MsgCancelUnbondingDelegation defines the SDK message for performing a cancel unbonding delegation for delegator

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  | amount is always less than or equal to unbonding delegation entry balance |
| `creation_height` | [int64](#int64) |  | creation_height is the height which the unbonding took place. |






<a name="cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse"></a>

### MsgCancelUnbondingDelegationResponse
MsgCancelUnbondingDelegationResponse

Since: cosmos-sdk 0.46






<a name="cosmos.staking.v1beta1.MsgCreateValidator"></a>

### MsgCreateValidator
MsgCreateValidator defines a SDK message for creating a new validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `commission` | [CommissionRates](#cosmos.staking.v1beta1.CommissionRates) |  |  |
| `min_self_delegation` | [string](#string) |  |  |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `pubkey` | [google.protobuf.Any](#google.protobuf.Any) |  |  |
| `value` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgCreateValidatorResponse"></a>

### MsgCreateValidatorResponse
MsgCreateValidatorResponse defines the Msg/CreateValidator response type.






<a name="cosmos.staking.v1beta1.MsgDelegate"></a>

### MsgDelegate
MsgDelegate defines a SDK message for performing a delegation of coins
from a delegator to a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgDelegateResponse"></a>

### MsgDelegateResponse
MsgDelegateResponse defines the Msg/Delegate response type.






<a name="cosmos.staking.v1beta1.MsgEditValidator"></a>

### MsgEditValidator
MsgEditValidator defines a SDK message for editing an existing validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `description` | [Description](#cosmos.staking.v1beta1.Description) |  |  |
| `validator_address` | [string](#string) |  |  |
| `commission_rate` | [string](#string) |  | We pass a reference to the new commission rate and min self delegation as it's not mandatory to update. If not updated, the deserialized rate will be zero with no way to distinguish if an update was intended. REF: #2373 |
| `min_self_delegation` | [string](#string) |  |  |






<a name="cosmos.staking.v1beta1.MsgEditValidatorResponse"></a>

### MsgEditValidatorResponse
MsgEditValidatorResponse defines the Msg/EditValidator response type.






<a name="cosmos.staking.v1beta1.MsgUndelegate"></a>

### MsgUndelegate
MsgUndelegate defines a SDK message for performing an undelegation from a
delegate and a validator.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `delegator_address` | [string](#string) |  |  |
| `validator_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="cosmos.staking.v1beta1.MsgUndelegateResponse"></a>

### MsgUndelegateResponse
MsgUndelegateResponse defines the Msg/Undelegate response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `completion_time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.staking.v1beta1.Msg"></a>

### Msg
Msg defines the staking Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateValidator` | [MsgCreateValidator](#cosmos.staking.v1beta1.MsgCreateValidator) | [MsgCreateValidatorResponse](#cosmos.staking.v1beta1.MsgCreateValidatorResponse) | CreateValidator defines a method for creating a new validator. | |
| `EditValidator` | [MsgEditValidator](#cosmos.staking.v1beta1.MsgEditValidator) | [MsgEditValidatorResponse](#cosmos.staking.v1beta1.MsgEditValidatorResponse) | EditValidator defines a method for editing an existing validator. | |
| `Delegate` | [MsgDelegate](#cosmos.staking.v1beta1.MsgDelegate) | [MsgDelegateResponse](#cosmos.staking.v1beta1.MsgDelegateResponse) | Delegate defines a method for performing a delegation of coins from a delegator to a validator. | |
| `BeginRedelegate` | [MsgBeginRedelegate](#cosmos.staking.v1beta1.MsgBeginRedelegate) | [MsgBeginRedelegateResponse](#cosmos.staking.v1beta1.MsgBeginRedelegateResponse) | BeginRedelegate defines a method for performing a redelegation of coins from a delegator and source validator to a destination validator. | |
| `Undelegate` | [MsgUndelegate](#cosmos.staking.v1beta1.MsgUndelegate) | [MsgUndelegateResponse](#cosmos.staking.v1beta1.MsgUndelegateResponse) | Undelegate defines a method for performing an undelegation from a delegate and a validator. | |
| `CancelUnbondingDelegation` | [MsgCancelUnbondingDelegation](#cosmos.staking.v1beta1.MsgCancelUnbondingDelegation) | [MsgCancelUnbondingDelegationResponse](#cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse) | CancelUnbondingDelegation defines a method for performing canceling the unbonding delegation and delegate back to previous validator.

Since: cosmos-sdk 0.46 | |

 <!-- end services -->



<a name="cosmos/tx/signing/v1beta1/signing.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/signing/v1beta1/signing.proto



<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor"></a>

### SignatureDescriptor
SignatureDescriptor is a convenience type which represents the full data for
a signature including the public key of the signer, signing modes and the
signature itself. It is primarily used for coordinating signatures between
clients.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signer |
| `data` | [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data) |  |  |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence of the account, which describes the number of committed transactions signed by a given address. It is used to prevent replay attacks. |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor.Data"></a>

### SignatureDescriptor.Data
Data represents signature data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `single` | [SignatureDescriptor.Data.Single](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single) |  | single represents a single signer |
| `multi` | [SignatureDescriptor.Data.Multi](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi) |  | multi represents a multisig signer |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Multi"></a>

### SignatureDescriptor.Data.Multi
Multi is the signature data for a multisig public key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bitarray` | [cosmos.crypto.multisig.v1beta1.CompactBitArray](#cosmos.crypto.multisig.v1beta1.CompactBitArray) |  | bitarray specifies which keys within the multisig are signing |
| `signatures` | [SignatureDescriptor.Data](#cosmos.tx.signing.v1beta1.SignatureDescriptor.Data) | repeated | signatures is the signatures of the multi-signature |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptor.Data.Single"></a>

### SignatureDescriptor.Data.Single
Single is the signature data for a single signer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mode` | [SignMode](#cosmos.tx.signing.v1beta1.SignMode) |  | mode is the signing mode of the single signer |
| `signature` | [bytes](#bytes) |  | signature is the raw signature bytes |






<a name="cosmos.tx.signing.v1beta1.SignatureDescriptors"></a>

### SignatureDescriptors
SignatureDescriptors wraps multiple SignatureDescriptor's.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signatures` | [SignatureDescriptor](#cosmos.tx.signing.v1beta1.SignatureDescriptor) | repeated | signatures are the signature descriptors |





 <!-- end messages -->


<a name="cosmos.tx.signing.v1beta1.SignMode"></a>

### SignMode
SignMode represents a signing mode with its own security guarantees.

This enum should be considered a registry of all known sign modes
in the Cosmos ecosystem. Apps are not expected to support all known
sign modes. Apps that would like to support custom  sign modes are
encouraged to open a small PR against this file to add a new case
to this SignMode enum describing their sign mode so that different
apps have a consistent version of this enum.

| Name | Number | Description |
| ---- | ------ | ----------- |
| SIGN_MODE_UNSPECIFIED | 0 | SIGN_MODE_UNSPECIFIED specifies an unknown signing mode and will be rejected. |
| SIGN_MODE_DIRECT | 1 | SIGN_MODE_DIRECT specifies a signing mode which uses SignDoc and is verified with raw bytes from Tx. |
| SIGN_MODE_TEXTUAL | 2 | SIGN_MODE_TEXTUAL is a future signing mode that will verify some human-readable textual representation on top of the binary representation from SIGN_MODE_DIRECT. It is currently not supported. |
| SIGN_MODE_DIRECT_AUX | 3 | SIGN_MODE_DIRECT_AUX specifies a signing mode which uses SignDocDirectAux. As opposed to SIGN_MODE_DIRECT, this sign mode does not require signers signing over other signers' `signer_info`. It also allows for adding Tips in transactions.

Since: cosmos-sdk 0.46 |
| SIGN_MODE_LEGACY_AMINO_JSON | 127 | SIGN_MODE_LEGACY_AMINO_JSON is a backwards compatibility mode which uses Amino JSON and will be removed in the future. |
| SIGN_MODE_EIP_191 | 191 | SIGN_MODE_EIP_191 specifies the sign mode for EIP 191 signing on the Cosmos SDK. Ref: https://eips.ethereum.org/EIPS/eip-191

Currently, SIGN_MODE_EIP_191 is registered as a SignMode enum variant, but is not implemented on the SDK by default. To enable EIP-191, you need to pass a custom `TxConfig` that has an implementation of `SignModeHandler` for EIP-191. The SDK may decide to fully support EIP-191 in the future.

Since: cosmos-sdk 0.45.2 |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/tx/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/v1beta1/tx.proto



<a name="cosmos.tx.v1beta1.AuthInfo"></a>

### AuthInfo
AuthInfo describes the fee and signer modes that are used to sign a
transaction.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer_infos` | [SignerInfo](#cosmos.tx.v1beta1.SignerInfo) | repeated | signer_infos defines the signing modes for the required signers. The number and order of elements must match the required signers from TxBody's messages. The first element is the primary signer and the one which pays the fee. |
| `fee` | [Fee](#cosmos.tx.v1beta1.Fee) |  | Fee is the fee and gas limit for the transaction. The first signer is the primary signer and the one which pays the fee. The fee can be calculated based on the cost of evaluating the body and doing signature verification of the signers. This can be estimated via simulation. |
| `tip` | [Tip](#cosmos.tx.v1beta1.Tip) |  | Tip is the optional tip used for transactions fees paid in another denom.

This field is ignored if the chain didn't enable tips, i.e. didn't add the `TipDecorator` in its posthandler.

Since: cosmos-sdk 0.46 |






<a name="cosmos.tx.v1beta1.AuxSignerData"></a>

### AuxSignerData
AuxSignerData is the intermediary format that an auxiliary signer (e.g. a
tipper) builds and sends to the fee payer (who will build and broadcast the
actual tx). AuxSignerData is not a valid tx in itself, and will be rejected
by the node if sent directly as-is.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | address is the bech32-encoded address of the auxiliary signer. If using AuxSignerData across different chains, the bech32 prefix of the target chain (where the final transaction is broadcasted) should be used. |
| `sign_doc` | [SignDocDirectAux](#cosmos.tx.v1beta1.SignDocDirectAux) |  | sign_doc is the SIGN_MODE_DIRECT_AUX sign doc that the auxiliary signer signs. Note: we use the same sign doc even if we're signing with LEGACY_AMINO_JSON. |
| `mode` | [cosmos.tx.signing.v1beta1.SignMode](#cosmos.tx.signing.v1beta1.SignMode) |  | mode is the signing mode of the single signer. |
| `sig` | [bytes](#bytes) |  | sig is the signature of the sign doc. |






<a name="cosmos.tx.v1beta1.Fee"></a>

### Fee
Fee includes the amount of coins paid in fees and the maximum
gas to be used by the transaction. The ratio yields an effective "gasprice",
which must be above some miminum to be accepted into the mempool.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount is the amount of coins to be paid as a fee |
| `gas_limit` | [uint64](#uint64) |  | gas_limit is the maximum gas that can be used in transaction processing before an out of gas error occurs |
| `payer` | [string](#string) |  | if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees. the payer must be a tx signer (and thus have signed this field in AuthInfo). setting this field does *not* change the ordering of required signers for the transaction. |
| `granter` | [string](#string) |  | if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does not support fee grants, this will fail |






<a name="cosmos.tx.v1beta1.ModeInfo"></a>

### ModeInfo
ModeInfo describes the signing mode of a single or nested multisig signer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `single` | [ModeInfo.Single](#cosmos.tx.v1beta1.ModeInfo.Single) |  | single represents a single signer |
| `multi` | [ModeInfo.Multi](#cosmos.tx.v1beta1.ModeInfo.Multi) |  | multi represents a nested multisig signer |






<a name="cosmos.tx.v1beta1.ModeInfo.Multi"></a>

### ModeInfo.Multi
Multi is the mode info for a multisig public key


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bitarray` | [cosmos.crypto.multisig.v1beta1.CompactBitArray](#cosmos.crypto.multisig.v1beta1.CompactBitArray) |  | bitarray specifies which keys within the multisig are signing |
| `mode_infos` | [ModeInfo](#cosmos.tx.v1beta1.ModeInfo) | repeated | mode_infos is the corresponding modes of the signers of the multisig which could include nested multisig public keys |






<a name="cosmos.tx.v1beta1.ModeInfo.Single"></a>

### ModeInfo.Single
Single is the mode info for a single signer. It is structured as a message
to allow for additional fields such as locale for SIGN_MODE_TEXTUAL in the
future


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `mode` | [cosmos.tx.signing.v1beta1.SignMode](#cosmos.tx.signing.v1beta1.SignMode) |  | mode is the signing mode of the single signer |






<a name="cosmos.tx.v1beta1.SignDoc"></a>

### SignDoc
SignDoc is the type used for generating sign bytes for SIGN_MODE_DIRECT.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body_bytes` | [bytes](#bytes) |  | body_bytes is protobuf serialization of a TxBody that matches the representation in TxRaw. |
| `auth_info_bytes` | [bytes](#bytes) |  | auth_info_bytes is a protobuf serialization of an AuthInfo that matches the representation in TxRaw. |
| `chain_id` | [string](#string) |  | chain_id is the unique identifier of the chain this transaction targets. It prevents signed transactions from being used on another chain by an attacker |
| `account_number` | [uint64](#uint64) |  | account_number is the account number of the account in state |






<a name="cosmos.tx.v1beta1.SignDocDirectAux"></a>

### SignDocDirectAux
SignDocDirectAux is the type used for generating sign bytes for
SIGN_MODE_DIRECT_AUX.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body_bytes` | [bytes](#bytes) |  | body_bytes is protobuf serialization of a TxBody that matches the representation in TxRaw. |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signing account. |
| `chain_id` | [string](#string) |  | chain_id is the identifier of the chain this transaction targets. It prevents signed transactions from being used on another chain by an attacker. |
| `account_number` | [uint64](#uint64) |  | account_number is the account number of the account in state. |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence number of the signing account. |
| `tip` | [Tip](#cosmos.tx.v1beta1.Tip) |  | Tip is the optional tip used for transactions fees paid in another denom. It should be left empty if the signer is not the tipper for this transaction.

This field is ignored if the chain didn't enable tips, i.e. didn't add the `TipDecorator` in its posthandler. |






<a name="cosmos.tx.v1beta1.SignerInfo"></a>

### SignerInfo
SignerInfo describes the public key and signing mode of a single top-level
signer.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `public_key` | [google.protobuf.Any](#google.protobuf.Any) |  | public_key is the public key of the signer. It is optional for accounts that already exist in state. If unset, the verifier can use the required \ signer address for this position and lookup the public key. |
| `mode_info` | [ModeInfo](#cosmos.tx.v1beta1.ModeInfo) |  | mode_info describes the signing mode of the signer and is a nested structure to support nested multisig pubkey's |
| `sequence` | [uint64](#uint64) |  | sequence is the sequence of the account, which describes the number of committed transactions signed by a given address. It is used to prevent replay attacks. |






<a name="cosmos.tx.v1beta1.Tip"></a>

### Tip
Tip is the tip used for meta-transactions.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated | amount is the amount of the tip |
| `tipper` | [string](#string) |  | tipper is the address of the account paying for the tip |






<a name="cosmos.tx.v1beta1.Tx"></a>

### Tx
Tx is the standard type used for broadcasting transactions.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body` | [TxBody](#cosmos.tx.v1beta1.TxBody) |  | body is the processable content of the transaction |
| `auth_info` | [AuthInfo](#cosmos.tx.v1beta1.AuthInfo) |  | auth_info is the authorization related content of the transaction, specifically signers, signer modes and fee |
| `signatures` | [bytes](#bytes) | repeated | signatures is a list of signatures that matches the length and order of AuthInfo's signer_infos to allow connecting signature meta information like public key and signing mode by position. |






<a name="cosmos.tx.v1beta1.TxBody"></a>

### TxBody
TxBody is the body of a transaction that all signers sign over.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `messages` | [google.protobuf.Any](#google.protobuf.Any) | repeated | messages is a list of messages to be executed. The required signers of those messages define the number and order of elements in AuthInfo's signer_infos and Tx's signatures. Each required signer address is added to the list only the first time it occurs. By convention, the first required signer (usually from the first message) is referred to as the primary signer and pays the fee for the whole transaction. |
| `memo` | [string](#string) |  | memo is any arbitrary note/comment to be added to the transaction. WARNING: in clients, any publicly exposed text should not be called memo, but should be called `note` instead (see https://github.com/cosmos/cosmos-sdk/issues/9122). |
| `timeout_height` | [uint64](#uint64) |  | timeout is the block height after which this transaction will not be processed by the chain |
| `extension_options` | [google.protobuf.Any](#google.protobuf.Any) | repeated | extension_options are arbitrary options that can be added by chains when the default options are not sufficient. If any of these are present and can't be handled, the transaction will be rejected |
| `non_critical_extension_options` | [google.protobuf.Any](#google.protobuf.Any) | repeated | extension_options are arbitrary options that can be added by chains when the default options are not sufficient. If any of these are present and can't be handled, they will be ignored |






<a name="cosmos.tx.v1beta1.TxRaw"></a>

### TxRaw
TxRaw is a variant of Tx that pins the signer's exact binary representation
of body and auth_info. This is used for signing, broadcasting and
verification. The binary `serialize(tx: TxRaw)` is stored in Tendermint and
the hash `sha256(serialize(tx: TxRaw))` becomes the "txhash", commonly used
as the transaction ID.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `body_bytes` | [bytes](#bytes) |  | body_bytes is a protobuf serialization of a TxBody that matches the representation in SignDoc. |
| `auth_info_bytes` | [bytes](#bytes) |  | auth_info_bytes is a protobuf serialization of an AuthInfo that matches the representation in SignDoc. |
| `signatures` | [bytes](#bytes) | repeated | signatures is a list of signatures that matches the length and order of AuthInfo's signer_infos to allow connecting signature meta information like public key and signing mode by position. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/tx/v1beta1/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/tx/v1beta1/service.proto



<a name="cosmos.tx.v1beta1.BroadcastTxRequest"></a>

### BroadcastTxRequest
BroadcastTxRequest is the request type for the Service.BroadcastTxRequest
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_bytes` | [bytes](#bytes) |  | tx_bytes is the raw transaction. |
| `mode` | [BroadcastMode](#cosmos.tx.v1beta1.BroadcastMode) |  |  |






<a name="cosmos.tx.v1beta1.BroadcastTxResponse"></a>

### BroadcastTxResponse
BroadcastTxResponse is the response type for the
Service.BroadcastTx method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_response` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) |  | tx_response is the queried TxResponses. |






<a name="cosmos.tx.v1beta1.GetBlockWithTxsRequest"></a>

### GetBlockWithTxsRequest
GetBlockWithTxsRequest is the request type for the Service.GetBlockWithTxs
RPC method.

Since: cosmos-sdk 0.45.2


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | height is the height of the block to query. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | pagination defines a pagination for the request. |






<a name="cosmos.tx.v1beta1.GetBlockWithTxsResponse"></a>

### GetBlockWithTxsResponse
GetBlockWithTxsResponse is the response type for the Service.GetBlockWithTxs method.

Since: cosmos-sdk 0.45.2


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `txs` | [Tx](#cosmos.tx.v1beta1.Tx) | repeated | txs are the transactions in the block. |
| `block_id` | [tendermint.types.BlockID](#tendermint.types.BlockID) |  |  |
| `block` | [tendermint.types.Block](#tendermint.types.Block) |  |  |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | pagination defines a pagination for the response. |






<a name="cosmos.tx.v1beta1.GetTxRequest"></a>

### GetTxRequest
GetTxRequest is the request type for the Service.GetTx
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [string](#string) |  | hash is the tx hash to query, encoded as a hex string. |






<a name="cosmos.tx.v1beta1.GetTxResponse"></a>

### GetTxResponse
GetTxResponse is the response type for the Service.GetTx method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [Tx](#cosmos.tx.v1beta1.Tx) |  | tx is the queried transaction. |
| `tx_response` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) |  | tx_response is the queried TxResponses. |






<a name="cosmos.tx.v1beta1.GetTxsEventRequest"></a>

### GetTxsEventRequest
GetTxsEventRequest is the request type for the Service.TxsByEvents
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `events` | [string](#string) | repeated | events is the list of transaction event type. |
| `pagination` | [cosmos.base.query.v1beta1.PageRequest](#cosmos.base.query.v1beta1.PageRequest) |  | **Deprecated.** pagination defines a pagination for the request. Deprecated post v0.46.x: use page and limit instead. |
| `order_by` | [OrderBy](#cosmos.tx.v1beta1.OrderBy) |  |  |
| `page` | [uint64](#uint64) |  | page is the page number to query, starts at 1. If not provided, will default to first page. |
| `limit` | [uint64](#uint64) |  | limit is the total number of results to be returned in the result page. If left empty it will default to a value to be set by each app. |






<a name="cosmos.tx.v1beta1.GetTxsEventResponse"></a>

### GetTxsEventResponse
GetTxsEventResponse is the response type for the Service.TxsByEvents
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `txs` | [Tx](#cosmos.tx.v1beta1.Tx) | repeated | txs is the list of queried transactions. |
| `tx_responses` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) | repeated | tx_responses is the list of queried TxResponses. |
| `pagination` | [cosmos.base.query.v1beta1.PageResponse](#cosmos.base.query.v1beta1.PageResponse) |  | **Deprecated.** pagination defines a pagination for the response. Deprecated post v0.46.x: use total instead. |
| `total` | [uint64](#uint64) |  | total is total number of results available |






<a name="cosmos.tx.v1beta1.SimulateRequest"></a>

### SimulateRequest
SimulateRequest is the request type for the Service.Simulate
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx` | [Tx](#cosmos.tx.v1beta1.Tx) |  | **Deprecated.** tx is the transaction to simulate. Deprecated. Send raw tx bytes instead. |
| `tx_bytes` | [bytes](#bytes) |  | tx_bytes is the raw transaction.

Since: cosmos-sdk 0.43 |






<a name="cosmos.tx.v1beta1.SimulateResponse"></a>

### SimulateResponse
SimulateResponse is the response type for the
Service.SimulateRPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gas_info` | [cosmos.base.abci.v1beta1.GasInfo](#cosmos.base.abci.v1beta1.GasInfo) |  | gas_info is the information about gas used in the simulation. |
| `result` | [cosmos.base.abci.v1beta1.Result](#cosmos.base.abci.v1beta1.Result) |  | result is the result of the simulation. |





 <!-- end messages -->


<a name="cosmos.tx.v1beta1.BroadcastMode"></a>

### BroadcastMode
BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.

| Name | Number | Description |
| ---- | ------ | ----------- |
| BROADCAST_MODE_UNSPECIFIED | 0 | zero-value for mode ordering |
| BROADCAST_MODE_BLOCK | 1 | BROADCAST_MODE_BLOCK defines a tx broadcasting mode where the client waits for the tx to be committed in a block. |
| BROADCAST_MODE_SYNC | 2 | BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for a CheckTx execution response only. |
| BROADCAST_MODE_ASYNC | 3 | BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns immediately. |



<a name="cosmos.tx.v1beta1.OrderBy"></a>

### OrderBy
OrderBy defines the sorting order

| Name | Number | Description |
| ---- | ------ | ----------- |
| ORDER_BY_UNSPECIFIED | 0 | ORDER_BY_UNSPECIFIED specifies an unknown sorting order. OrderBy defaults to ASC in this case. |
| ORDER_BY_ASC | 1 | ORDER_BY_ASC defines ascending order |
| ORDER_BY_DESC | 2 | ORDER_BY_DESC defines descending order |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.tx.v1beta1.Service"></a>

### Service
Service defines a gRPC service for interacting with transactions.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Simulate` | [SimulateRequest](#cosmos.tx.v1beta1.SimulateRequest) | [SimulateResponse](#cosmos.tx.v1beta1.SimulateResponse) | Simulate simulates executing a transaction for estimating gas usage. | POST|/cosmos/tx/v1beta1/simulate|
| `GetTx` | [GetTxRequest](#cosmos.tx.v1beta1.GetTxRequest) | [GetTxResponse](#cosmos.tx.v1beta1.GetTxResponse) | GetTx fetches a tx by hash. | GET|/cosmos/tx/v1beta1/txs/{hash}|
| `BroadcastTx` | [BroadcastTxRequest](#cosmos.tx.v1beta1.BroadcastTxRequest) | [BroadcastTxResponse](#cosmos.tx.v1beta1.BroadcastTxResponse) | BroadcastTx broadcast transaction. | POST|/cosmos/tx/v1beta1/txs|
| `GetTxsEvent` | [GetTxsEventRequest](#cosmos.tx.v1beta1.GetTxsEventRequest) | [GetTxsEventResponse](#cosmos.tx.v1beta1.GetTxsEventResponse) | GetTxsEvent fetches txs by event. | GET|/cosmos/tx/v1beta1/txs|
| `GetBlockWithTxs` | [GetBlockWithTxsRequest](#cosmos.tx.v1beta1.GetBlockWithTxsRequest) | [GetBlockWithTxsResponse](#cosmos.tx.v1beta1.GetBlockWithTxsResponse) | GetBlockWithTxs fetches a block with decoded txs.

Since: cosmos-sdk 0.45.2 | GET|/cosmos/tx/v1beta1/txs/block/{height}|

 <!-- end services -->



<a name="cosmos/upgrade/v1beta1/upgrade.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/upgrade/v1beta1/upgrade.proto



<a name="cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"></a>

### CancelSoftwareUpgradeProposal
CancelSoftwareUpgradeProposal is a gov Content type for cancelling a software
upgrade.
Deprecated: This legacy proposal is deprecated in favor of Msg-based gov
proposals, see MsgCancelUpgrade.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |






<a name="cosmos.upgrade.v1beta1.ModuleVersion"></a>

### ModuleVersion
ModuleVersion specifies a module and its consensus version.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name of the app module |
| `version` | [uint64](#uint64) |  | consensus version of the app module |






<a name="cosmos.upgrade.v1beta1.Plan"></a>

### Plan
Plan specifies information about a planned upgrade and when it should occur.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | Sets the name for the upgrade. This name will be used by the upgraded version of the software to apply any special "on-upgrade" commands during the first BeginBlock method after the upgrade is applied. It is also used to detect whether a software version can handle a given upgrade. If no upgrade handler with this name has been set in the software, it will be assumed that the software is out-of-date when the upgrade Time or Height is reached and the software will exit. |
| `time` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | **Deprecated.** Deprecated: Time based upgrades have been deprecated. Time based upgrade logic has been removed from the SDK. If this field is not empty, an error will be thrown. |
| `height` | [int64](#int64) |  | The height at which the upgrade must be performed. Only used if Time is not set. |
| `info` | [string](#string) |  | Any application specific upgrade info to be included on-chain such as a git commit that validators could automatically upgrade to |
| `upgraded_client_state` | [google.protobuf.Any](#google.protobuf.Any) |  | **Deprecated.** Deprecated: UpgradedClientState field has been deprecated. IBC upgrade logic has been moved to the IBC module in the sub module 02-client. If this field is not empty, an error will be thrown. |






<a name="cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"></a>

### SoftwareUpgradeProposal
SoftwareUpgradeProposal is a gov Content type for initiating a software
upgrade.
Deprecated: This legacy proposal is deprecated in favor of Msg-based gov
proposals, see MsgSoftwareUpgrade.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `title` | [string](#string) |  |  |
| `description` | [string](#string) |  |  |
| `plan` | [Plan](#cosmos.upgrade.v1beta1.Plan) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/upgrade/v1beta1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/upgrade/v1beta1/query.proto



<a name="cosmos.upgrade.v1beta1.QueryAppliedPlanRequest"></a>

### QueryAppliedPlanRequest
QueryCurrentPlanRequest is the request type for the Query/AppliedPlan RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name is the name of the applied plan to query for. |






<a name="cosmos.upgrade.v1beta1.QueryAppliedPlanResponse"></a>

### QueryAppliedPlanResponse
QueryAppliedPlanResponse is the response type for the Query/AppliedPlan RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | height is the block height at which the plan was applied. |






<a name="cosmos.upgrade.v1beta1.QueryAuthorityRequest"></a>

### QueryAuthorityRequest
QueryAuthorityRequest is the request type for Query/Authority

Since: cosmos-sdk 0.46






<a name="cosmos.upgrade.v1beta1.QueryAuthorityResponse"></a>

### QueryAuthorityResponse
QueryAuthorityResponse is the response type for Query/Authority

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="cosmos.upgrade.v1beta1.QueryCurrentPlanRequest"></a>

### QueryCurrentPlanRequest
QueryCurrentPlanRequest is the request type for the Query/CurrentPlan RPC
method.






<a name="cosmos.upgrade.v1beta1.QueryCurrentPlanResponse"></a>

### QueryCurrentPlanResponse
QueryCurrentPlanResponse is the response type for the Query/CurrentPlan RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `plan` | [Plan](#cosmos.upgrade.v1beta1.Plan) |  | plan is the current upgrade plan. |






<a name="cosmos.upgrade.v1beta1.QueryModuleVersionsRequest"></a>

### QueryModuleVersionsRequest
QueryModuleVersionsRequest is the request type for the Query/ModuleVersions
RPC method.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module_name` | [string](#string) |  | module_name is a field to query a specific module consensus version from state. Leaving this empty will fetch the full list of module versions from state |






<a name="cosmos.upgrade.v1beta1.QueryModuleVersionsResponse"></a>

### QueryModuleVersionsResponse
QueryModuleVersionsResponse is the response type for the Query/ModuleVersions
RPC method.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `module_versions` | [ModuleVersion](#cosmos.upgrade.v1beta1.ModuleVersion) | repeated | module_versions is a list of module names with their consensus versions. |






<a name="cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest"></a>

### QueryUpgradedConsensusStateRequest
QueryUpgradedConsensusStateRequest is the request type for the Query/UpgradedConsensusState
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `last_height` | [int64](#int64) |  | last height of the current chain must be sent in request as this is the height under which next consensus state is stored |






<a name="cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse"></a>

### QueryUpgradedConsensusStateResponse
QueryUpgradedConsensusStateResponse is the response type for the Query/UpgradedConsensusState
RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `upgraded_consensus_state` | [bytes](#bytes) |  | Since: cosmos-sdk 0.43 |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.upgrade.v1beta1.Query"></a>

### Query
Query defines the gRPC upgrade querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CurrentPlan` | [QueryCurrentPlanRequest](#cosmos.upgrade.v1beta1.QueryCurrentPlanRequest) | [QueryCurrentPlanResponse](#cosmos.upgrade.v1beta1.QueryCurrentPlanResponse) | CurrentPlan queries the current upgrade plan. | GET|/cosmos/upgrade/v1beta1/current_plan|
| `AppliedPlan` | [QueryAppliedPlanRequest](#cosmos.upgrade.v1beta1.QueryAppliedPlanRequest) | [QueryAppliedPlanResponse](#cosmos.upgrade.v1beta1.QueryAppliedPlanResponse) | AppliedPlan queries a previously applied upgrade plan by its name. | GET|/cosmos/upgrade/v1beta1/applied_plan/{name}|
| `UpgradedConsensusState` | [QueryUpgradedConsensusStateRequest](#cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateRequest) | [QueryUpgradedConsensusStateResponse](#cosmos.upgrade.v1beta1.QueryUpgradedConsensusStateResponse) | UpgradedConsensusState queries the consensus state that will serve as a trusted kernel for the next version of this chain. It will only be stored at the last height of this chain. UpgradedConsensusState RPC not supported with legacy querier This rpc is deprecated now that IBC has its own replacement (https://github.com/cosmos/ibc-go/blob/2c880a22e9f9cc75f62b527ca94aa75ce1106001/proto/ibc/core/client/v1/query.proto#L54) | GET|/cosmos/upgrade/v1beta1/upgraded_consensus_state/{last_height}|
| `ModuleVersions` | [QueryModuleVersionsRequest](#cosmos.upgrade.v1beta1.QueryModuleVersionsRequest) | [QueryModuleVersionsResponse](#cosmos.upgrade.v1beta1.QueryModuleVersionsResponse) | ModuleVersions queries the list of module versions from state.

Since: cosmos-sdk 0.43 | GET|/cosmos/upgrade/v1beta1/module_versions|
| `Authority` | [QueryAuthorityRequest](#cosmos.upgrade.v1beta1.QueryAuthorityRequest) | [QueryAuthorityResponse](#cosmos.upgrade.v1beta1.QueryAuthorityResponse) | Returns the account with authority to conduct upgrades

Since: cosmos-sdk 0.46 | GET|/cosmos/upgrade/v1beta1/authority|

 <!-- end services -->



<a name="cosmos/upgrade/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/upgrade/v1beta1/tx.proto
Since: cosmos-sdk 0.46


<a name="cosmos.upgrade.v1beta1.MsgCancelUpgrade"></a>

### MsgCancelUpgrade
MsgCancelUpgrade is the Msg/CancelUpgrade request type.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | authority is the address of the governance account. |






<a name="cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse"></a>

### MsgCancelUpgradeResponse
MsgCancelUpgradeResponse is the Msg/CancelUpgrade response type.

Since: cosmos-sdk 0.46






<a name="cosmos.upgrade.v1beta1.MsgSoftwareUpgrade"></a>

### MsgSoftwareUpgrade
MsgSoftwareUpgrade is the Msg/SoftwareUpgrade request type.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `authority` | [string](#string) |  | authority is the address of the governance account. |
| `plan` | [Plan](#cosmos.upgrade.v1beta1.Plan) |  | plan is the upgrade plan. |






<a name="cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse"></a>

### MsgSoftwareUpgradeResponse
MsgSoftwareUpgradeResponse is the Msg/SoftwareUpgrade response type.

Since: cosmos-sdk 0.46





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.upgrade.v1beta1.Msg"></a>

### Msg
Msg defines the upgrade Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SoftwareUpgrade` | [MsgSoftwareUpgrade](#cosmos.upgrade.v1beta1.MsgSoftwareUpgrade) | [MsgSoftwareUpgradeResponse](#cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse) | SoftwareUpgrade is a governance operation for initiating a software upgrade.

Since: cosmos-sdk 0.46 | |
| `CancelUpgrade` | [MsgCancelUpgrade](#cosmos.upgrade.v1beta1.MsgCancelUpgrade) | [MsgCancelUpgradeResponse](#cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse) | CancelUpgrade is a governance operation for cancelling a previously approvid software upgrade.

Since: cosmos-sdk 0.46 | |

 <!-- end services -->



<a name="cosmos/vesting/v1beta1/vesting.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/vesting/v1beta1/vesting.proto



<a name="cosmos.vesting.v1beta1.BaseVestingAccount"></a>

### BaseVestingAccount
BaseVestingAccount implements the VestingAccount interface. It contains all
the necessary fields needed for any vesting account implementation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_account` | [cosmos.auth.v1beta1.BaseAccount](#cosmos.auth.v1beta1.BaseAccount) |  |  |
| `original_vesting` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `delegated_free` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `delegated_vesting` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `end_time` | [int64](#int64) |  |  |






<a name="cosmos.vesting.v1beta1.ContinuousVestingAccount"></a>

### ContinuousVestingAccount
ContinuousVestingAccount implements the VestingAccount interface. It
continuously vests by unlocking coins linearly with respect to time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_vesting_account` | [BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount) |  |  |
| `start_time` | [int64](#int64) |  |  |






<a name="cosmos.vesting.v1beta1.DelayedVestingAccount"></a>

### DelayedVestingAccount
DelayedVestingAccount implements the VestingAccount interface. It vests all
coins after a specific time, but non prior. In other words, it keeps them
locked until a specified time.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_vesting_account` | [BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount) |  |  |






<a name="cosmos.vesting.v1beta1.Period"></a>

### Period
Period defines a length of time and amount of coins that will vest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `length` | [int64](#int64) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.vesting.v1beta1.PeriodicVestingAccount"></a>

### PeriodicVestingAccount
PeriodicVestingAccount implements the VestingAccount interface. It
periodically vests by unlocking coins during each specified period.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_vesting_account` | [BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount) |  |  |
| `start_time` | [int64](#int64) |  |  |
| `vesting_periods` | [Period](#cosmos.vesting.v1beta1.Period) | repeated |  |






<a name="cosmos.vesting.v1beta1.PermanentLockedAccount"></a>

### PermanentLockedAccount
PermanentLockedAccount implements the VestingAccount interface. It does
not ever release coins, locking them indefinitely. Coins in this account can
still be used for delegating and for governance votes even while locked.

Since: cosmos-sdk 0.43


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `base_vesting_account` | [BaseVestingAccount](#cosmos.vesting.v1beta1.BaseVestingAccount) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="cosmos/vesting/v1beta1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## cosmos/vesting/v1beta1/tx.proto



<a name="cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount"></a>

### MsgCreatePeriodicVestingAccount
MsgCreateVestingAccount defines a message that enables creating a vesting
account.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `start_time` | [int64](#int64) |  |  |
| `vesting_periods` | [Period](#cosmos.vesting.v1beta1.Period) | repeated |  |






<a name="cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse"></a>

### MsgCreatePeriodicVestingAccountResponse
MsgCreateVestingAccountResponse defines the Msg/CreatePeriodicVestingAccount
response type.

Since: cosmos-sdk 0.46






<a name="cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccount"></a>

### MsgCreatePermanentLockedAccount
MsgCreatePermanentLockedAccount defines a message that enables creating a permanent
locked account.

Since: cosmos-sdk 0.46


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






<a name="cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccountResponse"></a>

### MsgCreatePermanentLockedAccountResponse
MsgCreatePermanentLockedAccountResponse defines the Msg/CreatePermanentLockedAccount response type.

Since: cosmos-sdk 0.46






<a name="cosmos.vesting.v1beta1.MsgCreateVestingAccount"></a>

### MsgCreateVestingAccount
MsgCreateVestingAccount defines a message that enables creating a vesting
account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `from_address` | [string](#string) |  |  |
| `to_address` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |
| `end_time` | [int64](#int64) |  |  |
| `delayed` | [bool](#bool) |  |  |






<a name="cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse"></a>

### MsgCreateVestingAccountResponse
MsgCreateVestingAccountResponse defines the Msg/CreateVestingAccount response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="cosmos.vesting.v1beta1.Msg"></a>

### Msg
Msg defines the bank Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `CreateVestingAccount` | [MsgCreateVestingAccount](#cosmos.vesting.v1beta1.MsgCreateVestingAccount) | [MsgCreateVestingAccountResponse](#cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse) | CreateVestingAccount defines a method that enables creating a vesting account. | |
| `CreatePermanentLockedAccount` | [MsgCreatePermanentLockedAccount](#cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccount) | [MsgCreatePermanentLockedAccountResponse](#cosmos.vesting.v1beta1.MsgCreatePermanentLockedAccountResponse) | CreatePermanentLockedAccount defines a method that enables creating a permanent locked account.

Since: cosmos-sdk 0.46 | |
| `CreatePeriodicVestingAccount` | [MsgCreatePeriodicVestingAccount](#cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount) | [MsgCreatePeriodicVestingAccountResponse](#cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse) | CreatePeriodicVestingAccount defines a method that enables creating a periodic vesting account.

Since: cosmos-sdk 0.46 | |

 <!-- end services -->



<a name="tendermint/libs/bits/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## tendermint/libs/bits/types.proto



<a name="tendermint.libs.bits.BitArray"></a>

### BitArray



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `bits` | [int64](#int64) |  |  |
| `elems` | [uint64](#uint64) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

