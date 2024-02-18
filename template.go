package template

// GenerateRandomNumber generates a random number using the full range of the randomizer.
func (tr *Randomizer) GenerateRandomNumber() int {
	return tr.Rand.Int()
}
