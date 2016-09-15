package textmash

type ColTfIdf struct {
	TFs           map[string]float64
	Idfs          map[string]float64
	DocCount      int
	mustCalculate bool
}

func TestMe() {
	_ = 1
}
