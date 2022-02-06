package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	message string
	details string
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, %s", err.details)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (wantAmount float64, err error) {
	fodder, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		if fodder < 0 {
			return 0.0, errors.New("negative fodder")
		}

		return fodder * 2.0 / float64(cows), nil
	} else if err != nil {
		return 0.0, err
	}

	if fodder < 0 {
		return 0.0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	} else if cows < 0 {
		return 0.0, &SillyNephewError{details: fmt.Sprintf("there cannot be %d cows", cows)}
	}

	wantAmount = fodder / float64(cows)
	err = nil

	return
}
