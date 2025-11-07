package book

type Book struct {
	ID     int `gorm:"primary_key"`
	Title  string
	Author string
	ISBN   string `gorm:"unique"`
}
