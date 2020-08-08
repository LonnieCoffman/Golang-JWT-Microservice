## Golang JWT Authentication Microservice
Authentication service built using Go.  Consists of separate administration and client logins.  Admin login is role based and currently has SuperAdmin and Admin roles.  Uses Gorm, Viper, Gin, JWT-Go and Swagger

This was my first attempt at building a full project using Go and due to my experience with other languages I may have not done things in the way of a true "Gopher".  Any feedback is welcome.

### Features
- Versioning
- Swagger integration
- Admin role based JWT Middleware
- Client JWT middleware
- Fully separated Admin and Client login system
- Database migration with sample data