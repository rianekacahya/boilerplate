# Apps Service

## Document Outline

- [Description](#description)
- [Technology Stacks](#stacks)
- [Command Parameters](#command-parameters)

## <a name="description"></a>Description

Specifically handling multi merchant apps initialisations and settings

## <a name="stacks"></a>Stack Usage

- Go (v1.14.x+)
- Go-Chi (v4.0.3)
- PostgreSQL (v11.x.x+)
- Redis (v5.x.x+)

## <a name="available-endpoints"></a>Available Endpoints:

| Endpoint                     | Method | Notes                                                                                                                                                                              |
|------------------------------|--------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `/v1/oauth2/token`           | POST   | Generate new request token                                                                                                                                                         |
| `/v1/auth/login`             | POST   | Sign-in into application                                                                                                                                                           |

Details that includes request & response, can be referred to Postman documentation [here](https://documenter.getpostman.com/view/6817990/SzzdDgQ6)

## <a name="command-parameters"></a>Command Parameters
Example using the command

```shell script
go run cmd/main.go rest
```

List Available Parameter

| Parameter               | Description                                                                                 |
|-------------------------|---------------------------------------------------------------------------------------------|
| `rest`                  | Run REST-API Server                                                                         |
| `migrate:up`            | Execute Database Migration Up                                                               |
| `migrate:down`          | Execute Database Migration Down                                                             |