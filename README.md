# scripts BD

> sqlite3 db.sqlite

> Create TABLE categories (id string, name string, description string);

# grpc

[shell 1]
> go run cmd/grpcServer/main.go

[shell 2]
> evans -r repl (console grpc client)

> package pb

> service CategoryService

> call CreateCategory

´´´
name (TYPE_STRING) => meu name
description (TYPE_STRING) => my descrip
{
  "category": {
    "description": "my descrip",
    "id": "ae026098-6d9a-47a2-a31b-db4789d820e9",
    "name": "meu name"
  }
}
´´´


[Gera o course_category_grpc.pb.go]
> protoc --go_out=. --go-grpc_out=. proto/course_category.proto


Exemplo de uso do Evans para criar nova categoria e listar

evans -r repl

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


127.0.0.1:50051> package pb

pb@127.0.0.1:50051> service CategoryService

pb.CategoryService@127.0.0.1:50051> call CreateCategory
name (TYPE_STRING) => name 2
description (TYPE_STRING) => desc 2
{
  "category": {
    "description": "desc 2",
    "id": "22c832ad-11e6-4113-ab5f-fe78fc6dfbb4",
    "name": "name 2"
  }
}

// Lista Categories 

pb.CategoryService@127.0.0.1:50051> call ListCategories
{
  "categories": [
    {
      "description": "my descrip",
      "id": "ae026098-6d9a-47a2-a31b-db4789d820e9",
      "name": "meu name"
    },
    {
      "description": "desc 2",
      "id": "22c832ad-11e6-4113-ab5f-fe78fc6dfbb4",
      "name": "name 2"
    }
  ]
}

// Get Category By ID

pb.CategoryService@127.0.0.1:50051> call GetCategory
id (TYPE_STRING) => 22c832ad-11e6-4113-ab5f-fe78fc6dfbb4
{
  "category": {
    "description": "desc 2",
    "id": "22c832ad-11e6-4113-ab5f-fe78fc6dfbb4",
    "name": "name 2"
  }
}