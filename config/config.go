package config

// HandleDirtyRepoOption is an enum of options when switching branches in a dirty repo.
type HandleDirtyRepoOption uint

const (
  // ABORT abort switching branches and errors out.
  ABORT HandleDirtyRepoOption = iota
  // STASH stash changes before switching branches.
  STASH
)

var handleDirtyRepoOptions = []string{
  "ABORT", "STASH",
}

func (option HandleDirtyRepoOption) Ordinal() int {
  return int(option)
}

func (option HandleDirtyRepoOption) String() string {
  return handleDirtyRepoOptions[option]
}

func (option HandleDirtyRepoOption) Values() *[]string {
 return &handleDirtyRepoOptions
}

// Config is a struct that represents config for handling switching branches.
type Config struct {
  // The base branch to always checkout a new branch from.
  BaseBranch string `json:"baseBranch"`

  // Whether to pull latest commits from remote before checking out to branch.
  PullFirst bool `json:"pullFirst"`

  // Choice of action when the repo has uncommitted changes.
  HandleDirtyRepo HandleDirtyRepoOption `json:"handleDirtyRepo"`
}