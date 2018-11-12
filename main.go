package main

import (
  "log"

  "sjf/client/operations"
  "github.com/go-openapi/strfmt"
  "github.com/go-openapi/spec"

  apiclient "sjf/client"
  httptransport "github.com/go-openapi/runtime/client"
)

func main() {

  // make the request to get all items
  resp, err := apiclient.Default.Operations.All(operations.GetAccounts{})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)
}
