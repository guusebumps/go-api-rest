# Usa a imagem oficial do Golang
FROM golang:1.19

# Define o diretório de trabalho
WORKDIR /app

# Copia todos os arquivos do projeto
COPY . .

# Baixa as dependências
RUN go mod tidy

# Compila o aplicativo
RUN go build -o main .

# Expõe a porta 8080
EXPOSE 8080

# Comando para rodar o app
CMD ["/app/main"]
