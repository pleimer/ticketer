# Ticketer
Initial Commit

# Codegen
Typescript API client hooks: [Orval](https://orval.dev/overview)
Golang route generation: [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)

# Dep
Vite (install)

# Dev
Generate a new db schema stub (run from `server/` directory): `go run entgo.io/ent/cmd/ent new <Object name>`. Schema names should be singular.

Generate `ent` clients after new schemas are made: 
`go generate ./ent`

# Dev Setup
- create pg db `ticketer`
- run db migration
- codegen
