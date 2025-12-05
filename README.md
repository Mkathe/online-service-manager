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

Hubs
GET "/services" 
GET "/services/total-cost"
POST "/services"
PUT "/services/:id"
DELETE "/services/:id"

```

Use this command to open this project:

```bash
docker compose up --build
```