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

func main() {
	agency := anAgency{name: "agency1", customers: []string{"customer1", "customer2"}}
	agency.handleProjects()

	secondAgency := anotherAgency{name: "agency2", projects: []string{"project1", "project2"}}
	secondAgency.handleProjects()
}
