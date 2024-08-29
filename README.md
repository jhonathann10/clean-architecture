# clean-architecture
Desafio de clean-architecture full cycle.

1- Subir o docker-compose:
```bash
docker-compose up -d
```

2- Executar o migrate up através do Makefile para gerar a tabela.
```bash
make migrateup
```

3- Inicializar o serviço:
```bash
cd cmd/ordersystem
go run main.go wire_gen.go
```
4- Realizar request para adicionar um dado:

Web:
- Utilizar os arquivos `api/create_order.http` e `api/get_order.http`.

GraphQL:
```graphql
mutation createOrder {
  createOrder(input: {id:"EEEE", Price: 99.0, Tax: 11}){
    id
    Price
    Tax
    FinalPrice
  }
}

query queryOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

gRPC:
```grpc
evans --proto internal/infra/grpc/protofiles/order.proto --host localhost --port 50051
=> call CreateOrder
=> 2
=> 10.5
=> 0.5

=> call ListOrders
```

5- Para comprovar que tudo funcionou com o esperado, abrir o terminal interativo dentro do contêiner do serviço mysql:
```bash
docker-compose exec mysql bash
```

6- Acessar o MySQL e digitar a senha do banco:
```bash
mysql -uroot -p orders
```


