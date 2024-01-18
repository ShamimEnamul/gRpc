BIN_DIR := bin
PROTO_DIR := proto
SERVER_DIR := server
CLIENT_DIR := client

# Detect the operating system and set shell-related variables
ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	RM_F_CMD := Remove-Item -ErrorAction SilentlyContinue -Force
	RM_RF_CMD := ${RM_F_CMD} -Recurse
	SERVER_BIN := ${SERVER_DIR}.exe
	CLIENT_BIN := ${CLIENT_DIR}.exe
else
	SHELL := bash
	RM_F_CMD := rm -f
	RM_RF_CMD := ${RM_F_CMD} -r
	SERVER_BIN := ${SERVER_DIR}
	CLIENT_BIN := ${CLIENT_DIR}
endif

.DEFAULT_GOAL := help
.PHONY: greet calculator blog help clean clean_greet clean_calculator clean_blog rebuild test

# List of projects
PROJECTS := greet #calculator blog

# Targets to generate protobuf files and build binaries for each project
all: $(PROJECTS)

$(PROJECTS):
	@${CHECK_DIR_CMD}
	protoc -I$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. $@/${PROTO_DIR}/*.proto
	go build -o ${BIN_DIR}/$@/${SERVER_BIN} ./$@/${SERVER_DIR}
	go build -o ${BIN_DIR}/$@/${CLIENT_BIN} ./$@/${CLIENT_DIR}

# Launch tests for all projects
test:
	go test ./...

# Clean generated files and binaries
clean: $(addprefix clean_, $(PROJECTS))
	${RM_F_CMD} ssl/*.crt
	${RM_F_CMD} ssl/*.csr
	${RM_F_CMD} ssl/*.key
	${RM_F_CMD} ssl/*.pem
	${RM_RF_CMD} ${BIN_DIR}

# Clean generated files for a specific project
clean_%:
	${RM_F_CMD} $*/${PROTO_DIR}/*.pb.go

# Rebuild the entire project
rebuild: clean all

# Update package versions
bump: all
	go get -u ./...

# Display information related to the build
about:
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

# Show help
help:
	@${HELP_CMD}
