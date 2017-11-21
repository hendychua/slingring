package subcommands

import(
  "errors"
  "strings"
  "fmt"

  "github.com/hendychua/slingring/config"
)

// DeleteDimension is a command that deletes a Dimension.
type DeleteDimension struct{}

// Run deletes a Dimension by name.
// If such a Dimension does not exist, the command returns an error.
func (d DeleteDimension) Run(args []string) error {
  if len(args) != 1 {
    return errors.New("illegal usage: DeleteDimension takes in 1 argument")
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

  delete(data.Dimensions, name)

  return data.DataToGlobalDataJSONFile()
}