package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "text/template"
    "strings"
    "net/http"
)

type Friend struct {
    Fname string
}

type Person struct {
    UserName string
    Emails []string
    Friends []*Friend
}

var templates *template.Template

func main() {

    var allFiles []string
    var prefixPath string = "src/template"
    files, err := ioutil.ReadDir(prefixPath)
    if err != nil {
        fmt.Println(err)
    }
    for _, file := range files {
        filename := file.Name()
        if strings.HasSuffix(filename, ".tmpl") {
            allFiles = append(allFiles, prefixPath + "/" + filename)
        }
    }
    fmt.Println(allFiles)

    templates, err = template.ParseFiles(allFiles...)

    s1 := templates.Lookup("header.tmpl")
    s1.ExecuteTemplate(os.Stdout, "header", nil)
    fmt.Println()
    s2 := templates.Lookup("content.tmpl")
    s2.ExecuteTemplate(os.Stdout, "content", nil)
    fmt.Println()
    s3 := templates.Lookup("footer.tmpl")
    s3.ExecuteTemplate(os.Stdout, "footer", nil)
    fmt.Println()
    s3.Execute(os.Stdout, nil)


//    http.HandleFunc("/", handle)
//    err := http.ListenAndServe(":9090", nil)
//     if err != nil {
//         fmt.Println("error", err)
//     }
}

func handle (w http.ResponseWriter, r *http.Request){
    fmt.Println("hello")
    // t := template.New("welcome")
    // t, _ = t.ParseFiles("./template/welcome.html")
    // t.Execute(w, nil)
    fmt.Fprintf(w, "Hello")
}

