📡 Projeto gRPC com Go e SQLite
Este projeto demonstra uma implementação simples de um serviço gRPC em Go, utilizando SQLite como banco de dados e Evans como cliente CLI para testes interativos.

📦 Requisitos
Go 1.19+

protoc (Protocol Buffers Compiler)

SQLite

Evans (CLI gRPC)

🗃️ Script do Banco de Dados
Para criar a tabela de categorias manualmente, execute:

bash
Copiar
Editar
sqlite3 db.sqlite
CREATE TABLE categories (id string, name string, description string);
🚀 Executando o Servidor
Abra dois terminais:

🖥️ Terminal 1 - Iniciar o servidor gRPC
bash
Copiar
Editar
go run cmd/grpcServer/main.go
🧪 Terminal 2 - Acessar o cliente Evans
bash
Copiar
Editar
evans -r repl
Configurar o pacote e serviço:

bash
Copiar
Editar
> package pb
> service CategoryService
✍️ Criar Categoria
Chame o método CreateCategory e preencha os campos:

plaintext
Copiar
Editar
name (TYPE_STRING) => meu name
description (TYPE_STRING) => my descrip
Resposta esperada:

json
Copiar
Editar
{
  "category": {
    "description": "my descrip",
    "id": "ae026098-6d9a-47a2-a31b-db4789d820e9",
    "name": "meu name"
  }
}
🛠️ Gerar arquivos .proto
Compile os arquivos .proto:

bash
Copiar
Editar
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
📚 Exemplos Práticos com Evans
No terminal Evans:

bash
Copiar
Editar
> package pb
> service CategoryService
✅ Criar nova categoria:
Chame o método CreateCategory e preencha os campos:

plaintext
Copiar
Editar
name (TYPE_STRING) => name 2
description (TYPE_STRING) => desc 2
Resposta esperada:

json
Copiar
Editar
{
  "category": {
    "description": "desc 2",
    "id": "22c832ad-11e6-4113-ab5f-fe78fc6dfbb4",
    "name": "name 2"
  }
}
📋 Listar todas as categorias:
Chame o método ListCategories:

bash
Copiar
Editar
call ListCategories
Resposta esperada:

json
Copiar
Editar
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
🔍 Buscar categoria por ID:
Chame o método GetCategory e preencha com o ID da categoria:

plaintext
Copiar
Editar
id (TYPE_STRING) => 22c832ad-11e6-4113-ab5f-fe78fc6dfbb4
Resposta esperada:

json
Copiar
Editar
{
  "category": {
    "description": "desc 2",
    "id": "22c832ad-11e6-4113-ab5f-fe78fc6dfbb4",
    "name": "name 2"
  }
}
📡 Streaming de Categorias
Este projeto também demonstra o uso de gRPC com Server Streaming, através do método CreateCategoryStream.

O que faz?
Permite enviar várias categorias via streaming do cliente para o servidor. Ao finalizar o envio, o servidor retorna a lista completa de categorias criadas.

Esse padrão é útil em situações como:

Upload em lote de dados

Comunicação contínua entre serviços

Redução de overhead em múltiplas requisições

🧪 Como testar no Evans:
Chame o método CreateCategoryStream e preencha múltiplas categorias:

plaintext
Copiar
Editar
name (TYPE_STRING) => Categoria A
description (TYPE_STRING) => Primeira

name (TYPE_STRING) => Categoria B
description (TYPE_STRING) => Segunda
Finalize com Ctrl+D (EOF).

Resposta esperada:

json
Copiar
Editar
{
  "categories": [
    {
      "id": "...",
      "name": "Categoria A",
      "description": "Primeira"
    },
    {
      "id": "...",
      "name": "Categoria B",
      "description": "Segunda"
    }
  ]
}
📂 Estrutura de Pastas
plaintext
Copiar
Editar
.
├── cmd/
│   └── grpcServer/
│       └── main.go
├── internal/
│   ├── pb/             # Código gerado do .proto
│   ├── database/       # Acesso ao SQLite
│   └── services/       # Implementação dos serviços
├── proto/
│   └── course_category.proto
├── db.sqlite
├── go.mod
└── go.sum
