#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  ifeq ($(OS),Windows_NT)
	VERSION := $(shell git describe --exact-match 2>$null)
  else
	VERSION := $(shell git describe --exact-match 2>/dev/null)
  endif
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

LEDGER_ENABLED ?= true
BUILDDIR ?= $(CURDIR)/build
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)
STATIK = $(GOPATH)/bin/statik

export GO111MODULE = on

ifeq ($(OS),Windows_NT)
  BINARYNAME := fxcored.exe
else
  BINARYNAME := fxcored
endif

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (badgerdb,$(findstring badgerdb,$(FX_BUILD_OPTIONS)))
  build_tags += badgerdb
endif

#ifeq (cleveldb,$(findstring cleveldb,$(FX_BUILD_OPTIONS)))
#  build_tags += gcc cleveldb muslc
#endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
BUILD_TAGS_COMMA_SEP := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(BUILD_TAGS_COMMA_SEP)" \
		  -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Name=fxcore \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=fxcored \

ifeq (,$(findstring nostrip,$(FX_BUILD_OPTIONS)))
  ldflags += -w -s
endif

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(FX_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# Check for debug option
ifeq (debug,$(findstring debug,$(FX_BUILD_OPTIONS)))
  BUILD_FLAGS += -gcflags "all=-N -l"
endif

###############################################################################
###                                  Build                                  ###
###############################################################################

all: build lint test

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify
	go mod tidy
	@echo "--> Download go modules to local cache"
	go mod download

build: go.sum
	go build -mod=readonly -v $(BUILD_FLAGS) -o $(BUILDDIR)/bin/$(BINARYNAME) ./cmd/fxcored
	@echo "--> Done building."

build-win:
	@$(MAKE) build

build-linux:
	@GOOS=linux GOARCH=amd64 $(MAKE) build

INSTALL_DIR := $(shell go env GOPATH)/bin
install: build $(INSTALL_DIR)
	mv $(BUILDDIR)/bin/fxcored $(shell go env GOPATH)/bin/fxcored
	@echo "--> Run \"fxcored start\" or \"$(shell go env GOPATH)/bin/fxcored start\" to launch fxcored."

$(INSTALL_DIR):
	@echo "Folder $(INSTALL_DIR) does not exist"
	mkdir -p $@

run-local: install
	@./develop/run_fxcore.sh init

.PHONY: build build-win install docker go.sum run-local

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	@echo "--> Running linter"
	@which golangci-lint > /dev/null || echo "\033[91m install golangci-lint ...\033[0m" && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@which gocyclo > /dev/null || echo "\033[91m install gocyclo ...\033[0m" && go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@which gofumpt > /dev/null || echo "\033[91m install gofumpt ...\033[0m" && go install mvdan.cc/gofumpt@latest
	golangci-lint run -v --go=1.18 --timeout 10m
	find . -name '*.go' -type f -not -path "./build*" -not -name "statik.go" -not -name "*.pb.go" -not -name "*.pb.gw.go" | xargs gocyclo -over 15
	find . -name '*.go' -type f -not -path "./build*" -not -name "statik.go" -not -name "*.pb.go" -not -name "*.pb.gw.go" | xargs gofumpt -d

format: format-goimports
	find . -name '*.go' -type f -not -path "./build*" -not -name "statik.go" -not -name "*.pb.go" -not -name "*.pb.gw.go" | xargs gofumpt -w -l
	golangci-lint run --fix

format-goimports:
	@go install github.com/incu6us/goimports-reviser/v3@latest
	@find . -name '*.go' -type f -not -path './build*' -not -name 'statik.go' -not -name '*.pb.go' -not -name '*.pb.gw.go' -exec goimports-reviser -use-cache -rm-unused {} \;

.PHONY: format lint

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

test:
	@echo "--> Running tests"
	go test -mod=readonly ./...

test-count:
	go test -mod=readonly -cpu 1 -count 1 -cover ./... | grep -v 'types\|cli\|no test files'

.PHONY: test

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=v0.7
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
protoSwaggerVer=0.11.2
protoSwaggerName=ghcr.io/cosmos/proto-builder:$(protoSwaggerVer)
containerProtoGen=$(PROJECT_NAME)-proto-gen-$(protoVer)
containerProtoGenSwagger=$(PROJECT_NAME)-proto-gen-swagger-$(protoVer)
containerProtoFmt=$(PROJECT_NAME)-proto-fmt-$(protoVer)
containerProtoDoc=$(PROJECT_NAME)-proto-doc-$(protoVer)

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --rm --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace/proto bufbuild/buf:1.15.0 \
		format -w; fi

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./develop/protocgen.sh; fi
	@go mod tidy

proto-doc-gen:
	@echo "Generating Protobuf Doc"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoDoc); else docker run --rm --name $(containerProtoDoc) -v $(CURDIR):/workspace --workdir /workspace $(protoSwaggerName) \
    		sh ./develop/protoc-doc-gen.sh; fi

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoSwaggerName) \
		sh ./develop/protoc-swagger-gen.sh; fi

.PHONY: proto-format proto-gen proto-swagger-gen

statik: $(STATIK)
$(STATIK):
	@echo "Installing statik..."
	@go install github.com/rakyll/statik@latest

update-swagger-docs: proto-swagger-gen statik
	$(GOPATH)/bin/statik -src=docs/swagger-ui -dest=docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
	perl -pi -e "print \"host: fx-rest.functionx.io\nschemes:\n  - https\n\" if $$.==6 " ./docs/swagger-ui/swagger.yaml

.PHONY: statik update-swagger-docs

###############################################################################
###                                Releasing                                ###
###############################################################################

PACKAGE_NAME:=github.com/functionx/fx-core/v4
GOLANG_CROSS_VERSION  = v1.18
GOPATH ?= '$(HOME)/go'
release-dry-run:
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v ${GOPATH}/pkg:/go/pkg \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		--rm-dist --skip-validate --skip-publish

release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release --rm-dist --skip-validate

.PHONY: release-dry-run release