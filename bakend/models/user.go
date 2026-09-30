package models

// 1. Parent Entity
type Parent struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	Name         string        `json:"name"`
	Email        string        `json:"email" gorm:"unique;not null"`
	Password     string        `json:"-"`
	Children     []Child       `json:"children,omitempty"`     // 1 Parent -> Many Children
	Transactions []Transaction `json:"transactions,omitempty"` // 1 Parent -> Many Transactions
}

// 2. Child Entity
type Child struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`
	ParentID uint   `json:"parent_id"` // Foreign Key linking to Parent
	Wallet   Wallet `json:"wallet"`    // 1 Child -> 1 Wallet
}

// 3. Wallet Entity
type Wallet struct {
	ID           uint          `json:"id" gorm:"primaryKey"`
	ChildID      uint          `json:"child_id" gorm:"unique"` // Foreign key (unique for 1-to-1)
	Balance      float64       `json:"balance" gorm:"default:0"`
	Transactions []Transaction `json:"transactions,omitempty"` // 1 Wallet -> Many Transactions
	Expenses     []Expense     `json:"expenses,omitempty"`     // Wallet -> Many Expenses
}

// 4. Transaction Entity
type Transaction struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	WalletID uint    `json:"wallet_id"` // Foreign Key linking to Wallet
	ParentID uint    `json:"parent_id"` // Foreign Key linking to Parent
	Amount   float64 `json:"amount"`
	Type     string  `json:"type"` // e.g., "deposit", "allowance"
}

// 5. Expense Entity
type Expense struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	WalletID    uint    `json:"wallet_id"` // Foreign Key linking to Wallet
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
