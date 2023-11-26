# mdg-inventory-tools

[![Server CI](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/go.yml/badge.svg)](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/go.yml)

## Docker Setup

To build and run the Docker container:

```bash
docker build -t mdg-postgres .
docker volume create mdg_postgres_data
docker run -d --name mdg-database -e POSTGRES_DB=mdg -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -v mdg_postgres_data:/var/lib/postgresql/data -p 5432:5432 mdg-postgres
```

## Go Server

The Go server provides a RESTful API to interact with the inventory database.

To build and start the web server:

```bash
cd /api
make
./mdg-inventory-api
```
