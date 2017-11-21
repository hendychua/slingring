package subcommands

import (
  "fmt"
)

// AddProject is a command that adds a project path to a Dimension.
type AddProject struct{}

// Run adds a project path to a Dimension.
// If dimensionToWorkOn happens to be the currently set dimension, this function
// will also attempt to jump to the dimension for the newly added project.
func (a AddProject) Run(dimensionToWorkOn string, args []string) error {
  // TODO:
  fmt.Println("AddProject called with dimensionToWorkOn: '", dimensionToWorkOn, "', args: ", args)
  return nil
}