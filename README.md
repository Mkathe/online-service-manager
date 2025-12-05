# Online Service Manager

The service manager is created for a test assignment by Effective Mobile company.

## Features
- CRUD of services
- Ability to find the total cost of services within the date parameters
- Swagger is provided

### Usage Routes:

```
Routes

GET "/healthz"

Services
GET "/services" 
GET "/services/total-cost"
POST "/services"
PUT "/services/:id"
DELETE "/services/:id"

To open swagger:
localhost:port/swagger
```

Use this command to open this project:

```bash
docker compose up --build
```