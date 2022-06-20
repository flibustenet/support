//go:build OMIT

package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed template.html
var templateFile string

//go:embed *.html
var templateFiles embed.FS

func main() {
	fmt.Println(templateFile)

	t := template.Must(template.ParseFS(templateFiles, "*.html"))
	t.ExecuteTemplate(os.Stdout, "template.html", map[string]string{"Nom": "Toto"})
}
