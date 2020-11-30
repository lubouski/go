package main

import (
	"fmt"
	"log"
	"errors"
)

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid year")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 365 {
		return errors.New("invalid year")
	}
	d.day = day
	return nil
}

func main() {
	date := Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(115)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}


