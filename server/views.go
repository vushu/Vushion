package main

import(
    "net/http"
    "html/template"
    "bytes"
    //"io"
    //"io/ioutil"
)
var HTML_DIR = "tmpl/"
type content struct {
        MenuHTML    template.HTML
        ContentHTML template.HTML
}

func parseTemplate(file string, data interface{}) (out []byte, error error) {
        var buf bytes.Buffer // buffer til skrive html ned i
        t, err := template.ParseFiles(file)
        if err != nil {
                return nil, err
        }
        err = t.Execute(&buf, data)// kan skrive til buffer som skal bruges til at lave rå html
        if err != nil {
                return nil, err
        }
        return buf.Bytes(), nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    intro,_ := parseTemplate(HTML_DIR+"welcome.html",nil)
    t := makeBase(intro)
    w.Write(t) 
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
    a, _ := parseTemplate(HTML_DIR+"about.html",nil)
    w.Write(makeBase(a))
}

func gamesHandler(w http.ResponseWriter, r *http.Request){
    a, _ := parseTemplate(HTML_DIR+"games.html",nil)
    w.Write(makeBase(a))
}

func pixelartHandler(w http.ResponseWriter, r *http.Request){
    a, _ := parseTemplate(HTML_DIR+"pixelart.html",nil)
    w.Write(makeBase(a))
}

func testHandler(w http.ResponseWriter, r *http.Request){
    a, _ := parseTemplate(HTML_DIR+"submittest.html",nil)
    w.Write(makeBase(a))
}
func makeBase(cont []byte) (out []byte){
    menu, _ := parseTemplate(HTML_DIR+"menu.html",nil)
    base, _ := parseTemplate(HTML_DIR+"base.html",content {MenuHTML: template.HTML(menu), ContentHTML: template.HTML(cont)})
    return base;
}

