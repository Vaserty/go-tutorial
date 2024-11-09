package calculators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBMIReturnCorrectValueForCorrectHeightAndWeight(t *testing.T) {
	input := &BMICalculateInput{
		Weight: 108,
		Height: 1.86,
	}
	expectedResult := 31.22
	calculatedBMI, err := BMI(input)
	assert.Equal(t, calculatedBMI, expectedResult, "they should be equal")
	assert.Nil(t, err)
}

func TestBMIReturnErrorWhenInputIsIncorrect(t *testing.T) {
	input := &BMICalculateInput{
		Weight: -108,
		Height: 1.86,
	}
	_, err := BMI(input)
	assert.NotNil(t, err)

}
