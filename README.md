# User Service

This is a Go microservice for user authentication and authorization using JWT tokens. It includes endpoints for user signup, login, and CRUD operations. The data is stored in a PostgreSQL database, and the application is containerized using Docker and Docker Compose.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```sh
git clone https://github.com/yourusername/auth_and_auth_microservice.git
cd auth_and_auth_microservice
```

##  Build and Run Locally with Docker
docker-compose up --build

### SignUP POST

    url :http://localhost:8000/signup
    body-json :
        {
  "username": "amol",
  "password": "amol@123"
}

### Login POST
    url : http://localhost:8000/login
    body-json :
       {
        "username": "amol",
        "password": "amol@123"
    }

### Welcome GET
 Message
    url :http://localhost:8000/welcome

    Headers:
    Authorization : Bearer <KEY>


### Delete User DELETE
    url : http://localhost:8000/user/{id}
    body-json :
        {
  "username": "amol",
  "password": "Amol@123"
}









