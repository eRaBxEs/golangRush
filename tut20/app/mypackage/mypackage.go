package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name = "Ehioje"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

type Date struct { //Lower Case; All attribute of the struct are not accessible outside this package
	day   int
	month int
	year  int
}

// Setter methods
func (d *Date) SetDay(day int) error { // Upper Case; All these Set methods of Date are accessible outside the package
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 12) {
		return errors.New("incorrect month value")
	}
	d.month = m
	return nil
}

func (d *Date) SetYear(yr int) error {
	if (yr < 1875) || (yr > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = yr
	return nil
}

// Getters methods
func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}
