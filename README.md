
## Description
It is a simple web template project, which is written in Go with popular go frameworks.
It uses these frameworks. Gin/ Logrus/ Gqlgen/ Oauth2, OIDC, Casbin, GoReleaser

This template shows how to use RestAPI, Graphql, Oauth2, Authentication, Authorization, Database Connection, 
and build and packaging.

It will build and package in Docker Image. You can configure  
_.goreleaser.yaml_ file for further build and package automation. 

## Prepare Running environment
```bash
brew install goenv goreleaser
goenv install 1.20.1
goenv global 1.20.1
git clone git@github.com:gogo-boot/go-boot.git
cd go-boot && go mod tidy && go run .
```

## Configuration
### Web Application
you can configure the web application.
```bash
vi platform/config/config.yml
```
### Go Releaser
you can configure releasement
```bash
vi .goreleaser.yaml
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
docker run -p 8080:8080 gogo-boot/go-boot 
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
- Health check endpoint
- Server-Side Events
- HTML Template
- Rest API
  - OpenAPI
  - GraphQL
- Logging in Json Format
- Security
  - Oauth2
  - OIDC
  - User Authorization
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

## Http Endpoints & Integration Test
you can open *.http files with IntelliJ and send http request or test by clicking without typing

## Roadmap
- DB Connection 
- Implement OIDC Logout
- Add Unit Test
- CICD automation
- Refactorings 
- Documentation
- Publish Metrics
- Release 1.0

## Reference 
If you like to change configuration or extend it, you can read further document
below.

[Gin Web Framework](https://github.com/gin-gonic/gin)

[Go Graphql Framework / gqlgen](https://github.com/99designs/gqlgen)

[Authorization Code Flow](https://auth0.com/docs/get-started/authentication-and-authorization-flow/authorization-code-flow)

[Authorization Code Flow with Proof Key for Code Exchange (PKCE)](https://auth0.com/docs/get-started/authentication-and-authorization-flow/authorization-code-flow-with-proof-key-for-code-exchange-pkce)

[OIDC Best Practice](https://auth0.com/docs/quickstart/webapp/golang/interactive)

[Go Authentication Framework / OAuth2 for Go](https://github.com/golang/oauth2)

[Go Authentication Framework / Go OIDC](https://github.com/coreos/go-oidc)

[Go Authorization Framework / Casbin](https://github.com/casbin/casbin)

[Go build/packaging manager / GoReleaser](https://github.com/goreleaser/goreleaser)

[Go Config](https://github.com/gookit/config)

[Logrus Logging](https://github.com/sirupsen/logrus)

[Casbin Access Control Model](https://articles.wesionary.team/understanding-casbin-with-different-access-control-model-configurations-faebc60f6da5)