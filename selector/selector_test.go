package selector

import (
	"testing"

	"github.com/SeaSkyThe/GeneticPhraseFinder/population"
)

func TestNewSelector(t *testing.T) {
	selector := NewSelector()
	if selector == nil {
		t.Errorf("Expected selector to be not nil")
	}
}

func TestElitistSelection(t *testing.T) {
	selector := NewSelector()
	pop := population.GeneratePopulation(10, 10)
	elite := selector.ElitistSelection(pop)
	if len(elite) != 5 {
		t.Errorf("Expected elite to have 5 individuals, got %d", len(elite))
	}
}
