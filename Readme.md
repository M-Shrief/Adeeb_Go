# Go RESTful API with Huma and Chi.

Objectives: 
- Leveling up my skills in creating Go APIs and my coding conventions.
- This project should be a base for a bigger project in my mind soon.

Tech Stack:
- Go
- Huma and Chi
- Pgx and Sqlc

Characteristics:
- Adherent to open industry standards thanks to Huma.
- Interactive and Automated Documentation with Huma.
- Auto generation for OpenAPI & JSON Schema thanks to Huma.
- Validation for Requests and Responses with Huma automatic validation.
- Using structured logging with Zerolog.
- JWT Authentication and Authorization.
- Using stdlib net/http middlewares - mainly Chi's middlewares - for logging requests, rate limiting, CORS, ...etc.


## Folders structure
I'll default Go files, and general folders like `.gitignore`...etc.

- `internal` folder for packages that specific for the project:
    - `auth` holds JWT authentication and authorization config and functionalities. 
    -  `components` folder for the components that holds specific groups of routes and functionality.
        - `heartbeat` to holds the server's health check functionalities.
        - `users` to hold Users' functionalities: Signup, Login, Update and Delete.
    - `config` holds the functionality to read ENVs, certificates and PEM files.
    - `database` holds Pgx and SQLC config and functionalities.
        - `sql` holds sql queries for sqlc.
        - `index.go` for specific Pgx config and connecting to the database.
        - `utils.go` to hold useful functions for database calls.
        - `db.go`, `models.go` and `queries.sql.go` are auto-generated files by Sqlc.
- `kreya` is a folder Kreya data, Kreya is an app to test RESTful and gRPC APIs, it's like Postman.
- `logger` folder holds Zerolog config and functionalities.
- `router` folder holds Chi and Huma config and functionalities.
    - `api.go` holds Huma's config & functionalities
    - `router.go` holds Chi config, funcionalities and Middlewares.
- `Setup.md` for Setup instructions to start the project
- `Sqlc.yaml` for Sqlc config.