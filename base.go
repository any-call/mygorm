package mygorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 共享锁，其它事务 只能读取记录，不能修改记录
func ShareLock(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.Locking{
		Strength: "share",
	})
}

// 独占锁：其它事务不能读取记录
func ExclusiveLock(tx *gorm.DB) *gorm.DB {
	return tx.Clauses(clause.Locking{
		Strength: "update",
	})
}
