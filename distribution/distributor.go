package distribution

import (
	"challenge2016_distribution/distribution/permissions/models"
	"fmt"
)

type Distributor struct {
	Info            Info
	Subdistributors []Distribution
}

type Info struct {
	Name        string
	Permissions models.Specifics
	Parent      string
}

func (d Distributor) HaveNodes() bool {
	return len(d.Subdistributors) == 0
}

func (d Distributor) PermittedRegions() string {
	area := "include: \n"
	for i, region := range d.Info.Permissions.Include {
		area = area + " " + fmt.Sprintf("%d. ", i+1) + region + "\n"
	}
	area += "exclude: \n"
	for i, region := range d.Info.Permissions.Exclude {
		area = area + " " + fmt.Sprintf("%d. ", i+1) + region + "\n"
	}
	return area
}
