//Create a restful API in Go
package main

import "fmt"
import "github.com/gorilla/mux"
import "net/http"

func main(){
  router:=mux.NewRouter()
  log.Fatal(http.ListenAndServe(":12345", router))

}
