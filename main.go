package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	result, err := FizzBuzz(1, 100, 2, 5, PositiveDirection)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}
}

func FizzBuzz(from, to, fizz, buzz int, direction DirectionType) ([]string, error) {
	if err := validateValues(from, to, fizz, buzz, direction); err != nil {
		return []string{}, err
	}
	var result []string
	if direction == PositiveDirection {
		for i := from; i <= to; i++ {
			switch {
			case i%fizz == 0 && i%buzz == 0:
				result = append(result, "FizzBuzz")
			case i%fizz == 0:
				result = append(result, "Fizz")
			case i%buzz == 0:
				result = append(result, "Buzz")
			default:
				result = append(result, strconv.Itoa(i))
			}
		}
	} else {
		for i := from; i >= to; i-- {
			switch {
			case i%fizz == 0 && i%buzz == 0:
				result = append(result, "FizzBuzz")
			case i%fizz == 0:
				result = append(result, "Fizz")
			case i%buzz == 0:
				result = append(result, "Buzz")
			default:
				result = append(result, strconv.Itoa(i))
			}
		}
	}
	return result, nil
}

func validateValues(from, to, fizz, buzz int, direction DirectionType) error {
	if from <= 0 || to <= 0 || fizz <= 0 || buzz <= 0 {
		return AllValuesNeedToBeGreaterThanZeroError
	}
	if direction == PositiveDirection {
		if from > to {
			return FromGreaterThanToError
		}
		if fizz <= from || buzz <= from {
			return FizzOrBuzzLessOrEqualThanFromError
		}
		if fizz >= to || buzz >= to {
			return FizzOrBuzzGreaterOrEqualThanToError
		}
	} else {
		if from < to {
			return FromLessThanTo
		}
		if fizz >= from || buzz >= from {
			return FizzOrBuzzGreaterOrEqualThanFromError
		}
		if fizz <= to || buzz <= to {
			return FizzOrBuzzLessOrEqualThanToError
		}
	}
	if fizz == buzz {
		return FizzEqualToBuzzError
	}

	return nil
}
