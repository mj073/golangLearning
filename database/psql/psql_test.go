package psql

import (
	"fmt"
	"testing"
)

var (
	db *DB
	err error
	dbName = "myDB"
	configs = []DBConfig{
		{
			Address:  "localhost:5432",
			Name:     "postgres",
			Username: "postgres",
			Password: "admin",
		},
		{
			Address:  "localhost:5432",
			Name:     "",
			Username: "postgres",
			Password: "admin",
		},
	}

	tables = []Table{
		{
			Name: "Orders",
			Fields: []TableField{
				{
					Name: "Id",
					Attribute: FieldAttribute{
						Type:       "UUID",
						PrimaryKey: true,
						//Default:    "get_random_uuid()",
					},
				},
				{
					Name: "Food",
					Attribute: FieldAttribute{
						Type:    "TEXT",
						NotNull: true,
					},
				},
				{
					Name: "Quantity",
					Attribute: FieldAttribute{
						Type:    "INTEGER",
						NotNull: true,
					},
				},
				{
					Name: "Timestamp",
					Attribute: FieldAttribute{
						Type:    "DATE",
						Default: "now()",
					},
				},
			},
		},
	}
)

func TestNewDbConn(t *testing.T) {
	for _, config := range configs {
		fmt.Println("config:", config)
		if db, err = NewDbConn(config); err != nil {
			t.Fatalf("DB Connection Failed.. ERROR: %v", err)
		} else {
			t.Log("DB Connection Success")
			if db != nil {
				if err = db.Close(); err != nil {
					t.Fatalf("DB Connection Close Failed.. ERROR: %v", err)
				} else {
					t.Log("DB Close Success")
				}
			}
		}
	}
}

func testNewDBConn(c DBConfig, t *testing.T) {
	if db, err = NewDbConn(c); err != nil {
		t.Fatalf("DB Connection Failed.. ERROR: %v", err)
	} else {
		t.Log("DB Connection Success")
	}
}
func testCloseDB(t *testing.T) {
	if db != nil {
		if err = db.Close(); err != nil {
			t.Fatalf("DB Connection Close Failed.. ERROR: %v", err)
		} else {
			t.Log("DB Close Success")
		}
	}
}

func TestDB_Close(t *testing.T) {
	testNewDBConn(configs[1], t)
	testCloseDB(t)
}

func testCreateDB(t *testing.T) {
	if db != nil {
		if err = db.CreateDB(dbName); err != nil {
			t.Fatalf("Failed to create DB..ERROR: %v", err)
		} else {
			t.Log("DB Created Successfully")
		}
	}
}

func testDropDB(t *testing.T) {
	if db != nil {
		if err = db.DropDB(dbName); err != nil {
			t.Fatalf("Failed to drop DB..ERROR: %v", err)
		} else {
			t.Log("DB Dropped Successfully")
		}
	}
}

func TestDB_CreateDB(t *testing.T) {
	testNewDBConn(configs[1], t)
	defer testCloseDB(t)
	testCreateDB(t)
}

func TestDB_DropDB(t *testing.T) {
	testNewDBConn(configs[1], t)
	defer func() {
		testDropDB(t)
		testCloseDB(t)
	}()
	testCreateDB(t)
}

func testCreateTable(t *testing.T) {
	if db != nil {
		if err = db.CreateTable(dbName, tables); err != nil {
			t.Fatalf("Failed to create DB table..ERROR: %v", err)
		} else {
			t.Log("DB Tables Created Successfully")
		}
	}
}

func TestDB_CreateTable(t *testing.T) {
	testNewDBConn(configs[1], t)
	defer func() {
		testDropDB(t)
		testCloseDB(t)
	}()
	testCreateDB(t)
	testCreateTable(t)
}
