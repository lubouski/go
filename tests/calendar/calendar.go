package main

import (
	"fmt"
	"time"
)

type calendar struct {
	date	time.Time

}

func NewCalendar(Date time.Time) *calendar {
	return &calendar{
	date: Date,
	}
}

func (c *calendar) CurrentQuarter() int {
	_, month, _ := c.date.Date()
	if month < 4 {
		return 1
	} else if month < 7 {
		return 2
	} else if month < 10 {
		return 3
	} else {
		return 4
	}
}

/*func NewCalendar(c calendar) time.Time {
	
}
*/

func main() {
	parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", "11"))
/*	fmt.Printf("Parsed type is -> %T \n", parsed)
	year, month, day := parsed.Date()
	fmt.Println(year, int(month), day)
	quater := int(month)/3 + 1
	fmt.Println(quater)
*/
	my_calendar := NewCalendar(parsed)
	fmt.Println(my_calendar.CurrentQuarter())
}
