package main

import(
  "os"
  "fmt"
  "log"
  "strings"
  "path/filepath"

  "github.com/urfave/cli"

  "github.com/hendychua/slingring/utils"
  "github.com/hendychua/slingring/config"
  "github.com/hendychua/slingring/subcommands"
)

const appName = "slingring"
const appVersion = "0.0.1"

func main() {
  app := cli.NewApp()
  app.Name = appName
  app.Usage = "Manage feature branches across projects"
  app.Version = appVersion

  // Initialize global options here

  app.Flags = []cli.Flag{
    cli.StringFlag{
      Name: "dimension, d",
      Value: "",
      Usage: "`name` of the dimension to use",
    },
  }

  // Initialize subcommands here

  app.Commands = []cli.Command{
    {
      Name: "create",
      Usage: "create a new dimension",
      UsageText: fmt.Sprintf("%s create <name>", appName),
      Description: "name - name of dimension to create\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.CreateDimension{}
        return exitOneOnError(subcommand.Run(c.Args()))
      },
    },
    {
      Name: "list",
      Usage: "list all dimensions",
      UsageText: fmt.Sprintf("%s list", appName),
      Description: "List all the dimensions created. * denotes current dimension.\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.ListDimensions{}
        return exitOneOnError(subcommand.Run(c.Args()))
      },
    },
    {
      Name: "delete",
      Usage: "delete a dimension",
      UsageText: fmt.Sprintf("%s delete <name>", appName),
      Description: "name - name of dimension to delete\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.DeleteDimension{}
        return exitOneOnError(subcommand.Run(c.Args()))
      },
    },
    {
      Name: "describe",
      Usage: "describe a dimension",
      UsageText: fmt.Sprintf("%s describe <name>", appName),
      Description: "name - name of dimension to describe\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.DescribeDimension{}
        return exitOneOnError(subcommand.Run(c.Args()))
      },
    },
    {
      Name: "jump",
      Usage: "set current dimension",
      UsageText: fmt.Sprintf("%s jump [<name>]", appName),
      Description: "name - name of dimension to jump to. If name is not provided, this command unsets the current dimension.\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.JumpDimension{}
        return exitOneOnError(subcommand.Run(c.Args()))
      },
    },
    {
      Name: "add-project",
      Usage: "add project(s) to dimension",
      UsageText: fmt.Sprintf("%s add-project <path>...", appName),
      Description: "Add project path(s) to dimension.\n   path - directory to add to dimension.\n",
      Category: "Project",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.AddProject{}
        var dimensionToAddTo string
        if len(c.GlobalString("dimension")) > 0 {
          dimensionToAddTo = c.GlobalString("dimension")
        } else {
          data, err := config.GetGlobalData()
          if err != nil {
            return exitOneOnError(err)
          }

          if len(data.CurrentDimension) <= 0 {
            return exitOneOnError(fmt.Errorf("error: %s%s",
              "no current dimension is set and no dimension is specified with global option. ",
              "Use global option --dimension (-d) to specify a dimension to add project path to."))
          }

          dimensionToAddTo = data.CurrentDimension
        }
        return exitOneOnError(subcommand.Run(dimensionToAddTo, c.Args()))
      },
    },
    // TODO: remove-project command
  }

  app.Run(os.Args)
}

func init() {
  // Set up the global settings and data file if they do not exist.
  // For example, running the app for the first time.

  globalSettingsFilePath := config.GetGlobalSettingsFile()
  maybeCreateGlobalFile(globalSettingsFilePath, "settings")

  globalDataFilePath := config.GetGlobalDataFile()
  maybeCreateGlobalFile(globalDataFilePath, "data")
}

func maybeCreateGlobalFile(f string,fileType string) {
  if _, fileInfoErr := os.Stat(f); os.IsNotExist(fileInfoErr) {
    log.Printf("%s does not exist. Creating...\n", f)
    mkdirErr := os.MkdirAll(filepath.Dir(f), 0755)
    utils.Check(mkdirErr)

    var err error

    // reason for not passing data as parameter is so that we only create the structs
    // only when we have checked that the respective file does not exist and we have to create them.
    if strings.EqualFold(fileType, "settings") {
      err = writeDefaultGlobalSettings()
      utils.Check(err)
    } else if strings.EqualFold(fileType, "data") {
      err = writeDefaultGlobalData()
      utils.Check(err)
    } else {
      log.Fatalf("Unexpected global fileType '%s'\n", fileType)
    }

  } else {
    utils.Check(fileInfoErr)
  }
}

func writeDefaultGlobalSettings() error {
  defaultConfig := config.Config{BaseBranch: "develop", PullFirst: true, HandleDirtyRepo: config.AbortAll}
  return defaultConfig.ConfigToGlobalSettingsJSONFile()
}

func writeDefaultGlobalData() error {
  defaultData := config.Data{Dimensions:make(map[string]config.Dimension, 0), CurrentDimension: ""}
  return defaultData.DataToGlobalDataJSONFile()
}

func exitOneOnError(err error) error {
  if err != nil {
    return cli.NewExitError(err, 1)
  }
  return err
}