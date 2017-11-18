package subcommands

import(
  "fmt"
)

type DeleteDimension struct{}

func (d DeleteDimension) Run(args []string) error {
  fmt.Println("DeleteDimension called with args: ", args)
  return nil
}