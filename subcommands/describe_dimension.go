package subcommands

import (
  "fmt"
  "errors"
  "strings"

  "github.com/hendychua/slingring/config"
)

// DescribeDimension is a command that describes a Dimension.
type DescribeDimension struct{}

// Run describes a Dimension with <name> in detail.
// If such a Dimension does not exist, the command returns an error.
func (d DescribeDimension) Run(args []string) error {
  if len(args) != 1 {
    return errors.New("illegal usage: DescribeDimension takes in 1 argument")
  }

  name := strings.TrimSpace(strings.ToLower(args[0]))

  if len(name) <= 0 {
    return errors.New("illegal name for Dimension: Cannot be empty")
  }

  data, err := config.GetGlobalData()
  if err != nil {
    return err
  }

  dimension, ok := data.Dimensions[name]
  if !ok {
    return fmt.Errorf("fatal error: no Dimension named '%s'", name)
  }

  // TODO: formatting the output to align the left columns can be done in a better way.
  fmt.Printf("Name:               %s\n", dimension.Name)
  if strings.EqualFold(data.CurrentDimension, dimension.Name) {
    fmt.Printf("Current default:    yes\n\n")
  } else {
    fmt.Printf("Current default:    no\n\n")
  }
  if len(dimension.Projects) > 0 {
    fmt.Printf("Projects:\n")
    for _, project := range dimension.Projects {
      fmt.Printf("%s\n", project)
    }
  } else {
    fmt.Printf("(No projects added to dimension.)\n")
  }
  fmt.Println()

  return nil
}