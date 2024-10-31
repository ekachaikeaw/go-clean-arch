package bmi

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalulateMBI(t *testing.T) {
	service := BMIService{}

	bmi, err := service.CalulteBMI(70, 1.75)
	assert.Nil(t, err)
    assert.InDelta(t, 22.86, bmi, 0.01)
}