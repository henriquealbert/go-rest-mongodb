# Go + Fiber + MongoDB

Go 
[Fiber](https://docs.gofiber.io/) - Express-inspired web framework
MongoDB - NoSQL database

Very simple product API with Create and Read operations.

### How to run the app

```bash
# Clone this repository
$ git clone https://github.com/henriquealbert/go-rest-mongodb.git

# Go into the repository
$ cd go-rest-mongodb

# Install dependencies
$ go install

# Run the app
$ go run cmd/main.go
```

## API Routes

### Products

| Method | Route | Description |
| ------ | ----- | ----------- |
| GET | /api/v1/products | Get all products |
| POST | /api/v1/products | Add product |
