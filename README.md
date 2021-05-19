# Stone Challenge

Aqui podemos encontrar a resolução do desafio, onde era necessário criar uma API de transferencia entre contas Internas de um banco digital.


### APIs disponíveis:

#### `/accounts`

- `GET /accounts` - obtém a lista de contas
- `POST /accounts` - cria uma conta
- `GET /accounts/{account_id}` - obtém uma conta baseada no Id
- `UPDATE /accounts/{account_id}` - atualiza informações
- `DELETE /accounts/{account_id}` - deleta uma conta
- `GET /accounts/{account_id}/balance` - obtém o saldo da conta

#### `/login`

- `POST /login` - realiza login gerando um token que expira em 5 minutos (TO-DO colocar account id no payload do token)


#### `/transfers`

- `GET /transfers` - obtém a lista das transferências (TO-DO realizar busca para o usuário logado)



Para subir a aplicação, basta clonar o repositório e rodar o comando de build da docker:

> docker-compose up --build

A aplicação foi configurada para criar a base de dados utilizando Postgre e realizar as migrações de dados necessárias para subir corretamente.
O servidor e o banco ficam em pé até que a docker seja desligada.

