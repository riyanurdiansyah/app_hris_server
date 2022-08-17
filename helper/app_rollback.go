package helper

import "gorm.io/gorm"

func CommitOrRollback(db *gorm.DB) {
	err := recover()
	if err != nil {
		errorRollback := db.Rollback()
		PanicIfError(errorRollback.Error)
	} else {
		errorCommit := db.Commit()
		PanicIfError(errorCommit.Error)
	}
}
