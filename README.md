# Simple Go Web App
This is a small example for running a go web application.  
`fmt` is being used to the output and `net/http` to provide a HTTP server.
It returns the used URL path.

## Usage
Using [docker](https://docs.docker.com/engine/install/) you can simply run this app with:
``` sh
docker compose up -d
```
Stopping:
``` sh
docker compose down
```

Alternatively with [go](https://go.dev/doc/install) installed, the script can be run directly:
```
go run App/app.go
```

After starting the container or script you can access the site at `http://localhost:8000` or some path like `http://localhost:8000/test`.
