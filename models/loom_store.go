package models

import "time"

type Height struct {
	Name            string `gorm:"primary_key"`
	LastBlockHeight uint64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type NewValueSet struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BlockHeight uint64
	BlockTime   uint64
	TxIdx       uint
	TxHash      string `gorm:"not null; unique_index"`
	Name        string
	Value       string
	Contract    string
}
