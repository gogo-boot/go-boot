
## Description
It is simple web template project, which is written in Go with Gin framework.
It is designed for microservice.
The packaging is small only 8 Mib. As Gin is very fast framework. it inherits its character. 

## Deployment
Deployment test
```agsl
goreleaser release --clean --snapshot
```
Deployment on Prod.
you need to set github token 
```agsl
export GITHUB_TOKEN=xxxxxx
git tag 0.0.1
goreleaser release --clean
```

## Running

```agsl
docker run -p 8080:8080 jinwoo/go-web-template
```

## request.http file
you can open this file with IntelliJ and send http request by clicking without nasty typing