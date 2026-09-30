package routes

import (
	"test/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(r *fiber.App, db *gorm.DB) {
	r.Get("/", controllers.Ac)
	r.Post("/test", controllers.Testa)
	r.Post("/testb", controllers.Testb)
	r.Post("/register", controllers.Register(db))
	r.Post("/login", controllers.Login(db))
	r.Get("/children", controllers.ListChildren(db))
	r.Post("/children", controllers.CreateChild(db))
	r.Post("/credit", controllers.TransactionCredit(db))
	r.Post("/debit", controllers.TransactionDebit(db))
	r.Get("/balance", controllers.WalletBalance(db))
	r.Post("/logout", controllers.Logout)
}
