package textmash

func NewColTfIdf() *ColTfIdf {
	var ctfidf ColTfIdf
	ctfidf.TFs = make(map[string]float64)
	ctfidf.mustCalculate = false
	return &ctfidf
}

func (self *ColTfIdf) AddDoc(tokens []string) {

	self.mustCalculate = true
	for _, e := range tokens {
		_, ok := self.TFs[e]
		if ok {
			self.TFs[e]++
		} else {
			self.TFs[e] = 1
		}

	}
	self.DocCount++

}
