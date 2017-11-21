package subcommands

import (
  "fmt"
  "errors"
  "os"
  "path/filepath"

  "github.com/hendychua/slingring/config"
)

// AddProject is a command that adds project path(s) to a Dimension.
type AddProject struct{}

// Run adds project path(s) to a Dimension.
// If dimensionToWorkOn happens to be the currently set dimension, this function
// will also attempt to jump to the dimension for the newly added project(s).
func (a AddProject) Run(dimensionToWorkOn string, args []string) error {
  if len(args) <= 0 {
    return errors.New("illegal usage: not enough arguments for AddProject command")
  }

  data, err := config.GetGlobalData()
  if err != nil {
    return err
  }

  dimension, ok := data.Dimensions[dimensionToWorkOn]
  if !ok {
    return fmt.Errorf("fatal error: no Dimension named '%s'", dimensionToWorkOn)
  }

  for _, projectPath := range args {
    absProjectPath, absErr := filepath.Abs(projectPath)
    if absErr != nil {
      return absErr
    }

    if _, fileInfoErr := os.Stat(absProjectPath); os.IsNotExist(fileInfoErr) {
      return fmt.Errorf("error while adding '%s' to '%s'. '%s' does not exist", projectPath, dimensionToWorkOn, projectPath)
    } else if err != nil {
      return err
    }

    // absProjectPath exists
    newProject := config.Project{ProjectPath: absProjectPath}
    dimension.Projects = append(dimension.Projects, newProject)
  }

  data.Dimensions[dimensionToWorkOn] = dimension
  err = data.DataToGlobalDataJSONFile()
  if err != nil {
    return err
  }

  // TODO: jump to dimension for all the added projects.

  return nil
}