## Beep Saúde Ops Stats POC

Repositório de prova de conceito para processamento de eventos e dados da operação e transforma-los em métricas consumiveis pelo prometheus para 
exibição em torre de monitoramento

#### DEPENDENCIES

- Golang v1.23.0 (use gvm)

#### Running the application (DEV)

Running the code directly:
```bash
go run ./cmd/main.go 
```

```bash
make run
```

this should run some events scenarios and spin up the /metrics endpoint.

After the events finish being generated, call the metrics api to see the metrics generated from it

```
curl http://localhost:8080/metrics
```

#### Building the application

Using makefile (Recommended)
```bash
make build
```

Building manually from source:
```bash
go build -ldflags="-s -w" -trimpath -o ./bin/beep_ops_stats ./cmd/main.go
```




