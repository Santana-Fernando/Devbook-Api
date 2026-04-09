# 🚀 DevBook API (Go)

API desenvolvida em Go para gerenciamento de usuários, seguidores e publicações (estilo rede social).

---

## 📋 Pré-requisitos

Antes de rodar o projeto, você precisa ter instalado:

* [Go](https://go.dev/) (versão 1.20+ recomendada)
* [Git](https://git-scm.com/)
* Banco de dados (PostgreSQL recomendado)

---

## ⚙️ Configuração do ambiente

### 🔹 1. Clonar o repositório

```bash
git clone https://github.com/seu-usuario/devbook-api.git
cd devbook-api
```

---

### 🔹 2. Criar arquivo `.env`

Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=devbook

API_PORT=5000
JWT_SECRET=secreta
```

---

### 🔹 3. Instalar dependências

```bash
go mod tidy
```

---

## 🗄️ Banco de dados

Certifique-se de que o banco está rodando e que as tabelas existem:

* usuarios
* seguidores
* publicacoes

> Você pode usar migrations ou criar manualmente via SQL.

---

## ▶️ Rodando a aplicação

```bash
go run main.go
```

Ou, se estiver dentro de uma pasta `cmd/api`:

```bash
go run cmd/api/main.go
```

---

## 🧪 Testando a API

Você pode usar:

* Postman
* Insomnia
* curl

Exemplo:

```bash
curl http://localhost:5000/publicacoes
```

---

## 🏗️ Build da aplicação

Para gerar um executável:

```bash
go build -o app
```

Executar:

```bash
./app
```

---

## 📁 Estrutura do projeto (exemplo)

```
📦 src
 ┣ 📂 controllers
 ┣ 📂 repositorios
 ┣ 📂 modelos
 ┣ 📂 seguranca
 ┗ 📂 middlewares
```

---

## 🔐 Variáveis importantes

| Variável    | Descrição        |
| ----------- | ---------------- |
| DB_HOST     | Host do banco    |
| DB_USER     | Usuário do banco |
| DB_PASSWORD | Senha do banco   |
| DB_NAME     | Nome do banco    |
| API_PORT    | Porta da API     |
| JWT_SECRET  | Chave JWT        |

---

## 🚀 Tecnologias utilizadas

* Go (Golang)
* GORM (ORM)
* PostgreSQL
* JWT para autenticação

---

## 💡 Observações

* Não subir o arquivo `.env` para o repositório
* Use `.env.example` como base
* Configure corretamente o banco antes de rodar

---

## 📌 Autor

Desenvolvido por Fernando Rodrigues 👨‍💻
