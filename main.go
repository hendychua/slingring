package main

import(
  "os"
  "fmt"

  "github.com/urfave/cli"
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
        fmt.Println("create: ", c.Args())
        return nil
      },
    },
    {
      Name: "list",
      Usage: "list all dimensions",
      UsageText: fmt.Sprintf("%s list", appName),
      Description: "List all the dimensions created. * denotes current dimension.\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        fmt.Println("list: ", c.Args())
        return nil
      },
    },
    {
      Name: "delete",
      Usage: "delete a dimension",
      UsageText: fmt.Sprintf("%s delete <name>", appName),
      Description: "name - name of dimension to delete\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        fmt.Println("delete: ", c.Args())
        return nil
      },
    },
    {
      Name: "describe",
      Usage: "describe a dimension",
      UsageText: fmt.Sprintf("%s describe <name>", appName),
      Description: "name - name of dimension to describe\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        fmt.Println("describe: ", c.Args())
        return nil
      },
    },
    {
      Name: "current-dimension",
      Usage: "show or set current dimension",
      UsageText: fmt.Sprintf("%s current-dimension [<name>]", appName),
      Description: "If <name> is not provided, show the current dimension. If <name> is provided, set the current dimension to <name>.\n",
      Category: "Dimension",
      Action: func(c *cli.Context) error {
        fmt.Println("current-dimension: ", c.Args())
        return nil
      },
    },
    {
      Name: "add-project",
      Usage: "add a project to dimension",
      UsageText: fmt.Sprintf("%s add-project <path>", appName),
      Description: "Add a project path to dimension.\n   path - directory to add to dimension.\n",
      Category: "Project",
      Action: func(c *cli.Context) error {
        fmt.Println("add-project: ", c.Args())
        return nil
      },
    },
  }

  app.Run(os.Args)
}