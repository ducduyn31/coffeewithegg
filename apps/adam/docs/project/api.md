# Project Endpoints

## Get All Projects

```graphql
query GetAllProjects($input: ProjectFilter) {
  projects(filters: $input) {
    id
    key
    name
    description
    technologies {
      id
      key
      name
      description
    }
  }
}
```

input:

```json
{
  "input": {
    "name": "adam",
    "technologies": ["golang", "graphql"],
    "count": 10,
    "offset": 0
  }
}
```

## Upsert Project

```graphql
mutation UpsertProject($input: ProjectInput) {
  upsertProject(input: $input) {
    id
    key
    name
    description
    technologies {
      id
      key
      name
      description
    }
  }
}
```

input:

```json
{
  "input": {
    "key": "adam",
    "name": "Adam Project",
    "description": "This project is created as a gateway and coordinator for coffeewithegg (CWE) project. It provides the following features:\n- GraphQL endpoints for `Project` - `Project` refers to my personal projects, such as this one.\n- GraphQL endpoints for `Infrastucture` - `Infrastructure` refers to the deployment plan that I use to host my projects, such as this one.",
    "technologies": ["golang", "graphql"]
  }
}
```
