VERSION?="0.0.1"
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')


default: bin


dep:
	go mod download

fmt:
	@gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

bin: fmtcheck test
	sh -c "'$(CURDIR)/scripts/build.sh'"

test: fmtcheck
	sh -c "'$(CURDIR)/scripts/test.sh'"

cover:
	sh -c "'$(CURDIR)/scripts/cover.sh'"


.NOTPARALLEL:

.PHONY: bin fmt fmtcheck test cover dep