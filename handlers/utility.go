package handlers

import(
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	ttemplate "text/template"
)

// Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}){
	t, err := template.ParseFiles(templateFile)
	if err != nil{
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

// this allows XSS due to use of the text/template package
// text/template does not escape user input
// use html/template package when rendering templates 
func RenderUnsafeTemplate(w http.ResponseWriter, templateFile string, templateData interface{}){
	t, err := ttemplate.ParseFiles(templateFile)
	if err != nil{
		log.Printf("Error encountered while aprsing the template: ", err)
	}
	// here we tell web broswer to keep loading even if xss is detected
	w.Header().Set("X-XSS-Protection", "0")
	t.Execute(w, templateData)
}

func GenerateUUID() string{
	f, err := os.Open("/dev/urandom")
	if err != nil{
		log.Println("Encountered the following error when attempting to generate an UUID: ", err)
		return ""
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[8:10], b[10:])
	return uuid
}