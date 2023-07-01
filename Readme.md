# Exercise: Write a simple fizz-buzz REST server.

"The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"".

The output would look like this: ""1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...""."

Your goal is to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters: three integers int1, int2 and limit, and two strings str1 and str2.

- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

The server needs to be:

- Ready for production

- Easy to maintain by other developers

Bonus: add a statistics endpoint allowing users to know what the most frequent request has been. This endpoint should:

- Accept no parameter

- Return the parameters corresponding to the most used request, as well as the number of hits for this request

# Run locally
### Prerequisites
- Go >=1.20
- (optional) Docker

Compilation
```
go build -o lbc-fizzbuzz cmd/main.go
```

Run unit tests
```
go test ./...
```

Start the server
```
./lbc-fizzbuzz
```

Generate Docker image
```
docker build --tag docker-lbc-fizzbuzz .
```

Run from docker
```
docker run -p 8080:8080 docker-lbc-fizzbuzz
```
Example request
```
curl -v -d '{"firstInteger":3, "secondInteger":5, "limit":15, "firstString":"fizz", "secondString":"buzz"}' http://localhost:8080/fizzbuzz
```

Should return
```
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz"]
```


# Notes
- Using gin-gonic as a REST server Framework
- Default values set to listen on localhost:8080, depending on the environment we want to deploy, we might want to override those values via config files. 
- Docker image is big (~1.20GB), but could be optimized.
- Logs are as of today written in stdout + file, depending on the logs collector used, this is meant to change.
- "/stats" endpoint could be replaced in a "real world" by a tool dedicated to perform statistics (Kibana, Splunk...) instead of caching requests in the server.