package model

import "time"

type Height struct {
	Name            string `gorm:"primary_key"`
	LastBlockHeight uint64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type GeneratedCard struct {
	ID          uint64 `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BlockHeight uint64
	TxIdx       uint
	Owner       string
	CardID      string
	BoosterType uint8
	Contract    string
}
