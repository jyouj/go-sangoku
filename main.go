package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

// Mapping
var giKingMap = map[string]string{
  "0": "曹操",
  "1": "曹丕/魏の文帝",
  "2": "曹叡/魏の明帝",
  "3": "曹芳",
  "4": "曹髦",
  "5": "曹奐",
}

var shokuKingMap = map[string]string{
  "1": "劉備",
  "2": "劉禅",
}

var goKingMap = map[string]string{
  "0": "孫策",
  "1": "孫権/呉の大帝",
  "2": "孫亮",
  "3": "孫休",
  "4": "孫晧",
}


func main() {
  app := cli.NewApp()

  app.Name = "go-sangoku"
  app.Usage = "Display Emperor Name in Gi or Shoku or Go of China!"
  app.Version = "0.0.1"

  app.Action = func (c *cli.Context) error {
    str := c.Args().Get(0)
    switch {
    case c.Bool("gi"):
      fmt.Println(giKingMap[str])
    case c.Bool("shoku"):
      fmt.Println(shokuKingMap[str])
    case c.Bool("go"):
      fmt.Println(goKingMap[str])
    default:
      fmt.Println("漢王朝が滅亡し、群雄割拠の時代に移りました！")
    }
    return nil
  }

  app.Flags = []cli.Flag{
    cli.BoolFlag {
      Name: "gi",
      Usage: "Display Emperor Name in Gi Dynasty",
    },
    cli.BoolFlag {
      Name: "shoku, s",
      Usage: "Display Emperor Name in Shoku Dynasty",
    },
    cli.BoolFlag {
      Name: "go",
      Usage: "Display Emperor Name in Go Dynasty",
    },
  }

  app.Run(os.Args)
}
