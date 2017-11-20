package config

import (
  "encoding/json"
  "io/ioutil"
)

// Config is a struct that represents config for handling switching branches.
type Config struct {
  // The base branch to always checkout a new branch from.
  BaseBranch string `json:"baseBranch"`

  // Whether to pull latest commits from remote before checking out to branch.
  PullFirst bool `json:"pullFirst"`

  // Choice of action when the repo has uncommitted changes.
  HandleDirtyRepo HandleDirtyRepoOption `json:"handleDirtyRepo"`
}

// ConfigToGlobalSettingsJSONFile writes c to the global settings file.
func (c Config) ConfigToGlobalSettingsJSONFile() error {
  configJSON, err := c.toJSON()
  if err != nil {
    return err
  }

  err = ioutil.WriteFile(GetGlobalSettingsFile(), configJSON, 0644)
  return err
}

func (c Config) toJSON() ([]byte, error) {
  return json.MarshalIndent(c, "", "  ") // 2-spaces indentation
}