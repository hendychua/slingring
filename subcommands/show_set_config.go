package subcommands

import (
  "strings"
  "errors"
  "fmt"
  "math"
  "strconv"
   "reflect"

  "github.com/hendychua/slingring/config"
)

// ShowSetConfig is a command that shows or sets config.
type ShowSetConfig struct{}

// Run shows or sets config. If args is empty, it prints the config.
// If args is non-empty, it is expected to be of the format
// <prop> <value> <prop> <value>..., and the config will be set
// accordingly.
func (showSetConfig ShowSetConfig) Run(args []string) error {
  configStored, err := config.GetGlobalSettings()
  if err != nil {
    return err
  }

  if len(args) > 0 {
    if math.Mod(float64(len(args)), 2.0) != 0 {
      return errors.New("error: args is expected to be of the following format: <key> <value> <key> <value>... ")
    }

    var key string
    for i, arg := range args {
      if math.Mod(float64(i), 2.0) == 0 {
        key = arg
        continue
      } else {
        // value
        switch strings.ToLower(key) {
        case "basebranch":
          configStored.BaseBranch = arg

        case "pullfirst":
          value, err := strconv.ParseBool(arg)
          if err != nil {
            return err
          }
          configStored.PullFirst = value

        case "handledirtyrepo":
          value := config.HandleDirtyRepoOptionFromString(arg)
          if value == -1 {
            return fmt.Errorf("error: invalid HandleDirtyRepoOption '%s'", arg)
          }
          configStored.HandleDirtyRepo = config.HandleDirtyRepoOption(value)

        default:
          return fmt.Errorf("error: unrecognized config property '%s'", key)
        }
      }
    }

    writeToFileErr := configStored.ConfigToGlobalSettingsJSONFile()
    if writeToFileErr != nil {
      return writeToFileErr
    }

    fmt.Println("New config written.")
    fmt.Println()
  }

  s := reflect.ValueOf(configStored).Elem()
  typeOfT := s.Type()

  for i := 0; i < s.NumField(); i++ {
    f := s.Field(i)
    fmt.Printf("%s: %v\n", typeOfT.Field(i).Name, f.Interface())
  }

  fmt.Println()

  return nil
}