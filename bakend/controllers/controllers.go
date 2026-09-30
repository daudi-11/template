package controllers

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"test/models"
	"test/utilis"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Ac(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hurrey u have made it!",
	})
}

func Testb(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	if name == "" {
		name = c.Get("name")
	}
	if email == "" {
		email = c.Get("email")
	}

	return c.JSON(fiber.Map{
		"message": "Hello " + name + ", your email is " + email,
	})
}

func Testa(c *fiber.Ctx) error {
	name := c.FormValue("name")
	if name == "" {
		name = c.Get("name")
	}

	return c.SendString("Hello " + name)
}

/*
	func reg_user(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		if name == "" {
			name = c.Get("name")
		}
		if email == "" {
			email = c.Get("email")
		}

		return c.JSON(fiber.Map{
			"message": "Hello " + name + ", your email is " + email,
		})
	}
*/
func Register(db *gorm.DB) fiber.Handler {

	// This function returns the actual Fiber request handler.
	// The db connection is available inside the handler,
	// so we can use GORM when we need to save the user.
	return func(c *fiber.Ctx) error {

		// TEMPORARY STRUCT
		// This will hold the registration information
		// that we receive FROM THE USER.
		var request struct {
			Name     string `json:"name" form:"name"`
			Email    string `json:"email" form:"email"`
			Password string `json:"password" form:"password"`
		}

		// RECEIVE DATA FROM THE REQUEST BODY
		//
		// Fiber reads the body and puts the received
		// name, email, and password into request.
		//
		// Example:
		// request.Name
		// request.Email
		// request.Password
		parseErr := c.BodyParser(&request)
		if request.Name == "" {
			request.Name = c.Query("name")
		}
		if request.Email == "" {
			request.Email = c.Query("email")
		}
		if request.Password == "" {
			request.Password = c.Query("password")
		}
		if parseErr != nil && request.Name == "" && request.Email == "" && request.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Send name, email, and password in the request body or as query parameters",
			})
		}

		// CLEAN THE USER INPUT
		//
		// TrimSpace removes unnecessary spaces
		// from the beginning and end.
		request.Name = strings.TrimSpace(request.Name)
		request.Email = strings.TrimSpace(request.Email)

		// CHECK THAT REQUIRED INFORMATION WAS PROVIDED
		//
		// If any of these fields are empty,
		// stop the registration.
		if request.Name == "" ||
			request.Email == "" ||
			request.Password == "" {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Name, email, and password are required",
			})
		}

		// HASH THE PASSWORD
		//
		// We DO NOT want to save the user's real password
		// directly into PostgreSQL.
		//
		// Example:
		//
		// User enters:
		// "kingd1121"
		//
		// bcrypt produces:
		// "$2a$10$..........."
		//
		// The hash is what we will store.
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(request.Password),
			bcrypt.DefaultCost,
		)

		// CHECK IF BCRYPT FAILED
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to register user",
			})
		}

		// CREATE THE PARENT MODEL
		//
		// We now take the information we received
		// from the user and put it into our model.
		//
		// Notice that we store hashedPassword,
		// NOT the original password.
		parent := models.Parent{
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashedPassword),
		}

		// SAVE THE PARENT INTO POSTGRESQL
		//
		// db is our GORM database connection.
		//
		// GORM takes the parent model and creates
		// a new record in the corresponding table.
		if err := db.Create(&parent).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to register user",
			})
		}

		// SEND A RESPONSE BACK TO THE USER
		//
		// Registration was successful.
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User registered successfully",
			"user":    parent,
		})
	}
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Read JSON or form input; query values remain supported for the current Postman setup.
		var request struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		parseErr := c.BodyParser(&request)
		if request.Email == "" {
			request.Email = c.Query("email")
		}
		if request.Password == "" {
			request.Password = c.Query("password")
		}
		// Reject an unreadable body when no query values were supplied.
		if parseErr != nil && request.Email == "" && request.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Send email and password in the request body or as query parameters",
			})
		}

		request.Email = strings.TrimSpace(request.Email)
		if request.Email == "" || request.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email and password are required",
			})
		}

		var parent models.Parent
		// Find the account by email without returning the stored password hash.
		if err := db.Where("email = ?", request.Email).First(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Keep this response identical to a wrong-password response.
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Invalid email or password",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Login failed",
			})
		}

		// Compare the submitted password with the bcrypt hash saved at registration.
		if err := bcrypt.CompareHashAndPassword([]byte(parent.Password), []byte(request.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid email or password",
			})
		}

		// Sign a JWT for the authenticated account using its email as the subject.
		token, err := utilis.GenerateJWT(parent.Email)
		error := err
		fmt.Print(error)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to create token",
			})
		}
		c.Cookie(&fiber.Cookie{
			Name:     "access_token",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			SameSite: "Lax",
		})

		// Return only public account fields; the model's password is never included.
		return c.JSON(fiber.Map{
			"message": "Login successful",
			"token":   token,
		})
	}
}

func CreateChild(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		parent, err := authenticatedParent(c, db)
		if errors.Is(err, fiber.ErrUnauthorized) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "A valid session is required"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load account"})
		}

		var request struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Send valid JSON with name, email, and password"})
		}
		request.Name = strings.TrimSpace(request.Name)
		request.Email = strings.TrimSpace(request.Email)
		if request.Name == "" || request.Email == "" || request.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Name, email, and password are required"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not secure child password"})
		}

		child := models.Child{
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashedPassword),
			ParentID: parent.ID,
		}
		err = db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&child).Error; err != nil {
				return err
			}
			wallet := models.Wallet{ChildID: child.ID}
			return tx.Create(&wallet).Error
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create child"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Child created successfully",
			"child": fiber.Map{
				"id":        child.ID,
				"name":      child.Name,
				"email":     child.Email,
				"parent_id": child.ParentID,
			},
			"balance": 0,
		})
	}
}

func ListChildren(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		parent, err := authenticatedParent(c, db)
		if errors.Is(err, fiber.ErrUnauthorized) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "A valid session is required"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load account"})
		}

		children := make([]models.Child, 0)
		if err := db.Select("id", "name", "email", "parent_id").Where("parent_id = ?", parent.ID).Order("id DESC").Find(&children).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load children"})
		}

		childIDs := make([]uint, len(children))
		for index, child := range children {
			childIDs[index] = child.ID
		}

		balances := make(map[uint]float64, len(children))
		if len(childIDs) > 0 {
			var wallets []models.Wallet
			if err := db.Select("child_id", "balance").Where("child_id IN ?", childIDs).Find(&wallets).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load wallet balances"})
			}
			for _, wallet := range wallets {
				balances[wallet.ChildID] = wallet.Balance
			}
		}

		result := make([]fiber.Map, 0, len(children))
		for _, child := range children {
			result = append(result, fiber.Map{
				"id":        child.ID,
				"name":      child.Name,
				"email":     child.Email,
				"parent_id": child.ParentID,
				"balance":   balances[child.ID],
			})
		}

		return c.JSON(fiber.Map{"children": result})
	}
}

/*func zombie(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	if name == "" {
		name = c.Get("name")
	}
	if email == "" {
		email = c.Get("email")
	}

	return c.JSON(fiber.Map{
		"message": "Hello " + name + ", your email is " + email,
	})
}

receive from user

func Register(c *fiber.Ctx) error {
    name := c.FormValue("name")
    email := c.FormValue("email")
    password := c.FormValue("password")

    return c.JSON(fiber.Map{
        "name":     name,
        "email":    email,
        "password": password,
    })
}


*/

var errInsufficientBalance = errors.New("insufficient balance")

func authenticatedParent(c *fiber.Ctx, db *gorm.DB) (models.Parent, error) {
	email, valid := utilis.SessionEmail(c)
	if !valid {
		return models.Parent{}, fiber.ErrUnauthorized
	}

	var parent models.Parent
	if err := db.Select("id", "email").Where("email = ?", email).First(&parent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Parent{}, fiber.ErrUnauthorized
		}
		return models.Parent{}, err
	}
	return parent, nil
}

type transactionRequest struct {
	ChildID uint    `json:"child_id"`
	Amount  float64 `json:"amount"`
}

func transactionHandler(db *gorm.DB, transactionType string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		parent, err := authenticatedParent(c, db)
		if errors.Is(err, fiber.ErrUnauthorized) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "A valid session is required"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load account"})
		}

		var request transactionRequest
		if err := c.BodyParser(&request); err != nil || request.ChildID == 0 || request.Amount <= 0 || math.IsInf(request.Amount, 0) || math.IsNaN(request.Amount) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "child_id and a positive amount are required"})
		}

		var updatedBalance float64
		var savedTransaction models.Transaction
		err = db.Transaction(func(tx *gorm.DB) error {
			var child models.Child
			if err := tx.Where("id = ? AND parent_id = ?", request.ChildID, parent.ID).First(&child).Error; err != nil {
				return err
			}

			wallet := models.Wallet{ChildID: child.ID}
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "child_id"}},
				DoNothing: true,
			}).Create(&wallet).Error; err != nil {
				return err
			}
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("child_id = ?", child.ID).First(&wallet).Error; err != nil {
				return err
			}

			updatedBalance = wallet.Balance
			if transactionType == "debit" {
				if request.Amount > wallet.Balance {
					return errInsufficientBalance
				}
				updatedBalance -= request.Amount
			} else {
				updatedBalance += request.Amount
			}

			if err := tx.Model(&wallet).Update("balance", updatedBalance).Error; err != nil {
				return err
			}
			savedTransaction = models.Transaction{
				WalletID: wallet.ID,
				ParentID: parent.ID,
				Amount:   request.Amount,
				Type:     transactionType,
			}
			return tx.Create(&savedTransaction).Error
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Child not found for this account"})
		}
		if errors.Is(err, errInsufficientBalance) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Insufficient balance"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Transaction failed"})
		}

		return c.JSON(fiber.Map{
			"message":     transactionType + " successful",
			"child_id":    request.ChildID,
			"amount":      request.Amount,
			"balance":     updatedBalance,
			"transaction": savedTransaction,
		})
	}
}

func TransactionCredit(db *gorm.DB) fiber.Handler {
	return transactionHandler(db, "credit")
}

func TransactionDebit(db *gorm.DB) fiber.Handler {
	return transactionHandler(db, "debit")
}

func WalletBalance(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		parent, err := authenticatedParent(c, db)
		if errors.Is(err, fiber.ErrUnauthorized) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "A valid session is required"})
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load account"})
		}

		childID := c.QueryInt("child_id")
		if childID <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "A valid child_id is required"})
		}

		var child models.Child
		if err := db.Select("id").Where("id = ? AND parent_id = ?", childID, parent.ID).First(&child).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Child not found for this account"})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load balance"})
		}

		var wallet models.Wallet
		if err := db.Where("child_id = ?", child.ID).First(&wallet).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(fiber.Map{"child_id": child.ID, "balance": 0})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not load balance"})
		}
		return c.JSON(fiber.Map{"child_id": child.ID, "balance": wallet.Balance})
	}
}

func Logout(c *fiber.Ctx) error {
	utilis.LogoutSession(c)
	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}
