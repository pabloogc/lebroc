# lebroc
Graphql App written in Go using [gocraft](https://github.com/gocraft/web) to handle http requests. 

All data is fake and automtacilly generated with [mockaroo](https://www.mockaroo.com).

This app uses a real MongoDB hosted in a public test server in Mongolab. Don't destroy it pls.

To execute simply run `go run main.go`, the server will handle incoming requests on `localhost:3000/gql?query= ...`.

Consider using the [Graphiql App](https://github.com/skevy/graphiql-app) to explore the squema, perform complex queries and so on.

![sample using app](http://i.imgur.com/LVIwcSS.png)
