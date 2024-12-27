# Temperature by zipcode

Getting weather of location by zipcode in an awesome tool

## Configuration

```bash
go mod tidy
```

### Running

Configure your [Weather API](https://weatherapi.com/) key on `.env` file like the example `.env.example`, it is required.

```bash
go run cmd/main.go
```
#### Or
```bash
docker compose up -d
```

The server will be running on port `:8080`

## How to use

### Local

```bash
curl http://localhost:8080/28928728/temperature
```

### GCP
```bash
curl https://golang-temp-by-zipcode-ea5pytw5bq-uc.a.run.app/28928728/temperature
```

The path parameter zipcode (28928728 on example), is required