package main

import (
 "net/http"
 "os"
 "net/url"
)

func redirect_and_store(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "http://devopsdeflope.ru/"+r.URL.Path[1:], http.StatusFound)

  values := make(url.Values)
  ga_cookie :="0"

  for i:= range r.Cookies() {
    if r.Cookies()[i].Name == "_ga" {
      ga_cookie = r.Cookies()[i].Value
    }
  }

  values.Set("v","1")
  values.Set("tid","UA-41332661-1")
  values.Set("cid",ga_cookie)
  values.Set("t","event")
  values.Set("ea","Download")
  values.Set("el",r.URL.Path)
  values.Set("ec","Podcast")
  http.PostForm("http://www.google-analytics.com/collect",url.Values(values))
}

func serve() (func (http.ResponseWriter, *http.Request)) {

  return func (w http.ResponseWriter, r *http.Request) {
    redirect_and_store(w,r)
  }
}

func main() {
  http.HandleFunc("/mp3/", serve())
  srv := http.Server{Addr: "0.0.0.0:8081"}
  srv.ListenAndServe()
  os.Exit(0)
}
