package model

type Ping struct {
	ID    string `gorm:"type:char(10);not null;primaryKey"`
	Value string `gorm:"type:char(10);not null"`
}
