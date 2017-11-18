package subcommands

import (
  "fmt"
)

type CurrentDimension struct{}

func (c CurrentDimension) Run(args []string) error {
  fmt.Println("CurrentDimension called with args: ", args)
  return nil
}