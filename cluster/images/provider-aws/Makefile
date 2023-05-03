# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/imagelight.mk

# ====================================================================================
# XPKG

-include ../../../build/makelib/xpkg.mk

-include ../../../build/makelib/k8s_tools.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) Family base image to build: $(IMAGE)
	$(INFO) Building image $${IMAGE}; \
	$(MAKE) BUILD_ARGS="--load" IMAGE=$${IMAGE} XPKG_REG_ORGS=$(XPKG_REG_ORGS) img.build.shared; \
	$(OK) docker build $${IMAGE};

# make BRANCH_NAME=main SUBPACKAGES=cloudplatform publish
img.publish: $(UP)
	@$(INFO) Subpackages to publish: $(SUBPACKAGES)
	@for g in $(SUBPACKAGES); do \
		PACKAGE=$(XPKG_REG_ORGS)/$(PROJECT_NAME)-$${g}:$(VERSION); \
		if [[ $${g} == "monolith" ]]; then PACKAGE=$(XPKG_REG_ORGS)/$(PROJECT_NAME):$(VERSION); fi; \
		$(INFO) Pushing package $${PACKAGE}; \
		$(UP) xpkg push $${PACKAGE} \
				$(foreach p,$(XPKG_LINUX_PLATFORMS),--package $(XPKG_OUTPUT_DIR)/$(p)/$(PROJECT_NAME)-$${g}-$(VERSION).xpkg ) || $(FAIL); \
		$(OK) Done processing $${PACKAGE}. Moving to the next package...; \
	done || $(FAIL)

ifneq (,$(filter $(APIGROUP),monolith config))
    ARG_XPKG_AUTH_EXT := "XPKG_AUTH_EXT=$(ROOT_DIR)/package/auth.yaml"
endif

DEP_CONSTRAINT := >= 0.0.0
ifeq (-,$(findstring -,$(VERSION)))
    DEP_CONSTRAINT = >= 0.0.0-0
endif

img.build.shared: $(UP)
	@cp Dockerfile $(IMAGE_TEMP_DIR) || $(FAIL)
	@cp terraformrc.hcl $(IMAGE_TEMP_DIR) || $(FAIL)
	@docker buildx build $(BUILD_ARGS) \
		--platform $(IMAGE_PLATFORMS) \
		--build-arg TERRAFORM_VERSION=$(TERRAFORM_VERSION) \
		--build-arg TERRAFORM_PROVIDER_SOURCE=$(TERRAFORM_PROVIDER_SOURCE) \
		--build-arg TERRAFORM_PROVIDER_VERSION=$(TERRAFORM_PROVIDER_VERSION) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_NAME=$(TERRAFORM_PROVIDER_DOWNLOAD_NAME) \
		--build-arg TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX=$(TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX) \
		--build-arg CROSSPLANE_PROVIDER_VERSION=$(VERSION) \
		-t $(IMAGE) \
		$(IMAGE_TEMP_DIR) || $(FAIL)
	# @$(MAKE) UP=$(UP) XPKG_REG_ORGS=$(XPKG_REG_ORGS) XPKG_REG_ORGS_NO_PROMOTE=$(XPKG_REG_ORGS_NO_PROMOTE) XPKGS=$(PROJECT_NAME)-$(APIGROUP) XPKG_DIR=$(IMAGE_TEMP_DIR)/package XPKG_EXAMPLES_DIR=$(IMAGE_TEMP_DIR)/examples $(ARG_XPKG_AUTH_EXT) build.xpkg || $(FAIL)

build.xpkg: do.build.xpkgs

img.promote:
	@$(INFO) Skipping image promotion from $(FROM_IMAGE) to $(TO_IMAGE)
	@echo Promote is deferred to xpkg machinery
	@$(OK) Image promotion skipped for $(FROM_IMAGE) to $(TO_IMAGE)
