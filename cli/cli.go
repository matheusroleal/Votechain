package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli"
)

type SendTxResponse struct {
    Txid       string
}

type GetNewAddressResponse struct {
    Address     string
}

func SendVote(c *cli.Context) error{
  if len(c.Args()) != 2 {
      return fmt.Errorf("To and amount must be specified")
  }
  var from string
  if c.String("from") == "" {
    return fmt.Errorf("From adress must be specified")
  } else {
      from = c.String("from")
  }
  memo := c.String("to")

  var res types.SendTxResponse
  err := Call("sendtx", map[string]string{
      "to": to,
      "from": from,
  }, &res)

  out, err := json.MarshalIndent(res, "","  ")
  if err != nil {
    return err
  }
  fmt.Println(string(out))
  return nil
}

func GetNewAddress(c *cli.Context){
  var res types.SendTxResponse
  err := Call("getnewaddress", map[string]string{}, &res)

  out, err := json.MarshalIndent(res, "","  ")
  if err != nil {
    return err
  }
  fmt.Println(string(out))
  return nil
}

func Call(cmd string, options map[string]string, out interface{}){
  vals := make(url.Values)
  for k, v := range options {
      vals.Set(k, v)
  }

  resp, err := http.PostForm("http://127.0.0.1:1234/" + cmd, vals)
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
