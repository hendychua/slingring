package subcommands

import (
  "fmt"
)

type CreateDimension struct{}

func (c CreateDimension) Run(args []string) error {
  fmt.Println("CreateDimension called with args: ", args)
  return nil
}