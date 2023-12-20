package main

type Agency interface {
	handleProjects() []string
}

type anAgency struct {
	name      string
	customers []string
}

type anotherAgency struct {
	name     string
	projects []string
}

func (a anAgency) handleProjects() []string {
	return []string{"project1", "project2"}
}

func (a anotherAgency) handleProjects() []string {
	return []string{"project3", "project4"}
}

func agencyProjects(a Agency) []string {
	return a.handleProjects()
}

func main() {
	agency := anAgency{name: "agency1", customers: []string{"customer1", "customer2"}}
	agencyProjects(agency)

	secondAgency := anotherAgency{name: "agency2", projects: []string{"project1", "project2"}}
	agencyProjects(secondAgency)
}
