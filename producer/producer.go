package producer

import (
	"challenge2016_distribution/distribution"
)

type Producer struct {
	Name         string
	Distributors []distribution.Distributor
}

func (p Producer) Contributes() bool {
	return true
}

func CreateProducer(name string, distributors []distribution.Distributor) Producer {
	return Producer{Name: name, Distributors: distributors}
}

func (p Producer) HaveNodes() bool {
	return len(p.Distributors) != 0
}

func (p Producer) PermittedRegions() string {
	return "World"
}

func (p Producer) Finances() bool {
	return true
}

func (p Producer) AddDistributors(distributors []distribution.Distributor) {
	p.Distributors = append(p.Distributors, distributors...)
}
