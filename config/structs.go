package config

// Project is a struct that represents a project.
type Project struct {
  ProjectPath string
}

// Dimension is a struct that represents a group of projects.
type Dimension struct {
  Name string
  Projects []Project
}