package main

import (
	"flag"
	"fmt"
	"log"
  "net"
	"net/url"
  "net/http"
  "net/http/httputil"
  //"io/ioutil"
	//"os"
  //"reflect"
)

func main() {

	var input_url = flag.String("url", "default", "URL")
	flag.Parse()

	fmt.Println("URL:")
  fmt.Printf("  %s \n\n", *input_url)

	u, err := url.Parse(*input_url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Host:")
	fmt.Printf("  %s\n\n", u.Host)

  i, err := net.LookupIP(u.Host)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("IPAddress:")
//  fmt.Println(reflect.TypeOf(i))

  for _, v := range i{
    fmt.Printf("  %s\n", v)
  }
  fmt.Println("")

  fmt.Println("ResolveHostName:")

//  for _, v := range i{
//    uh, err := net.LookupIP(v)
//    if err != nil {
//      log.Fatal(err)
//    }
//    fmt.Printf("  %s:", uh)
//  }

//  req, _ := http.NewRequest("GET", *input_url, nil)
//
//  dump, _ := httputil.DUmpRequestOut(req, true)
//  fmt.Printf("%s", dump)
//



}
