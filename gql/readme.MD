*gql server*

To run this application, call `make go_run_gql` on root

```
gql
 ├── webstatic              # Frontend Code, taken from https://github.com/graphql/graphiql/tree/main/examples/graphiql-cdn
 ├── gqlserver              # GQL implementation code
 |    ├──  gqltype.go       # Contains gql object description
 |    ├──  handler.go       # Main handler to translate http connection to gql
 |    ├──  resolver.go      # Contains all function description and bridge to outside of gql application
 |    └──  schema.go        # Contains gql schema description
 └── main.go                # Main program of the app
```
