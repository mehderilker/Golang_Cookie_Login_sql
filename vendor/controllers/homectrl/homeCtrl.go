package homectrl

import (
	"net/http"
	"models"
	"services/display"

	"services/menuservice"
	"io"
	"fmt"
	"strings"
	"crypto/sha1"
	"os"
	"path/filepath"
	"github.com/nu7hatch/gouuid"
	"services/photoservice"
	"services/layoutservice"
)

func Index(w http.ResponseWriter,r *http.Request){
	var data models.Home
	data.Page = layoutservice.GetSharedData("Anasayfa","Anasayfa",r)

	display.View(w,r,"homeIndex",data)
}
func GetCookiew(w http.ResponseWriter,r *http.Request) *http.Cookie{
	c,err := r.Cookie("session")

	if err != nil {
		sID,_ := uuid.NewV4()
		c = &http.Cookie{
			Name:"session",
			Value:sID.String(),
		}
		http.SetCookie(w,c)
	}
	return c
}

func IndexPost(w http.ResponseWriter,r *http.Request){

	var data models.Home

	data.Page.Category = menuservice.GetMenuList()

	var model models.Photo

	c := GetCookiew(w,r)

		mf,fh,err := r.FormFile("nf")
		if err != nil {
			fmt.Println(err)
			display.View(w,r,"homeIndex",data)
			return
		}
		defer mf.Close()

		ext := strings.Split(fh.Filename,".")[1]
		h := sha1.New()
		io.Copy(h,mf)

		model.Pname.Scan(fmt.Sprintf("%x",h.Sum(nil))+"."+ext)

		wd,err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			display.View(w,r,"homeIndex",data)
			return
		}

		path := filepath.Join(wd,"public","pics",model.Pname.String)
		nf,err := os.Create(path)

		if err != nil {
			fmt.Println(err)
			display.View(w,r,"homeIndex",data)
			return
		}
		var response models.Response
		response = photoservice.Insert(model)
		data.Status = response.Status

		defer nf.Close()

		mf.Seek(0,0)
		io.Copy(nf,mf)

		c =AppendValue(w,c,model.Pname.String)

	//data.FileNames = strings.Split(c.Value,"|")
	data.FileName = model.Pname.String
	display.View(w,r,"homeIndex",data)


}


func AppendValue(w http.ResponseWriter, cookie *http.Cookie,fname string)*http.Cookie{

	s := cookie.Value

	if !strings.Contains(s,fname){
		s += "|" + fname
	}

	cookie.Value = s
	http.SetCookie(w,cookie)

	return cookie
}