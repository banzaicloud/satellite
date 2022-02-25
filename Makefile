# Project variables
PACKAGE = github.com/banzaicloud/satellite
BINARY_NAME ?= satellite
DOCKER_IMAGE ?= banzaicloud/$(BINARY_NAME)
TAG ?= dev-$(shell git log -1 --pretty=format:"%h")

LD_FLAGS = -X "main.version=$(TAG)"
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS=$(shell go list ./... | grep -v /vendor)

# Build variables
BUILD_DIR ?= build
VERSION ?= $(shell git symbolic-ref -q --short HEAD || git describe --tags --exact-match)
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
LDFLAGS += -X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH} -X main.buildDate=${BUILD_DATE}
export GO111MODULE=on

# Dependency versions
GOLANGCI_VERSION = 1.40.0
GOLANG_VERSION = 1.17

.PHONY: _no-target-specified
_no-target-specified:
	$(error Please specify the target to make - `make list` shows targets.)

.PHONY: list
list:
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

all: clean fmt vet docker push

clean: ## Clean the working area and the project
	rm -rf bin/ ${BUILD_DIR}/ vendor/

fmt:
	@gofmt -w ${GOFILES_NOVENDOR}

vet:
	@go vet -composites=false ./...

docker:
	docker build --rm -t $(DOCKER_IMAGE):$(TAG) .

push:
	docker push $(DOCKER_IMAGE):$(TAG)

build: ## Build a binary
ifeq (${VERBOSE}, 1)
	go env
endif
ifneq (${IGNORE_GOLANG_VERSION_REQ}, 1)
	@printf "${GOLANG_VERSION}\n$$(go version | awk '{sub(/^go/, "", $$3);print $$3}')" | sort -t '.' -k 1,1 -k 2,2 -k 3,3 -g | head -1 | grep -q -E "^${GOLANG_VERSION}$$" || (printf "Required Go version is ${GOLANG_VERSION}\nInstalled: `go version`" && exit 1)
endif

	@$(eval GENERATED_BINARY_NAME = ${BINARY_NAME})
	@$(if $(strip ${BINARY_NAME_SUFFIX}),$(eval GENERATED_BINARY_NAME = ${BINARY_NAME}-$(subst $(eval) ,-,$(strip ${BINARY_NAME_SUFFIX}))),)
	go build ${GOARGS} -tags "${GOTAGS}" -ldflags "${LDFLAGS}" -o ${BUILD_DIR}/${GENERATED_BINARY_NAME} ${PACKAGE}

build-all: check-fmt check-misspell lint vet build

check-fmt:
	PKGS="${GOFILES_NOVENDOR}" GOFMT="gofmt" ./scripts/fmt-check.sh

check-misspell: install-misspell
	PKGS="${GOFILES_NOVENDOR}" MISSPELL="misspell" ./scripts/misspell-check.sh

misspell: install-misspell
	misspell -w ${GOFILES_NOVENDOR}

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

.PHONY: lint
lint: bin/golangci-lint ## Run linter
	bin/golangci-lint run
