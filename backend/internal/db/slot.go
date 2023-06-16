package db

import "gorm.io/gorm"

// SaveToDB saves the current slot to the database.
func (slot *Slot) SaveToDB() error {
	if slot == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Save(slot).Error
	}); err != nil {
		return err
	}

	return nil
}

// LoadFromDB loads the current slot object with information from the database slot
// with given `id`.
func (slot *Slot) LoadFromDB(id string) error {
	if slot == nil {
		return ErrNilPointer
	}

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.First(slot, "id = ?", id).Error
	}); err != nil {
		return err
	}

	return nil
}

// Reload saves the current slot object to the database and loads up the updated version.
func (slot *Slot) Reload() error {
	if slot == nil {
		return ErrNilPointer
	}

	if err := slot.SaveToDB(); err != nil {
		return err
	}

	if err := slot.LoadFromDB(*slot.ID); err != nil {
		return err
	}

	return nil
}

// AllSlots retrieves a list of all available slots from the database.
func AllSlots() ([]Slot, error) {
	var slots []Slot

	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("is_free = ?", true).Find(&slots).Error
	}); err != nil {
		return nil, err
	}

	return slots, nil
}

// DeleteSlot deletes a slot with id from database.
func DeleteSlot(id string) error {
	if err := Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM slots WHERE id = ?`, id).Error
	}); err != nil {
		return err
	}

	return nil
}
