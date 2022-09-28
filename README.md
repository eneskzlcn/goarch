### GOARCH

### About Project
Goarch is a cli tool that allows user to create a folder
structure for golang projects by given architecture. 

### How To Install
After you download the latest binary from releases section, you 
can add it to you PATH to use from anywhere in you computer.

### Usage

```
goarch <architecture>
```
The only command in that tool is the name of the architecture you
want to create.
- **ms:** Stands for domain driven designed microservice architecture
that you can see the result directory structure from below `Microservice Structure Look Like` section. Calling
`goarch ms` in the directory you want to create the architecture is enough to make it work.

- **nlb** Stands for n-layered backend architecture that you can see the result directory
structure from below `N-Layered Backend Architecture Looks Like` section. Calling `goarch nlb`
in the directory you want to create the architecture is enough to make it work.

- **nlw** Stands for n-layered web application architecture that you can see the result directory
  structure from below `N-Layered Web Architecture Looks Like` section. Calling `goarch nlb`
  in the directory you want to create the architecture is enough to make it work.

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

``` N-Layered Web Application Looks Like
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
 ┃ ┣ 📂core
 ┃ ┃ ┣ 📂cache
 ┃ ┃ ┃ ┣ 📜cache.go
 ┃ ┃ ┃ ┗ 📜gcache_adapter.go
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
 ┃ ┃ ┣ 📂mail
 ┃ ┃ ┃ ┣ 📜gomail_adapter.go
 ┃ ┃ ┃ ┗ 📜mail.go
 ┃ ┃ ┣ 📂postgres
 ┃ ┃ ┃ ┣ 📜mock_postgres.go
 ┃ ┃ ┃ ┣ 📜postgres.go
 ┃ ┃ ┃ ┗ 📜postgres_test_test.go
 ┃ ┃ ┣ 📂rabbitmq
 ┃ ┃ ┃ ┣ 📜config.go
 ┃ ┃ ┃ ┣ 📜config_test_test.go
 ┃ ┃ ┃ ┗ 📜db.go
 ┃ ┃ ┣ 📂router
 ┃ ┃ ┃ ┣ 📜mux_router_adapter.go
 ┃ ┃ ┃ ┗ 📜router.go
 ┃ ┃ ┣ 📂server
 ┃ ┃ ┃ ┗ 📜server.go
 ┃ ┃ ┣ 📂session
 ┃ ┃ ┃ ┣ 📜college_session_adapter.go
 ┃ ┃ ┃ ┗ 📜session.go
 ┃ ┃ ┗ 📂util
 ┃ ┃ ┃ ┗ 📂ctxutil
 ┃ ┃ ┃ ┃ ┗ 📜ctxutil.go
 ┃ ┣ 📂entity
 ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂repository
 ┃ ┃ ┣ 📜repository.go
 ┃ ┃ ┗ 📜repository_test_test.go
 ┃ ┣ 📂service
 ┃ ┃ ┣ 📜repository.go
 ┃ ┃ ┗ 📜repository_test_test.go
 ┃ ┗ 📂web
 ┃ ┃ ┣ 📂template
 ┃ ┃ ┃ ┣ 📂include
 ┃ ┃ ┃ ┃ ┣ 📜README.md
 ┃ ┃ ┃ ┃ ┣ 📜header.gohtml
 ┃ ┃ ┃ ┃ ┗ 📜layout.gohtml
 ┃ ┃ ┃ ┗ 📜home.gohtml
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜home.go
 ┃ ┃ ┣ 📜template.go
 ┃ ┃ ┗ 📜web.go
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