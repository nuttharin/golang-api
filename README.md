# golang-api

## Getting Started
The project uses Docker to run the API service and MySQL database. Upon startup, the API performs automatic database migration and data seeding.

### Prerequisite
Make sure you have the following installed on your system:
- Docker

### This api uses 
- golang
- Mysql
- Unit Test (Mockery)
- docker

#### 1. Run Application
```bash
# build docker
- docker-compose up --build


# api 
- api run at localhost:8001
```

####  Change Configuration
- config.json
