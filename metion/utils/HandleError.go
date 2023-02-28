package utils

import (
	"bizd/metion/model"
	"gorm.io/gorm"
)

func Try(tryFunc func(), catchFunc func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catchFunc(err)
		}
	}()
	tryFunc()
}

func JudgeCommit(tx *gorm.DB) {
	if err := tx.Commit().Error; err != nil {
		panic(&model.MyError{Code: 400, Message: err.Error()})
	}
}
