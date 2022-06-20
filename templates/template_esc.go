//go:build OMIT

package main

import (
	"html/template"
	"os"
)

func main() {
	t := template.Must(template.New("t.html").Parse(`
<html>
<p>{{.Para}}</p>

<a title='{{.Title}}' 
   onClick="return confirm({{.Confirm}})" >link</a>

<script>
var s = {{.Struct}};
</script>
</html>
`))
	type S struct {
		A string
		I int
	}
	t.ExecuteTemplate(os.Stdout, "t.html", map[string]interface{}{
		"Para":    "< & > Texte avec échappement automatique < & >",
		"Title":   `c'est "entre guillemet" ?`,
		"Confirm": `c'est "entre guillemet" ?`,
		"Struct":  []S{{`c'est "entre guillemet" ?`, 5}, {"X", 0}},
	})

}
