# Go Chuck Norrisify
## Description
This is a simple RESTful API/web service that integrates with several third party services to provide humour across multiple languages.

## Installation & Run
### Development
* docker-compose

```bash
  docker-compose up -d
```

*kubernetes

- Edit the values.yaml to have appropriate configuration for your deployment

```bash
  helm install -n develop -f k8s/develop/development.values.yaml --name go-chuck-norrisify helm-charts/web-app-k8s
```

### Production
* docker-compose

```bash
  docker-compose -f production.docker-compose.yaml up -d
```
* kubernetes

- Edit the values.yaml to have appropriate configuration for your deployment

```bash
  helm install -n production -f k8s/production/production.values.yaml
```

## Structure
```
├── app  // Slim docker image
│   └── Dockerfile
├── config.ini
├── config.ini.sample
├── controllers // Request controllers based on MVC model
│   ├── init.go
│   └── joke.go
├── docker-compose.yaml
├── docker-entrypoint.sh
├── Dockerfile
├── main.go
├── models // Data models the app interacts with
│   ├── joke.go
│   └── name.go
├── production.docker-compose.yaml
├── README.md
├── test
└── utils // Helper functions used throughout the app
    ├── config.go
    ├── errors.go
    └── http.go
```

## API

#### /
* `GET` : Generate a joke 

## Todo

- [ ] Add ability to pass parameter to translate joke using Google Translate
- [ ] Make convenient wrappers for creating API handlers.
- [ ] Write the tests for all APIs.
- [x] Organize the code with packages
- [ ] Make docs with GoDoc
- [ ] Building a deployment process 
