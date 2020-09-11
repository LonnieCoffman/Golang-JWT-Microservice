## Golang JWT Authentication Microservice
Authentication service built using Go that utilizes access tokens and refresh tokens.  Consists of separate administration and client logins.  Admin login is role based and currently has SuperAdmin and Admin roles.  Custom middleware verifies tokens and roles, checks user status and passes logged in data to controllers. Uses Gorm, Viper, Gin, JWT-Go and Swagger

This was my first attempt at building a full project using Go and due to my experience with other languages I may have not done things in the way of a true "Gopher".  Any feedback is welcome.

### Features
- Versioning
- Access Tokens and Refresh Tokens
- Impersonate a client
- Swagger integration
- Admin role based JWT Middleware
- Client JWT middleware
- Fully separated Admin and Client login system
- MySQL Database migration with sample data
