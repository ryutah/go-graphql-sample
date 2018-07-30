# GraphQL CRUD examle
CRUD example using GraphQL

## Request example
### List
#### Query
```
query {
  list {
    id,
    name,
    info,
    price
  }
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "query { list { id, name, info, price } }"}' \
  http://localhost:8080/product
```


### Get
#### Query
```
query {
  product(id: 1) {
    id,
    name,
    info,
    price
  }
}
```

#### Curl 
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "query { product(id: 1) { id, name, info, price } }"}' \
  http://localhost:8080/product
```


### Post
#### Query
```
mutation {
  create(name: "foo", info: "bar", price: 1.23) {
    id,
    name,
    info,
    price
  }
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "mutation { create(name: \"foo\", info: \"bar\", price: 1.23) { id, name, info, price } }"}' \
  http://localhost:8080/product
```


### Put
#### Query
```
mutation {
  update(id: 1, price: 1.23) {
    id,
    name,
    info,
    price
  }
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "mutation { update(id: 1, price: 1.23) { id, name, info, price } }"}' \
  http://localhost:8080/product
```


### Delete
#### Query
```
mutation {
  delete(id: 1) {
    id,
    name,
    info,
    price
  }
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "mutation { delete(id: 1) { id, name, info, price } }"}' \
  http://localhost:8080/product
```


### Composite
#### Query
```
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
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "query { list { id } product(id: 1) { id, name, info, price } }"}' \
  http://localhost:8080/product
```


### Aliases
#### Query
```
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
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "query { first: product(id: 1) { id, name, price } second: product(id: 2) { id, name, price } }"}' \
  http://localhost:8080/product
```


### Fragment
#### Query
```
query {
  first: product(id: 1) {
    ...custom
  }
  second: product(id: 2) {
    ...custom
  }
}
fragment custom on Product {
  name,
  info,
  price
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "query { first: product(id: 1) { ...custom } second: product(id: 2) { ...custom } } fragment custom on Product { name, info, price }"}' \
  http://localhost:8080/product
```


### Arguments
#### Query
```
query get($id: Int!) {
  product(id: $id) {
    name
  }
}
```

#### Variables
```json
{
  "id": 1
}
```

#### Curl
```
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"query": "query get($id: Int!) { product(id: $id) { name } }", "variables": {"id": 1}}' \
  http://localhost:8080/product
```

