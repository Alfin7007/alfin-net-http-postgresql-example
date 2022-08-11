package users

type Core struct {
	UserID   int
	Name     string
	Email    string
	Password string
}

type Data interface {
	InsertUser(Core) error
	FindUser(Email string) (Core, error)
	SelectUser(id int) (Core, error)
}

type Bussiness interface {
	Register(Core) error
	Login(Core) (id int, token string, err error)
	GetData(id int) (Core, error)
}
