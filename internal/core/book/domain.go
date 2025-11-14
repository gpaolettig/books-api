package book

type Book struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `gorm:"unique" json:"isbn"`
}
