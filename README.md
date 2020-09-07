# Preparação do ambiente

## Postgres Database
Execute o arquivo `script.sql` em seu banco PostgresSQL

## Sem docker
1. Clone o repositório  no `$GOPATH/src` 
2. Entre na pasta `$ cd tdez`
3. Executar o projeto `$ go run main.go`
4. Execute `$ cp .env.example .env`
5. Configure seu .env 


## Com docker
1. Clone o repositório `$GOPATH/src` 
2. Execute `$ docker build . -t go-dock`
3. Execute `$ docker run -p 8080:8080 go-dock` 
4. Execute `$ cp .env.example .env`
5. Configure seu .env 



# Documentação
o acesso da documentação é feita através do pacote Swaggo (https://github.com/swaggo/swag). Pode ser acessada através da rota `docs/index.html`
como por exemplo: http://localhost:8080/docs/index.html

user:tdez
password:t10


# Observações
Como falei no primeiro contato que tive com a recrutadora ainda não estou habituado 
com TDD, uma habilidade que devo desenvolver. Além do mais, não desenvolvi um
service de email como demonstrado no model.jpg, pois não encontrei uma api free para
que pudesse implementar.
