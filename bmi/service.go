package bmi

import "errors"

type BMIService struct{}

func (bms *BMIService) CalulteBMI(weight, height float64) (float64, error) {
	if height <= 0 {
		return 0, errors.New("height must greater than 0")
	}

	return weight / (height * height), nil
}
