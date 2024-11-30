# define var
BINARY_NAME=search
GO_FILES=main.go
BIN_DIR=bin

# target platform
PLATFORMS=windows/amd64 darwin/amd64 darwin/arm64 linux/amd64 linux/arm64

all: $(PLATFORMS)

# rule：build different platform
$(PLATFORMS):
	@IFS=/; for platform in $@; do \
		OS=$${platform%%/*}; \
		ARCH=$${platform##*/}; \
		if [ "$$OS" = "windows" ]; then \
			OUTPUT=$(BIN_DIR)/$(BINARY_NAME)_$${OS}_$${ARCH}.exe; \
		else \
			OUTPUT=$(BIN_DIR)/$(BINARY_NAME)_$${OS}_$${ARCH}; \
		fi; \
		echo "Building for $$OS/$$ARCH..."; \
		GOOS=$$OS GOARCH=$$ARCH go build -o "$$OUTPUT" $(GO_FILES); \
	done

# clean all binary file
clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)*

.PHONY: all clean
