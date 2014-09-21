package main

import (
  "github.com/go-martini/martini"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "crypto/md5"
  "encoding/hex"
  "github.com/erlendr/store"
)

func main() {
  m := martini.Classic()

  m.Get("/", func() string {
    return "Welcome to Uploader."
  })

  m.Post("/", func(res http.ResponseWriter, req *http.Request) string {
    // Fetch image from form
    file, header, err := req.FormFile("image")
    if(err != nil) {
      println("Error while receiving upload: ", err)
      panic(err)
    }

    //Read image data
    data, err := ioutil.ReadAll(file)
    if(err != nil) {
      println("Error while reading file: ", err)
      panic(err)
    }

    //Fetch last four chars from filename as extension
    extension := header.Filename[(len(header.Filename)-4):]
    
    h := md5.New()
    filename := hex.EncodeToString(h.Sum([]byte(header.Filename + string(req.ContentLength))))

    err = ioutil.WriteFile("temp/" + filename + extension, data, 0644)
    if err != nil {
      panic(err)
    }

    println("Filename:", header.Filename)

    store.Upload(filename + extension)

    res.WriteHeader(200)

    out, err := json.Marshal(filename + extension)
    if(err != nil) {
      panic(err)
    }
    return string(out)
  })

  m.Run()
}