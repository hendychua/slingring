package subcommands

import (
  "strings"
  "errors"
  "io/ioutil"
  "log"
  "fmt"

  "github.com/hendychua/slingring/utils"
  "github.com/hendychua/slingring/config"
)

// CreateDimension is a command that creates a new Dimension.
type CreateDimension struct{}

// Run creates a new Dimension if a Dimension with the name does not exist.
// If such a Dimension exist, return an error.
func (c CreateDimension) Run(args []string) error {
  if len(args) != 1 {
    return errors.New("illegal usage: CreateDimension takes in 1 argument")
  }

  name := strings.TrimSpace(strings.ToLower(args[0]))

  if len(name) <= 0 {
    return errors.New("illegal name for Dimension: Cannot be empty")
  }

  dataJSONContents, err := ioutil.ReadFile(config.GetGlobalDataFile())
  utils.Check(err)
  data := config.Data{}
  config.DataFromJSON(dataJSONContents, &data)
  log.Println("data read from JSON file: ", data)

  if data.HasDimensionNamed(name) {
    return fmt.Errorf("fatal error: already has Dimension named '%s'", name)
  }

  newDimension := config.Dimension{Name: name, Projects: make([]config.Project, 0)}
  data.Dimensions[name] = newDimension

  data.DataToGlobalDataJSONFile()

  return nil
}