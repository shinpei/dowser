package main

import (
"github.com/go-martini/martini"
"net/http"
"fmt"
"encoding/json"
"log"
"github.com/shinpei/dowser"
)

type Hoge struct {
  Body string;
}

func main () {
     m := martini.Classic()
     d := dowser.Create();
     d.Initiate("");

     h := &Hoge {"Hige"};
     b, err := json.Marshal(h)
     if err != nil { log.Fatal(err)}
     fmt.Println(string(b));

     m.Get("/", func(res http.ResponseWriter, req *http.Request) {
       fmt.Fprintf(res, "<html><head><style>form, input, textarea {font-size:13pt;font-family:\"Courier New\", Courier, monospace}</style><title>Dowser</title></head><body>")
       fmt.Fprintf(res, "<form method=POST>")
     	fmt.Fprintf(res, "Document:<textarea cols=60 rows=15 name=\"document\">(paste your document)</textarea><br/>\n");
       fmt.Fprintf(res, "Query: <input type=text name=\"query\"></input>\n");
       fmt.Fprintf(res, "</form>");
       fmt.Fprintf(res, "</body></html>");
       res.WriteHeader(200);
     });

     m.Post("/", func(res http.ResponseWriter, req *http.Request) {
       fmt.Fprintf(res, "<html><head><title>Dowser</title></head><body>")
       req.ParseForm();
       query := req.PostFormValue("query");
       rawDocument := req.PostFormValue("document");

       fmt.Fprintf(res, "query:'%s'<br>", query);
       doc := d.Encode2Document(rawDocument);
       d.UpdateIndex(doc);
       dResponse := d.Search(query);
       fmt.Fprintf(res, "result:%s<br>", dResponse.ToString());
       res.WriteHeader(200);
     });
     m.Run();
}
