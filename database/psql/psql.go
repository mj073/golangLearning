package psql

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	logger "log"
	"strings"
)

var log = logger.Default()

type DBConfig struct {
	Address string
	Name string
	Username string
	Password string
}

type DB struct {
	db *sql.DB
}

type Table struct {
	Name string
	Fields []TableField
}

type TableField struct {
	Name string
	Attribute FieldAttribute
}

type FieldAttribute struct {
	Type string
	PrimaryKey bool
	ForeignKey bool
	References string
	NotNull bool
	Unique bool
	Default interface{}
	Check string

}

func NewDbConn(config DBConfig) (*DB, error) {
	return newDBConn(config, false)
}

func NewSecureDbConn(config DBConfig) (*DB, error) {
	return newDBConn(config, true)
}

func newDBConn(config DBConfig, secure bool) (db *DB, err error) {
	sslmode := "disable"
	if secure {
		sslmode = "enable"
	}
	dbName := ""
	if config.Name != "" {
		dbName = fmt.Sprintf("/%s", config.Name)
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s%s?sslmode=%s",
		config.Username, config.Password, config.Address, dbName, sslmode)
	if dbP ,errO := sql.Open("pgx", connStr); errO == nil {
		if dbP != nil {
			if errO = dbP.Ping(); errO != nil {
				log.Printf("DB Connectivity failed..ERROR: %v", errO)
			} else {
				db = &DB{
					db: dbP,
				}
			}
		}
	}
	return
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) CreateDB(name string) error {
	//query := fmt.Sprintf("select 'create database %v' where not exists (select from pg_database where datname ='%v')", name, name)
	query := fmt.Sprintf("create database %v", name)
	_, err := db.db.Query(query)
	return err
}

func (db *DB) DropDB(name string) error {
	query := fmt.Sprintf("drop database %v", name)
	_, err := db.db.Query(query)
	return err
}

func (db *DB) CreateTable(dbname string, tables []Table) error {
	//query := fmt.Sprintf("select 'create database %v' where not exists (select from pg_database where datname ='%v')", name, name)
	/*query := fmt.Sprintf("\\c %s", dbname)
	if _, err := db.db.Query(query); err != nil {
		return err
	}*/
	query := ""
	for _, table := range tables {
		query = fmt.Sprintf("create table %v (", table.Name)
		for _, field := range table.Fields {
			query = fmt.Sprintf("%s %s", query, field.Name)
			attr := field.Attribute
			if attr.Type != "" {
				query = fmt.Sprintf("%s %s", query, attr.Type)
			}
			if attr.PrimaryKey {
				query = fmt.Sprintf("%s PRIMARY KEY", query)
			}
			if attr.NotNull {
				query = fmt.Sprintf("%s NOT NULL", query)
			}
			if attr.Unique {
				query = fmt.Sprintf("%s UNIQUE", query)
			}
			if c := attr.Check; c != "" {
				query = fmt.Sprintf("%s CHECK %s", query, c)
			}
			if attr.ForeignKey {
				if attr.References != "" {
					query = fmt.Sprintf("%s FOREIGN KEY REFERENCES %s", query, attr.References)
				}
			}
			if attr.Default != nil {
				query = fmt.Sprintf("%s DEFAULT %v", query, attr.Default)
			}
			query = fmt.Sprintf("%s,", query)
		}
		query = strings.TrimSuffix(query, ",")
		query = fmt.Sprintf("%s);", query)
		log.Println("query:", query)
		if _, err := db.db.Query(query); err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) Select() {

}

func (db *DB) Insert() {

}

func (db *DB) Update() {

}

func (db *DB) Delete() {

}

