# Bookstore API

A concurrent RESTful API for an online bookstore built with Go. The API manages books, authors, customers, orders, and generates periodic sales reports. \
(*In final project*)

## Features

- **Book Management**: CRUD operations and search functionality
- **Author Management**: Complete author lifecycle with book relationship handling
- **Customer Management**: Customer profile handling
- **Order Processing**: Order creation with stock management
- **Sales Reports**: Automatic daily sales report generation
- **Data Persistence**: All data is persisted to JSON files
- **Concurrent Operations**: Handles multiple requests safely
- **Request Timeouts**: Implements context-based timeout handling

## Project Structure

```
bookstore/
├── handlers/     # HTTP request handlers
├── models/           # Data models
├── store/            # Data storage implementations
├── service/         # Report Generation
├── data/                 # Persisted data files
├── output-reports/       # Generated sales reports
└── main.go              # Application entry point
```


## Installation

1. Clone the repository


2. Install dependencies:
```bash
go mod download
```

## Running the Application

1. Create required directories:
```bash
mkdir data output-reports
```

2. Start the server:
```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Books
- `GET /books` - List/search books
- `GET /books/{id}` - Get a specific book
- `POST /books` - Create a new book
- `PUT /books/{id}` - Update a book
- `DELETE /books/{id}` - Delete a book

#### Book Search Parameters
- `title`: Search by book title
- `author`: Search by author name
- `genres`: Filter by genres (comma-separated)
- `min_price`: Minimum price
- `max_price`: Maximum price
- `in_stock`: Filter by stock availability (true/false)

### Authors
- `GET /authors` - List all authors
- `GET /authors/{id}` - Get a specific author
- `POST /authors` - Create a new author
- `PUT /authors/{id}` - Update an author
- `DELETE /authors/{id}` - Delete an author (fails if author has books)

### Customers
- `GET /customers/{id}` - Get a customer
- `POST /customers` - Create a new customer
- `PUT /customers/{id}` - Update a customer
- `DELETE /customers/{id}` - Delete a customer (fails if customer has orders)

### Orders
- `GET /orders` - List all orders
- `GET /orders/{id}` - Get an order
- `POST /orders` - Create a new order
- `DELETE /orders/{id}` - Delete an order

### Reports
- `GET /reports` - List all reports
- `GET /reports?start_date=2024-01-01&end_date=2024-01-31` - List reports in date range
- `GET /reports/{id}` - Get a specific report

## API Examples

### Creating an Author
```http
POST /authors
Content-Type: application/json

{
    "first_name": "first test",
    "last_name": "last test",
    "bio": "test bio"
}
```

### Creating a Book
```http
POST /books
Content-Type: application/json

{
    "title": "GO",
    "author": {
        "id": 1
    },
    "genres": ["Fiction", "Adventure"],
    "published_at": "2024-01-15T00:00:00Z",
    "price": 29.00,
    "stock": 100
}
```

### Creating a Customer
```http
POST /customers
Content-Type: application/json

{
    "name": "test cust",
    "email": "cust@example.com",
    "address": {
        "street": "test street",
        "city": "test city",
        "state": "test state",
        "postal_code": "12345",
        "country": "test country"
    }
}
```

### Creating an Order
```http
POST /orders
Content-Type: application/json

{
    "customer": {
        "id": 1,
        "name": "test cust",
        "email": "cust@example.com",
        "address": {
            "street": "test street",
            "city": "test city",
            "state": "test state",
            "postal_code": "12345",
            "country": "test country"
        }
    },
    "items": [
        {
            "book": {
                "id": 1,
                "title": "world",
                "author": {
                    "id": 2,
                    "first_name": "B",
                    "last_name": "B",
                    "bio": "B"
                },
                "genres": [
                    "history",
                    "fantasy"
                ],
                "published_at": "2006-01-02T15:04:05Z",
                "price": 40.8,
                "stock": 70
            },
            "quantity": 2
        }
    ]
}
```

### Searching Books
```http
GET /books?title=great&min_price=20&max_price=50&in_stock=true
```

### Getting Reports
```http
GET /reports?start_date=2025-01-01&end_date=2025-01-31
```

## Data Persistence

All data is automatically saved to JSON files in the `data/` directory:
- `books.json`
- `authors.json`
- `customers.json`
- `orders.json`

Sales reports are saved in the `output-reports/` directory with timestamps in their filenames.

## Error Handling

The API returns appropriate HTTP status codes and error messages:
- 200: Success
- 201: Resource created
- 400: Bad request / Invalid input
- 404: Resource not found
- 500: Server error

## Logging

The application logs all HTTP requests including:
- Request method and path
- Response status
- Request duration

## Features Implementation Details

### Data Persistence
- Data is automatically saved to JSON files when modified
- Data is loaded when the server starts
- Each entity type has its own file

### Concurrency Handling
- Thread-safe operations using mutexes
- Request timeout handling using contexts
- Safe concurrent access to shared resources

### Sales Reports
- Generated daily
- Include total revenue and order counts
- List top-selling books
- Filterable by date range
