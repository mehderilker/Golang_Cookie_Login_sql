package photoservice

import (
	"models"
	"services/database"
	"fmt"
	"services/convert"
)

func Insert(model models.Photo)models.Response{
	query := "INSERT photo SET PhotoName=?,Date=NOW()"
	res,err :=database.DB.Exec(query,model.Pname)

	var response models.Response
	if err != nil{
		fmt.Println(err)
		response.Message = "Bir hata oluştu.Lütfen tekrar deneyiniz."
		return  response
	}
	response.Status =true
	response.Message = "İşlem başarılı.Fotoğraf başarı ile kayıt edildi."
	lastid,_ :=res.LastInsertId()
	response.ID = int(lastid)
	return response
}



func DeleteInfo(ID int)models.Response{
	query := "DELETE FROM photo WHERE ID = ?"
	_ ,err := database.DB.Exec(query,ID)

	var response models.Response

	if err != nil {
		fmt.Println(err)
		response.Message = "Silinirken Bir Hata Oluştu"
		return response
	}

	response.Status = true
	response.Message = "KAyıt Başarılı ile  silindi."

	return response
}

func GetList() []models.Photo{
	query := "SELECT ID,PhotoName,Date FROM photo ORDER BY ID DESC"
	rows,err := database.DB.Query(query)
	defer rows.Close()
	if err != nil{
		fmt.Println(err)
	}
	var list []models.Photo
	for rows.Next(){
		var model models.Photo
		rows.Scan(&model.ID,&model.Pname,&model.Date)
		model.Datestr = convert.ToDateString(model.Date)
		list =append(list,model)
	}
	return list
}

func GetPhotoInfo(ID int)models.Photo{

	var model models.Photo
	query := "SELECT ID,PhotoName,Date FROM photo WHERE ID = ?"
	err := database.DB.QueryRow(query,ID).Scan(&model.ID,&model.Pname,&model.Date)

	if err != nil {
		fmt.Println(err)
	}

	return model
 }

