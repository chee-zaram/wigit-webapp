package handlers

import (
	"errors"

	"github.com/wigit-gh/webapp/internal/db"
)

// DBConnector servers as a global link to the database.
var DBConnector *db.DB

var (
	// ErrEmailNotProvided indicates email not provided during signin.
	ErrEmailNotProvided = errors.New("Email required")
	// ErrInvalidUser indicates the user does not exist during signin.
	ErrInvalidUser = errors.New("Invalid Email")
	// ErrInternalServer indicates some server side error occured that can't be handled.
	ErrInternalServer = errors.New("Something went wrong!")
	// ErrFailedToAddUserToDB indicates that an error occured when adding user to database.
	ErrFailedToAddUserToDB = errors.New("Failed to add user to database")
	// ErrDuplicateUser indicates that User with email already exists during signup.
	ErrDuplicateUser = errors.New("User with email already exists")
	// ErrInvalidFirstName indicates no First Name was provided for user on sign up.
	ErrInvalidFirstName = errors.New("Valid First Name required")
	// ErrInvalidLastName indicates no Last Name was provided for user on sign up.
	ErrInvalidLastName = errors.New("Valid Last Name required")
	// ErrInvalidAddress indicates no Address was provided for user on sign up.
	ErrInvalidAddress = errors.New("Valid Address required")
	// ErrNoPhone indicates no Phone number was provided for user on sign up.
	ErrNoPhone = errors.New("Valid Phone required")
	// ErrInvalidPhone indicates wrong number of digits were passed for user on sign up.
	ErrInvalidPhone = errors.New("Invalid number of digits in Phone field")
	// ErrInvalidPass indicates the user did not provide a Password on sign up or a valid password on signin.
	ErrInvalidPass = errors.New("Valid Password required")
	// ErrPassMismatch indicates the user entered did not repeat the password correctly.
	ErrPassMismatch = errors.New("Password Mismatch")
	// ErrPassTooShort indicates the user entered a password less than 8 characters long on sign up.
	ErrPassTooShort = errors.New("Password must be at least 8 characters long")
	// ErrPassTooLong indicates the user entered a password longer than 45 characters during sign up.
	ErrPassTooLong = errors.New("Password must not exceed 45 characters")
	// ErrInvalidProductID indicates that the id provided is invalid.
	ErrInvalidProductID = errors.New("Invalid Product ID")
	// ErrInvalidCategory indicates that the category of products provided is invalid.
	ErrInvalidCategory = errors.New("Not a valid category")
)
