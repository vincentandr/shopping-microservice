# shopping-microservice (In Progress)

![Untitled Diagram](https://user-images.githubusercontent.com/42005057/147475016-bc2f7406-ed5d-4d1f-8471-b9ba47caef6b.png)

## Start via Docker commands
**Create `data` folder for mongo DB replica set first:**
```
shopping-microservice/
  ...
  data/
    mongo1/
    mongo2/
    mongo3/
```
Then start the docker containers by using the command:
```
docker-compose up -d
```

## API

Call the APIs using API platform, e.g. Postman.

### Catalog APIs
```
/products                   GET - Get all products 
/products/search            GET - Search products by name query parameter
```

### Cart APIs
```
/cart/{userId}              GET - Get cart items
/cart/{userId}              DELETE - Delete all cart items
/cart/{userId}/{productId}  PUT - Add or replace cart item with qty query parameter
/cart/{userId}/{productId}  DELETE - Delete a particular cart item by id
/cart/checkout/{userId}     GET - Checkout, returns an order ID to be used for payment
```

### Payment APIs
```
/payment                    GET - Get all orders
/payment/{userId}           GET - Get all orders by user ID
/payment/{orderId}          PUT - Change order status to paid
```

## Tests
HTTP handlers test
```
go test -v ./bff/internal/handler/...
```

Cart helper functions test
```
go test -v ./cart/
```

## Miscellaneous
### RabbitMQ Management / Admin

Access RabbitMQ admin web page for monitoring purpose.
```
localhost:15672
```

### Proto compiler command

```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <filename>
```

## Possible improvements

- Purely asynchronous communication between microservices (RabbitMQ pub/sub instead of GRPC, and return the async response using websocket / the client periodically polls for result)
- Caching results for certain requests
