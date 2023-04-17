
## Description
It is simple web template project, which is written in Go with popular go frameworks.
It uses these frameworks. Gin/ GORM/ Graphql-Go/ Oauth2 for Go/ Casbin/ GoReleaser

This template shows how to use RestAPI, Graphql, Oauth2, Authentication, Authorization, Database Connection, 
and build and packaging.

It will be built and packaged in Docker Image. If you don't want docker image build, 
you can configure _.goreleaser.yaml_ file.

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

## Source code start point
you can start review from _main.go_ file and follow further for more detail.

## Remove Redundant
This uses several frameworks. RestAPI, Graphql, Oauth2, GORM... 
If you don't need of it, you can delete the directory and update _main.go_ file.
```agsl
go mod tidy
```
then the unused framework will be removed and the build package will be slimmer.

## Feature
- Rest API
- GraphQL
- DB Connection 
- Logging in Json Format
- Oauth2
- Configuration
- Multi architect build 
- Dockerizing

## Integration Test
you can open *.http files with IntelliJ and send http request or test by clicking without typing

## Roadmap
- Implement OpenAPI auto RestAPI Generation 
- Implement Role based authorization
- Implement Logging
- Implement OIDC
- Unit Test
- Magics

## Reference 
If you like to change configuration or extend it, you can read further document
from the following reference site.

[Gin Web Framework](https://github.com/gin-gonic/gin)

[Go ORM Framework / GORM](https://github.com/go-gorm/gorm)

[Go Graphql Framework / graphql-go](https://github.com/graph-gophers/graphql-go)

[Go Authentification Framework / OAuth2 for Go](https://github.com/golang/oauth2)

[Go Authorization Framework / Casbin](https://github.com/casbin/casbin)

[Go build/packaging manager / GoReleaser](https://github.com/goreleaser/goreleaser)

[Go Config](https://github.com/gookit/config)

[Logrus Logging](https://github.com/sirupsen/logrus)