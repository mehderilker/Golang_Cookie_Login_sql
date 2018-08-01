package models

type SimplePage struct {
	Title string
	Description string
	Category []Part
	GirisYapmismi bool
}

type Home struct {
	Page SimplePage
	FileName string
	Status bool
	Response Response
	PostOlduMu bool
	User User

}
type Part struct {
	Category string
}

type PhotoPage struct {
	Page SimplePage
	PhotoList []Photo
}

type PhotoUpdate struct {
	Page SimplePage
	PhotoInfo Photo
}
type PhotoDelete struct {
	Page SimplePage
	PhotoInfo Photo
	Response Response
}

