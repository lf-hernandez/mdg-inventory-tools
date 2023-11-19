# mdg-inventory-tools

[![Server CI](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/go.yml/badge.svg)](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/go.yml)

## Docker Setup

To build and run the Docker container:

```bash
docker build -t mdg-postgres .
docker run -p 5432:5432 --name mdg-database -d mdg-postgres
```

## Go Server

The Go server provides a RESTful API to interact with the inventory database.

To build and start the web server:

```bash
cd /api
make
./api
```
