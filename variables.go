package main

import "errors"

type DirectionType string

const (
	NegativeDirection DirectionType = "negative"
	PositiveDirection DirectionType = "positive"
)

var (
	FromGreaterThanToError                = errors.New("'from' can't be greater than 'to' for positive direction")
	FizzOrBuzzLessOrEqualThanFromError    = errors.New("fizz/buzz value can't be less than or equal to 'from' for positive direction")
	FizzOrBuzzGreaterOrEqualThanToError   = errors.New("fizz/buzz value can't be greater than or equal to 'to' for positive direction")
	FromLessThanTo                        = errors.New("'from' can't be less than 'to' for negative direction")
	FizzOrBuzzGreaterOrEqualThanFromError = errors.New("fizz/buzz value can't be greater than or equal to 'from' for negative direction")
	FizzOrBuzzLessOrEqualThanToError      = errors.New("fizz/buzz value can't be less than or equal to 'to' for negative direction")
	FizzEqualToBuzzError                  = errors.New("'fizz' can't be equal to 'buzz'")
	AllValuesNeedToBeGreaterThanZeroError = errors.New("all values need to ge greater than zero")
)
