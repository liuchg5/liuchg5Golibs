package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/liuchg5Golibs/log"
)

type Mysql struct {
	l      log.Logger
	host   string
	port   string
	user   string
	pw     string
	option string

	dbMap map[string]*sql.DB
}

func NewMysql(
	l log.Logger,
	host string,
	port string,
	user string,
	pw string,
	//dbName string,
	option string,
) *Mysql {
	return &Mysql{
		l:      l,
		host:   host,
		port:   port,
		user:   user,
		pw:     pw,
		option: option,
	}
}

//
func (p *Mysql) GetDB(dbName string) (db *sql.DB, err error) {
	var ok bool
	if db, ok = p.dbMap[dbName]; ok {
		return
	}
	db, err = p.connect(dbName)
	if err != nil {
		p.l.Error("err=", err)
		return
	}
	return
}

//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
func (p *Mysql) connect(dbName string) (db *sql.DB, err error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?%s`, p.user, p.pw, p.host, p.port, dbName, p.option)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		p.l.Error("err=", err)
		return
	}
	err = db.Ping()
	if err != nil {
		p.l.Error("err=", err)
		return
	}
	return
}
