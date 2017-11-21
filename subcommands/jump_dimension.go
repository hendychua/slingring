package subcommands

import (
  "fmt"
  "errors"
  "strings"

  "github.com/hendychua/slingring/config"
)

// JumpDimension is a command that jumps to another dimension and setting it
// as the current Dimension.
type JumpDimension struct{}

// Run sets the current Dimension to Dimension with <name>.
func (j JumpDimension) Run(args []string) error {
  if len(args) != 1 {
    return errors.New("illegal usage: JumpDimension takes in 1 argument")
  }

  name := strings.TrimSpace(strings.ToLower(args[0]))

  if len(name) <= 0 {
    return errors.New("illegal name for Dimension: Cannot be empty")
  }

  data, err := config.GetGlobalData()
  if err != nil {
    return err
  }

  if data.HasDimensionNamed(name) == false {
    return fmt.Errorf("fatal error: no Dimension named '%s'", name)
  }

  // TODO: for every Project in the dimension:
  // Check if repo is dirty. follow HandleDirtyRepoOption protocol.
  // If project has branch named <name>, checkout to it.
  // If project does not have branch named <name>, checkout -b to it.
  // return errors when it happens.

  data.CurrentDimension = name
  return data.DataToGlobalDataJSONFile()
}