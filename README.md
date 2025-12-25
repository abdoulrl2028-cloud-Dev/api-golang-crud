# API Go CRUD

Uma API RESTful completa para operações CRUD (Create, Read, Update, Delete) de usuários desenvolvida em Go.

## Estrutura do Projeto

```
api-golang-crud/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── handler/                 # Handlers HTTP
│   ├── service/                 # Lógica de negócio
│   ├── repository/              # Acesso a dados
│   ├── model/                   # Estruturas de dados
│   └── db/                      # Conexão com banco de dados
├── migrations/                  # Migrações do banco de dados
├── .env.example                 # Exemplo de configuração
├── go.mod                       # Dependências do projeto
└── README.md                    # Documentação
```

## Requisitos

- Go 1.21 ou superior
- PostgreSQL 12 ou superior
- Git

## Configuração

### 1. Clone o repositório

```bash
git clone https://github.com/abdoulrl2028-cloud-Dev/api-golang-crud.git
cd api-golang-crud
```

### 2. Configure as variáveis de ambiente

Copie o arquivo `.env.example` para `.env`:

```bash
cp .env.example .env
```

Edite o arquivo `.env` com suas configurações:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=crud_api
API_PORT=:8080
```

### 3. Crie o banco de dados

```bash
createdb crud_api
```

### 4. Execute as migrações

```bash
psql -U postgres -d crud_api -f migrations/001_create_users_table.sql
```

### 5. Instale as dependências

```bash
go mod download
```

## Rodando a Aplicação

```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## Endpoints da API

### Health Check
- **GET** `/health` - Verifica o status da API

### Usuários
- **GET** `/users` - Lista todos os usuários
- **POST** `/users` - Cria um novo usuário
- **GET** `/users/{id}` - Obtém um usuário por ID
- **PUT** `/users/{id}` - Atualiza um usuário
- **DELETE** `/users/{id}` - Deleta um usuário

## Exemplos de Requisições

### Criar um usuário

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "phone": "(11) 99999-9999"
  }'
```

### Listar todos os usuários

```bash
curl http://localhost:8080/users
```

### Obter um usuário específico

```bash
curl http://localhost:8080/users/1
```

### Atualizar um usuário

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva Atualizado",
    "email": "joao.novo@example.com",
    "phone": "(11) 98888-8888"
  }'
```

### Deletar um usuário

```bash
curl -X DELETE http://localhost:8080/users/1
```

## Arquitetura

A aplicação segue a arquitetura de camadas:

- **Handler (Apresentação)**: Recebe requisições HTTP e retorna respostas
- **Service (Lógica de Negócio)**: Implementa as regras de negócio
- **Repository (Acesso a Dados)**: Interage com o banco de dados
- **Model (Entidades)**: Define as estruturas de dados

## Dependências

- `gorilla/mux` - Roteador HTTP
- `lib/pq` - Driver PostgreSQL
- `joho/godotenv` - Gerenciador de variáveis de ambiente

## Estrutura de Resposta

Todas as respostas da API seguem este formato:

```json
{
  "success": true,
  "message": "Descrição da operação",
  "data": {
    "id": 1,
    "name": "João Silva",
    "email": "joao@example.com",
    "phone": "(11) 99999-9999",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z"
  },
  "error": null
}
```

## Códigos de Status HTTP

- `200 OK` - Requisição bem-sucedida
- `201 Created` - Recurso criado com sucesso
- `400 Bad Request` - Requisição inválida
- `404 Not Found` - Recurso não encontrado
- `500 Internal Server Error` - Erro no servidor

## Desenvolvimento

### Estrutura de Pastas Explicada

- **cmd/api/** - Contém o arquivo main.go que é o ponto de entrada da aplicação
- **internal/handler/** - Define os handlers HTTP que recebem as requisições
- **internal/service/** - Implementa a lógica de negócio da aplicação
- **internal/repository/** - Define como os dados são persistidos no banco
- **internal/model/** - Define as estruturas de dados (User, etc)
- **internal/db/** - Gerencia a conexão com o banco de dados

### Padrão de Projeto

A aplicação usa o padrão de **Dependency Injection** para facilitar testes e manutenção.

## Testes

Para rodar os testes:

```bash
go test ./...
```

## Troubleshooting

### Erro de conexão com banco de dados

Verifique se:
- PostgreSQL está rodando
- As credenciais em `.env` estão corretas
- O banco de dados foi criado

### Erro de dependências

Execute:
```bash
go mod tidy
```

## Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## Autor

**abdoulrl2028-cloud-Dev**

---

Desenvolvido com ❤️ em Go
