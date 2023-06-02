package db

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

// SaveToDB saves the current user to the database.
func (user *User) SaveToDB() error {
	if user == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadByEmail loads the latest version of user with given `email` from the database.
func (user *User) LoadByEmail(email string) (int, error) {
	if user == nil {
		return 0, ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("email = ?", email).Preload("Orders.Items").Preload("Bookings.Slot").First(user).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusBadRequest, err
	} else if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// Reload saves and loads up the current user from the database.
func (user *User) Reload() error {
	if user == nil {
		return ErrNilPointer
	}

	if err := user.SaveToDB(); err != nil {
		return err
	}

	if _, err := user.LoadByEmail(*user.Email); err != nil {
		return err
	}

	return nil
}

// LoadByID loads the user with given `id` into the current object.
func (user *User) LoadByID(id string) error {
	if user == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Preload("Orders").Preload("Bookings").First(user, "id = ?", id).Error
	}); err != nil {
		return err
	}

	return nil
}

// UpdateResetToken updates the `reset_token` column for the current user.
func (user *User) UpdateResetToken(token string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Model(user).Update("reset_token", token).Error
	}); err != nil {
		return err
	}

	if err := user.LoadByID(*user.ID); err != nil {
		return err
	}

	return nil
}

// UpdateInfo updates the user with the given info and returns the updated user.
func (user *User) UpdateInfo(email, addr, phone, firstName, lastName string) error {
	if user == nil {
		return ErrNilPointer
	}

	user.Email = &email
	user.Address = &addr
	user.Phone = &phone
	user.FirstName = &firstName
	user.LastName = &lastName

	if err := user.Reload(); err != nil {
		return err
	}

	return nil
}

// UpdateRole updates the role of the current user to `newRole`.
func (user *User) UpdateRole(newRole string) error {
	user.Role = &newRole
	if err := user.Reload(); err != nil {
		return err
	}

	return nil
}

// Admins gets slice of all admins in the database.
func Admins() ([]User, error) {
	var admins []User

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("first_name asc").
			Where("role = 'admin'").
			Preload("Orders").
			Preload("Bookings").
			Find(&admins).Error
	}); err != nil {
		return nil, err
	}

	return admins, nil
}

// Customers gets slice of all customers in the database.
func Customers() ([]User, error) {
	var customers []User

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("first_name asc").
			Where("role = 'customer'").
			Preload("Orders").
			Preload("Bookings").
			Find(&customers).Error
	}); err != nil {
		return nil, err
	}

	return customers, nil
}

// DeleteUser deletes the user with email from the database.
func DeleteUser(email string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM users WHERE email = ?`, email).Error
	}); err != nil {
		return err
	}

	return nil
}
