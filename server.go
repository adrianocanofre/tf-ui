package main

import (
    "html/template"
    "net/http"
    "path/filepath"
    "log"
    "encoding/json"
    "io/ioutil"
    "os"
)



func server(w http.ResponseWriter, r *http.Request) {
    // Build path to template
    tmplPath := filepath.Join("./static", "index.html")
    // Load template from disk
    tmpl := template.Must(template.ParseFiles(tmplPath))

    file, _ := os.OpenFile(output_path, os.O_RDWR|os.O_APPEND, 0666)
    b, _ := ioutil.ReadAll(file)
    resources2 := Vpc{}
    json.Unmarshal([]byte(b), &resources2)

    //fmt.Println(resources2)
    // Inject data into template
    tmpl.Execute(w, &resources2)
}



func start() {

    http.HandleFunc("/", server)
    http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

    log.Println("Listening on :8000...")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
      log.Fatal(err)
    }
}
