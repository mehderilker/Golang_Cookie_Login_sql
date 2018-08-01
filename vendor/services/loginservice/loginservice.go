package loginservice

import (
	"models"
	"services/database"
	"fmt"
)

func  GetloginDB(Username string,Password string)models.Response{
	var logs models.User
	query := "SELECT ID FROM users WHERE Username = ? AND Password = ?"
	err := database.DB.QueryRow(query,Username,Password).Scan(&logs.ID)

	if err != nil{
		fmt.Println(err)


	}
	var response models.Response
	response.Message = "Giriş başarısız.Lütfen  tekrar deneyiniz."
	if logs.ID > 0 {
		response.Name = Username
		response.ID=logs.ID
		response.Status=true
		response.Message="Giriş Başarılı"
	}

	return response


}
