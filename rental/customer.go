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

func Charge(each Rental) float64 {
	result := 0.0
	switch each.Movie().PriceCode() {
	case REGULAR:
		result += 2
		if each.DaysRented() > 2 {
			result += float64(each.DaysRented()-2) * 1.5
		}
	case NEW_RELEASE:
		result += float64(each.DaysRented()) * 3.0
	case CHILDRENS:
		result += 1.5
		if each.DaysRented() > 3 {
			result += float64(each.DaysRented()-3) * 1.5
		}
	}
	return result
}

func (rcvr Customer) Statement() string {
	totalAmount := 0.0
	frequentRenterPoints := 0
	result := fmt.Sprintf("%v%v%v", "Rental Record for ", rcvr.Name(), "\n")
	for _, r := range rcvr.rentals {
		thisAmount := Charge(r)

		frequentRenterPoints++
		if r.Movie().PriceCode() == NEW_RELEASE && r.DaysRented() > 1 {
			frequentRenterPoints++
		}
		result += fmt.Sprintf("\t%v\t%.1f\n", r.Movie().Title(), thisAmount)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %.1f\n", totalAmount)
	result += fmt.Sprintf("You earned %v frequent renter points", frequentRenterPoints)
	return result
}
