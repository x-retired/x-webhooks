DOCKER       = docker
VERSION      = 0.1.0
DOCKER_IMAGE = xiexianbin/webhooks
DOCKER_RUN   = $(DOCKER) run --rm --interactive --tty --volume $(CURDIR):/src

.PHONY: all build build-preview help serve

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

all: dev ## run dev model

build: ## Build docker image
	$(DOCKER) build . --tag ${DOCKER_IMAGE}:$(VERSION)

dev: ## run dev model
	$(DOCKER_RUN) $(DOCKER_IMAGE):$(VERSION) bee run
