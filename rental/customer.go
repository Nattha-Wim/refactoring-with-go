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
func (c Customer) AddRental(arg Rental) {
	c.rentals = append(c.rentals, arg)
}
func (c Customer) Name() string {
	return c.name
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

// func RegularCharge(daysRented int) float64 {
// 	result := 2.0
// 	if daysRented > 2 {
// 		result += float64(daysRented-2) * 1.5
// 	}
// 	return result
// }

// func NewReleaseCharge(daysRented int) float64 {
// 	result := 0.0
// 	result += float64(daysRented) * 3.0
// 	return result
// }

// func ChildrensCharge(daysRented int) float64 {

// 	result := 1.5
// 	if daysRented > 3 {
// 		result += float64(daysRented-3) * 1.5
// 	}
// 	return result
// }

// func (r Rental) Charge() float64 { // Charge belong to Rental so we can convert Charge to Rental's methor
// return r.Movie().Price.Charge(r.daysRented)
// switch r.Movie().PriceCode() {
// case REGULAR:
// 	return r.Movie().Charger.Charge(r.daysRented)
// case NEW_RELEASE:
// 	return r.Movie().Charger.Charge(r.daysRented)
// case CHILDRENS:
// 	return r.Movie().Charger.Charge(r.daysRented)
// case 0:
// 	return r.Movie().Charger.Charge(r.daysRented)
// }
// }
