# GoXpress 

This project was created to help Node.js developers adapt to the **Go** ecosystem. As a Node developer transitioning to Go, this repository provides structured comments and patterns that relate **Express.js** concepts to their Gin framework equivalents, making the learning curve smoother and more intuitive.


### **Go** + **Express** = **GoXpress** 


## Features

- **Authentication & Authorization** — JWT token-based auth with role-based middleware (admin/user)
- **Product Management** — Full CRUD operations with admin-only write access
- **Shopping Cart** — Add, update, remove items with automatic price calculation
- **Order Processing** — Checkout with atomic stock updates and transaction management

## Architecture

<img width="1341" height="872" alt="image" src="https://github.com/user-attachments/assets/91b5c0db-b1ab-454c-9c01-2507675fd34a" />


The project follows a clean three-layer architecture pattern:

```
Handler Layer (HTTP) → Service Layer (Business Logic) → Repository Layer (Database)
```

### Modules

| Module | Description |
|--------|-------------|
| **Auth** | User registration, login, role management |
| **Product** | Product catalog with inventory tracking |
| **Cart** | Shopping cart operations |
| **Order** | Checkout and order processing |

## Tech Stack

| Category | Technology |
|----------|------------|
| **Language** | Go 1.x |
| **Framework** | Gin |
| **Database** | PostgreSQL |
| **ORM** | GORM |
| **Authentication** | JWT |
| **Password Hashing** | bcrypt |

## Prerequisites

- Go 1.x or higher
- PostgreSQL database

## Getting Started

### Installation

1. **Clone the repository**

```bash
git clone <repository-url>
cd goxpress
```

2. **Set up environment variables**

Create a `.env` file in the root directory:

```env
DATABASE_CONFIG=host=localhost user=postgres password=yourpassword dbname=goxpress port=5432 sslmode=disable
APP_PORT=8080
JWT_SECRET=your_jwt_secret_key
```

3. **Install dependencies**

```bash
go mod download
```

4. **Run the application**

```bash
go run cmd/main.go
```

The server will start at `http://localhost:8080`

## API Endpoints

### Authentication

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| `POST` | `/api/v1/auth/login` | User login | Public |
| `POST` | `/admin/register` | Admin registration | Public |
| `POST` | `/user/register` | User registration | Public |

#### Login Request

```json
{
  "username": "john_doe",
  "password": "password123"
}
```

#### Register Request

```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

### Products

| Method | Endpoint | Description | Access |
|--------|----------|-------------|--------|
| `GET` | `/api/v1/product` | List all products | Public |
| `GET` | `/api/v1/product/:id` | Get product by ID | Public |
| `POST` | `/api/v1/product` | Create product | Admin only |
| `PUT` | `/api/v1/product/:id` | Update product | Admin only |
| `DELETE` | `/api/v1/product/:id` | Delete product | Admin only |

#### Create/Update Product Request

```json
{
  "name": "Product Name",
  "description": "Product description",
  "price": 10000,
  "stock": 50
}
```

### Cart

All cart endpoints require user authentication.

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/cart` | Get user's cart |
| `POST` | `/api/v1/cart/add` | Add item to cart |
| `PUT` | `/api/v1/cart/item/:id` | Update cart item |
| `DELETE` | `/api/v1/cart/item/:id` | Remove cart item |

#### Add to Cart Request

```json
{
  "product_id": 1,
  "quantity": 2
}
```

#### Update Cart Item Request

```json
{
  "quantity": 5
}
```

### Orders

All order endpoints require user authentication.

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/v1/order/checkout` | Process checkout |

## Database Models

### Product

```go
type Product struct {
    ID          uint
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   time.Time
    Name        string
    Stock       int
    Price       float64
    Description string
}
```

### Transaction

```go
type Transaction struct {
    ID         uint
    UserID     uint
    TotalPrice float64
    Payment    Payment
    OrderItems []OrderItem
}
```

### OrderItem

```go
type OrderItem struct {
    ID            uint
    TransactionID uint
    ProductID     uint
    Quantity      int
    PriceAtTime   float64
}
```

### Cart

```go
type Cart struct {
    ID        uint
    UserID    uint
    CartItems []CartItem
}
```

### CartItem

```go
type CartItem struct {
    ID        uint
    CartID    uint
    ProductID uint
    Quantity  int
    Price     float64
}
```

## Project Structure

```
goxpress/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── auth/                # Authentication module
│   │   ├── handlers/        # HTTP handlers
│   │   ├── services/        # Business logic
│   │   ├── repository/      # Data access layer
│   │   └── auth.go          # Route setup
│   ├── product/             # Product module
│   │   ├── handlers/
│   │   ├── services/
│   │   ├── repository/
│   │   ├── model/           # Product model
│   │   └── product.go
│   ├── cart/                # Cart module
│   │   ├── handlers/
│   │   ├── services/
│   │   ├── repository/
│   │   ├── model/
│   │   └── cart.go
│   ├── order/               # Order module
│   │   ├── handlers/
│   │   ├── services/
│   │   ├── repository/
│   │   ├── model/
│   │   └── order.go
│   ├── db/                  # Database connection
│   │   └── db.go
│   └── middleware/          # Auth middleware
│       └── middleware.go
├── pkg/
│   └── util/                # Utility functions
│       └── jwt.go           # JWT helper
├── .env                     # Environment variables
├── go.mod                   # Go module file
└── go.sum                   # Go dependencies
```

## Security Features

### JWT Authentication

- JWT tokens are stored in HTTP-only cookies for security
- Tokens include user role information for authorization
- Token expiration is enforced

### Password Security

- Passwords are hashed using bcrypt with default cost
- Plain text passwords are never stored in the database

### Role-Based Access Control

The application implements middleware for role-based authorization:

```go
middleware.RoleMiddleware("admin")  // Admin-only routes
middleware.RoleMiddleware("user")   // User-only routes
```

### Protected Routes

- Product creation, updates, and deletion require admin role
- Cart operations require user authentication
- Order processing requires user authentication

## Architecture Details

### Three-Layer Pattern

**Handler Layer**
- Receives HTTP requests
- Validates request data
- Calls service layer methods
- Returns HTTP responses

**Service Layer**
- Contains business logic
- Coordinates between repositories
- Handles transactions
- Implements validation rules

**Repository Layer**
- Direct database operations
- GORM query implementations
- Data persistence logic

### Cross-Module Dependencies

The Cart and Order modules consume Product repository services for:
- Price lookups during cart operations
- Stock validation before checkout
- Atomic stock updates during order processing

### Bootstrap Process

All modules are initialized during application startup:

1. Load environment variables
2. Connect to PostgreSQL database
3. Initialize Gin router
4. Set up module repositories
5. Wire services with dependencies
6. Register HTTP routes
7. Start server

## Development

### Running Tests

```bash
go test ./...
```

### Building for Production

```bash
go build -o goxpress cmd/main.go
./goxpress
```

### Database Migrations

The application uses GORM's AutoMigrate feature. Models are automatically migrated when the application starts.


---

**Built with Go and Gin Framework**
