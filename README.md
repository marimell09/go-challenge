# Stone Challenge

Aqui podemos encontrar a resolução do desafio, onde era necessário criar uma API de transferencia entre contas internas de um banco digital.

### APIs disponíveis:

#### `/accounts`

- `GET /accounts` - obtém a lista de contas
- `POST /accounts` - cria uma conta
- `GET /accounts/{account_id}` - obtém uma conta baseada no id
- `UPDATE /accounts/{account_id}` - atualiza informações
- `DELETE /accounts/{account_id}` - deleta uma conta
- `GET /accounts/{account_id}/balance` - obtém o saldo da conta

#### `/login`

- `POST /login` - realiza login gerando um token que expira em 5 minutos

#### `/transfers`

- `GET /transfers` - obtém a lista das transferências do usuário logado
- `POST /transfers` - realiza uma nova transferência entre contas

Obs: A rota de transferência só deixa realizar o request corretamente caso haja saldo suficiente na conta e exista o id de destino da conta exista.
Caso uma dessas informações não seja passada corretamente, um erro é exibido.

# Docker
Para subir a aplicação, basta clonar o repositório e rodar o comando de build da docker:

> docker-compose up --build

A aplicação foi configurada para criar a base de dados utilizando Postgre e realizar as migrações de dados necessárias para subir corretamente.
Como o compose tenta subir as aplicações ao mesmo tempo, a migração e o servidor restartam caso ocorra algum erro na conexão com a base de dados.
O servidor e o banco ficam em pé até que a docker seja desligada.

[Postman Collection] (https://www.getpostman.com/collections/c7bcc0395f386133692b)

### TODO

- Testes unitários
- Refatoração do design
- Refatoração de erros
- Melhorias de encapsulamento
- Melhorias no docker compose para não expor dados do banco
- Configurar Swagger para rotas existentes
