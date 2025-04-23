## 📦 gRPC com Go

API gRPC para gerenciamento de categorias, com chamadas simples e streaming, escrita em Go com banco SQLite.

---

### 🔗 Bibliotecas utilizadas

- [`google.golang.org/grpc`](https://pkg.go.dev/google.golang.org/grpc) — Framework gRPC para Go  
- [`evans`](https://github.com/ktr0731/evans) — Cliente gRPC interativo via linha de comando  
- [`protoc`](https://grpc.io/docs/protoc-installation/) — Compilador Protocol Buffers

---

### 🗂️ Estrutura do projeto

```
.
├── cmd/
│   └── grpcServer/      # Entrypoint do servidor gRPC
├── database/            # Lógica de persistência com SQLite
├── pb/                  # Arquivos gerados via protoc
├── proto/               # Arquivos .proto com definição da API
└── go.mod / go.sum
```

---

### ✨ Instruções para rodar o projeto

**1. Criar banco SQLite:**
```bash
sqlite3 db.sqlite
```

**2. Criar tabela `categories`:**
```sql
CREATE TABLE categories (
  id TEXT,
  name TEXT,
  description TEXT
);
```

**3. Gerar os arquivos gRPC:**
```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```

**4. Iniciar o servidor gRPC:**
```bash
go run cmd/grpcServer/main.go
```

---

### 🧪 Usando o cliente Evans

Abra outro terminal:

```bash
evans -r repl
```

#### Selecionar serviço e pacote

```bash
package pb
service CategoryService
```

---

### ✅ Exemplo: Criar categoria

```bash
call CreateCategory
name (TYPE_STRING) => GoLang
description (TYPE_STRING) => Linguagem do Google
```

**Resposta:**
```json
{
  "category": {
    "id": "uuid-gerado",
    "name": "GoLang",
    "description": "Linguagem do Google"
  }
}
```

---

### 📄 Listar todas as categorias

```bash
call ListCategories
```

**Resposta:**
```json
{
  "categories": [
    {
      "id": "...",
      "name": "...",
      "description": "..."
    }
  ]
}
```

---

### 🔍 Buscar categoria por ID

```bash
call GetCategory
id (TYPE_STRING) => <id-da-categoria>
```

---

### 🌊 Streaming: Criar streaming de categorias `CategoryStream`

```bash
call CategoryStream
```

```bash
name (TYPE_STRING) => GoLang
description (TYPE_STRING) => Linguagem do Google
name (TYPE_STRING) => GoLang
description (TYPE_STRING) => Linguagem do Google
name (TYPE_STRING) => GoLang
description (TYPE_STRING) => Linguagem do Google
```

> Pressione Ctrl + D para encerrar o streaming.

**Resposta:**
```json
{
  "category": {
    "id": "ae026098-6d9a-47a2-a31b-db4789d820e9",
    "name": "meu name",
    "description": "my descrip"
  }
}
{
  "category": {
    "id": "22c832ad-11e6-4113-ab5f-fe78fc6dfbb4",
    "name": "name 2",
    "description": "desc 2"
  }
}
```

> A resposta retorna um stream contínuo de categorias. Cada item é recebido separadamente, ideal para casos com grande volume de dados.

---

## Gerar os arquivos Go a partir do arquivo .proto:

```bash
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```


