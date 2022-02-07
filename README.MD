# CAR Informer Service
## CRUD service for checking cars parameters from external API

### Endpoints:
 
/ - service description

### Migrations:

_https://github.com/golang-migrate/migrate_

Create table users:

`./migrate -path migrations -database "postgres://{user}:{password}@localhost:5432/mydb?sslmode=disable" up`

Drop database users:

`./migrate -path migrations -database "postgres://{user}:{password}@localhost:5432/mydb?sslmode=disable" down`

Lunch services:

docker-compose up -d

Stop/remove:

docker-compose stop / docker-compose down