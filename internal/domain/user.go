package domain

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"goweb/internal/model/dbmodel"
	"goweb/pkg/mysql"
)

type userDomain struct {
}

var UserDomain = new(userDomain)

func (d *userDomain) Insert(_ context.Context, user *dbmodel.User) (int64, error) {
	db, err := mysql.GetDefaultDb()
	if err != nil {
		return 0, err
	}

	db = db.Create(user)
	if db.Error != nil {
		return 0, db.Error
	}

	return user.UserID, nil
}

func (d *userDomain) SelectOneByEmail(_ context.Context, email string) (*dbmodel.User, error) {
	db, err := mysql.GetDefaultDb()
	if err != nil {
		return nil, err
	}

	user := new(dbmodel.User)
	db = db.Where(fmt.Sprintf("%s = ?", dbmodel.TCUserEmail), email).First(user)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return nil, nil
		}
		return nil, db.Error
	}

	return user, nil
}

func (d *userDomain) SelectOneByUserID(_ context.Context, userID int64) (*dbmodel.User, error) {
	db, err := mysql.GetDefaultDb()
	if err != nil {
		return nil, err
	}

	user := new(dbmodel.User)
	db = db.Where(fmt.Sprintf("%s = ?", dbmodel.TCUserID), userID).First(user)
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return nil, nil
		}
		return nil, db.Error
	}

	return user, nil
}
