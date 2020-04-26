package workplace

type Workplace struct {
	Name     string
	Location string
	// Ranking to priorize the best workplaces, if few users are present
	Ranking uint

	CurrentOwner string
}
