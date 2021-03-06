package config

import (
  "strings"
)

// HandleDirtyRepoOption is an enum of options when switching branches in a dirty repo.
type HandleDirtyRepoOption uint

const (
  // AbortAll abort switching branches and errors out.
  AbortAll HandleDirtyRepoOption = iota

  // AbortContinue abort switching branches in one project and
  // continues to switch branches in other projects in the Dimension.
  AbortContinue

  // Stash stash changes before switching branches.
  Stash
)

var handleDirtyRepoOptions = []string{
  "AbortAll","AbortContinue", "Stash",
}

// Ordinal returns the int value for the enum.
func (option HandleDirtyRepoOption) Ordinal() int {
  return int(option)
}

func (option HandleDirtyRepoOption) String() string {
  return handleDirtyRepoOptions[option]
}

// Values returns the list of possible values in the enum.
func (option HandleDirtyRepoOption) Values() *[]string {
 return &handleDirtyRepoOptions
}

// HandleDirtyRepoOptionFromString returns the int value
// for the given string. Returns -1 if the string is not a valid
// HandleDirtyRepoOption value.
func HandleDirtyRepoOptionFromString(s string) int {
  for i, val := range handleDirtyRepoOptions {
    if strings.EqualFold(val, s) {
      return i
    }
  }

  return -1
}