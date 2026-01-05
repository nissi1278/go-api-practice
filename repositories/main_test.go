package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テスト全体で使用するsql.DB
var testDB *sql.DB
var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbPort     = 3310
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?parseTime=true", dbUser, dbPassword, dbPort, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {
	cmd := exec.Command("mysql",
		"-h", "127.0.0.1",
		"-P", "3310",
		"-u", "docker",
		"sampledb",
		"--password=docker",
		"-e",
		"source ./testdata/setupDB.sql",
	)

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}

	if err := cleanupDB(); err != nil {
		return err
	}
	if err := setupTestData(); err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql",
		"-h", "127.0.0.1",
		"-P", "3310",
		"-u", "docker",
		"sampledb",
		"--password=docker",
		"-e",
		"source ./testdata/cleanupDBsql",
	)

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func teardown() {
	// テスト後の後処理で、テスト時に挿入したデータを削除
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()
	teardown()
}
