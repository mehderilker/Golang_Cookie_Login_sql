package display

import (
	"net/http"
	"fmt"
	"encoding/json"
)

// View Girilen template'i sayfaya basar.
func View(w http.ResponseWriter,r *http.Request , tpl string, data interface{}){
	var shouldReset = true
	if shouldReset{
		LoadTemplates()
	}

	err := tmpl.ExecuteTemplate(w,tpl,data)
	if err != nil {
		fmt.Println(err)
	}
}

func Json (w http.ResponseWriter,data interface{}){
	w.Header().Set("Content-Type","application/json")
	jData,_ := json.Marshal(data)
	w.Write(jData)
	return
}