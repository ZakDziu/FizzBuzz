package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueValidation(t *testing.T) {
	//positive direction
	_, err1 := FizzBuzz(3, 2, 1, 1, PositiveDirection)
	assert.ErrorIs(t, err1, FromGreaterThanToError)

	_, err2 := FizzBuzz(2, 100, 3, 1, PositiveDirection)
	assert.ErrorIs(t, err2, FizzOrBuzzLessOrEqualThanFromError)
	_, err3 := FizzBuzz(2, 100, 1, 3, PositiveDirection)
	assert.ErrorIs(t, err3, FizzOrBuzzLessOrEqualThanFromError)
	_, err4 := FizzBuzz(2, 100, 2, 3, PositiveDirection)
	assert.ErrorIs(t, err4, FizzOrBuzzLessOrEqualThanFromError)
	_, err5 := FizzBuzz(2, 100, 3, 2, PositiveDirection)
	assert.ErrorIs(t, err5, FizzOrBuzzLessOrEqualThanFromError)

	_, err6 := FizzBuzz(1, 100, 101, 2, PositiveDirection)
	assert.ErrorIs(t, err6, FizzOrBuzzGreaterOrEqualThanToError)
	_, err7 := FizzBuzz(1, 100, 2, 101, PositiveDirection)
	assert.ErrorIs(t, err7, FizzOrBuzzGreaterOrEqualThanToError)
	_, err8 := FizzBuzz(1, 100, 100, 2, PositiveDirection)
	assert.ErrorIs(t, err8, FizzOrBuzzGreaterOrEqualThanToError)
	_, err9 := FizzBuzz(1, 100, 2, 100, PositiveDirection)
	assert.ErrorIs(t, err9, FizzOrBuzzGreaterOrEqualThanToError)

	//negative direction
	_, err10 := FizzBuzz(2, 5, 3, 3, NegativeDirection)
	assert.ErrorIs(t, err10, FromLessThanTo)

	_, err11 := FizzBuzz(10, 1, 11, 2, NegativeDirection)
	assert.ErrorIs(t, err11, FizzOrBuzzGreaterOrEqualThanFromError)
	_, err12 := FizzBuzz(10, 1, 2, 11, NegativeDirection)
	assert.ErrorIs(t, err12, FizzOrBuzzGreaterOrEqualThanFromError)
	_, err13 := FizzBuzz(10, 1, 10, 2, NegativeDirection)
	assert.ErrorIs(t, err13, FizzOrBuzzGreaterOrEqualThanFromError)
	_, err14 := FizzBuzz(10, 1, 2, 10, NegativeDirection)
	assert.ErrorIs(t, err14, FizzOrBuzzGreaterOrEqualThanFromError)

	_, err15 := FizzBuzz(10, 5, 4, 7, NegativeDirection)
	assert.ErrorIs(t, err15, FizzOrBuzzLessOrEqualThanToError)
	_, err16 := FizzBuzz(10, 5, 7, 4, NegativeDirection)
	assert.ErrorIs(t, err16, FizzOrBuzzLessOrEqualThanToError)
	_, err17 := FizzBuzz(10, 5, 5, 7, NegativeDirection)
	assert.ErrorIs(t, err17, FizzOrBuzzLessOrEqualThanToError)
	_, err18 := FizzBuzz(10, 5, 7, 5, NegativeDirection)
	assert.ErrorIs(t, err18, FizzOrBuzzLessOrEqualThanToError)

	//for both direction tests
	_, err19 := FizzBuzz(1, 10, 5, 5, PositiveDirection)
	assert.ErrorIs(t, err19, FizzEqualToBuzzError)
	_, err20 := FizzBuzz(10, 1, 5, 5, NegativeDirection)
	assert.ErrorIs(t, err20, FizzEqualToBuzzError)
}

func TestFizzBuzz(t *testing.T) {
	res, err := FizzBuzz(-10, 10, 2, 5, PositiveDirection)
	assert.NoError(t, err)
	assert.Equal(t, "FizzBuzz", res[0])
	assert.Equal(t, "-9", res[1])
	assert.Equal(t, "Fizz", res[4])
	assert.Equal(t, "Buzz", res[5])
	assert.Equal(t, "FizzBuzz", res[10])
	assert.Equal(t, "1", res[11])
	assert.Equal(t, "FizzBuzz", res[20])
	//you can see the result after uncomment below line
	//fmt.Println(res)

	res, err = FizzBuzz(10, -10, 2, 5, NegativeDirection)
	assert.NoError(t, err)
	assert.Equal(t, "FizzBuzz", res[0])
	assert.Equal(t, "9", res[1])
	assert.Equal(t, "Fizz", res[2])
	assert.Equal(t, "Buzz", res[5])
	assert.Equal(t, "Fizz", res[6])
	assert.Equal(t, "1", res[9])
	assert.Equal(t, "FizzBuzz", res[10])
	assert.Equal(t, "Fizz", res[14])
	assert.Equal(t, "Buzz", res[15])
	assert.Equal(t, "FizzBuzz", res[20])
	//you can see the result after uncomment below line
	//fmt.Println(res)

}
