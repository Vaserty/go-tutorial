package calculators

import (
	"math"

	"github.com/go-playground/validator/v10"
)

type BMICalculateInput struct {
	Weight float64 `validate:"required,gte=0"`
	Height float64 `validate:"required,gte=0"`
}

func BMI(input *BMICalculateInput) (float64, error) {
	validate := validator.New()
	errors := validate.Struct(input)
	if errors != nil {
		return 0, errors
	}
	bmi := input.Weight / math.Pow(input.Height, 2)
	return RoundFloat(bmi, 2), nil
}
