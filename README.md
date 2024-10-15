# 

## Running migrations

```
export POSTGRES_HOST=localhost \
	&& export POSTGRES_USER=postgres \
	&& export POSTGRES_PASSWORD=mysecretpassword \
	&& export POSTGRES_DB=feedback_db \
make migrate-up
```

### Rolling back migrations

```
export POSTGRES_HOST=localhost \
    && export POSTGRES_USER=postgres \
    && export POSTGRES_PASSWORD=mysecretpassword \
    && export POSTGRES_DB=feedback_db \
make migrate-down   
```