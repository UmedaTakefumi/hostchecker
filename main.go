package main

import (
	"flag"
	"fmt"
	"log"
  "net"
	"net/url"
	//  "os"
  //"reflect"
)

func main() {

	var input_url = flag.String("url", "default", "URL")
	flag.Parse()

	fmt.Printf("URL: %s \n", *input_url)

	u, err := url.Parse(*input_url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Host: %s\n", u.Host)

  i, err := net.LookupIP(u.Host)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("IPAddress:")
//  fmt.Println(reflect.TypeOf(i))

  for k, v := range i{
    fmt.Printf("  %d: %s\n", k, v)
  }


//  uh, err := net.LookupIP(i)
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Printf("A recoard: %s\n", uh)


}
