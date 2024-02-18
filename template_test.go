package template_test

import (
	"testing"

	template "github.com/go-random/template"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomNumber(t *testing.T) {
	// Setup
	randomizer := template.NewRandomizer()
	randomNumber := randomizer.GenerateRandomNumber()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "Generated number should not be 0")
}
