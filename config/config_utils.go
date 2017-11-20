package config

import (
  "path"

  "github.com/mitchellh/go-homedir"

  "github.com/hendychua/slingring/utils"
)

const appSettingsDir = ".slingring"
const appGlobalSettingsFile = "globalSettings.json"
const appGlobalDataFile = "globalData.json"

// GetGlobalSettingsDir returns the global settings directory.
func GetGlobalSettingsDir() string {
  return path.Join(getUserHomeDir(), appSettingsDir)
}

// GetGlobalSettingsFile returns the global settings file.
func GetGlobalSettingsFile() string {
  return path.Join(GetGlobalSettingsDir(), appGlobalSettingsFile)
}

// GetGlobalDataFile returns the global data file.
func GetGlobalDataFile() string {
  return path.Join(GetGlobalSettingsDir(), appGlobalDataFile)
}

func getUserHomeDir() string {
  userHomeDir, err := homedir.Dir()
  utils.Check(err)
  return userHomeDir
}