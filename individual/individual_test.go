package individual

import "testing"

func TestNewIndividual(t *testing.T) {
	ind := NewIndividual("abc123", 0)
	if ind.Genes != "abc123" {
		t.Errorf("Expected ind.Genes to be \"abc123\", got %s", ind.Genes)
	}
	if ind.Fitness != 0 {
		t.Errorf("Expected ind.Fitness to be 0, got %f", ind.Fitness)
	}
}

func TestCalculateFitness(t *testing.T) {
	ind := NewIndividual("abc123", 0)
	ind.CalculateFitness("abc123")
	if ind.Fitness != 1.0 + EPSILON {
		t.Errorf("Expected ind.Fitness to be 1.0, got %f", ind.Fitness)
    }
}
