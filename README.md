<!-- Migration scripts -->

## Create migration scripts

```bash
    migrate create -ext sql -dir db/migration -seq init_schema
```

## Create Database in the docker exec -it

```bash
    createdb --username=root --owner=root simple_bank
```
