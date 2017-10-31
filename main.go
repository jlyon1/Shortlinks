package main
import (
  "fmt"
  "docs/database"
)

func connectDB(db database.DB){
  fmt.Printf("%v\n",db.Connect())
  fmt.Printf("%v\n",db.Disconnect())

}

func main(){
  test := &database.Redis{}
  test.IP = "localhost"
  test.Port = "6379"
  test.DB = 0
  test.Password = ""
  connectDB(test)
}
