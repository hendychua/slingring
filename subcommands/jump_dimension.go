package subcommands

import (
  "fmt"
  "errors"
  "strings"

  "github.com/hendychua/slingring/config"
  "github.com/hendychua/slingring/gitutils"
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

  globalConfig, err := config.GetGlobalSettings()
  if err != nil {
    return err
  }

  err = jumpDimensionForProjects(globalConfig, data.Dimensions[name])
  if err != nil {
    return err
  }

  err = setCurrentDimension(data, name)

  if err == nil {
    fmt.Printf("Jumped to '%s'\n", name)
  }

  return err
}

func jumpDimensionForProjects(c *config.Config, dimension config.Dimension) error {
  dirtyProjects, err := getDirtyProjectsInDimension(dimension)
  if err != nil {
    return err
  }

  if len(dirtyProjects) > 0 {
    fmt.Printf("The following projects in '%s' have uncommitted changes:\n", dimension.Name)
    for project := range dirtyProjects {
      fmt.Printf("    %s\n", project.ProjectPath)
    }

    if c.HandleDirtyRepo == config.AbortAll {
      return fmt.Errorf("aborting dimension jump because option HandleDirtyRepo is set to AbortAll")
    } else if c.HandleDirtyRepo == config.AbortContinue {
      fmt.Println("Dimension jump will be skipped for dirty repos because option HandleDirtyRepo is set to AbortContinue.")
    } else if c.HandleDirtyRepo == config.Stash {
      fmt.Println("Changes in dirty repos will be stashed before dimension jump because option HandleDirtyRepo is set to Stash.")
    }
  }

  branchName := dimension.Name

  for _, project := range dimension.Projects {
    if _, ok := dirtyProjects[project]; ok {
      // dirty
      if c.HandleDirtyRepo == config.AbortContinue {
        continue
      } else if c.HandleDirtyRepo == config.Stash {
        output, err := gitutils.GitStash(project.ProjectPath)
        if err != nil {
          return err
        }
        fmt.Printf("Stash successful for '%s'. Results: %s\n", project.ProjectPath, output)
      }
    }

    // checkout to the base branch to create new branches from.
    _, err := gitutils.GitCheckout(project.ProjectPath, c.BaseBranch)
    if err != nil {
      return err
    }

    if c.PullFirst {
      err = gitutils.GitPull(project.ProjectPath)
      if err != nil {
        return err
      }
    }

    err = gitutils.GitCheckoutCreate(project.ProjectPath, branchName)
    if err != nil {
      return err
    }

    fmt.Printf("Checked out to branch '%s' for project '%s'\n", branchName, project.ProjectPath)
  }

  return nil
}

func getDirtyProjectsInDimension(dimension config.Dimension) (map[config.Project]bool, error) {
  dirtyProjects := make(map[config.Project]bool, 0)
  for _, project := range dimension.Projects {
    if clean, err := gitutils.IsClean(project.ProjectPath); err != nil {
      return nil, err
    } else if !clean {
      dirtyProjects[project] = true
    }
  }

  return dirtyProjects, nil
}

func setCurrentDimension(d *config.Data, dimension string) error {
  d.CurrentDimension = dimension
  return d.DataToGlobalDataJSONFile()
}