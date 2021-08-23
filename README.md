# NodesApi

Just a REST API with CRUD operations.
In this case it's a Node entity.
The Node can be of 'Store' type or 'Pickup' type.

Also there is a service that determines the node closest to a given point (Latitud & Longitud)

## Development

This api was developed using
- Go
- Gin-gonic
- Swaggo
- Mongo DB
- Docker

## Installation


First of all, you need to install [Docker](https://www.docker.com/products)


Clone the project
````
git clone https://github.com/gastitan/NodesApi
`````

Go to the new directory
`````
cd NodesApi/
`````

Run the project
`````
docker-compose up -d --build
`````

The previous code basically have 3 important things
1. Download MongoDB image
2. Create an image of NodesApi
3. Build

For more details see [docker-compose.yaml](https://github.com/gastitan/NodesApi/blob/master/docker-compose.yaml) & [Dockerfile](https://github.com/gastitan/NodesApi/blob/master/Dockerfile)

## Let's play
API documentation with Swagger 2.0

http://localhost:8080/swagger/index.html
