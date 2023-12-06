package dsn

import "fmt"

func GenDsn(driver,ip, username, password, dbname string, port int) string {
	return fmt.Sprintf("Driver={%s};Server=%s;Port=%d;Database=%s;Uid=%s;Pwd=%s",
					driver, ip, port, dbname, username, password)
}