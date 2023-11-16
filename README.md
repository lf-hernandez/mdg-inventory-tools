# mdg-inventory-tools

Postgres database and set of tooling for managing inventory.

## Docker setup

To build and run Docker container, run:

```bash
docker build -t mdg-postgres .
docker run -p 5432:5432 --name mdg-database -d mdg-postgres
```
