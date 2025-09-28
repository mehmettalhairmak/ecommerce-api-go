# 🛒 E-Commerce API Go

[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org/doc/go1.25)
[![Stripe](https://img.shields.io/badge/Stripe-v82.5.1-purple.svg)](https://stripe.com/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A **demo** e-commerce payment API built with Go and Stripe integration. This project demonstrates secure payment processing, RESTful API design, and Docker containerization for learning and development purposes.

## ✨ Features

- 💳 **Stripe Payment Integration** - Secure payment processing demo
- 🚀 **RESTful API** - Clean API design patterns
- 🐳 **Docker Support** - Easy deployment and containerization
- 🏥 **Health Check** - System status monitoring
- 🛡️ **Secure Configuration** - Environment variables for sensitive data
- 📦 **Minimal Dependencies** - Lightweight and focused
- ⚡ **High Performance** - Built with Go's performance advantages
- 🧪 **Demo Purpose** - Perfect for learning and experimentation

## 🛍️ Demo Products

The API includes 3 sample products for demonstration:

| Product        | Price   |
| -------------- | ------- |
| 👕 Demo Shirt  | $155.00 |
| 👖 Demo Pants  | $260.00 |
| 🩳 Demo Shorts | $300.00 |

> **Note**: These are fictional products for testing and development purposes only.

## 🚀 Quick Start

### Prerequisites

- Go 1.25+
- Docker (optional)
- Stripe account and API keys (for testing)

### 📦 Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/ecommerce-api-go.git
   cd ecommerce-api-go
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env
   # Edit .env file and add your STRIPE_SECRET_KEY
   ```

4. **Run the application**

   ```bash
   go run server.go
   ```

   The server will start at http://localhost:4242

### 🐳 Running with Docker

```bash
# Using Docker Compose
docker-compose up --build

# Or manually
docker build -t ecommerce-api .
docker run -p 4242:4242 --env-file .env ecommerce-api
```

## 📡 API Endpoints

### POST /create-payment-intent

Creates a new payment intent for demo purposes.

**Request Body:**

```json
{
  "product_id": "Demo Shirt",
  "first_name": "John",
  "last_name": "Doe",
  "address_1": "123 Main St",
  "address_2": "Apt 4B",
  "city": "New York",
  "state": "NY",
  "zip": "10001",
  "country": "US"
}
```

**Response:**

```json
{
  "clientSecret": "pi_xxxxx_secret_xxxxx"
}
```

**Curl Example:**

```bash
curl -X POST http://localhost:4242/create-payment-intent \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "Demo Shirt",
    "first_name": "John",
    "last_name": "Doe",
    "address_1": "123 Main St",
    "city": "New York",
    "state": "NY",
    "zip": "10001",
    "country": "US"
  }'
```

### GET /health

Checks server status.

**Response:**

```
Server is up and running!
```

## ⚙️ Configuration

### Environment Variables

| Variable            | Description           | Required |
| ------------------- | --------------------- | -------- |
| `STRIPE_SECRET_KEY` | Stripe secret API key | ✅       |

### Stripe Configuration

1. Go to [Stripe Dashboard](https://dashboard.stripe.com/)
2. Get your secret key from API keys section
3. Add it to `.env` file:
   ```
   STRIPE_SECRET_KEY=sk_test_xxxxx...
   ```

> **Important**: Use test keys for this demo project!

## 🏗️ Project Structure

```
ecommerce-api-go/
├── 📁 .env.example          # Environment variables template
├── 📁 .gitignore           # Git ignore file
├── 📁 docker-compose.yml   # Docker Compose configuration
├── 📁 Dockerfile          # Docker image definition
├── 📁 go.mod              # Go module file
├── 📁 go.sum              # Go dependencies checksums
├── 📁 README.md           # Project documentation
└── 📁 server.go           # Main application file
```

## 🧪 Testing

```bash
# Health check
curl http://localhost:4242/health

# Create payment intent
curl -X POST http://localhost:4242/create-payment-intent \
  -H "Content-Type: application/json" \
  -d '{"product_id": "Demo Shirt", "first_name": "Test", "last_name": "User", "address_1": "123 Test St", "city": "Test City", "state": "TS", "zip": "12345", "country": "US"}'
```

## 🚀 Production Deployment

> **Warning**: This is a demo project. For production use, implement proper authentication, validation, error handling, and security measures.

### Docker Production

```bash
# Production build
docker build -t ecommerce-api-production .

# Run with production environment
docker run -d \
  -p 80:4242 \
  -e STRIPE_SECRET_KEY=sk_live_xxxxx \
  --name ecommerce-api \
  ecommerce-api-production
```

### Security Notes

- 🔐 Use live Stripe keys only in production
- 🌐 Always use HTTPS in production
- 🔒 Secure your environment variables
- 📊 Add proper logging and monitoring

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 🆘 Support

For questions and support:

- 🐛 [Issues](https://github.com/mehmettalhairmak/ecommerce-api-go/issues)
- 💬 [Discussions](https://github.com/mehmettalhairmak/ecommerce-api-go/discussions)
- 📧 Email: mehmetirmaakk@gmail.com

## 📚 Resources

- [Stripe Go SDK](https://github.com/stripe/stripe-go)
- [Go Documentation](https://golang.org/doc/)
- [Docker Documentation](https://docs.docker.com/)

---

⭐ If you found this demo helpful, please give it a star!

**Made with ❤️ for learning purposes**
