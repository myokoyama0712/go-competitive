package lake_counting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleLakeCounting(t *testing.T) {
	input := [][]rune{
		[]rune{'W', '.', '.', '.', '.', '.', '.', '.', '.', 'W', 'W', '.'},
		[]rune{'.', 'W', 'W', 'W', '.', '.', '.', '.', '.', 'W', 'W', 'W'},
		[]rune{'.', '.', '.', '.', 'W', 'W', '.', '.', '.', 'W', 'W', '.'},
		[]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'W', 'W', '.'},
		[]rune{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'W', '.', '.'},
		[]rune{'.', '.', 'W', '.', '.', '.', '.', '.', '.', 'W', '.', '.'},
		[]rune{'.', 'W', '.', 'W', '.', '.', '.', '.', '.', 'W', 'W', '.'},
		[]rune{'W', '.', '.', '.', 'W', '.', '.', '.', '.', '.', 'W', '.'},
		[]rune{'.', 'W', '.', 'W', '.', '.', '.', '.', '.', '.', 'W', '.'},
		[]rune{'.', '.', 'W', '.', '.', '.', '.', '.', '.', '.', 'W', '.'},
	}
	actual := countLake(input)
	assert.Equal(t, 3, actual)
}
