package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func LoadTemplates(templateDir string) (*template.Template, error) {
	tmpl := template.New("main")
	tmpl.Funcs(template.FuncMap{
		"dict": dict,
	})

	tmpl.ParseGlob(fmt.Sprintf("%s/*.html", templateDir))
	return tmpl, nil
}

func dict(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call: number of arguments must be even")
	}

	dict := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func RenderPage(tmpl template.Template, pageTemplate string, data any, w http.ResponseWriter) error {
	var b bytes.Buffer
	err := tmpl.ExecuteTemplate(&b, "page_til", data)
	if err != nil {
		// TODO: error page
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	b.WriteTo(w)
	return nil
}
