package gauss

import (
	"errors"
	"log"
	"strconv"
)

// Operation -
type Operation struct {
	Values []interface{}
}

func calculatePrimes() (primes []int64) {
	var notPrimes = make([]bool, (3400*3400)+1)
	for i := 2; i < 3400; i++ {
		for j := i; j < 3400; j++ {
			notPrimes[i*j] = true
		}
	}

	for i := 2; i < len(notPrimes); i++ {
		if !notPrimes[i] {
			primes = append(primes, int64(i))
		}
	}

	return
}

// Add -
func (o *Operation) Add(flag bool) (sum interface{}, err error) {
	var (
		intSum   int64
		floatSum float64
	)

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to sum less than 2 operands")
	}

	for _, num := range o.Values {

		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		if flag {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			intSum += number
		} else {
			s := num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			floatSum += number
		}
	}

	if flag {
		sum = intSum
	} else {
		sum = floatSum
	}

	return
}

// Sub -
func (o *Operation) Sub(flag bool) (sub interface{}, err error) {
	var (
		intSub   int64
		floatSub float64
		firt     bool = true
	)

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to subtract less than 2 operands")
	}

	for _, num := range o.Values {

		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		if flag {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			if firt {
				intSub = number
			} else {
				intSub -= number
			}

			firt = false

		} else {
			s := num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			if firt {
				floatSub = number
			} else {
				floatSub -= number
			}

			firt = false
		}
	}

	if flag {
		sub = intSub
	} else {
		sub = floatSub
	}

	return
}

// Prod -
func (o *Operation) Prod(flag bool) (prod interface{}, err error) {
	var (
		intProd   int64   = 1
		floatProd float64 = 1
	)

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to productory less than 2 operands")
	}

	for _, num := range o.Values {

		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		if flag {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			intProd *= number
		} else {
			s := num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			floatProd *= number
		}
	}

	if flag {
		prod = intProd
	} else {
		prod = floatProd
	}

	return
}

// DecomposePrimes
func (o *Operation) DecomposePrimes() (deco [][]int64, err error) {
	deco = make([][]int64, len(o.Values))
	for i, num := range o.Values {
		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		number, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		primes := calculatePrimes()
		j := 0
		v := primes[0]

		for v <= number {
			for (number % v) == 0 {
				deco[i] = append(deco[i], v)
				number = number / v
			}
			j++
			v = primes[j]
		}
	}

	return
}


// PointDistance -
func (o *Operation) PointDistance() (c *float64, err error) {
	if len(o.Values) != 4 {
		return nil, errors.New("It's necessary specify at lest and only 4 operands")
	}

	s1, ok := o.Values[0].(string)
	if !ok {
		return nil, errors.New("Type value error: not a string")
	}

	x1, err := strconv.ParseFloat(s1, 64)
	if err != nil {
		log.Fatal(err)
	}

	s2, ok := o.Values[1].(string)
	if !ok {
		return nil, errors.New("Type value error: not a string")
	}

	y1, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		log.Fatal(err)
	}

	s3, ok := o.Values[2].(string)
	if !ok {
		return nil, errors.New("Type value error: not a string")
	}

	x2, err := strconv.ParseFloat(s3, 64)
	if err != nil {
		log.Fatal(err)
	}

	s4, ok := o.Values[3].(string)
	if !ok {
		return nil, errors.New("Type value error: not a string")
	}

	y2, err := strconv.ParseFloat(s4, 64)
	if err != nil {
		log.Fatal(err)
	}

	c2 := ((x1 - x2) * (x1 - x2)) + ((y1 - y2) * (y1 - y2))

	sqrt := c2 / 2
	temp := float64(0)

	for sqrt != temp {
		temp = sqrt
		sqrt = (c2/temp + temp) / 2
	}

	c = &sqrt

	return
}
