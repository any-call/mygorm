package mygorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ShareLock(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.Locking{
		Strength: "share",
	})
}

func ExclusiveLock(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.Locking{
		Strength: "update",
	})
}
