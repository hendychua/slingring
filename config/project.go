package config

import (
  "fmt"
)

// Project is a struct that represents a project.
type Project struct {
  ProjectPath string
}

func (p Project) String() string {
  return fmt.Sprintf("ProjectPath: '%s'", p.ProjectPath)
}