package subcommands

import (
  "fmt"
  "errors"
  "strings"

  "github.com/hendychua/slingring/config"
)

// JumpDimension is a command that jumps to another dimension and setting it
// as the current Dimension.
// If no argument is provided, this command unsets the current dimension.
type JumpDimension struct{}

// Run sets the current Dimension to Dimension with <name>.
// If no argument is provided, this command unsets the current dimension.
func (j JumpDimension) Run(args []string) error {
  data, err := config.GetGlobalData()
  if err != nil {
    return err
  }

  if len(args) <= 0 {
    return setCurrentDimension(data, "")
  }

  name := strings.TrimSpace(strings.ToLower(args[0]))

  if len(name) <= 0 {
    return errors.New("illegal name for Dimension: Cannot be empty")
  }

  if data.HasDimensionNamed(name) == false {
    return fmt.Errorf("fatal error: no Dimension named '%s'", name)
  }

  // TODO: for every Project in the dimension:
  // Check if repo is dirty. follow HandleDirtyRepoOption protocol.
  // If project has branch named <name>, checkout to it.
  // If project does not have branch named <name>, checkout -b to it.
  // return errors when it happens.

  err = setCurrentDimension(data, name)

  if err == nil {
    fmt.Printf("Jumpted to '%s'\n", name)
  }

  return err
}

func setCurrentDimension(d *config.Data, dimension string) error {
  d.CurrentDimension = dimension
  return d.DataToGlobalDataJSONFile()
}