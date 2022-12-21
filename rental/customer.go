package rental

import "fmt"

type Customer struct {
	name    string
	rentals []Rental
}

func NewCustomer(name string) (rcvr Customer) {
	return Customer{
		name:    name,
		rentals: []Rental{},
	}
}
func (rcvr Customer) AddRental(arg Rental) {
	rcvr.rentals = append(rcvr.rentals, arg)
}
func (rcvr *Customer) Name() string {
	return rcvr.name
}

func RegularCharge(r Rental) float64 {
	result := 0.0
	result += 2
	if r.DaysRented() > 2 {
		result += float64(r.DaysRented()-2) * 1.5
	}
	return result
}

func NewReleaseCharge(r Rental) float64 {
	result := 0.0
	result += float64(r.DaysRented()) * 3.0
	return result
}

func ChildrensCharge(r Rental) float64 {
	result := 0.0
	result += 1.5
	if r.DaysRented() > 3 {
		result += float64(r.DaysRented()-3) * 1.5
	}
	return result
}

func (r Rental) Charge() float64 { // Charge belong to Rental so we can convert Charge to Rental's methor
	result := 0.0
	switch r.Movie().PriceCode() {
	case REGULAR:
		return RegularCharge(r)
	case NEW_RELEASE:
		return NewReleaseCharge(r)
	case CHILDRENS:
		return ChildrensCharge(r)
	}
	return result
}

func GetPoint(r Rental) int {
	//frequentRenterPoints
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		return 2
	}
	return 1
}

func GetTotalAmount(rentals []Rental) float64 {
	result := 0.0
	for _, r := range rentals {
		result += r.Charge()
	}
	return result
}

func GetTotalPoint(rentals []Rental) int {
	// frequentRenterPoints := 0
	points := 0
	for _, r := range rentals {
		points += GetPoint(r)
	}
	return points
}

func (c Customer) Statement() string {
	points := GetTotalPoint(c.rentals)

	totalAmount := GetTotalAmount(c.rentals)

	result := fmt.Sprintf("Rental Record for %v\n", c.Name())

	for _, r := range c.rentals {
		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), r.Charge())
	}

	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", points)
	return result
}
