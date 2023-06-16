package db

// GetSchemas returns all the models for which schemas are to be created.
func GetSchemas() (*User, *Order, *Booking, *Slot, *Item, *Product, *Service) {
	return &User{}, &Order{}, &Booking{}, &Slot{}, &Item{}, &Product{}, &Service{}
}
