
## Build
```bash
go build .
```
## Run
Build and run
```bash
go run .
```

or just run the compiled binary
```bash
./go-boot
```

## Graphql 
Graphql Endpoint
```text
localhost:8080/graphql/
```

Example of graphql 
```graphql
mutation createTodo {
  createTodo(input: { text: "todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}
```

```graphql
query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}
```