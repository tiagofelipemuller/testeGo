# Use a imagem oficial do Go como base
FROM golang:latest

# Crie e defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copie o arquivo go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixe as dependências do Go
RUN go mod download

# Copie todo o código fonte para o diretório de trabalho
COPY . .

# Compile a aplicação Go
RUN go build -o app .

# Comando padrão para iniciar a aplicação quando o contêiner for iniciado
CMD ["./app"]
