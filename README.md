# TODO
Simple REST API application serving as TODO list. 

## Getting Started
### Prerequisites
- Docker
- Docker Compose
- (Optional) Postman or any other API testing tool
- (Optional) K6 for load testing
- (Optional) PgAdmin or any other PostgreSQL client
### Running the application
  1. Clone the repository
     ```shell
     git clone https://github.com/tobiaszkonieczny/todo.git
        ```
  2. Navigate to the project directory
  3. Start the application using Docker Compose
     ```shell
     docker-compose up -d --build
     ```
## REST API Endpoints
- Everything is shown under /swagger/index.html

## Tech Stack
- Go
- Gin
- Gorm
- PostgreSQL
- Docker
- Docker Compose
- Swagger
- Nginx
- JWT
- K6
- Git for version control
- Websocket (for real-time updates)

## General architecture & setup
Everything is containerized using Docker. The application and the database are running in separate containers. Docker Compose is used to manage the multi-container application.

Nginx is used as simple API gateway. It forwards the requests to the application container. It also handles HTTPS secure connection and serves static files (Swagger UI).

The application is connected to PostgreSQL database using Gorm ORM. The database connection parameters are set using environment variables in **docker-compose.yml**.



## Database
Database is initialized via script located in **/db/init.sql**. The script is executed when the database container is started for the first time.

### Database diagram
![Diagram](/docs/diagram%20.png)


## Features

### - Rest API
- Create, Read, Update, Delete TODO items
- Swagger documentation
- Based on Gin
    + Better performance than net/http
    + Middleware support
    + Easy routing
    + Shorter code
    + JSON validation suppor
    + Websocket support
    + File upload support
- Authentication with JWT
- Used Gorm as ORM
    + Auto migration
    + Easier database operations
    + Database agnostic (can be used with other databases)
    + Relationships support
- Implemented API request logs (using middleware)
- Full API testing with Bruno
- Performance testing with K6
    + to run the tests, navigate to /testing and run: `k6 run --insecure-skip-tls-verify k6_test.js`