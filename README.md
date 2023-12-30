# mdg-inventory-tools

[![Server CI](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/go.yml/badge.svg)](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/go.yml) [![Client CI](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/react.yml/badge.svg)](https://github.com/lf-hernandez/mdg-inventory-tools/actions/workflows/react.yml)

MDG Inventory Tools is a comprehensive inventory management system, consisting of a Go web api server and a Vite/React client. It allows users to effectively manage inventory data.

## Docker compose development setip

To build and run the project for development:

```bash
docker compose -f compose.dev.yml up -d
```

This will spin up all 3 services: db, frontend, and backend and supports hot reload on code changes in the frontend.

## Database Docker Setup

To build and run the Docker container:

```bash
docker build -t mdg-postgres .
docker volume create mdg_postgres_data
docker run -d --name mdg-database -e POSTGRES_DB=mdg -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -v mdg_postgres_data:/var/lib/postgresql/data -p 5432:5432 mdg-postgres
```

## [Go Server](api)

The Go server provides a RESTful API to interact with the inventory database.

To build and start the web server:

```bash
cd /api
make
```

## [Vite/React Client](client)

The client is a web application that allows users to manage inventory.

To run:

```bash
cd /client
npm install
npm run dev
```
