# blacklist-example

- Swagger API server is generated to serve HTTP routes.
- The example handler is in the main.go and serves an in memory example


## Requirements
- golang (new enough version to support go modules)

Run locally...
```
- go run main.go
```
```
http://127.0.0.1:8080/v1/blacklist
```

Run in kubernetes...
```
kubectl apply -f kubernetes/
```


The API should result with...
```
curl http://127.0.0.1:8080/v1/blacklist
{"domains":["piratebay.org","piratebay.io","piratebay.co.uk"]}
```

### Considerations

Currently as this is a toy project the default port is hard coded
```
//Set a default port
server.Port = 8080
```
