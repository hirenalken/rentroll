// This code loads a csv file with business definitions.  Its format is:
//
// 		Company Designation, Company Name, Default Occupancy Type, Parking Permit In Use
//
// The code will try to load a Phonebook company with the designation supplied. It will pull
// default information from the Phonebook entry. For now, the information only includes the
// company name.
//
// Examples:
// 		REH,,4,0
// 		BBBB,Big Bob's Barrel Barn,4,0
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"phonebook/lib"
	"rentroll/rlib"
)
import _ "github.com/go-sql-driver/mysql"

// App is the global application structure
var App struct {
	dbdir    *sql.DB // phonebook db
	dbrr     *sql.DB //rentroll db
	DBDir    string  // phonebook database
	DBRR     string  //rentroll database
	DBUser   string  // user for all databases
	Report   int64   // if testing engine, which report/action to perform
	BizFile  string  // name of csv file with new biz info
	AsmtFile string  // name of csv file with assessment types
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	verPtr := flag.Bool("v", false, "prints the version to stdout")
	bizPtr := flag.String("b", "", "add business via csv file")
	asmtPtr := flag.String("a", "", "add assessment types via csv file")
	flag.Parse()
	if *verPtr {
		fmt.Printf("Version:    %s\nBuild Time: %s\n", getVersionNo(), getBuildTime())
		os.Exit(0)
	}
	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.BizFile = *bizPtr
	App.AsmtFile = *asmtPtr
}

func main() {
	readCommandLineArgs()

	var err error
	//----------------------------
	// Open RentRoll database
	//----------------------------
	// s := fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", DBUser, DBRR)
	s := rlib.RRGetSQLOpenString(App.DBUser, App.DBRR)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", App.DBRR, App.DBUser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = lib.GetSQLOpenString(App.DBUser, App.DBDir)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open: Error = %v\n", err)
		os.Exit(1)
	}
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("dbdir.Ping: Error = %v\n", err)
		os.Exit(1)
	}

	rlib.InitDBHelpers(App.dbrr, App.dbdir)

	if len(App.BizFile) > 0 {
		rlib.LoadBusinessCSV(App.BizFile)
	}
	if len(App.AsmtFile) > 0 {
		rlib.LoadAsessmentTypesCSV(App.AsmtFile)
	}

}
