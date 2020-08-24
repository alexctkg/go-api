# tdez 
Teste backend Alex Colussi

# Preparação do ambiente

## Postgres Database
Execute o arquivo `script.sql` em seu banco PostgresSQL

## Sem docker
1. Clonar o repositório em `$GOPATH/src` com o nome tdez
2. Entre na pasta `$ cd tdez`
3. Execute o comando `$ go build`
4. Executar o projeto `$ go run main.go`
5. Execute `$ cp .env.example .env`
6. Configure seu .env 


## Com docker
1. Clonar o repositório em `$GOPATH/src` com o nome tdez
2. Execute `$ docker build . -t go-dock`
3. Execute `$ docker run -p 8080:8080 go-dock` 
4. Execute `$ cp .env.example .env`
5. Configure seu .env 



# Documentação
o acesso da documentação é feita através do pacote Swaggo (https://github.com/swaggo/swag). Pode ser acessada através da rota `docs/index.html`
como por exemplo: http://localhost:8080/docs/index.html

user:tdez
password:t10


