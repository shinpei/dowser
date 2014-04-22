package main

import (
"github.com/go-martini/martini"
"net/http"
"fmt"
"github.com/shinpei/dowser"
)

func main () {
     m := martini.Classic()
     m.Get("/", func(res http.ResponseWriter, req *http.Request) {
     	fmt.Fprintf(res, "Hello world\n");
       res.WriteHeader(200);
     });
     d := dowser.Create();
     d.Initiate("");
     m.Run();
}
