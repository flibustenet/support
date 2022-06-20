//go:build OMIT

package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.New("").Parse(`
Bonjour {{.Nom}}
`))
	t.Execute(os.Stdout, map[string]interface{}{
		"Nom": "Alice",
	})

	type S struct {
		Nom string
	}
	t.Execute(os.Stdout, S{
		Nom: "Bob",
	})
}
