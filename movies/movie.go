package movies

type Movie struct {
	Name string
	Crew []Crew
}

type Crew interface {
	Contributes() bool
}

func CreateMovie(name string) *Movie {
	movie := Movie{}
	movie.Name = name
	movie.Crew = make([]Crew, 0)
	return &movie
}

func (m *Movie) AddProducer(prod Crew) {
	m.Crew = append(m.Crew, prod)
}

func (m *Movie) AddCrew(role Crew) {
	m.Crew = append(m.Crew, role)
}

func (m *Movie) GetCrew() []Crew {
	return m.Crew
}
