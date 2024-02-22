package ipv4_test

import (
	"strings"
	"testing"

	"github.com/go-random/ipv4"
	"github.com/stretchr/testify/assert"
)

func TestIpv4(t *testing.T) {
	// Setup
	randomizer := ipv4.NewRandomizer()
	randomIpv4 := randomizer.GenerateRandomIpv4()

	//Act
	parts := strings.Split(randomIpv4.String(), ".")

	//Assert
	expectedParts := 4
	assert.Equal(t, expectedParts, len(parts), "An Ipv4 address should consist of four 8-bit numbers")
}
