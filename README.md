
**Device Manager Service**
=====================  

The Device Manager Service is a RESTful API that provides endpoints for managing devices.   
The service is built using GoLang. 

**Frameworks and Tools:**
------------  
[Gin Framework](https://github.com/gin-gonic/gin) was selected as the to handle the HTTP Layer

[Postgres](https://github.com/postgres) was selected to Database engine in the persistence layer.

[Mockery](https://github.com/vektra/mockery) to create the Mocks

[Testify](https://github.com/stretchr/testify) for the UnitTests

[Pure Go (pq)](https://github.com/lib/pq) Postgres drive to Database SQL

[goose](https://github.com/pressly/goose) Database Migration tool


**Endpoints**
------------  

### Create Device

* **POST /v1/devices**
+ Creates a new device
+ Request Body:
```json
{
  "name": "iPhone 13",
  "brand": "Apple"
}
```
+ Response: Created device object (JSON)
```json
  {
  "device_id":  "3289239a-4b0a-4a39-8202-e77031157132",
  "name":  "iPhone 13",
  "brand":  "Apple",
  "is_active":  true,
  "created_at":  "2024-07-17T09:31:56.131691+01:00",
  "updated_at":  "2024-07-17T09:31:56.131691+01:00"
  }
```
### Update Device

* **PUT /v1/devices/:id**
+ Updates an existing device
+ Request Body: Device object (JSON)
```json
    {
    "name":"iPhone 13",
    "brand":"Apple",
    "is_active":  true
    }
```

+ Response: Updated device object (JSON)
```json
  {
  "device_id":  "3289239a-4b0a-4a39-8202-e77031157132",
  "name":  "iPhone 13",
  "brand":  "Apple",
  "is_active":  true,
  "created_at":  "2024-07-17T09:31:56.131691+01:00",
  "updated_at":  "2024-07-18T09:31:56.131691+01:00"
  }
```

### Patch Device

* **PATCH /v1/devices/:id**
+ Partially updates an existing device
+ Request Body: Device object (JSON)  
 ```json 
{
  "is_active":  false
}
```
+ Response: Updated device object (JSON)  
  ```json
  204 No Content
  ```

### Delete Device

* **DELETE /v1/devices/:id**
+ Deletes an existing device
+ Response: 
 ```json
  204 No Content
  ```

### Get Device

* **GET /devices/:id**
+ Retrieves a device by ID
+ Response: 
```json
  {
  "device_id":  "3289239a-4b0a-4a39-8202-e77031157132",
  "name":  "iPhone 13",
  "brand":  "Apple",
  "is_active":  true,
  "created_at":  "2024-07-17T09:31:56.131691+01:00",
  "updated_at":  "2024-07-18T09:31:56.131691+01:00"
  }
```

### Get All Devices

* **GET /devices?page=1&limit=10&brand=Apple**
+ Retrieves a list of all devices
+ Response: List of device objects (JSON)
```json
[
  {
  "device_id":  "3289239a-4b0a-4a39-8202-e77031157132",
  "name":  "iPhone 13",
  "brand":  "Apple",
  "is_active":  true,
  "created_at":  "2024-07-17T09:31:56.131691+01:00",
  "updated_at":  "2024-07-18T09:31:56.131691+01:00"
  }
]
```
**Prerequisites**
---------------  

To run the project, you need to have Docker installed on your machine. The development Postgres database is configured in the `docker-compose` file.

**Technical Stack**: Go
-----------------  

### HTTP Layer

The HTTP layer is built using Gin, a popular Go web framework.

### Persistence Layer

The persistence layer uses Postgres as the database management system.

The `\db\migrations\` contains the database schema, and the `pq` library is used to run database migrations.

**Getting Started**
-------------------  

To get started with the project, simply run `docker-compose up` to start the development environment. Then, you can use a tool like `curl` or a REST client to interact with the API endpoints.

### Build/Run Scripts
    docker-compose -f docker/docker-compose.yaml -d
    go run cmd/api/*.go
