
## Description
It is a simple web template project, which is written in Go with popular go frameworks.
It uses these frameworks. Gin/ Logrus/ Gqlgen/ Oauth2, OIDC for Go/ Casbin/ GoReleaser

This template shows how to use RestAPI, Graphql, Oauth2, Authentication, Authorization, Database Connection, 
and build and packaging.

It will build and package the binary in Docker Image. You can configure  
_.goreleaser.yaml_ file for further build and package automation. 

## Prepare Running environment
```bash
brew install goenv goreleaser
goenv install 1.20.1
goenv global 1.20.1
git clone git@github.com:mgcos1231/go-boot.git
cd go-boot && go mod tidy && go run .
```
## Build 
```bash
go build .
```
## Deployment
Deployment test
```bash
goreleaser release --clean --snapshot --skip-publish
```
Deployment on Prod.
you need to set a GitHub token 
```bash
export GITHUB_TOKEN=xxxxxx
git tag 0.0.1
goreleaser release --clean
```

## Running

```bash
docker run -p 8080:8080 mgcos1231/go-boot 
```

## Source code start point
you can start review from _main.go_ file and follow further for more detail.

## Update Dependencies and organize dependencies
This uses several frameworks. RestAPI, Graphql, Oauth2, GORM... 
If you don't need of it, you can delete the directory and update _main.go_ file.
```bash
go get -u
go mod tidy
```
then the unused framework will be removed and the build package will be slimmer.

## Feature
- HTML Template
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

```bash
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
- implement Html Streaming
- DB Connection 
- implement Health check endpoint
- Implement Role based authorization
- Add Unit Test
- CICD automation
- Bug fix : 
  - critical : login per context safe / per user safe
- Refactorings / Magics
- Documentation
- Release

## Reference 
If you like to change configuration or extend it, you can read further document
from the following reference site.

[Gin Web Framework](https://github.com/gin-gonic/gin)

[Go ORM Framework / GORM](https://github.com/go-gorm/gorm)

[Go Graphql Framework / graphql-go](https://github.com/graph-gophers/graphql-go)

[Go Authentication Framework / OAuth2 for Go](https://github.com/golang/oauth2)

[Go Authentication Framework / Go OIDC](https://github.com/coreos/go-oidc)

[Go Authorization Framework / Casbin](https://github.com/casbin/casbin)

[Go build/packaging manager / GoReleaser](https://github.com/goreleaser/goreleaser)

[Go Config](https://github.com/gookit/config)

[Logrus Logging](https://github.com/sirupsen/logrus)