# MDG Inventory Tools

[![API-CI/CD](https://github.com/rwx-solutions/mdg-inventory-manager/actions/workflows/api-cicd.yml/badge.svg)](https://github.com/rwx-solutions/mdg-inventory-manager/actions/workflows/api-cicd.yml) [![Client-CI/CD](https://github.com/rwx-solutions/mdg-inventory-manager/actions/workflows/client-cicd.yml/badge.svg)](https://github.com/rwx-solutions/mdg-inventory-manager/actions/workflows/client-cicd.yml)
## Docker Compose Development Setup

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Clone the repository.
- Create .env file at project root with the following environment variables:

```bash
POSTGRES_DB=<db_name>
POSTGRES_USER=<user>
POSTGRES_PASSWORD=<password>
VITE_API_URL=http://backend:8000
JWT_SECRET=<shh_some_secret>
DATABASE_URL=postgresql://postgres:postgres@postgres:5432/<db_name>?sslmode=disable
PORT=8000
CORS_ORIGINS=http://frontend:5173, frontend:5173
```

### Running the Services

1. **Start Services**: Navigate to the root directory of the project and run the following command to start all services (PostgreSQL database, Go backend, Vite/React frontend):

   ```bash
   docker compose -f compose.dev.yml up -d
   ```

2. **Verify Services**: Ensure that all services are up and running by executing:

   ```bash
   docker ps
   ```

3. **Access Services**:
   - Backend: Accessible at `http://localhost:8000`
   - Frontend: Accessible at `http://localhost:5173`
   - Database: Accessible at `localhost:5432`

### Hot Reloading

- The frontend service is configured for hot reloading. Any changes made in the source code will be immediately reflected in the running application.

## Database Docker Setup

To manually set up the PostgreSQL database in Docker:

1. **Build the PostgreSQL Container**:

   ```bash
   docker build -t mdg-postgres .
   ```

2. **Create a Persistent Volume**:

   ```bash
   docker volume create mdg_postgres_data
   ```

3. **Run the PostgreSQL Container**:

   ```bash
   docker run -d --name mdg-database -e POSTGRES_DB=mdg -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -v mdg_postgres_data:/var/lib/postgresql/data -p 5432:5432 mdg-postgres
   ```

## [Go Server](api)

The Go server offers a RESTful API for interacting with the inventory database.

### Building and Running the Server

1. Navigate to the `api` directory:

   ```bash
   cd api
   ```

2. Build and start the server:

   ```bash
   make
   ```

## [Vite/React Client](client)

The client application provides a user interface for inventory management.

### Running the Client

1. Navigate to the `client` directory:

   ```bash
   cd client
   ```

2. Install dependencies and start the client:

   ```bash
   npm install
   npm run dev
   ```

---

**Note**: Ensure that all the necessary env vars and configurations are correctly set up as per the project requirements.
