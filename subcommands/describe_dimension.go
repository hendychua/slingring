package subcommands

import (
  "fmt"
)

type DescribeDimension struct{}

func (d DescribeDimension) Run(args []string) error {
  fmt.Println("DescribeDimension called with args: ", args)
  return nil
}