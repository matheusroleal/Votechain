package main

import (
  "fmt"
  "log"
  "os"
  "encoding/json"
  "net/url"
  "net/http"

  "github.com/urfave/cli"
  types "github.com/matheusroleal/Votechain/types"
)

func SendVote(c *cli.Context) error{
  if c.NumFlags() != 2 {
      return fmt.Errorf("To and amount must be specified")
  }
  var from string
  if c.String("from") == "" {
    return fmt.Errorf("From adress must be specified")
  } else {
      from = c.String("from")
  }
  to := c.String("to")

  var res types.SendTxResponse
  err := Call("sendvote", map[string]string{
      "to": to,
      "from": from,
  }, &res)
  if err != nil {
    return err
  }

  out, err := json.MarshalIndent(res, "","  ")
  if err != nil {
    return err
  }
  log.Println(string(out))
  return nil
}

func GetNewAddress(c *cli.Context) error{
  var res types.GetNewAddressResponse

  err := Call("getnewaddress", map[string]string{}, &res)
  if err != nil {
    return err
  }

  out, err := json.MarshalIndent(res, "","  ")
  if err != nil {
    return err
  }
  log.Println(string(out))
  return nil
}

func Call(cmd string, options map[string]string, out interface{}) error{
  vals := make(url.Values)
  for k, v := range options {
      vals.Set(k, v)
  }
  url := os.Getenv("VOTECHAIN_HOST")+":"+os.Getenv("VOTECHAIN_PORT")+"/"
  resp, err := http.PostForm(url+ cmd, vals)
  if err != nil {
      return err
  }
  err = json.NewDecoder(resp.Body).Decode(out)
  if err != nil {
    return err
  }
  return nil
}

func main() {
  app := cli.NewApp()
  app.Name = "votechain-cli"
  app.Usage = "rpc client for votechain"

  app.Commands = []cli.Command{
      {
        Name:    "sendvote",
        Usage:   "send a transaction",
        Flags: []cli.Flag {
            cli.StringFlag{
                Name: "from",
                Value: "",
                Usage: "specify from address",
            },
            cli.StringFlag{
                Name: "to",
                Value: "",
                Usage: "specify to address",
            },
        },
        Action:  SendVote,
      },
      {
        Name:   "getnewaddress",
        Usage:  "get new address",
        Action: GetNewAddress,
      },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
