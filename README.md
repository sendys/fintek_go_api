# Product CRUD API with Image Upload

A complete product management system built with Go, Gin Framework, GORM, and MySQL featuring full CRUD operations and image upload functionality.

## Features

- ✅ **Complete CRUD Operations** - Create, Read, Update, Delete products
- ✅ **Image Upload** - Upload and manage product images
- ✅ **Authentication** - JWT-based authentication for protected endpoints
- ✅ **Pagination** - Efficient pagination for product listings
- ✅ **Filtering & Search** - Filter by category, status, and search functionality
- ✅ **File Management** - Automatic image cleanup and validation
- ✅ **Database Migration** - Auto-migration with GORM
- ✅ **Input Validation** - Comprehensive request validation
- ✅ **Error Handling** - Proper error responses and logging

## Tech Stack

- **Backend**: Go 1.24.4
- **Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: MySQL
- **Authentication**: JWT
- **File Upload**: Multipart form handling
- **UUID**: Google UUID for unique identifiers

## Project Structure

```
backend/
├── config/
│   └── database.go          # Database configuration
├── controllers/
│   ├── auth_controller.go   # Authentication controller
│   ├── user_controller.go   # User management
│   └── product_controller.go # Product CRUD operations
├── middlewares/
│   └── jwt_middleware.go    # JWT authentication middleware
├── models/
│   ├── user.go             # User model
│   └── product.go          # Product model and DTOs
├── routes/
│   ├── auth_routes.go      # Authentication routes
│   ├── user_routes.go      # User routes
│   └── product_routes.go   # Product routes
├── utils/
│   ├── token.go            # JWT token utilities
│   └── image_upload.go     # Image upload utilities
├── uploads/
│   └── products/           # Product images storage
├── docs/
│   └── PRODUCT_API.md      # API documentation
├── main.go                 # Application entry point
├── test_product_api.go     # API testing script
└── README.md              # This file
```

## Installation & Setup

### Prerequisites

- Go 1.24.4 or higher
- MySQL 5.7 or higher
- Git

### 1. Clone the Repository

```bash
git clone <repository-url>
cd backend
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Database Setup

1. Create a MySQL database:
```sql
CREATE DATABASE fintek_shared;
```

2. Update database configuration in `config/database.go`:
```go
dsn := "root:bismillah@tcp(127.0.0.1:3306)/fintek_shared?charset=utf8mb4&parseTime=True&loc=Local"
```

### 4. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8081`

### 5. Database Migration

The application automatically creates the required tables on startup:
- `users` table
- `products` table

## API Endpoints

### Public Endpoints (No Authentication Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products with pagination |
| GET | `/products/{id}` | Get product by ID |
| GET | `/products/categories` | Get all product categories |
| GET | `/uploads/products/{filename}` | Access uploaded images |

### Protected Endpoints (Authentication Required)

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/products` | Create new product |
| PUT | `/products/{id}` | Update product |
| DELETE | `/products/{id}` | Delete product |
| POST | `/products/{id}/image` | Upload product image |

## Usage Examples

### 1. Create a Product

```bash
curl -X POST http://localhost:8081/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token" \
  -d '{
    "name": "iPhone 15",
    "description": "Latest iPhone model",
    "price": 999.99,
    "stock": 50,
    "category": "Electronics",
    "brand": "Apple",
    "sku": "IPHONE15-001"
  }'
```

### 2. Upload Product Image

```bash
curl -X POST http://localhost:8081/products/{product_id}/image \
  -H "Authorization: Bearer your_jwt_token" \
  -F "image=@/path/to/image.jpg"
```

### 3. Get All Products with Filters

```bash
curl "http://localhost:8081/products?page=1&limit=10&category=Electronics&search=iPhone"
```

### 4. Update Product

```bash
curl -X PUT http://localhost:8081/products/{product_id} \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token" \
  -d '{
    "price": 899.99,
    "stock": 30
  }'
```

### 5. Delete Product

```bash
curl -X DELETE http://localhost:8081/products/{product_id} \
  -H "Authorization: Bearer your_jwt_token"
```

## Testing

### Run the Test Script

```bash
go run test_product_api.go
```

This will test all CRUD operations including:
- Creating products
- Uploading images
- Retrieving products
- Updating products
- Deleting products

### Manual Testing

You can also test the API using tools like:
- Postman
- Insomnia
- curl commands

## Image Upload Specifications

### Supported Formats
- JPG/JPEG
- PNG
- GIF
- WEBP

### File Size Limit
- Maximum: 10MB

### Storage Location
- Directory: `uploads/products/`
- Naming: `{UUID}_{timestamp}.{extension}`
- Access URL: `http://localhost:8081/uploads/products/{filename}`

## Database Schema

### Products Table

```sql
CREATE TABLE products (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  created_at DATETIME,
  updated_at DATETIME,
  deleted_at DATETIME,
  uuid CHAR(36) UNIQUE,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  price DECIMAL(10,2) NOT NULL,
  stock INT NOT NULL DEFAULT 0,
  category VARCHAR(255) NOT NULL,
  brand VARCHAR(255),
  sku VARCHAR(255) UNIQUE,
  image_path VARCHAR(500),
  image_url VARCHAR(500),
  status VARCHAR(50) DEFAULT 'active',
  created_by CHAR(36),
  updated_by CHAR(36)
);
```

## Authentication

The API uses JWT (JSON Web Tokens) for authentication. To access protected endpoints:

1. Login to get a JWT token
2. Include the token in the Authorization header:
   ```
   Authorization: Bearer <your_jwt_token>
   ```

## Error Handling

The API returns consistent error responses:

```json
{
  "error": "Error message",
  "details": "Detailed error information"
}
```

Common HTTP status codes:
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `500` - Internal Server Error

## Security Features

- JWT authentication for protected routes
- File type validation for image uploads
- File size limits
- SQL injection prevention with GORM
- Input validation and sanitization
- Secure file naming with UUIDs

## Performance Considerations

- Database indexing on frequently queried fields
- Pagination for large datasets
- Efficient image storage and retrieval
- Soft deletes for data integrity

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

For questions or issues, please create an issue in the repository or contact the development team.

---

**Happy Coding! 🚀**