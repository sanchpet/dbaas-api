# DBaaS API Demo

> Database as a Service (DBaaS) is a cloud computing service model that allows users to access and manage databases without the complexities associated with traditional database management.

This repository contains a **demo example** for creating DBaaS instances in k8s. For now i desided to host only demo version of the app because the full version is not ready enough to be published.

## Includes:

- The idiomatic structure based on the resource-oriented design.
- The usage of Docker, Docker compose and Alpine images on development.
- Healthcheck and CRUD API implementations with OpenAPI specifications.
- The usage of [Zerolog](https://github.com/rs/zerolog) as the centralized Syslog logger.
- The usage of [Validator.v10](https://github.com/go-playground/validator) as the form validator.

## 🚀 Endpoints

| Name            | HTTP Method | Route              |
|-----------------|-------------|--------------------|
| Health          | GET         | /livez             |
|                 |             |                    |
| List Instances  | GET         | /v1/instances      |
| Create Instance | POST        | /v1/instances      |
| Get Instance    | GET         | /v1/instances/{id} |
| Delete Instance | DELETE      | /v1/instances/{id} |

💡 [swaggo/swag](https://github.com/swaggo/swag) : `swag init -g cmd/api/main.go -o .swagger -ot yaml`

## 📦 Container image sizes

- API
    - Development environment: 1.24GB
    - Production environment: 41.5MB ; 💡`docker build -f prod.Dockerfile . -t myapp_app`

## 📁 Project structure

```shell
dbaas-api/
├── Dockerfile
├── LICENSE
├── README.md
├── api
│   ├── resource
│   │   ├── common
│   │   │   ├── err
│   │   │   │   └── err.go
│   │   │   └── log
│   │   │       └── log.go
│   │   ├── health
│   │   │   └── handler.go
│   │   └── instance
│   │       ├── handler.go
│   │       ├── k8sapi.go
│   │       └── model.go
│   └── router
│       ├── middleware
│       │   ├── content_type.go
│       │   ├── content_type_test.go
│       │   ├── request_id.go
│       │   ├── request_id_test.go
│       │   └── requestlog
│       │       ├── handler.go
│       │       └── log_entry.go
│       └── router.go
├── bin
│   ├── gofumpt
│   ├── govulncheck
│   ├── staticcheck
│   └── swag
├── cmd
│   └── api
│       └── main.go
├── config
│   └── config.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── prod.Dockerfile
└── util
    ├── ctx
    │   └── ctx.go
    ├── logger
    │   └── logger.go
    ├── test
    │   └── test.go
    └── validator
        ├── validator.go
        └── validator_test.go
```
