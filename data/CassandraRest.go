// GENERATED FILE so do not edit or will be overwritten upon next generate
package data

import (
  "sjf/models"
  "sjf/restapi/operations"
    middleware "github.com/go-openapi/runtime/middleware"
    "github.com/gocql/gocql"
    "os"
    "log"
)

var session *gocql.Session

func SetUp() {
  var err error
  log.Println("Connecting to Cassandra DB on ", os.Getenv("CASSANDRA_SERVICE_HOST"))
  cluster := gocql.NewCluster(os.Getenv("CASSANDRA_SERVICE_HOST"))
  cluster.Keyspace = "test"
  cluster.Consistency = gocql.One
  session, err = cluster.CreateSession()
  if ( err != nil ) {
	  log.Fatal("Have you set the env var CASSANDRA_SERVICE_HOST. Connection to Cannandra failed", err)
  } else {
	  log.Println("Connection to Cannandra established")
  }
}

func Stop() {
    log.Println("Stopping Cassandra")
    session.Close()
}

func Search(params operations.GetAccountsParams) middleware.Responder {

    var id int64
    var name string
    ret := &operations.GetAccountsOK{}
    ret.Payload = make([]*models.GetAccountsOKBodyItems,1)
    if err := session.Query(`SELECT id,name FROM accounts WHERE id = ? and name = ? `,params.ID,params.Name).Consistency(gocql.One).Scan(&id,&name); err != nil {
        // No data
        log.Println("No data? ", err)
        return operations.NewGetAccountsBadRequest()
    }
    payLoad := new(models.GetAccountsOKBodyItems)
    
    ret.Payload[0] = payLoad
    payLoad.ID = &id
    payLoad.Name = &name
    return operations.NewGetAccountsOK().WithPayload(ret.Payload)
}
