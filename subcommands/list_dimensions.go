package subcommands

import (
  "fmt"
  "strings"
  "sort"

  "github.com/hendychua/slingring/config"
)

// ListDimensions is a command that lists all the dimensions created.
type ListDimensions struct{}


// Run lists all the dimensions created in alphabetical order.
// The current dimension has an annotation of * followed by a space before the dimension name.
func (l ListDimensions) Run(args []string) error {
  data, err := config.GetGlobalData()
  if err != nil {
    return err
  }

  var padding = ""
  if len(data.CurrentDimension) > 0 {
    // a default dimension is set
    padding = "  "
  }

  sortedNames := make([]string, len(data.Dimensions))
  i := 0
  for dimensionName := range data.Dimensions {
    sortedNames[i] = dimensionName
    i++
  }
  sort.Strings(sortedNames)

  var defaultDimensionSet = false
  for _, dimensionName := range sortedNames {
    if strings.EqualFold(data.CurrentDimension, dimensionName) {
      fmt.Printf("* %s\n", dimensionName)
      defaultDimensionSet = true
    } else {
      fmt.Printf("%s%s\n", padding, dimensionName)
    }
  }

  if defaultDimensionSet {
    fmt.Println("(* denotes current dimension)")
  } else {
    fmt.Println("(no current dimension set)")
  }

  return nil
}