### GOARCH

### About Project
Goarch is a cli tool that allows user to create an empty
folder structure for given architecture. 

### Usage
After you download the binary and attach it to your PATH you can
directly use `goarch ms` to create a microservice structure
shown in the section `Microservice Strucuture Look Like` or you can
can use `goarch nlb` to create a n-layered backend application that uses
fiber as a server technology.

### About .arch files
These are the files used as template to create new files with that template content.



### Microservice Structure Look like

```
📦try
 ┣ 📂.cd
 ┃ ┣ 📂deployment-artifacts
 ┃ ┃ ┣ 📜deployment.yaml
 ┃ ┃ ┣ 📜pv-claim.yaml
 ┃ ┃ ┣ 📜pv.yaml
 ┃ ┃ ┗ 📜service.yaml
 ┃ ┗ 📜cd.yaml
 ┣ 📂.dev
 ┃ ┣ 📜dev.yaml
 ┃ ┣ 📜local.yaml
 ┃ ┣ 📜prod.yaml
 ┃ ┗ 📜qa.yaml
 ┣ 📂cmd
 ┃ ┗ 📜main.go
 ┣ 📂internal
 ┃ ┣ 📂client
 ┃ ┃ ┣ 📂httpclient
 ┃ ┃ ┃ ┣ 📜httpclient.go
 ┃ ┃ ┃ ┗ 📜httpclient_test_test.go
 ┃ ┃ ┗ 📂rabbitclient
 ┃ ┃ ┃ ┣ 📜rabbitclient.go
 ┃ ┃ ┃ ┗ 📜rabbitclient_test_test.go
 ┃ ┣ 📂config
 ┃ ┃ ┣ 📜config.go
 ┃ ┃ ┣ 📜config_test_test.go
 ┃ ┃ ┗ 📜db.go
 ┃ ┣ 📂domain
 ┃ ┃ ┣ 📜domain.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜handler_test_test.go
 ┃ ┃ ┣ 📜repository.go
 ┃ ┃ ┣ 📜repository_test.go
 ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┗ 📜service_test_test.go
 ┃ ┣ 📂mocks
 ┃ ┃ ┗ 📂domain
 ┃ ┃ ┃ ┗ 📜mock_domain_repository.go
 ┃ ┗ 📂util
 ┃ ┃ ┗ 📂ctxutil
 ┃ ┃ ┃ ┗ 📜ctxutil.go
 ┣ 📂logger
 ┃ ┣ 📜logger.go
 ┃ ┣ 📜zap_logger_adapter.go
 ┃ ┗ 📜zap_logger_adapter_test.go
 ┣ 📂postgres
 ┃ ┣ 📜mock_postgres.go
 ┃ ┣ 📜postgres.go
 ┃ ┗ 📜postgres_test_test.go
 ┣ 📂rabbitmq
 ┃ ┗ 📜rabbitmq.go
 ┣ 📂seed
 ┃ ┣ 📂cmd
 ┃ ┃ ┗ 📜main.go
 ┃ ┣ 📜create_seed.sql
 ┃ ┣ 📜drop_seed.sql
 ┃ ┣ 📜readme.md
 ┃ ┣ 📜seed.go
 ┃ ┗ 📜seed_test_test.go
 ┣ 📜go.mod
 ┗ 📜main.go
```

### N-Layered Backend Architecture Look Like
```
📦try
 ┣ 📂.cd
 ┃ ┣ 📂deployment-artifacts
 ┃ ┃ ┣ 📜deployment.yaml
 ┃ ┃ ┣ 📜pv-claim.yaml
 ┃ ┃ ┣ 📜pv.yaml
 ┃ ┃ ┗ 📜service.yaml
 ┃ ┗ 📜cd.yaml
 ┣ 📂.dev
 ┃ ┣ 📜dev.yaml
 ┃ ┣ 📜local.yaml
 ┃ ┣ 📜prod.yaml
 ┃ ┗ 📜qa.yaml
 ┣ 📂cmd
 ┃ ┗ 📜main.go
 ┣ 📂internal
 ┃ ┣ 📂api
 ┃ ┃ ┣ 📜api.go
 ┃ ┃ ┗ 📜api_test_test.go
 ┃ ┣ 📂client
 ┃ ┃ ┣ 📂httpclient
 ┃ ┃ ┃ ┣ 📜httpclient.go
 ┃ ┃ ┃ ┗ 📜httpclient_test_test.go
 ┃ ┃ ┗ 📂rabbitclient
 ┃ ┃ ┃ ┣ 📜rabbitclient.go
 ┃ ┃ ┃ ┗ 📜rabbitclient_test_test.go
 ┃ ┣ 📂config
 ┃ ┃ ┣ 📜config.go
 ┃ ┃ ┣ 📜config_test_test.go
 ┃ ┃ ┗ 📜db.go
 ┃ ┣ 📂core
 ┃ ┃ ┣ 📂client
 ┃ ┃ ┃ ┣ 📂httpclient
 ┃ ┃ ┃ ┃ ┣ 📜httpclient.go
 ┃ ┃ ┃ ┃ ┗ 📜httpclient_test_test.go
 ┃ ┃ ┃ ┗ 📂rabbitclient
 ┃ ┃ ┃ ┃ ┣ 📜rabbitclient.go
 ┃ ┃ ┃ ┃ ┗ 📜rabbitclient_test_test.go
 ┃ ┃ ┣ 📂logger
 ┃ ┃ ┃ ┣ 📜logger.go
 ┃ ┃ ┃ ┣ 📜zap_logger_adapter.go
 ┃ ┃ ┃ ┗ 📜zap_logger_adapter_test.go
 ┃ ┃ ┣ 📂postgres
 ┃ ┃ ┃ ┣ 📜mock_postgres.go
 ┃ ┃ ┃ ┣ 📜postgres.go
 ┃ ┃ ┃ ┗ 📜postgres_test_test.go
 ┃ ┃ ┣ 📂rabbitmq
 ┃ ┃ ┃ ┣ 📜config.go
 ┃ ┃ ┃ ┣ 📜config_test_test.go
 ┃ ┃ ┃ ┗ 📜db.go
 ┃ ┃ ┣ 📂server
 ┃ ┃ ┃ ┗ 📜server.go
 ┃ ┃ ┗ 📂util
 ┃ ┃ ┃ ┗ 📂ctxutil
 ┃ ┃ ┃ ┃ ┗ 📜ctxutil.go
 ┃ ┣ 📂entity
 ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂repository
 ┃ ┃ ┣ 📜repository.go
 ┃ ┃ ┗ 📜repository_test_test.go
 ┃ ┗ 📂service
 ┃ ┃ ┣ 📜service.go
 ┃ ┃ ┗ 📜service_test_test.go
 ┣ 📂seed
 ┃ ┣ 📂cmd
 ┃ ┃ ┗ 📜main.go
 ┃ ┣ 📜create_seed.sql
 ┃ ┣ 📜drop_seed.sql
 ┃ ┣ 📜readme.md
 ┃ ┣ 📜seed.go
 ┃ ┗ 📜seed_test_test.go
 ┣ 📜go.mod
 ┗ 📜main.go
```