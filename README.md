# GoLang API Rest

A users Rest API using Go with the MVC architectural pattern, along with a login that verifies the user credentials.
The create user method encrypts with jwt before saving in MongoDB, then the user is able to login.

### Main tools

![MongoDB](https://img.shields.io/badge/MongoDB-black.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Docker](https://img.shields.io/badge/docker-black.svg?style=for-the-badge&logo=docker&logoColor=white)

### Other tools inside this project

- Gin
- Testify
- Gomock
- Zap

### Status

![Build status](https://github.com/BrunoPolaski/go-crud/actions/workflows/build.yml/badge.svg)

### Initializing without Docker

If you don't have docker installed, you can install Mongo and MongoExpress, create a collection and start to use it by running the application with:

```bash
go mod download
go run main.go
```

### Initializing with Docker

To start the app, mongo and mongo-express, make sure you have docker installed and simply run in the terminal:

```bash
docker compose up
```

And you will be ready to go!

After this, you can create a user with:

```bash
curl --json '{
	"email": "test@gmail.com",
	"name": "John",
	"age": 22,
	"password": "Chewbacca@777"
}' localhost:8080/createUser
```

And log in the server with:

```bash
curl --json '{
	"email": "test@gmail.com",
	"name": "John",
	"age": 22,
	"password": "Chewbacca@777"
}' localhost:8080/login
```

### Testing

To test the app, you must run the following command in the terminal:

```bash
go test -v ./...
```
