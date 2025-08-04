# Golang-Api ( api(gin), gorm, mysql, swagger, docker )


## Features
- RESTful API built with Gin (CRUD)
- Database ORM with GORM
- Auto DB migration
- MySQL integration
- Swagger documentation
- Docker
- Unit Test (Mockery)
  
### Prerequisite
Make sure you have the following installed on your system:
- Docker
- go


#### 1. Run Application
```bash
# Run Docker Compose (API automatic database migration and data seeding)
- docker-compose up --build

# API Endpoints
-  localhost:8001

# Setup Swagger
- go install github.com/swaggo/swag/cmd/swag@latest
- swag init

# Can change configuration
- config.json

# Unit-test (example user-service)
- go test ./services/user
```

####  Api Docs
- your_url/swagger/index.html ( localhost:8001/swagger/index.html )
