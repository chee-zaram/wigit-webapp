package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wigit-ng/webapp/backend/internal/db"
)

// UpdateUser binds to the request body sent to update a user's information.
type UpdateUser struct {
	Email     *string `json:"email" binding:"required,email,max=45"`
	Phone     *string `json:"phone" binding:"required,min=8,max=11"`
	Address   *string `json:"address" binding:"required,max=255"`
	FirstName *string `json:"first_name" binding:"required,max=45"`
	LastName  *string `json:"last_name" binding:"required,max=45"`
}

// cleanUp removes whitespaces from the request body.
func (user *UpdateUser) cleanUp() {
	if user == nil {
		return
	}

	*user.Phone = strings.TrimSpace(*user.Phone)
	*user.Address = strings.TrimSpace(*user.Address)
	*user.FirstName = strings.TrimSpace(*user.FirstName)
	*user.LastName = strings.TrimSpace(*user.LastName)
}

// CustomerDeleteUser Deletes the current user from the database.
//
//	@Summary	Allows the current user delete their account.
//	@Tags		users
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		email			path		string					true	"User's email"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/users/{email} [delete]
func CustomerDeleteUser(ctx *gin.Context) {
	user, err := validateDeleteUserRequest(ctx)
	if err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	if err := db.DeleteUser(*user.Email); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully",
	})
}

// validateDeleteUserRequest validates the request sent to delete a user.
func validateDeleteUserRequest(ctx *gin.Context) (*db.User, error) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		return nil, ErrUserCtx
	}

	email := ctx.Param("email")
	if email == "" {
		return nil, ErrEmailParamNotSet
	}

	if email != *user.Email {
		return nil, errors.New("Cannot perform operation on another user's account")
	}

	return user, nil
}

// CustomerPutUser Updates a user's information in the database.
//
//	@Summary	Allows the current user update their account information.
//	@Tags		users
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		user			body		UpdateUser				true	"User information"
//	@Param		email			path		string					true	"User email"
//	@Success	200				{object}	map[string]interface{}	"data,msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/users/{email} [put]
func CustomerPutUser(ctx *gin.Context) {
	user, newUser, err := validatePutUserRequest(ctx)
	if err != nil {
		AbortCtx(ctx, http.StatusBadRequest, err)
		return
	}

	newUser.cleanUp()
	if err := user.UpdateInfo(
		*newUser.Email,
		*newUser.Address,
		*newUser.Phone,
		*newUser.FirstName,
		*newUser.LastName); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User updated successfully",
		"data": user,
	})
}

// validatePutUserRequest validates data sent to the `users` endpoint.
// It is used when updating a user's information.
func validatePutUserRequest(ctx *gin.Context) (*db.User, *UpdateUser, error) {
	_user, exists := ctx.Get("user")
	user, ok := _user.(*db.User)
	if !exists || !ok {
		return nil, nil, ErrUserCtx
	}

	email := ctx.Param("email")
	if email == "" {
		return nil, nil, ErrEmailParamNotSet
	}

	if email != *user.Email {
		return nil, nil, errors.New("Cannot perform operation on another user's account")
	}

	newUser := new(UpdateUser)
	if err := ctx.ShouldBind(newUser); err != nil {
		return nil, nil, err
	}

	return user, newUser, nil
}

// AdminGetUserOrdersBookings Get orders and bookings
//
//	@Summary	Allows admin get all orders and bookings belonging to user with email.
//	@Tags		admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		email			path		string					true	"User's email"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/admin/users/{email}/orders_bookings [get]
func AdminGetUserOrdersBookings(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrEmailParamNotSet)
		return
	}

	user := new(db.User)
	if code, err := user.LoadByEmail(email); err != nil {
		AbortCtx(ctx, code, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"orders":   user.Orders,
			"bookings": user.Bookings,
		},
	})
}

// SuperAdminUpdateRole Update a user role.
//
//	@Summary	Allows super_admin update another user's role
//	@Tags		super_admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		email			path		string					true	"User's email"
//	@Param		new_role		path		string					true	"User's new role"
//	@Success	200				{object}	map[string]interface{}	"msg,data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/super_admin/users/{email}/{new_role} [put]
func SuperAdminUpdateRole(ctx *gin.Context) {
	email := ctx.Param("email")
	role := ctx.Param("new_role")
	if email == "" || role == "" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("email or role param not set"))
		return
	}

	if role != "admin" && role != "customer" {
		AbortCtx(ctx, http.StatusBadRequest, errors.New("Invalid user role"))
		return
	}

	user := new(db.User)
	if code, err := user.LoadByEmail(email); err != nil {
		AbortCtx(ctx, code, err)
		return
	}

	if err := user.UpdateRole(role); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "User role updated successfully",
		"data": user,
	})
}

// SuperAdminDeleteUser Delete a user account.
//
//	@Summary	Allows super_admin delete another user's account.
//	@Tags		super_admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		email			path		string					true	"User's email"
//	@Success	200				{object}	map[string]interface{}	"msg"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/super_admin/users/{email} [delete]
func SuperAdminDeleteUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrEmailParamNotSet)
		return
	}

	if err := db.DeleteUser(email); err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "User deleted successfully",
	})
}

// SuperAdminGetAdmins Gets all admins in the database.
//
//	@Summary	Allows super_admin retrieve all admins in the database.
//	@Tags		super_admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/super_admin/users/admins [get]
func SuperAdminGetAdmins(ctx *gin.Context) {
	admins, err := db.Admins()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": admins,
	})
}

// SuperAdminGetCustomers Get all customers
//
//	@Summary	Allows super_admin retrieve all customers in the database.
//	@Tags		super_admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/super_admin/users/customers [get]
func SuperAdminGetCustomers(ctx *gin.Context) {
	customers, err := db.Customers()
	if err != nil {
		AbortCtx(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}

// SuperAdminGetUser Get a user with email
//
//	@Summary	Allows super_admin retrieve another user with given email
//	@Tags		super_admin
//	@Produce	json
//	@Param		Authorization	header		string					true	"Bearer <token>"
//	@Param		email			path		string					true	"User's email"
//	@Success	200				{object}	map[string]interface{}	"data"
//	@Failure	400				{object}	map[string]interface{}	"error"
//	@Failure	500				{object}	map[string]interface{}	"error"
//	@Router		/super_admin/users/{email} [get]
func SuperAdminGetUser(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		AbortCtx(ctx, http.StatusBadRequest, ErrEmailParamNotSet)
		return
	}

	user := new(db.User)
	if code, err := user.LoadByEmail(email); err != nil {
		AbortCtx(ctx, code, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
