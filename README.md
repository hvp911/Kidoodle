
# Content API

Content API is a content retrieval API. Content API Stores contents in encrypted format per devices based on protection systems.

# Development environment

## Docker Desktop

Install latest version of Docker using Docker Desktop.

## Starting database

In the root of the directory, run this command to start the PostgreSQL database (on port 5432):

```bash
$ docker-compose up postgresql -d
```

## Connecting to database

During development, you can connect to and experiment with the PostgreSQL database by running this command:

```bash
$ docker-compose exec postgresql psql "postgres://postgres:password123@localhost:5432/content?sslmode=disable"
```

Then, on the psql CLI, test as follows:

```psql
content=# \dt
```

If everything went well, you should get this result:

```psql
               List of relations
 Schema |       Name        | Type  |  Owner
--------+-------------------+-------+----------
 public | contents          | table | postgres
 public | devices           | table | postgres
 public | protection_system | table | postgres
(3 rows)
```

To exit the PostgreSQL session, type `\q` and press `ENTER`.

# Deploying and running back-end microservice

## Spinning up microservice on IDEA (GoLand)

- Start postgresql db
```bash
$ docker-compose up postgresql -d
```
- go to api.go file under `\service\cmd\api\api.go`
- run little play button near main method.

## Building and running back-end app

```bash
$ docker-compose up --build
```

## Postman collection
- content_postman_collection.json

## Swagger Clients generation
