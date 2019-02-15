package food

type Food struct {
	Id string
}

func NewFood(id string) *Food {
	f := Food{}
	f.Id = id
	return &f
}
