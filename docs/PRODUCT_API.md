# Product CRUD API Documentation

This document describes the complete Product CRUD API with image upload functionality built with Go, Gin Framework, GORM, and MySQL.

## Base URL
```
http://localhost:8081
```

## Authentication
Protected endpoints require JWT authentication. Include the JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

## Endpoints

### 1. Create Product (Protected)
**POST** `/products`

Creates a new product.

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "name": "Product Name",
  "description": "Product description",
  "price": 99.99,
  "stock": 100,
  "category": "Electronics",
  "brand": "Brand Name",
  "sku": "SKU123",
  "status": "active"
}
```

**Response:**
```json
{
  "message": "Product created successfully",
  "data": {
    "id": "uuid",
    "name": "Product Name",
    "description": "Product description",
    "price": 99.99,
    "stock": 100,
    "category": "Electronics",
    "brand": "Brand Name",
    "sku": "SKU123",
    "image_url": "",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 2. Upload Product Image (Protected)
**POST** `/products/{id}/image`

Uploads an image for a specific product.

**Headers:**
- `Content-Type: multipart/form-data`
- `Authorization: Bearer <token>`

**Form Data:**
- `image`: Image file (JPG, JPEG, PNG, GIF, WEBP, max 10MB)

**Response:**
```json
{
  "message": "Image uploaded successfully",
  "image_url": "http://localhost:8081/uploads/products/filename.jpg"
}
```

### 3. Get All Products (Public)
**GET** `/products`

Retrieves all products with pagination and filtering.

**Query Parameters:**
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)
- `category` (optional): Filter by category
- `status` (optional): Filter by status (default: active)
- `search` (optional): Search in name and description

**Example:**
```
GET /products?page=1&limit=10&category=Electronics&search=phone
```

**Response:**
```json
{
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": "uuid",
      "name": "Product Name",
      "description": "Product description",
      "price": 99.99,
      "stock": 100,
      "category": "Electronics",
      "brand": "Brand Name",
      "sku": "SKU123",
      "image_url": "http://localhost:8081/uploads/products/image.jpg",
      "status": "active",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 50,
    "total_pages": 5
  }
}
```

### 4. Get Product by ID (Public)
**GET** `/products/{id}`

Retrieves a single product by its ID.

**Response:**
```json
{
  "message": "Product retrieved successfully",
  "data": {
    "id": "uuid",
    "name": "Product Name",
    "description": "Product description",
    "price": 99.99,
    "stock": 100,
    "category": "Electronics",
    "brand": "Brand Name",
    "sku": "SKU123",
    "image_url": "http://localhost:8081/uploads/products/image.jpg",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 5. Update Product (Protected)
**PUT** `/products/{id}`

Updates an existing product. All fields are optional.

**Headers:**
- `Content-Type: application/json`
- `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "name": "Updated Product Name",
  "description": "Updated description",
  "price": 149.99,
  "stock": 50,
  "category": "Updated Category",
  "brand": "Updated Brand",
  "sku": "NEWSKU123",
  "status": "active"
}
```

**Response:**
```json
{
  "message": "Product updated successfully",
  "data": {
    "id": "uuid",
    "name": "Updated Product Name",
    "description": "Updated description",
    "price": 149.99,
    "stock": 50,
    "category": "Updated Category",
    "brand": "Updated Brand",
    "sku": "NEWSKU123",
    "image_url": "http://localhost:8081/uploads/products/image.jpg",
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 6. Delete Product (Protected)
**DELETE** `/products/{id}`

Soft deletes a product and removes its image file.

**Headers:**
- `Authorization: Bearer <token>`

**Response:**
```json
{
  "message": "Product deleted successfully"
}
```

### 7. Get Product Categories (Public)
**GET** `/products/categories`

Retrieves all unique product categories.

**Response:**
```json
{
  "message": "Categories retrieved successfully",
  "data": [
    "Electronics",
    "Clothing",
    "Books",
    "Home & Garden"
  ]
}
```

## Image Upload Specifications

### Supported Formats
- JPG/JPEG
- PNG
- GIF
- WEBP

### File Size Limit
- Maximum: 10MB

### Storage
- Images are stored in `uploads/products/` directory
- Files are renamed with UUID + timestamp for uniqueness
- Accessible via: `http://localhost:8081/uploads/products/{filename}`

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request data",
  "details": "Validation error details"
}
```

### 401 Unauthorized
```json
{
  "error": "Unauthorized"
}
```

### 404 Not Found
```json
{
  "error": "Product not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error",
  "details": "Error details"
}
```

## Database Schema

### Product Model
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
  updated_by CHAR(36),
  INDEX idx_products_deleted_at (deleted_at),
  INDEX idx_products_uuid (uuid),
  INDEX idx_products_category (category),
  INDEX idx_products_status (status)
);
```

## Example Usage with cURL

### Create Product
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

### Upload Image
```bash
curl -X POST http://localhost:8081/products/{product_id}/image \
  -H "Authorization: Bearer your_jwt_token" \
  -F "image=@/path/to/image.jpg"
```

### Get All Products
```bash
curl -X GET "http://localhost:8081/products?page=1&limit=10&category=Electronics"
```

### Update Product
```bash
curl -X PUT http://localhost:8081/products/{product_id} \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token" \
  -d '{
    "price": 899.99,
    "stock": 30
  }'
```

### Delete Product
```bash
curl -X DELETE http://localhost:8081/products/{product_id} \
  -H "Authorization: Bearer your_jwt_token"
```