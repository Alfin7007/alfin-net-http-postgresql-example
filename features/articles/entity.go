package articles

type Core struct {
	ArticleID int
	Title     string
	Detail    string
	UserID    int
	User      User
}
type User struct {
	UserID int
	Name   string
}

type Data interface {
	InsertData(Core) error
	SelectAll() ([]Core, error)
	SelectData(ID int) (Core, error)
}

type Bussiness interface {
	InsertArticle(Core) error
	GetAll() ([]Core, error)
	GetArticle(ID int) (Core, error)
}
