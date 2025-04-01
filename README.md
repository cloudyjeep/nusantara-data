# 🚀 Go API with Fiber


## Available API 
> Public

    > get list category
    GET /category

    > create category
    POST /category/:name

    > delete category
    DELETE /category:/name
    

> Protected

    > create product
    POST /product
    
    > get list product
    GET /product
    
    > get detail product
    GET /product/:id
    
    > update product
    PUT /product/:id
    
    > delete product
    DELETE /product/:id

> Authorization: Bearer abcd



## Available Commands

### 1. Build the Application
```sh
make build
```
This will compile the Go project and place the binary inside the `bin/` directory with optimizations.

### 2. Build and Run the Application
```sh
make build-run
```
This will first build the application and then run the generated binary.

### 3. Run in Development Mode
```sh
make dev
```
Runs the application with `go run`, suitable for development.

### 4. Clean Build Artifacts
```sh
make clean
```
Removes the `bin/` directory and clears compiled binaries.

### 5. Format the Code
```sh
make fmt
```
Runs `go fmt` to format Go code.

### 6. Lint the Code
```sh
make lint
```
Runs `go vet` to analyze and catch potential issues in the code.

### 7. Run Tests
```sh
make test
```
Runs all unit tests within the project.

## Custom Environment Variables
```sh
make build PORT_ENV=3000
```
This allows for configuration adjustments without modifying the Makefile itself.

---
