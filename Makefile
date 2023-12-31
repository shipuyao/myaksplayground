REGISTRY ?= shipudemo.azurecr.io
IMAGE_NAME := myaksdocs
IMAGE_VERSION ?= latest

DOCS_IMAGE := $(REGISTRY)/$(IMAGE_NAME):$(IMAGE_VERSION)

# The architecture of the image
ARCH ?= amd64

.PHONY: build
build:
	docker build \
		--platform="linux/$(ARCH)" \
		--tag=$(DOCS_IMAGE) .

.PHONY: push
push:
	docker push $(DOCS_IMAGE)
