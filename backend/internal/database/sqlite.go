package database

import (
	"database/sql"
	"os"
	"path/filepath"
)

type SQLiteDB struct {
	Path string
}

// create a database file path
func (db *SQLiteDB) createDatabaseFile() {
	dir := filepath.Dir(db.Path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	_, err := os.Create(db.Path)
	if err != nil {
		panic(err)
	}
}

// try connect to a sqlite database, then ping for connectivity
func (s *SQLiteDB) connect() *sql.DB {
	db, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

// create a database
func (s *SQLiteDB) Connect() *sql.DB {
	//alreadyExisted := true
	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		//alreadyExisted = false
		s.createDatabaseFile()
	}
	db := s.connect()
	/*
	if !alreadyExisted {
		schemaPath := "core/sql/schema.sql"
		s.logger.Info("running schema creation")
		file, err := os.ReadFile(schemaPath)
		if err != nil {
			panic(err) // schema should always be there
		}
		sql := string(file)
		err = d.Execute(context.TODO(), sql)
		if err != nil {
			panic(err)
		}
	}
	*/
	return db
}