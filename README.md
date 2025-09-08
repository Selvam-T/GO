## 1. Docker as a self-contained Go development environment
#### Trial setup

### Dockerfile
```
FROM golang:1.25

WORKDIR /usr/src/app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        nano && \
    rm -rf /var/lib/apt/lists/*

CMD ["bash"]
```
### You can then build and run the Docker image:
```
docker build -t go-learning .
docker run -it -v "$(pwd)/workspace":/usr/src/app go-learning

```
## Inside the container create first program

```
# Create a new Go file
nano main.go

# Example content
package main
import "fmt"
func main() {
    fmt.Println("Hello from Docker Go!")
}
```

### Run it without creating go.mod
```
go run main.go
```

### Or create a module for more advanced experiments
```
go mod init myapp
go run main.go
```

## 2. Development environment for a learning project

### Folder structure

```
go-backend/
│
├─ Dockerfile
├─ docker-compose.yml      # optional, for easier volume/network management
├─ workspace/              # persistent volume for your Go code
│   ├─ main.go             # starter main file (optional, can be empty)
│   ├─ go.mod              # can be initialized inside container
│   ├─ handlers/           # API handlers (e.g., HTTP routes)
│   ├─ models/             # data structures / structs
│   └─ utils/              # utility packages
```

### Dockerfile

```
# Use official Go image
FROM golang:1.25

# Set working directory inside the container
WORKDIR /usr/src/app

# Install bash utilities (optional, for convenience)
RUN apt-get update && apt-get install -y \
    nano \
    curl \
    git \
    && rm -rf /var/lib/apt/lists/*

# Default command: start a bash shell
CMD ["bash"]
```

### docker-compose.yml

```
version: "3.9"

services:
  go-learning:
    build: .
    container_name: go-learning
    volumes:
      - ./workspace:/usr/src/app
    tty: true
    stdin_open: true
```

### Build and Run with docker-compose:
```
docker-compose run go-learning
```

### Inside the container, initialize Go modules if needed:
```
cd /usr/src/app
go mod init foodpanda-backend
```

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

## 3. Core Focus Areas of study

#### 1. REST APIs with net/http or frameworks (Gin/Fiber).
#### 2. JSON handling (encoding/json for requests/responses).
#### 3. Database connectivity (Postgres/MySQL with database/sql, sqlx, or GORM).
#### 4. Concurrency (goroutines & channels) → for handling requests at scale.
#### 5. Context (context package) → cancellation, deadlines for API requests.
#### 6. Error handling (if err != nil idioms, wrapping).
#### 7. Testing (testing package, table-driven tests).

## 4. Topics learned

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
- from browser connect to http://localhost:8080/hello

### > server-project.go

- in-memory CRUD API
- Reading/writing JSON 
- Handling GET, POST, DELETE 
- Updating the in-memory slice safely 
- Proper error handling and loggin