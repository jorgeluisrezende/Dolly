package main
import (
  "os/exec"
  "fmt"
  "strconv"
  "log"
  "io/ioutil"
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
    
    s := fmt.Sprint(Err) + ": Out - Error on Pull"
    writeFile(s+"\n", "error")

   }else{

     s:= string(Out)
     writeFile(s+"\n", "success")     
   }
}

func writeFile(in string, logName string){
  fout := []byte(in);
  ioutil.WriteFile("./"+logName+"-logs.txt", fout, 0644)
}
