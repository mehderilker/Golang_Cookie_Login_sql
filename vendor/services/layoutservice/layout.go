package layoutservice

import (
	"models"
	"services/hmacservice"
	"net/http"
	"services/menuservice"
)

func GetSharedData(title,description string ,r *http.Request)models.SimplePage{
	var data models.SimplePage
	data.GirisYapmismi = hmacservice.IsAuth(r)
	data.Category = menuservice.GetMenuList()
	data.Title = title
	data.Description=description
	return data
}