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



Para subir a aplicação, basta clonar o repositório e rodar o comando de build da docker:

> docker-compose up --build

