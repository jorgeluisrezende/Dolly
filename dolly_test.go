package main
import (
  "os/exec"
  "fmt"
  "strconv"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

////////////////////////////////////////////////////////////////////////////////
func CreateServer(p *int, h *string){
  routes := mux.NewRouter().StrictSlash(true)
  routes.HandleFunc("/"+*h, deployhook)
  http.ListenAndServe(":"+strconv.Itoa(*p), routes)
  log.Println("listening on {*p}")
}

func deployhook(w http.ResponseWriter, r *http.Request){
  gitPullCmd()
}

func gitPullCmd(){
   Out, Err := exec.Command("git", "pull", "origin", "master").Output()

   if Err != nil{
    
    fmt.Println(fmt.Sprint(Err) + ": Out - Error on Pull")

   }else{

     s:= string(Out)
     fmt.Println(s)
       
   }
}