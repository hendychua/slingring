package subcommands

import (
  "fmt"
  "errors"
  "path/filepath"

  "github.com/hendychua/slingring/config"
)

// RemoveProject is a command that removes project path(s) from a Dimension.
type RemoveProject struct{}

// Run removes project path(s) from a Dimension.
// If the project path does not belong to a Dimension, this method simply ignores it
// instead of returning an error.
// This command does not modify the project path in any way. It simply removes it from the Dimension.
func (r RemoveProject) Run(dimensionToWorkOn string, args []string) error {
  if len(args) <= 0 {
    return errors.New("illegal usage: not enough arguments for RemoveProject command")
  }

  data, err := config.GetGlobalData()
  if err != nil {
    return err
  }

  dimension, ok := data.Dimensions[dimensionToWorkOn]
  if !ok {
    return fmt.Errorf("fatal error: no Dimension named '%s'", dimensionToWorkOn)
  }

  absPathsToRemove := make(map[string]bool, 0)
  for _, projectPath := range args {
    absProjectPath, absErr := filepath.Abs(projectPath)
    if absErr != nil {
      return absErr
    }

    absPathsToRemove[absProjectPath] = true
  }

  var newProjectsArray []config.Project
  for _, project := range dimension.Projects {
    if _, ok := absPathsToRemove[project.ProjectPath]; !ok {
      // not in absPathToRemove, so we add it to newProjectsArray
      newProjectsArray = append(newProjectsArray, project)
    }
  }

  dimension.Projects = newProjectsArray

  data.Dimensions[dimensionToWorkOn] = dimension
  err = data.DataToGlobalDataJSONFile()

  if err == nil {
    fmt.Printf("Projects remaining in dimension '%s': %s.\n", dimension.Name, dimension.Projects)
  }

  return err
}