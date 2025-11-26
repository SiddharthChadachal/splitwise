package splitwise

type Participant struct {
	ID   string
	Name string
}

type Item struct {
	ID       string
	Name     string
	Price    float64
	SharedBy []string // Participant IDs
}

type Bill struct {
	ID           string
	Items        []Item
	Participants []Participant
	TaxPercent   float64
}
