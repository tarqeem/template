package utl

import (
	"embed"
	"html/template"
	"io"
	"log"
	"path"
)

var Views embed.FS
var TemplateFuncs template.FuncMap

func getFSFilesRecursively(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := getFSFilesRecursively(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}
	return
}

type TemplateExecutor interface {
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type DebugTemplateExecutor struct {
	Glob []string
}

func (e DebugTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	t, err := GetTemplates()
	if err != nil {
		return err
	}
	return t.ExecuteTemplate(wr, name, data)
}

type ReleaseTemplateExecutor struct {
	Template *template.Template
}

func (e ReleaseTemplateExecutor) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return e.Template.ExecuteTemplate(wr, name, data)
}

func GetTemplates() (*template.Template, error) {
	files, err := getFSFilesRecursively(&Views, "pages")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	ts := template.Must(template.New("").Funcs(TemplateFuncs).ParseFiles(files...))

	return ts, err

}
