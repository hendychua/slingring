package subcommands

import (
  "fmt"
)

type ListDimensions struct{}

func (l ListDimensions) Run(args []string) error {
  fmt.Println("ListDimension called with args: ", args)
  return nil
}