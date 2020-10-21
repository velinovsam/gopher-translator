# Gopher Translator

A simple web service relying on Gin, which helps with translation of english words and sentences to gopher language.

## Requirements
* **Go 1.15**

## Starting it up


```sh 
# Starts the service
$ go run main.go
```

Passing a ```-port``` param specifies the port the service will use. If omitted, defaults to 1234.



## Endpoints

The service accepts POST and GET requests in json format and returns json upon success.

POST /word  
POST /sentence  
GET /history

## Testing
Tests use the standard Go testing library and can be run by executing  

```sh 
$ go test ./...
```
