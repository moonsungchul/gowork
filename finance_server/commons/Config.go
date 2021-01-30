package commons

type Config struct {
	MySQL_Host   string `json:mysql_host`
	MySQL_Port   int16  `json:mysql_port`
	MySQL_User   string `json:mysql_user`
	MySQL_Passwd string `json:mysql_passwd`
	MySQL_Dbname string `json:mysql_dbname`
}
