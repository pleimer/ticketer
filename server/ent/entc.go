//go:build ignore
// +build ignore

package main

import (
	"log"
	"os"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	file, err := os.Create("../../internal/api/tickets_v2.json")
	if err != nil {
		log.Fatalf("Error opening file: %w", err)
		return
	}
	defer file.Close() // Ensure the file is closed when we're done

	ex, err := entoas.NewExtension(
		entoas.WriteTo(file),
		entoas.SimpleModels(),
		entoas.AllowClientUUIDs(),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
