# Go Project Template: Automatically compile as you code.
### Step 1: Copy file

1. Copy `.env.example` file to `.env` 
2. Set GID and UID inside `.env` to your current User and Group IDs. To get these numbers, just type `id` in linux.

The rest of the environment file is described below:
|Name|Default|Information|
|---|---|---|
|UID|1000|User ID of current account|
|GID|1000|Group ID of current account|
|WEB_PORT|3000|Port number to host application under Golang for final node build/* assets.
|TRACE_PORT|3001|Port number for which pprof will run under. `ENVIRONMENT` variable must be set to `local`|
|ENVIRONMENT|local|Can be used anywhere in the application. Currently only used to decide to run a secondary thread in Golang for another tracing webhost (pprof).
|GO_BINARY|"go-web-template"|Name of binary produced by CompileDaemon. If you use this binary in production, you should disable `-race` mode in `./docker/golang/entrypoint.sh` and rebuild with `docker-compose build`
|PORT|8080|Port used to host webpack development server for live reloading. Do not use this server for production applications. Node container should not need to run in production because you should have a pipeline building the project in `./client/build/*` and running that under `WEB_PORT` using Golang above.
|TYPESCRIPT|true|If you remove the client directory using `rm -rf client/*` and re-run `docker-compose up -d`, this will reinstantiate the project with Create React App and use either typescript mode or not.


### Step 2: Run
To start monitoring and automatically building code, just run:
```
docker-compose up -d
```

### Step 3: Code
You can start changing `main.go` or adding and linking files. You may also hack away on the client and view the result at `http://localhost:<PORT>` To see logs, run:
```
docker-compose logs -f
```

### Tips:
1. Race mode is enabled by default. To change this, update `./docker/golang/entrypoint.sh`