package subcommands

import (
  "fmt"
)

type AddProject struct{}

func (a AddProject) Run(currentDimension string, args []string) error {
  fmt.Println("AddProject called with currentDimension: '", currentDimension, "', args: ", args)
  return nil
}