## 1. Docker Setup

### Go Development Container (go-service)

- Mounts host folder ./workspace to /workspace for live coding.
- Provides Go environment and tools (go, git, nano, etc.).
- Connects to Postgres using host db from Compose.

### Postgres Database Container (db)

- Uses official postgres:15 image.
- Persists data in host folder ./db_data.
- Initializes database, user, and optional tables via /init/init.sql.
- Exposes port 5432 for connectivity.

### Docker Compose

- docker compose up -d starts both containers.
- docker compose exec go-service bash enters Go environment.
- psql -h db -U <DB_USER> -d <DB_NAME> connects to database from Go container.

### Makefile Commands

- make build → build images
- make up → start containers
- make shell → enter Go container
- make logs → view logs for all services
- make down → stop containers

### Folder structure

```
project-root/
├── docker-compose.yml      # Docker Compose configuration for Go and Postgres
├── Dockerfile              # Go development container image
├── Makefile                # Build, up, shell, logs, and cleanup commands
├── README.md               # Project overview and instructions
├── db_data/                # Host folder for Postgres persistent data
├── workspace/              # Host folder mounted into Go container for code
└── init/                   # Database initialization scripts
    ├── init.sh             # Shell script to run SQL/init tasks
    └── init.sql            # SQL script to create database, users, tabless
```

### Dockerfile

#### The Dockerfile installs:

- Install:
    - golang base image → Go compiler and tools
    - nano, curl, git → useful development utilities
    - postgresql-client → allows connecting to the Postgres database from the container via psql
- CMD: 
    - Keeps the container running so developers can docker exec into it.
    - No Go application is started automatically; you run Go programs manually inside the container.

#### Why there is no Dockerfile for Postgres

- The Postgres container uses the official postgres:15 image from Docker Hub.
- This image already contains a ready-to-run Postgres server.
- Database initialization (creating the user, database, and optional tables) is handled by placing scripts in:

```
./init → /docker-entrypoint-initdb.d
```
- On first container startup, Postgres automatically runs these scripts; no custom Dockerfile is needed.

### Inside the container, initialize Go modules if needed:
```
cd /usr/src/app
go mod init foodpanda-backend
```

### Test DB connection inside postgres container:

```
export PGPASSWORD=${DB_PASSWORD}
psql -h db -U ${DB_USER} -d ${DB_NAME}
```
where, -h localhost, if you are connecting from the same container  
or, -h db, if connecting from another container, and if service name in .yml is 'db'  

### Alternative: jump directly into the Postgres containers
```
docker compose exec db psql -U ${DB_USER} -d ${DB_NAME}
```

### Useful psql commands

\l .....................List all databases  
\c go_db ...............Connect to go_db (you’re already connected)  
\dt	....................List all tables in the current DB  
\d table_name ..........Describe table structure  
SELECT * FROM orders; ..Query all rows from the orders table  
\q ....................Quit psql  

#### go.mod explained
- Tells Go this is a module and not just a random set of .go files.
- Prepares your project to add dependencies with go get later.
- Makes builds reproducible, because dependencies and versions are tracked automatically.
- Then, any dependency you go get or go mod tidy will also generate go.sum.
- This way, your project is fully self-contained inside Docker, and your host stays clean.

### Run main.go created above

```
go run main.go
```

## 2. Core Focus Areas of study

#### 1. REST APIs with net/http or frameworks (Gin/Fiber).
#### 2. JSON handling (encoding/json for requests/responses).
#### 3. Database connectivity (Postgres/MySQL with database/sql, sqlx, or GORM).
#### 4. Concurrency (goroutines & channels) → for handling requests at scale.
#### 5. Context (context package) → cancellation, deadlines for API requests.
#### 6. Error handling (if err != nil idioms, wrapping).
#### 7. Testing (testing package, table-driven tests).

## 3. Topics learned

### > json_parsing.go

- Variables.
- Functions → main.
- Structs.
- JSON encoding/decoding.
- Mini-task (parse JSON, return structured output).

### > sliceAndMap.go

- Slices.
- Maps.

### > conditional.go

- if / else
- switch

### > loops.go

- for
- while (or like it)
- infinite for

### > pointers.go

- function argument pass array by value
- function argument pass array by address
- function argument accepting a slice

### > receiver_type.go

- value receiver type method
- pointer receiver type method

### > err_handling1.go

- built-in interface type called error for handling errors.

### > interface.go

- implement an interface
- named interface
- interface value
- assertion

### > hello-server.go

- http.ListenAndServe()
- http.HandleFunc()
- http.ResponseWriter
- *http.Request
- io.WriteString()

- 1. from browser connect to http://localhost:8080/hello
- 2. curl commands:
    - GET > curl http://localhost:8080/orders
    - POST > curl -X POST http://localhost:8080/orders -d '{"id":1,"item":"Burger","quantity":2}' -H "Content-Type: application/json"
    - DELETE > curl -X DELETE http://localhost:8080/orders -d '{"id":41}'

### > server-project.go

- in-memory CRUD API
- Reading/writing JSON 
- Handling GET, POST, DELETE 
- Updating the in-memory slice safely 
- Proper error handling and loggin

### > connectToDb.go

- Reads environment variables for database connection
- Opens a connection to the Postgres database
- Checks the connection

#### You need a Postgres driver because Go’s standard database/sql package only defines generic database interfaces — it cannot communicate with Postgres by itself.  
#### Inside the container, run:  
```
go get github.com/lib/pq
go mod tidy
```