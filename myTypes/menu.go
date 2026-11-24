package myTypes

type MenuItemType struct {
	Title    string
	Slug     string
	Children []MenuItemType
}
