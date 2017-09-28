//JSON file handling in GO
package main

import "fmt"
import "os"
import "encoding/json"

type Config struct {
	Database struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"database"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func LoadConfig(filename string) (Config, error) {
  var config Config
  configFile, err := os.Open(filename)
  defer configFile.Close()
  if(err!=nil){
    return config, err
  }
  jsonParser:=json.NewDecoder(configFile)
  err=jsonParser.Decode(&config)
  return config, err
}
func main(){
  config, _ := LoadConfig("config.json")
  fmt.Println(config)

}
