package main

import (
	"encoding/base64"
	"log"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) < 4 || len(os.Args)%2 != 0 {
		log.Fatal("must pass args like this: SECRETNAME K1 V1 K2 V2 ...")
	}
	tmpl, err := template.New("secret").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	var secret Secret
	secret.Name = os.Args[1]
	secret.Secrets = make(map[string]string)
	for i := 2; i < len(os.Args); i += 2 {
		val := base64.StdEncoding.EncodeToString([]byte(os.Args[i+1]))
		secret.Secrets[os.Args[i]] = val
	}

	err = tmpl.Execute(os.Stdout, secret)
	if err != nil {
		log.Fatalf("template execute: %v", err)
	}
}

var templateString = `
apiVersion: v1
kind: Secret
metadata:
  name: {{ .Name }}
type: Opaque
data:
{{- range $k, $v := .Secrets }}
  {{ $k }}: {{ $v -}}
{{ end }}
`

type Secret struct {
	Name    string
	Secrets map[string]string
}
