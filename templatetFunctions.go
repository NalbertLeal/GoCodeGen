package main

import (
	"bytes"
	"text/template"
)

func genFromTemplate(temaplateName string, content string, variables map[string]string) (*bytes.Buffer, error) {
	buf := &bytes.Buffer{}
	t := template.Must(template.New(temaplateName).Parse(content))
	err := t.Execute(buf, variables)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
