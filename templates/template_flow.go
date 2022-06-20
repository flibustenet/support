//go:build OMIT

package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.New("").Parse(`
Liste des véhicules :

{{range .Vehicules}}
  Véhicule : {{.}} 
  {{if eq . "Vélo"}} avec deux roues {{else}} avec quatre roues {{end}}
{{end}}
`))
	t.Execute(os.Stdout, map[string]interface{}{
		"Vehicules": []string{"Voiture", "Vélo"},
	})
}
