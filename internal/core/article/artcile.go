package article

type Category struct {
	ID   int
	Name string
}

type DBModel struct {
	ID         int
	Title      string
	Content    string
	Categories []Category
}
