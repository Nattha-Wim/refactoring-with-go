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
	return r.movie.Charge(r.daysRented)
}
