package main

import (
  "flag"
  "fmt"
  "crypto/md5"
  "encoding/hex"
)
func main() {
  FlagsUsing()
}

func FlagsUsing(){
  var start = flag.Bool("s", false, "It starts the server listener")
  var port = flag.Int("p", 8001, "It sets a webserver listener port")
  var hash = flag.String("hash", "", "It sets the hash that you got in -gh (You can put in it another hash that you want)")
  var gHash = flag.String("gh", "", "It gives you a md5 string hash that will be the endpoint to your hook")
  flag.Parse()
  if *start {
     CreateServer(port, hash)

  }
  if *gHash != ""{
    fmt.Println(generateHash(*gHash) + " - put it in the hash flag to up out your webhook server")
  }
}

func generateHash (text string) string{
   hasher := md5.New()
   hasher.Write([]byte(text))
   return hex.EncodeToString(hasher.Sum(nil))
}
