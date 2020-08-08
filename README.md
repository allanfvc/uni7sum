# Sample REST API with Golang

## How to Run
```
go run main.go
```

## Building docker image
```
docker build -t allanfvc/uni7sum:<VERSION_TAG> . 
```

## Running docker image
```
docker run -it -p 8080:8080 allanfvc/uni7sum:<VERSION_TAG>
```

## Pushing to DockerHub
```
docker login 
docker push allanfvc/uni7sum:<VERSION_TAG>
```

## Docker image
Docker image are available at [DockerHub](https://hub.docker.com/r/allanfvc/uni7sum).
