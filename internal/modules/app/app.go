package app

type App struct {
	Name    string
	Version string
}

func NewAppInfo() App {
	return App{
		Name:    Name,
		Version: Version,
	}
}
