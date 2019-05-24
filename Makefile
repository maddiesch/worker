GOCMD := go

.PHONY: test
test:
	$(GOCMD) test -v ./...

.PHONY: watch
watch:
	rerun --clear --no-notify --pattern {**/*.go} -x $(MAKE) test
