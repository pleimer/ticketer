# Ticketer
A simple ticket management tool


# Run locally

1. Copy the sample `env` file
```bash
cp .env.example .env
```

2. Adjust variables in the new `.env` according to your Nylas account configuration and source them:
```bash
source .env
```

3. Run services with Docker Compose

```bash
docker compose up -d
```

# Development

## Codegen

Codegen is done in three steps:
1. Generate OpenAPI spec from entgo schema
2. Generate server stubs from OpenAPI schemas with [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
3. Generate typescript react-query client stubs using [orval](https://orval.dev/)

Step 1 and 2 do not work seemlessly - the full OpenAPI spec is not generated from the entgo schema. This generation was used to build the initial OpenAPI schema, but further adaptations were made to it after
due to limitaions in the `entoas` and `orval` code generators.

See the `Makefile` for relevant commands.

## DB Migrations

DB migrations run automaticall as part of the entgo startup sequence. In a production
environment, this will be moved to versioned migrations that run apart from server
initialization.

Migrations can still be run separately with `make db-migrate`

## Testing

I did not get around to adding integration test suites for this project. There is not much need for unit tests, but integration tests and e2e tests would be most useful.