# Variáveis
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOIMPORTS=goimports
GOFMT=gofmt
GOLINT=golint

# Nome do executável
BINARY_NAME=main

# Diretório onde estão os arquivos do projeto
SRC_DIR="/home/lucaspereira/projects/rest-api/cmd"

# Comando para organizar imports
imports:
	$(GOIMPORTS) -w $(SRC_DIR)/

# Comando para organizar o código
format:
	$(GOFMT) -w $(SRC_DIR)/

# Comando para verificar o estilo do código
lint:
	$(GOLINT) ./...

# Comando para instalar dependências
deps:
	$(GOGET) -u ./...

# Comando para executar testes
test:
	$(GOTEST) -v ./...

# Comando para compilar o projeto
build:
	$(GOBUILD) -o $(BINARY_NAME) $(SRC_DIR)/...

# Comando para executar o projeto
run:
	$(GOBUILD) -o $(BINARY_NAME) $(SRC_DIR)/...
	./$(BINARY_NAME)

# Comando para limpar arquivos compilados
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Comando padrão ao digitar apenas `make` no terminal
default: build
