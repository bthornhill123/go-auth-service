## About
This authentication service can inspire a more build out authentication system. It allows clients access to common authentication use cases.
## Installation

1. In project root, startup all necessary services via Docker
   ```sh
   docker-compose up
   ```
2. Utilize example requests in [endpoints.http](./endpoints.http) to access endpoints
   - VSCode's "REST Client" extension allows easy testing via the endpoint.http file
   - The requests can also be manually copied and tested in alternitive REST clients (e.g. Postman)

## Design Choices
### GoLang 
- Go's performance makes it a good candidate for an authentication service
- Not having work experience in Go, I wanted to prove that I can pick up new technologies quickly
### Go Fiber
- While Go has several popular web frameworks, I chose Fiber because it is inspired by Express 
- For new GoLang developers, using a framework with similar patterns to Express may ease onboarding if they have developed Node.js web applications
### Gorm
- I chose to use the defacto ORM in the Go ecosystem for this project
- While ORM's introduce abstractions and performance hits, they also allow a lot of flexibility 
- Switching database systems (PostgreSQL -> SQL Server) will be simpler since the schemas and database-access logic is all contained in the Go source code
### PostgreSQL
- Powerful, open-source database technology



