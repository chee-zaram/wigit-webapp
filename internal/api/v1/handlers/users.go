package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wigit-gh/webapp/internal/db"
	"gorm.io/gorm"
)

// CustomerDeleteUser deletes the current user from the database.
func CustomerDeleteUser(ctx *gin.Context) {
	user, _, err := validateUserParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM users WHERE id = ?`, *user.ID).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully",
	})
}

// CustomerPutUser updates a user's information in the database.
func CustomerPutUser(ctx *gin.Context) {
	user, newUser, err := validateUserParams(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUserInfo(user, newUser)
	if err := user.SaveToDB(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := getUserByID(*user.ID)
	if err != nil {
		updatedUser = user
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User updated successfully",
		"data": updatedUser,
	})
}

// validateUserParams validates data sent to the `users` endpoint.
// It is used during updating and deletion of a user or information.
func validateUserParams(ctx *gin.Context) (*db.User, *db.User, error) {
	_user, exists := ctx.Get("user")
	if !exists {
		return nil, nil, errors.New("User not set in context")
	}
	user := _user.(*db.User)
	id := ctx.Param("user_id")
	if id == "" {
		return nil, nil, errors.New("User ID must be set")
	}

	if id != *user.ID {
		return nil, nil, errors.New("Cannot perform operation on another user's account")
	}

	newUser := new(db.User)
	if err := ctx.ShouldBind(newUser); err != nil {
		return nil, nil, err
	}

	return user, newUser, nil
}

// updateUserInfo updates a user's information on put request.
func updateUserInfo(user, newUser *db.User) {
	// TODO: validate lengths of strings
	user.Email = newUser.Email

	if newUser.FirstName != nil && *newUser.FirstName != "" {
		user.FirstName = newUser.FirstName
	}

	if newUser.LastName != nil && *newUser.LastName != "" {
		user.LastName = newUser.LastName
	}

	if newUser.Address != nil && *newUser.Address != "" {
		user.Address = newUser.Address
	}

	if newUser.Phone != nil && *newUser.Phone != "" {
		user.Phone = newUser.Phone
	}
}

// AdminGetUserOrdersBookings gets all orders and booking belonging to a user
// with given email.
func AdminGetUserOrdersBookings(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrEmailParamNotSet.Error()})
		return
	}

	user, code, err := getUserFromDB(email)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"orders":   user.Orders,
			"bookings": user.Bookings,
		},
	})
}

// getUserFromDB gets the user with `email` from the database.
func getUserFromDB(email string) (*db.User, int, error) {
	dbUser := new(db.User)

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Where("email = ?", email).Preload("Orders.Items").Preload("Bookings.Slot").First(dbUser).Error
	}); err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, http.StatusBadGateway, ErrInvalidUser
	} else if err != nil {
		return nil, http.StatusInternalServerError, ErrInternalServer
	}

	return dbUser, http.StatusOK, nil
}

// SuperAdminUpdateRole updates the role of a user using super admin privileges.
func SuperAdminUpdateRole(ctx *gin.Context) {
	email := ctx.Param("email")
	role := ctx.Param("new_role")
	if email == "" || role == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email or role param not set"})
		return
	}

	if role != "admin" && role != "customer" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid user role"})
		return
	}

	user, code, err := getUserFromDB(email)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := updateUserRole(user, role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User role updated successfully",
		"user": dbUser,
	})
}

// updateUserRole updates a given user's role and returns the updated user.
func updateUserRole(user *db.User, role string) (*db.User, error) {
	user.Role = &role
	if err := user.SaveToDB(); err != nil {
		return nil, ErrInternalServer
	}

	dbUser := new(db.User)
	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.First(dbUser, "id = ?", *user.ID).Error
	}); err != nil {
		return nil, ErrInternalServer
	}

	return dbUser, nil
}

// SuperAdminDeleteUser deletes a user account.
func SuperAdminDeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrEmailParamNotSet.Error()})
		return
	}

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Exec(`DELETE FROM users WHERE email = ?`, email).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully",
	})
}

// SuperAdminGetAdmins gets all admins in the database.
func SuperAdminGetAdmins(ctx *gin.Context) {
	var admins []db.User

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("first_name asc").Where("role = 'admin'").Preload("Orders").Preload("Bookings").Find(&admins).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": admins,
	})
}

// SuperAdminGetCustomers retrieves all customers.
func SuperAdminGetCustomers(ctx *gin.Context) {
	var customers []db.User

	if err := db.Connector.Query(func(tx *gorm.DB) error {
		return tx.Order("first_name asc").Where("role = 'customer'").Preload("Orders").Preload("Bookings").Find(&customers).Error
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}

// SuperAdminGetUser gets a user with given email from database.
func SuperAdminGetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": ErrEmailParamNotSet.Error()})
		return
	}

	user, code, err := getUserFromDB(email)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
