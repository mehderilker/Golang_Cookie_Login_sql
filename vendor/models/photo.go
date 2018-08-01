package models

import (
	"time"
	"database/sql"
)

type Photo struct {
	ID int
	Pname sql.NullString
	Date time.Time
	Datestr string
	Status bool
	PhotoAttr PhotoAttr
}

type PhotoAttr struct {
	X,Y int
	Blur,Brightness float64
	Width int
	Height int
	FileName string
	Path string

}

