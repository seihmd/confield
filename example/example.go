package main

import (
	"fmt"
	"os"

	"github.com/seihmd/confield"
	yaml "gopkg.in/yaml.v2"
)

type dbSetting struct {
	Password confield.F
	User     confield.F
	Name     confield.F
	Host     confield.F
}

func main() {
	yml := `
password: $CONFIELD_DBPASS|mypass
user: $CONFIELD_DBUSER|root
name: $CONFIELD_DBNAME|db_dev
host: $CONFIELD_DBHOST|localhost
`
	conf := dbSetting{}
	yaml.Unmarshal([]byte(yml), &conf)

	os.Setenv("CONFIELD_DBPASS", "dbpass")
	os.Setenv("CONFIELD_DBUSER", "dbuser")
	defer func() {
		os.Setenv("CONFIELD_DBPASS", "")
		os.Setenv("CONFIELD_DBUSER", "")
	}()

	fmt.Println(conf.Password.String()) // dbpass
	fmt.Println(conf.User.String())     // dbuser
	fmt.Println(conf.Name.String())     // db_dev
	fmt.Println(conf.Host.String())     // localhost
}
