package mygorm

import (
	"strings"

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

func InSetWithOR(field string, values []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(values) == 0 || strings.TrimSpace(field) == "" {
			return db
		}

		seen := make(map[string]struct{}, len(values))
		normalized := make([]string, 0, len(values))
		for _, raw := range values {
			v := strings.TrimSpace(raw)
			if v == "" {
				continue
			}
			if _, ok := seen[v]; ok {
				continue
			}
			seen[v] = struct{}{}
			normalized = append(normalized, v)
		}

		if len(normalized) == 0 {
			return db
		}

		conds := make([]string, len(normalized))
		args := make([]any, len(normalized))
		for i, v := range normalized {
			conds[i] = "FIND_IN_SET(?, " + field + ")"
			args[i] = v
		}

		// 使用括号包裹 OR 子句，避免与外层 AND 条件组合时发生优先级歧义。
		return db.Where("("+strings.Join(conds, " OR ")+")", args...)
	}
}
