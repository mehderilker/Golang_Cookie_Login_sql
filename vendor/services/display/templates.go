package display

import (
	"os"
	"path/filepath"
	"html/template"
)

var tmpl *template.Template

//template Dosyalarını Yükler.
func LoadTemplates(){
	path ,err := GetAllFilePathsInDirectory("template")
	if err != nil {
		return
	}

	tmpl = template.Must(template.New("").ParseFiles(path...))

}

// Klasötdeki Tüm Template Dosyalarını Çeker

func GetAllFilePathsInDirectory(dirpath string)([]string,error){
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo,err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir(){
			paths = append(paths,path)
		}
		return nil
	})
	if err!=nil {
		return nil,err
	}

	return paths,nil
}