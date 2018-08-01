package photoctrl

import (
	"net/http"
	"models"
	"services/photoservice"
	"services/display"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
	"os"
	"services/fileservice"
	"services/convert"
	"services/hmacservice"
	"services/layoutservice"
)

func Index(w http.ResponseWriter , r *http.Request){
	var data models.PhotoPage
	data.Page = layoutservice.GetSharedData("Photo","Photo",r)
	data.PhotoList = photoservice.GetList()

	fmt.Println("PhotoCtrl İndex Bölümü")
	display.View(w,r,"Photos",data)

	}


func Update(w http.ResponseWriter , r *http.Request){
		vars := mux.Vars(r)
		IDstr := vars["ID"]
 		ID , err := strconv.Atoi(IDstr)
 		if err !=nil {
 			fmt.Println(err)
 		}
		var data models.PhotoUpdate
		data.Page = layoutservice.GetSharedData("Photo","Photo",r)
		data.PhotoInfo=photoservice.GetPhotoInfo(ID)
		FileName := data.PhotoInfo.Pname.String
		width,height:= fileservice.PhotoAttr(FileName)
		data.PhotoInfo.PhotoAttr.X=width
		data.PhotoInfo.PhotoAttr.Y=height



		display.View(w,r,"photoEdit",data)
}

func DeletePhoto(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var idStr = r.FormValue("ID")
	var id,_ = strconv.Atoi(idStr)
	var photoInfo = photoservice.GetPhotoInfo(id)
	var response = photoservice.DeleteInfo(id)
	if response.Status{
		if photoInfo.Pname.String != ""{
			os.Remove("./public/pics/"+photoInfo.Pname.String)
		}
	}
	display.Json(w,response)
}

func UpdatePost(w http.ResponseWriter , r *http.Request){
	r.ParseForm()
	var ID = convert.ToInt(r.FormValue("ID"))
	var data models.PhotoUpdate
	data.Page = layoutservice.GetSharedData("Photo","Photo",r)
	data.Page.GirisYapmismi = hmacservice.IsAuth(r)
	data.PhotoInfo=photoservice.GetPhotoInfo(ID)
	FileName := data.PhotoInfo.Pname.String
	width,height:= fileservice.PhotoAttr(FileName)
	data.PhotoInfo.PhotoAttr.X=width
	data.PhotoInfo.PhotoAttr.Y=height

	var attr models.PhotoAttr
	attr.FileName = FileName
	fmt.Println(attr.FileName)
	attr.Width = convert.ToInt(r.FormValue("widthEdit"))
	attr.Height = convert.ToInt(r.FormValue("heightEdit"))
	attr.Blur= convert.ToFloat64(r.FormValue("BlurEdit"))
	attr.Brightness = convert.ToFloat64(r.FormValue("BrightnessEdit"))
	data.PhotoInfo.PhotoAttr.Blur=attr.Blur
	data.PhotoInfo.PhotoAttr.Brightness = attr.Brightness


	if r.Method == http.MethodPost{
		fileservice.NewPhotoAttrSave(attr)

		http.Redirect(w,r,"/editpic/duzenle/"+convert.ToString(ID),301)
		return
		}

	display.View(w,r,"photoEdit",data)


}
