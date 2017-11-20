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