package main

import (
	"fmt"
	"os"
	"text/template"
)

const (
	projectTemplate = `{
    "folders":
    [
        {
            "path": "{{ .ProjectDir }}"
        }
    ],
    "settings": {
        "GoSublime": {
            "env": {
                "GOPATH": "{{ .ProjectDir }}:{{ .ProjectDir }}/vendor",
                "PATH": "$PATH:$HOME/src/go/bin"
            }
        }
    }
}
`
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("writing sublime-project to %s...\n", wd)

	tmpl := template.Must(template.New("project").Parse(projectTemplate))

	out, err := os.Create("project.sublime-project")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, struct{ ProjectDir string }{wd})
	if err != nil {
		panic(err)
	}

}
