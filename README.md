
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
docker run -p 8080:8080 mgcos1231/go-boot 
```

## Source code start point
you can start review from _main.go_ file and follow further for more detail.

## Update Dependencies and organize dependencies
This uses several frameworks. RestAPI, Graphql, Oauth2, GORM... 
If you don't need of it, you can delete the directory and update _main.go_ file.
```agsl
go get -u
go mod tidy
```
then the unused framework will be removed and the build package will be slimmer.

## Feature
- Rest API
  - OpenAPI
  - GraphQL
- Logging in Json Format
- Security
  - Oauth2
  - OIDC
- Configuration
- Multi architect build 
- Dockerizing

## Generate OpenAPI Code

```azure
# Install by npm
# npm install @openapitools/openapi-generator-cli -g
# Install by brew
# brew install openapi-generator
  
curl -O https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml
openapi-generator-cli generate -i petstore.yaml -g go-gin-server -o ./opeapi-gen \
--global-property=apiDocs=false,modelDocs=false \
--additional-properties=apiPath=openapi,packageName=openapi
```

adjust router group in the main.go file.

## Integration Test
you can open *.http files with IntelliJ and send http request or test by clicking without typing

## Roadmap
- DB Connection 
- Implement Role based authorization
- Unit Test
- Bug fix : 
  - critical : login per context safe / per user safe
- Refactorings / Magics
- Release

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