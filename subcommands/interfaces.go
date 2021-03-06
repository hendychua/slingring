package subcommands

// SubCommand is a command that does not need a Dimension context to work.
type SubCommand interface {
  Run(args []string) error
}

// DimensionContextCommand is a command that requires a Dimension context to work.
// For example, adding a project to a Dimension, or setting Dimension-specific config.
type DimensionContextCommand interface {
  Run(currentDimension string, args []string) error
}