package main

import (
  "fmt"
  "net/http"
  "os"
  "path"
)


func handler(w http.ResponseWriter, r *http.Request) {
  
  fmt.Fprintf(w, "<html></br><img src='/images/FRONTAL.jpg' ><br>Soy alumno de la UOC</br></html>")
  
}

func main() {
// rootdir comprueba que existe el directorio y no da error
rootdir, err := os.Getwd()
  if err != nil {
    rootdir = "No aparece"
  }


  
  http.Handle("/images/", http.StripPrefix("/images",
        http.FileServer(http.Dir(path.Join(rootdir, "images/")))))
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
