
# Graphql practice

## graphql fundamentals

### Schema Definition Language

GraphQL has its own type system that’s used to define the schema of an API. The syntax for writing schemas is called Schema Definition Language (SDL).

The schema is one of the most important concepts when working with a GraphQL API. It specifies the capabilities of the API and defines how clients can request the data. It is often seen as a contract between the server and client.

Generally, a schema is simply a collection of GraphQL types. However, when writing the schema for an API, there are some special root types:

1. type Query { ... }
2. type Mutation { ... }
3. type Subscription { ... }
4. type Input

#### General types

For example below we defined two types, person and posts and if they are related we express them in each others defenition meaning a person can have many posts so we put posts: [Post!]!
and on the other end in post we put auther: Person!
meaning each post is written by a Person

```json

type Person {
  name: String!
  age: Int!
  posts: [Post!]!
}

type Post {
  title: String!
  author: Person!
}

```

#### Query types

As shown above we defined our types now for the client to be able to [query](#query-data-from-api) these data we have to create a query type for example we define allPersons Query type with last attribute

```json

type Query {
  allPersons(last: Int): [Person!]!
}

```

#### Mutations types

Now that we can query the data we might want to be able to update the data of backend  server through the API ([mutaitions](#modify-data-with-mutations)) to deine a way to mutate data throuhg the API we define a mutations type

```json

type Mutation {
  createPerson(name: String!, age: Int!): Person!
}

```

#### Subscriptions types

Finally to implement the [subscription](#subscribe-to-data) function (the function in which the client is modified by the server on data update) we define a subscription type

```json

type Subscription {
  newPerson: Person!
}

```

#### Input type

In a GraphQL schema, input is a special type used for defining structured inputs to queries or mutations. It allows you to group multiple fields into a single object that can be passed as an argument.

```json

input InputTypeName {
  fieldName1: FieldType1
  fieldName2: FieldType2
}


```

For example below is an implementation of an input type representing a user

```json
input CreateUserInput {
  name: String!
  email: String!
  age: Int
}

type Mutation {
  createUser(input: CreateUserInput!): User
}

```

### Query data from API

In graphql we explicitly say what we need and we only get that nothing more nothing less and thats one of the advantages of graphql over REST it prevents over or under fetching
an exaple of a query is below

```json

{
  allPersons {
    name
  }
}

```

the respone we recive is below

```json

{
  "allPersons": [
    { "name": "Johnny" },
    { "name": "Sarah" },
    { "name": "Alice" }
  ]
}

```

#### Query with parameters

we can query data with specific parameters for example the query below will return the last 2 persons

```json

{
  allPersons(last: 2) {
    name
  }
}

```

### Modify data with mutations

You can create, update, delete data of the backend server by querying API this is called mutations

To modify data we use **mutation** keyword in our query and pass in arguments for exaple to create a person we query bellow

```json

mutation {
  createPerson(name: "Bob", age: 36)
}

```

### Subscribe to data

Another important requirement for many applications today is to have a realtime connection to the server in order to get immediately informed about important events. For this use case, GraphQL offers the concept of subscriptions.

When a client subscribes to an event, it will initiate and hold a steady connection to the server. Whenever that particular event then actually happens, the server pushes the corresponding data to the client. Unlike queries and mutations that follow a typical “request-response-cycle”, subscriptions represent a stream of data sent over to the client.

Subscriptions are written using the same syntax as queries and mutations. Here’s an example where we subscribe on events happening on the Person type:

```json

subscription {
  newPerson {
    name
    age
  }
}

```

### Graphql server implementation architecture

1. GraphQL server with a connected database
2. GraphQL server that is a thin layer in front of a number of third party or legacy systems and integrates them through a single GraphQL API
3. A hybrid approach of a connected database and third party or legacy systems that can all be accessed through the same GraphQL API

#### GraphQL server with a connected database

This architecture will be the most common for greenfield projects. In the setup, you have a single (web) server that implements the GraphQL specification. When a query arrives at the GraphQL server, the server reads the query’s payload and fetches the required information from the database. This is called resolving the query. It then constructs the response object as described in the official specification and returns it to the client.

#### GraphQL server that is a thin layer in front of a number of third party or legacy systems and integrates them through a single GraphQL API

Another major use case for GraphQL is the integration of multiple existing systems behind a single, coherent GraphQL API. This is particularly compelling for companies with legacy infrastructures and many different APIs that have grown over years and now impose a high maintenance burden. One major problem with these legacy systems is that they make it practically impossible to build innovative products that need access to multiple systems.

In that context, GraphQL can be used to unify these existing systems and hide their complexity behind a nice GraphQL API. This way, new client applications can be developed that simply talk to the GraphQL server to fetch the data they need. The GraphQL server is then responsible for fetching the data from the existing systems and package it up in the GraphQL response format.

#### A hybrid approach of a connected database and third party or legacy systems that can all be accessed through the same GraphQL API

Finally, it’s possible to combine the two approaches and build a GraphQL server that has a connected database but still talks to legacy or third—party systems.

### Resolver functions

When the server receives a query, it will call all the functions for the fields that are specified in the query’s payload. It thus resolves the query and is able to retrieve the correct data for each field. Once all resolvers returned, the server will package data up in the format that was described by the query and send it back to the client.

## gplgen

The package to create production grade graphql server

### Prerequsites

You can install the package to use locally with

```console
go get github.com/99designs/gqlgen
go install  github.com/99designs/gqlgen

```

and run the rest of the commands with gqlgen command line or use go run to run the requiered commands as done below

### Setup gqlgen repository

to start with it first get the package and run init

```console
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init

```

This will generate files and code to create a graphql server using gplgen

### Define our schema

after that we modift graph\schema.graphqls to modify our schema and use the command bellow to generate go code and models

```console

go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen generate

```

### Implement resolvers

now that we generated our models and some code we need to implement out resolvers (function that will get the data querried from the client to return it to the client)
The resolvers are in graph\schema.resolvers.go

## Use the gqlgen

Now you can use the playground in [http://localhost:port](http://localhost:port) where port is in /graph/server.go and defaultPort variable it is set by default to 8080
And the graphql endpoint is located in /query path changeable in /graph/server.go
