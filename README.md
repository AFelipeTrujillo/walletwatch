# WalletWatch API

A simple expense tracker API built with Go, Gin, and GORM.

## Features

- Create, read, update, and delete expenses
- Categorize expenses
- Track spending by category
- SQLite database for easy setup

## Tech Stack

- **Go** - Programming language
- **Gin** - Web framework
- **GORM** - ORM library
- **SQLite** - Database

## Getting Started

1. Clone the repository
```bash
git clone https://github.com/AFelipeTrujillo/walletwatch.git
cd walletwatch
```

2. Install dependencies
```bash
go mod download
```

3. Set up environment variables
```bash
cp .env.example .env
```

4. Run the application
```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

- `GET /api/expenses` - Get all expenses
- `GET /api/expenses/:id` - Get expense by ID
- `POST /api/expenses` - Create new expense
- `PUT /api/expenses/:id` - Update expense
- `DELETE /api/expenses/:id` - Delete expense
- `GET /api/expenses/stats` - Get spending statistics by category

## License

MIT