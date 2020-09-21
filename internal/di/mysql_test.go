package di

import (
	"github.com/stretchr/testify/assert"
	"goweb/internal/conf"
	"goweb/pkg/mysql"
	"testing"
)

type MySQLInfo struct {
	Version string
}

func TestMySQL(t *testing.T) {
	if conf.Conf.MySQL == nil {
		return
	}

	db, err := mysql.GetDefaultDb()
	assert.Nil(t, err)

	var info MySQLInfo
	err = db.Raw("select version() as `version`").Scan(&info).Error
	assert.Nil(t, err)
	t.Log(info.Version)
}
