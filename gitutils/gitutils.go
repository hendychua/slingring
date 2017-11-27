package gitutils

import (
  "os/exec"
  "strings"
  "fmt"
)

const gitCommand = "git"

// IsClean method checks whether a directory is a clean repo.
// A dirty repo is one with one of the following:
// - modified but unstaged changes
// - staged but uncommited changes
// Untracked files are not considered dirty.
func IsClean(dir string) (bool, error) {
  cmd := exec.Command(gitCommand, "status", "--porcelain")
  cmd.Dir = dir
  statusOut, err := cmd.Output()
  if err != nil {
    return false, err
  }

  statusOutString := strings.TrimSpace(string(statusOut))
  if len(statusOutString) > 0 {
    const untrackedFilesPrefix = "?? "
    for _, line := range strings.Split(statusOutString, "\n") {
      if strings.HasPrefix(strings.TrimSpace(line), untrackedFilesPrefix) == false {
        return false, nil
      }
    }
  }

  return true, nil
}

// GitPull runs "git pull" in the directory.
func GitPull(dir string) error {
  cmd := exec.Command(gitCommand, "pull")
  cmd.Dir = dir
  _, err := cmd.Output()
  return err
}

// GitCheckout switches to brancName in the directory.
// If the branch does not exist, it returns an error.
func GitCheckout(dir string, branchName string) (string, error) {
  cmd := exec.Command(gitCommand, "checkout", branchName)
  cmd.Dir = dir
  output, err := cmd.CombinedOutput()
  return string(output), err
}

// GitCheckoutCreate switches to brancName in the directory.
// If the branch does not exist, it will be created.
func GitCheckoutCreate(dir string, branchName string) error {
  checkoutOutput, err := GitCheckout(dir, branchName)
  if err == nil {
    return nil
  }

  // An error occurred with the previous command.
  // It is possible the branch does not exist. Create it.

  cmd := exec.Command(gitCommand, "checkout", "-b", branchName)
  cmd.Dir = dir
  checkoutNewOutput, err := cmd.CombinedOutput()
  if err == nil {
    return nil
  }

  return fmt.Errorf("Unable to check out '%s'. " +
    "Error executing 'checkout': '%s'. " +
    "Error executing 'checkout -b': '%s'", branchName,
    checkoutOutput, string(checkoutNewOutput))
}

// GitStash runs git stash in the directory.
// The first return value is the output of running the command,
// i.e. where the stash occurs.
func GitStash(dir string) (string, error) {
  cmd := exec.Command(gitCommand, "stash")
  cmd.Dir = dir
  output, err:= cmd.CombinedOutput()
  return string(output), err
}