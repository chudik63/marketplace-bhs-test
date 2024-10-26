# BHS Marketplace
A test project for the Marketplace-Hive.
# Installation
```
git clone https://github.com/chudik63/marketplace-bhs-test.git
cd marketplace-bhs-test
```
# Main Commands
Build and start all containers:
```
make build
```
Start all containers:
```
make up
```
Stop all containers:
```
make down
```

Start only the PostgreSQL container for development::
```
make dev-up
```
Stop the PostgreSQL container::
```
make dev-down
```

# Usage
## API Endpoints
- User Registration: POST /sign-up
```
curl -X POST http://localhost:8080/sign-up \
     -H "Content-Type: application/json" \
     -d '{
           "username": "user",
           "password": "password"
         }'
```
- User Authentication: POST /sign-in
```
curl -X POST http://localhost:8080/sign-in \
     -H "Content-Type: application/json" \
     -d '{
           "username": "user",
           "password": "password"
         }'
```
- User Signing Out: POST /sign-out
```
curl -X POST http://localhost:8080/sign-out \
     --cookie "access_token=YOUR_ACCESS_TOKEN; refresh_token=YOUR_REFRESH_TOKEN"
```
- Create Asset: POST /marketplace/assets
```
curl -X POST http://localhost:8080/marketplace/assets \
     -H "Content-Type: application/json" \
     --cookie "access_token=YOUR_ACCESS_TOKEN; refresh_token=YOUR_REFRESH_TOKEN" \
     -d '{
           "name": "Asset",
           "description": "This is an example asset",
           "price": 100.0
         }'

```
- Delete Asset: DELETE /marketplace/assets/:id
```
curl -X DELETE http://localhost:8080/marketplace/assets/1 \
     --cookie "access_token=YOUR_ACCESS_TOKEN; refresh_token=YOUR_REFRESH_TOKEN"
```
- Buy Asset: PATCH /marketplace/assets/:id
```
curl -X PATCH http://localhost:8080/marketplace/assets/1 \
     --cookie "access_token=YOUR_ACCESS_TOKEN; refresh_token=YOUR_REFRESH_TOKEN"
```

# Technologies Used
- Golang (Gin, Viper, Gorm)
- PostgreSQL
- Swagger
- Docker