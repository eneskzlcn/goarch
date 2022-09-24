### GOARCH

### About Project
That project is a cli tool that allows user to create an empty
folder structure for given architecture and technologies. For now,
only supports microservices and for GOARCH = "arm64".

### About .arch files
These are the files used as template to create new files with that template content.



### Microservice Structure Look like

```
📦root
┣ 📂.cd
┃ ┣ 📂deployment-artifacts
┃ ┃ ┣ 📜pv-claim.yaml
┃ ┃ ┣ 📜pv.yaml
┃ ┃ ┗ 📜service.yaml
┃ ┗ 📜cd.yaml
┣ 📂.dev
┃ ┣ 📜dev.yaml
┃ ┣ 📜prod.yaml
┃ ┣ 📜qa.yaml
┃ ┗ 📜test.yaml
┣ 📂cmd
┃ ┗ 📜main.go
┣ 📂internal
┃ ┣ 📂client
┃ ┃ ┣ 📂httpclient
┃ ┃ ┃ ┣ 📜httpclient.go
┃ ┃ ┃ ┗ 📜httpclient_test.go
┃ ┃ ┣ 📂kafkaclient
┃ ┃ ┃ ┣ 📜kafkaclient.go
┃ ┃ ┃ ┗ 📜kafkaclient_test.go
┃ ┃ ┗ 📂rabbitmqclient
┃ ┃ ┃ ┣ 📜rabbitmqclient.go
┃ ┃ ┃ ┗ 📜rabbitmqclient_test.go
┃ ┣ 📂config
┃ ┃ ┣ 📜config.go
┃ ┃ ┣ 📜config_test.go
┃ ┃ ┗ 📜db.go
┃ ┣ 📂domain
┃ ┃ ┣ 📜domain.go
┃ ┃ ┣ 📜domain_test.go
┃ ┃ ┣ 📜handler.go
┃ ┃ ┣ 📜handler_test.go
┃ ┃ ┣ 📜repository.go
┃ ┃ ┣ 📜repository_test.go
┃ ┃ ┣ 📜request.go
┃ ┃ ┣ 📜request_test.go
┃ ┃ ┣ 📜response.go
┃ ┃ ┣ 📜response_test.go
┃ ┃ ┣ 📜service.go
┃ ┃ ┗ 📜service_test.go
┃ ┣ 📂mocks
┃ ┃ ┣ 📜mock_user_repository.go
┃ ┃ ┗ 📜mock_user_repository_test.go
┃ ┗ 📂util
┃ ┃ ┗ 📂ctxutil
┃ ┃ ┃ ┣ 📜ctxutil.go
┃ ┃ ┃ ┗ 📜ctxutil_test.go
┗ 📂seed
┃ ┣ 📂cmd
┃ ┃ ┗ 📜main.go
┃ ┣ 📜create-seed.sql
┃ ┣ 📜drop-seed.sql
┃ ┣ 📜seed.go
┃ ┗ 📜seed_test.go
```