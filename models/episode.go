package models

type Episode struct {
	ID            int          `json:"id" gorm:"primary_key:auto_increment"`
	Title         string       `json:"title" from:"title"  gorm:"type: varchar(255)"`
	Thumbnailfilm string       `json:"thumbnailfilm" gorm:"type: varchar(255)"`
	Linkfilm      string       `json:"linkfilm" gorm:"type:text"`
	FilmID        int          `json:"-"`
	Film          FilmResponse `json:"film"`
}
