package mysqlerr

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func CheckMySQLError(err error) error {
	mySqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return err
	}
	fmt.Println(mySqlErr.Number)
	switch mySqlErr.Number {
	case 1062:
		return DuplicateKeyError{}
	}
	return err
}
