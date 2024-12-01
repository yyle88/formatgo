COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/erero/blob/aacef44379ac6c5e3c831d1df6b47de10d731a88/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
