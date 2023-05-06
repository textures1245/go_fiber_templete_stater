# Sample Go Pattern

# Start

```bash
delete go.mod
delete go.some
go mod init payso-{module_name}
```

## Prepare ENV

### Production

```bash
export $(cat .env | xargs)
```

### Staging

```bash
export $(cat .env_stg | xargs)
```

### Run

```bash
go run .
```
