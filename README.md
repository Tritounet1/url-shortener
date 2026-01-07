# URL shortener

## Configure .env

```
cp .env.example .env
```

and change it.

```
MONGODB_URI=...
JWT_SECRET_KEY=...
PORT=...
```

## Run dev

Start the application :

```
go run main
```

Start the application with [air]((https://github.com/air-verse/air)) :

```sh
air
```

For automatically hot reload the app when it changed.

## Structure of code

```
├─📁 models
|  ├─ url.go
|  ├─ user.go
|  └─ visitor.go
|
├─📁 routes
|  ├─ 📁 auth
|  │   ├─ login.go
|  │   └─ register.go
|  └─ default.go
|  └─ routes.go
|  └─ url.go
│
├─📁 services
|  └─ pooler.go
│
├─📁 utils
|  └─ client.go
|  └─ utils.go
|
├─🔑 .env
├─🐳 docker-compose.yml
├─📦 main.go
├─📝 TODO.md
```

## Run database

```sh
docker compose up --build -d
```

## Container shell access and viewing MongoDB logs

The `docker exec` command allows you to run commands inside a Docker container. 

The following command line will give you a bash shell inside your `mongo` container:

```sh
docker exec -it backend-mongo-1 bash
```

The MongoDB Server log is available through Docker's container log

```sh
docker logs backend-mongo-1
```

## Links

- [Mongo official image](https://hub.docker.com/_/mongo)
- [Getting start with mongodb in Go](https://www.mongodb.com/docs/drivers/go/current/get-started/)
- [For emojis](https://emojipedia.org/)