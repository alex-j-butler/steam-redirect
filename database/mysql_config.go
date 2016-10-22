package database

import "fmt"

type MySQLConfig struct {
	Address   string
	Username  string
	Password  string
	Database  string
	Charset   string
	ParseTime bool
}

func (c MySQLConfig) getConnectString() string {
	var str string
	if c.ParseTime {
		str = "True"
	} else {
		str = "False"
	}

	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=%s&loc=Local",
		c.Username,
		c.Password,
		c.Address,
		c.Database,
		c.Charset,
		str,
	)
}

func (c MySQLConfig) getDialect() string {
	return "mysql"
}
