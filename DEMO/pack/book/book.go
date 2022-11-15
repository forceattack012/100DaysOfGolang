package book

type book struct {
	Name string
}

func New() book {
	return book{Name: "test"}
}

func new() book {
	return book{}
}
