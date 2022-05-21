package slotMachine_test

import (
	"testing"

	"github.com/darvoid/gRPC-slotMachine/slotMachine"
)

type indexTest struct {
	in       int
	in2      int
	expected string
	name     string
}

var indexTests = []indexTest{
	{100, 10, "Total de Jogadas: 100\nWinning chance: 10\nTotal Wins: 0\nLast Element: 0\n", "setup 10%"},
	{200, 30, "Total de Jogadas: 200\nWinning chance: 30\nTotal Wins: 0\nLast Element: 0\n", "setup 30%"},
	{300, 20, "Total de Jogadas: 300\nWinning chance: 20\nTotal Wins: 0\nLast Element: 0\n", "setup 20%"},
}

//tests Game startup
func TestSetup(t *testing.T) {
	for _, test := range indexTests {
		t.Run(test.name, func(t *testing.T) {

			game, err := slotMachine.Setup(test.in, test.in2)
			if err != nil {
				t.Errorf("Failed to setup: %v\n", test.name)

			}
			result := game.OutputCheckGameState()
			if result != test.expected {
				t.Errorf("Expected result not given for test: %v\n", test.name)
			}
		})

	}
}
