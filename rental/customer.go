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

func (r Rental) Charge() float64 { // Charge belong to Rental so we can convert Charge to Rental's methor
	result := 0.0
	switch r.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if r.DaysRented() > 2 {
			result += float64(r.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(r.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if r.DaysRented() > 3 {
			result += float64(r.DaysRented()-3) * 1.5
		}
	}
	return result
}

func GetPoint(r Rental) int {
	frequentRenterPoints := 0
	frequentRenterPoints++
	if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
		frequentRenterPoints++
	}
	return frequentRenterPoints
}

func (rcvr Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("Rental Record for %v\n", rcvr.Name())
	for _, r := range rcvr.rentals {
		thisAmount := r.Charge()

		frequentRenterPoints += GetPoint(r)

		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), thisAmount)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
