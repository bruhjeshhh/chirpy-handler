# Chirpy Handler

A Go-based web API that provides a Twitter-like microblogging service called "chirps". This is a backend REST API with user authentication, JWT tokens, and PostgreSQL database integration.

## Features

### User Management
- User registration with email and password
- Secure authentication with JWT tokens
- Profile updates (email and password)
- Premium "Chirpy Red" membership system via webhooks

### Chirp System
- Create and share short messages (max 140 characters)
- Retrieve all chirps with optional filtering
- Get individual chirps by ID
- Delete your own chirps
- Automatic profanity filtering

### Security
- Argon2ID password hashing
- JWT access tokens (1-hour expiration)
- Refresh tokens (60-day expiration)
- Token revocation support
- API key authentication for webhooks

### Admin Features
- Metrics dashboard for server monitoring
- Database reset functionality
- Health check endpoints

## Tech Stack

- **Language**: Go 1.25.6
- **Web Framework**: Standard library `net/http`
- **Database**: PostgreSQL
- **Database Driver**: `github.com/lib/pq`
- **ORM/Query Builder**: SQLC for type-safe SQL queries
- **Authentication**: JWT tokens (`github.com/golang-jwt/jwt/v5`)
- **Password Hashing**: Argon2ID (`github.com/alexedwards/argon2id`)
- **Environment Variables**: `github.com/joho/godotenv`
- **UUID Generation**: `github.com/google/uuid`

## Installation

### Prerequisites
- Go 1.25.6 or higher
- PostgreSQL database
- SQLC (for code generation, if modifying schema)

### Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd chirpy-handler
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   DB_URL="postgres://username:password@localhost:5432/chirpy?sslmode=disable"
   secrett="your-jwt-secret-key"
   POLKA_KEY="your-webhook-api-key"
   ```

4. **Set up the database**
   ```bash
   # Create PostgreSQL database
   createdb chirpy
   
   # Run migrations (execute SQL files in order)
   psql chirpy < sql/schema/001_users.sql
   psql chirpy < sql/schema/002_chirps.sql
   psql chirpy < sql/schema/003_tokens.sql
   ```

5. **Generate SQLC code** (only if modifying SQL queries)
   ```bash
   sqlc generate
   ```

6. **Run the server**
   ```bash
   go run .
   ```

The server will start on port 8080.

## API Endpoints

### Health & Admin
- `GET /api/healthz` - Health check
- `GET /admin/metrics` - Admin metrics dashboard
- `POST /admin/reset` - Reset database (admin only)

### Authentication
- `POST /api/users` - User registration
- `POST /api/login` - User login (returns access and refresh tokens)
- `POST /api/refresh` - Refresh access token
- `POST /api/revoke` - Revoke refresh token
- `PUT /api/users` - Update user profile (requires authentication)

### Chirps
- `POST /api/chirps` - Create chirp (requires authentication)
- `GET /api/chirps` - Get all chirps (supports query parameters)
  - `author_id` - Filter by author
  - `sort` - Sort order (`asc` or `desc`)
- `GET /api/chirps/{chirpID}` - Get specific chirp
- `DELETE /api/chirps/{chirpID}` - Delete chirp (requires authentication, author only)

### Webhooks
- `POST /api/polka/webhooks` - Process membership upgrades (requires API key)

## Usage Examples

### User Registration
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'
```

### Create Chirp
```bash
curl -X POST http://localhost:8080/api/chirps \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <access_token>" \
  -d '{"body": "Hello, world!"}'
```

### Get Chirps
```bash
curl http://localhost:8080/api/chirps?sort=desc
```

## Project Structure

```
chirpy-handler/
├── main.go                    # Main application entry point
├── app.go                     # Health check handler
├── metrics.go                 # Admin metrics functionality
├── chirps.go                  # Chirp CRUD operations
├── user.go                    # User authentication and management
├── refreshandler.go           # Token refresh and revocation
├── respond.go                 # HTTP response utilities
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── sqlc.yaml                  # SQLC configuration
├── .env                       # Environment variables
├── .gitignore                 # Git ignore rules
├── index.html                 # Simple welcome page
├── assets/                    # Static assets
├── internal/
│   ├── auth/                  # Authentication logic
│   └── database/             # Database layer
└── sql/
    ├── schema/               # Database migrations
    └── queries/              # SQL queries for SQLC
```

## Configuration

### Environment Variables
- `DB_URL`: PostgreSQL connection string
- `secrett`: JWT secret key for token signing
- `POLKA_KEY`: API key for webhook authentication

### Server Configuration
- **Port**: 8080
- **Static Files**: Served from `/app/` prefix
- **JWT Access Token Expiry**: 1 hour
- **JWT Refresh Token Expiry**: 60 days

## Development

### Running Tests
```bash
go test ./...
```

### Code Generation
If you modify the SQL queries or schema:
```bash
sqlc generate
```

### Linting and Formatting
```bash
go fmt ./...
go vet ./...
```

## Security Notes

- Passwords are hashed using Argon2ID
- JWT tokens are used for authentication
- Input validation is performed on all endpoints
- Profanity filtering is applied to chirp content
- API keys are used for webhook authentication

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License.