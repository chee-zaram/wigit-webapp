package db

const (
	// Pending is the pending `status`. Used for both orders and bookings.
	Pending = "pending"
	// Paid is the status `paid`. Use for both orders and bookings.
	Paid = "paid"
	// Fulfilled is the status `fulfilled`. Used for bookings.
	Fulfilled = "fulfilled"
	// Cancelled is the status `cancelled`. Used for bookings.
	Cancelled = "cancelled"
	// Shipped is the status `shipped`. Used for orders.
	Shipped = "shipped"
	// Delivered is the status `delivered`. Used for orders.
	Delivered = "delivered"
)
