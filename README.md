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

### Running && testing

You can run the project using

```bash
go run main.go
```

Or, if you want to test it, use

```bash
go test -v ./...
```
