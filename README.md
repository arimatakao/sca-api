<div align="center">

# Spy Cat Agency API (SCA-API) 🐱

Spy Cat Agency (SCA) is a management application designed to streamline the spy work processes of an agency using cats. It allows SCA to manage their agents (cats), track missions, and handle targets with detailed data entry and updates. The application features CRUD operations for cats and missions, including target assignment and note management, along with database. Built with Golang + Gin framework + PostgreSQL.

Task description here - [TASK.md](./TASK.md)

</div>

# How to run

Before running the application, ensure you create `config.yaml` file:

```yaml
app:
  port: 8080
  base_path: "/api"
  db_url: "postgresql://user:password@127.0.0.1:5432/mycats?sslmode=disable"
```

*see `config_docker.yaml` and `config_example.yaml`*

`127.0.0.1:8081` - pgadmin

## Native

Use `make`:

```sh
make run
```

Or build binary file and run:

```sh
make build
./bin/sca-api -config ./config.yaml
```

## Docker

Run in docker:

```sh
docker compose up -d
```

*this command will use `config_docker.yaml` file as config*