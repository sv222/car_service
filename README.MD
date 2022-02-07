# CAR Informer Service
## CRUD service for checking cars parameters from external API

### Migrations:

_https://github.com/golang-migrate/migrate_

`./migrate -path migrations -database "postgres://{user}:{password}@localhost:5432/mydb?sslmode=dis
able" up`

`./migrate -path migrations -database "postgres://{user}:{password}@localhost:5432/mydb?sslmode=dis
able" down`