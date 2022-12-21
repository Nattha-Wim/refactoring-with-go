package rental

type Rental struct {
	movie      Movie
	daysRented int
}

func NewRental(movie Movie, daysRented int) (rcvr Rental) {
	return Rental{
		movie:      movie,
		daysRented: daysRented,
	}

}
func (r Rental) DaysRented() int {
	return r.daysRented
}
func (r Rental) Movie() Movie {
	return r.movie
}
func (r Rental) Charge() float64 { // Charge belong to Rental so we can convert Charge to Rental's methor

	return r.Movie().Price.Charge(r.daysRented)
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
}
