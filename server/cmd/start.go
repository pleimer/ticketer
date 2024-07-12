package cmd

import "fmt"

type Start struct {
}

func (s *Start) Execute(args []string) error {
	fmt.Println("start server")

	return nil
}
