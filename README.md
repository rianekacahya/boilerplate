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

Details that includes request & response, can be referred to Apiary documentation [here](https://app.apiary.io/boilerplate10)

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