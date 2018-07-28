# GraphQL CRUD examle
CRUD example using GraphQL

## Request example
### List
```
curl -X POST -d '
  query {
    list {
      id,
      name,
      info,
      price
    }
  }' \
  http://localhost:8080/product
```

### Get
```
curl -X POST -d '
  query {
    product(id: 1) {
      id,
      name,
      info,
      price
    }
  }' \
  http://localhost:8080/product
```

### Post
```
curl -X POST -d '
  mutation {
    create(name: "foo", info: "bar", price: 1.23) {
      id,
      name,
      info,
      price
    }
  }' \
  http://localhost:8080/product
```

### Put
```
curl -X POST -d '
  mutation {
    update(id: 1, price: 1.23) {
      id,
      name,
      info,
      price
    }
  }' \
  http://localhost:8080/product
```

### Delete
```
curl -X POST -d '
  mutation {
    delete(id: 1) {
      id,
      name,
      info,
      price
    }
  }' \
  http://localhost:8080/product
```

### Composite
```
curl -X POST -d '
  query {
    list {
      id
    }
    product(id: 1) {
      id,
      name,
      info,
      price
    }
  }' \
  http://localhost:8080/product
```

### Aliases
```
curl -X POST -d '
  query {
    first: product(id: 1) {
      id,
      name,
      price
    }
    second: product(id: 2) {
      id,
      name,
      price
    }
  }' \
  http://localhost:8080/product
```
