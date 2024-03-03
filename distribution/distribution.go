package distribution

type Distribution interface {
	HaveNodes() bool //returns true if the component has children/subdistributors
	PermittedRegions() string
}
