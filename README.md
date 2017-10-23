# confield

**confield** is named string type of Go for helping to write self-explanatory configuration containing $ENVIRONMENT_VARIABLE.

### pros
You can write enviroment variable and default value in yaml/json text.

### cons
You can not write configuration value containing "|"; the charactor used as a separator.

## Example

```go
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

	fmt.Println(conf.Password.String()) // "dbpass" from env var
	fmt.Println(conf.User.String())     // "dbuser" from env var
	fmt.Println(conf.Name.String())     // "db_dev" as default value
	fmt.Println(conf.Host.String())     // "localhost" as default value
}
```
