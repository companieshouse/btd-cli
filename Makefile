bin ?= btd-cli

.EXPORT_ALL_VARIABLES:
GO111MODULE = on

.PHONY: all
all: clean build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build:
	go build

.PHONY: test
test: test-unit test-integration

.PHONY: test-unit
test-unit:
	$(warning $@ unimplemented)

.PHONY: test-integration
test-integration:
	$(warning $@ unimplemented)

.PHONY: lint
lint:
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck -checks all ./...

.PHONY: package
package:
ifndef version
	$(error No 'version' value was specified; aborting)
endif
	zip $(bin)-$(version).zip $(bin)

.PHONY: clean
clean:
	$(RM) $(bin) *.zip
