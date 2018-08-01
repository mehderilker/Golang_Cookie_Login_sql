package menuservice

import "models"

func GetMenuList()(list []models.Part){
	var data models.Part
	data.Category = "anasayfa"
	list = append(list,data)
	data.Category = "editpic"
	list = append(list,data)
	data.Category = "hakkimizda"
	list = append(list,data)
	data.Category = "iletişim"
	list = append(list,data)
	return list
}