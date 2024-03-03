package movies

type Person struct {
	Role string
}

func (p Person) Contributes() bool {
	return true
}
