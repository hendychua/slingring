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

  err = JumpDimensionForProjects(globalConfig, data.Dimensions[name])
  if err != nil {
    return err
  }

  err = setCurrentDimension(data, name)

  if err == nil {
    fmt.Printf("Jumped to '%s'\n", name)
  }

  return err
}

// JumpDimensionForProjects switches the branch to a branch with the dimension's name,
// aka jumping dimension.
// Before jumping dimension for each project, this method will check the config on how to handle dirty repos.
// It also checks the config for whether to pull first before jumping and the base branch
// to create new branches (if necessary) from.
// By default, if nothing is passed in for projects list, thie method operates on each project in
// dimension. If projects is a non-empty list, this method operates on each project in the list.
func JumpDimensionForProjects(c *config.Config, dimension config.Dimension, projects ...config.Project) error {
  branchName := dimension.Name

  var projectsList []config.Project
  if len(projects) == 0 {
    projectsList = dimension.Projects
  } else {
    projectsList = projects
  }

  dirtyProjects, err := getDirtyProjectsIfNotOnBranch(branchName, projectsList...)
  if err != nil {
    return err
  }

  if len(dirtyProjects) > 0 {
    fmt.Printf("The following projects in '%s' have uncommitted changes:\n", dimension.Name)
    for project := range dirtyProjects {
      fmt.Printf("    %s\n", project.ProjectPath)
    }

    if c.HandleDirtyRepo == config.AbortAll {
      return fmt.Errorf("error: aborting dimension jump because option HandleDirtyRepo is set to AbortAll")
    } else if c.HandleDirtyRepo == config.AbortContinue {
      fmt.Println("Dimension jump will be skipped for dirty repos because option HandleDirtyRepo is set to AbortContinue.")
    } else if c.HandleDirtyRepo == config.Stash {
      fmt.Println("Changes in dirty repos will be stashed before dimension jump because option HandleDirtyRepo is set to Stash.")
    }
  }

  for _, project := range projectsList {
    if isCurrent, err := gitutils.IsCurrentBranch(project.ProjectPath, branchName); err != nil {
      return err
    } else if isCurrent {
      // already on the dimension. don't have to checkout.
      fmt.Printf("'%s' already on '%s'. Nothing to do.\n", project.ProjectPath, branchName)
      continue
    }

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

func getDirtyProjectsIfNotOnBranch(branchName string, projects ...config.Project) (map[config.Project]bool, error) {
  dirtyProjects := make(map[config.Project]bool, 0)
  for _, project := range projects {
    if isCurrent, err := gitutils.IsCurrentBranch(project.ProjectPath, branchName); err != nil {
      return nil, err
    } else if isCurrent {
      // if the project is already on the branch, we don't really need to care if it is dirty
      // because we won't attempt to switch branch in it.
      continue
    }

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