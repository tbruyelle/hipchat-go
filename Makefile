SRC_DIR=./hipchat
VETARGS=-asmdecl -atomic -bool -buildtags -copylocks -methods -nilfunc -printf -rangeloops -shift -structtags -unsafeptr

all: test testrace vet

default: test

# test runs the unit tests and vets the code
test:
	go test -v $(SRC_DIR) $(TESTARGS) -timeout=30s -parallel=4

# testrace runs the race checker
testrace:
	go test -race $(SRC_DIR) $(TESTARGS)

# vet runs the Go source code static analysis tool `vet` to find
# any common errors
vet:
	@go tool vet 2>/dev/null ; if [ $$? -eq 3 ]; then \
		go get golang.org/x/tools/cmd/vet; \
	fi
	@echo "go tool vet $(VETARGS) $(SRC_DIR) "
	@go tool vet $(VETARGS) $(SRC_DIR) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: all default test vet
