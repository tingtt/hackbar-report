GO ?= go
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
MODULE_NAME ?= $(shell head -n1 go.mod | cut -f 2 -d ' ')

.PHONY: test
test:
	$(GO) test ./... -parallel 10

APP_NAME ?= $(MODULE_NAME)
BUILD_DIR ?= .build/$(GOOS)-$(GOARCH)

.PHONY: build
build:
	mkdir -p $(BUILD_DIR)/
	if [ $(GOOS) = "windows" ] ; then \
		GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build \
			-o .binary cmd/report/main.go ; \
		mv .binary ./$(BUILD_DIR)/$(APP_NAME).exe ; \
	else \
		GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build \
			-o .binary cmd/report/main.go ; \
		mv .binary ./$(BUILD_DIR)/$(APP_NAME) ; \
	fi ;

TAG	?=	$(shell git tag | tail -n1)
PACKAGE_NAME ?= $(APP_NAME)-$(TAG).$(GOOS)-$(GOARCH)
PACKAGE_DIR ?= ./.packages/$(TAG)/$(PACKAGE_NAME)

.PHONY: package
package:
	mkdir -p $(PACKAGE_DIR)/
	cp README.md \
		$(PACKAGE_DIR)/
	if [ $(GOOS) = "windows" ] ; then \
		cp ./$(BUILD_DIR)/$(APP_NAME).exe $(PACKAGE_DIR)/ ; \
	else \
		cp ./$(BUILD_DIR)/$(APP_NAME) $(PACKAGE_DIR)/ ; \
	fi
	cd ./.packages/$(TAG) ; \
	if [ $(GOOS) = "windows" ] ; then \
		zip -r $(PACKAGE_NAME).zip ./$(PACKAGE_NAME) ; \
	else \
		tar cvf $(PACKAGE_NAME).tar.gz ./$(PACKAGE_NAME) ; \
	fi ; \
	rm -r ./$(PACKAGE_NAME)

.PHONY: build-and-package-all
build-and-package-all:
	$(GO) tool dist list | grep 'aix\|darwin\|freebsd\|illumos\|linux\|netbsd\|openbsd\|windows' | while read line ; \
	do \
		printf GOOS= > ./.build.env ; \
		echo $$line | cut -f 1 -d "/" >> ./.build.env ; \
		printf GOARCH= >> ./.build.env ; \
		echo $$line | cut -f 2 -d "/" >> ./.build.env ; \
		. ./.build.env ; \
		make build GOOS=$$GOOS GOARCH=$$GOARCH ; \
		make package GOOS=$$GOOS GOARCH=$$GOARCH ; \
	done
	rm ./.build.env

.PHONY: clean
clean:
	-rm -r .binary .build/ .packages/ .build.env