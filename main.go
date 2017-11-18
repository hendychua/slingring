package main

import(
  "os"
  "fmt"
  "log"
  "strings"
  "path"
  "path/filepath"
  "io/ioutil"

  "github.com/urfave/cli"
  "github.com/mitchellh/go-homedir"

  "github.com/hendychua/slingring/config"
  "github.com/hendychua/slingring/subcommands"
)

const appName = "slingring"
const appVersion = "0.0.1"
const appSettingsDir = ".slingring"
const appGlobalSettingsFile = "globalSettings.json"
const appGlobalDataFile = "globalData.json"

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
        return subcommand.Run(c.Args())
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
        return subcommand.Run(c.Args())
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
        return subcommand.Run(c.Args())
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
        return subcommand.Run(c.Args())
      },
    },
    {
      Name: "current-dimension",
      Usage: "show or set current dimension",
      UsageText: fmt.Sprintf("%s current-dimension [<name>]", appName),
      Description: "If <name> is not provided, show the current dimension. If <name> is provided, set the current dimension to <name>.\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.CurrentDimension{}
        return subcommand.Run(c.Args())
      },
    },
    {
      Name: "add-project",
      Usage: "add a project to dimension",
      UsageText: fmt.Sprintf("%s add-project <path>", appName),
      Description: "Add a project path to dimension.\n   path - directory to add to dimension.\n",
      Category: "Project",
      Action: func(c *cli.Context) error {
        subcommand := subcommands.AddProject{}
        // TODO: if global option "dimension" is not provided, should get the current dimension context. If it is also not set, error out.
        return subcommand.Run(c.GlobalString("dimension"), c.Args())
      },
    },
  }

  app.Run(os.Args)
}

func init() {
  // Set up the global settings and data file if they do not exist.
  // For example, running the app for the first time.

  userHomeDir, err := homedir.Dir()
  check(err)

  globalSettingsDir := path.Join(userHomeDir, appSettingsDir)

  globalSettingsFilePath := path.Join(globalSettingsDir, appGlobalSettingsFile)
  maybeCreateGlobalFile(globalSettingsFilePath, "settings")

  globalDataFilePath := path.Join(globalSettingsDir, appGlobalDataFile)
  maybeCreateGlobalFile(globalDataFilePath, "data")
}

func maybeCreateGlobalFile(f string, fileType string) {
  if _, fileInfoErr := os.Stat(f); os.IsNotExist(fileInfoErr) {
    log.Printf("%s does not exist. Creating...\n", f)
    mkdirErr := os.MkdirAll(filepath.Dir(f), 0755)
    check(mkdirErr)

    var data []byte
    var err error

    // reason for not passing data as parameter is so that we only create the structs
    // only when we have checked that the respective file does not exist and we have to create them.
    if strings.EqualFold(fileType, "settings") {
      data, err = getDefaultGlobalSettings()
      check(err)
    } else if strings.EqualFold(fileType, "data") {
      data, err = getDefaultGlobalData()
      check(err)
    } else {
      log.Fatalf("Unexpected global fileType '%s'\n", fileType)
    }

    writeFileErr := ioutil.WriteFile(f, data, 0644)
    check(writeFileErr)

  } else {
    check(fileInfoErr)
  }
}

func getDefaultGlobalSettings() ([]byte, error) {
  defaultConfig := config.Config{BaseBranch: "develop", PullFirst: true, HandleDirtyRepo: config.ABORT}
  return defaultConfig.ToJSON()
}

func getDefaultGlobalData() ([]byte, error) {
  defaultData := config.Data{Dimensions:make([]config.Dimension, 0), CurrentDimension: -1}
  return defaultData.ToJSON()
}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}