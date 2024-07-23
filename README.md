# Ticketer
A simple ticket management tool

This service watches for updates coming into an email inbox configured with Nylas in a temporal workflow. Messages that begin a new thread will result in a new ticket being created. 

For this project, no thread or message information is stored in the application DB. Rather, all messaging state is queried from the mail server. This does make for some fragile interations such as replying to a thread through the UI. This can be improved by putting this operation in a temporal workflow of its own, making adjustments for transactionality and even tracking some messageing/threading data in the application db.

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

4. Visit `localhost:8080`

# Creating Tickets

Tickets are created when a new email arrives in the inbox of the account configured with Nylas. To open a ticket,
send a message to the address with an alternate email. After 60s, a new ticket will be created and a notification sent to the originator.

Messages can be added to the ticket by replying to the email thread created for the ticket.

The Admin email (configured with Nylas) can reply to the thread through the ticket page in the UI.

Adjust status, priority, assignee by visiting the ticket page.

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

I did not get around to adding integration test suites for this project. There is minimal need for unit tests, but integration tests and e2e tests are necessary.