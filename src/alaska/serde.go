package alaska

type Serde interface {
	ProcessLine(line string) []string
}
