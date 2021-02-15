package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func do(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		panic(err)
	}

	return strings.Replace(string(out), "\n", "", 1)

}

func getPackages() string {
	db, err1 := sql.Open("sqlite3", "/var/cache/dnf/packages.db")

	if err1 != nil {
		panic(err1)
	}

	row, err2 := db.Query("SELECT count(*) FROM installed")

	if err2 != nil {
		panic(err2)
	}

	db.Close()

	for row.Next() {
		var name string
		err := row.Scan(&name)

		if err != nil {
			panic(err)
		}

		return name
	}

	return "Error"

}

func main() {

	user := do("whoami")
	host := do("hostname")
	osname := "Fedora"                 // might add other systems in the future
	kernel := do("uname", "-sr")
	uptime := do("uptime", "-p")[3:]
	shell := do("basename", os.Getenv("SHELL"))
	WM := os.Getenv("WM")

	// __--COLORS--__
	reset := "\033[0m"

	// ample/coffee theme
	c1 := "\033[0m\033[1m\033[38;2;192;103;9m"
	c2 := "\033[0m\033[1m\033[38;2;226;78;78m"
	font1 := "\033[0m\033[1m\033[38;2;226;78;78m"
	font2 := "\033[0m\033[1m\033[38;2;192;103;9m"
	// ample/coffee theme

	// --__COLORS__--

	fmt.Printf(
`%s      ______    
%s     /   __ \   %s%s%s@%s%s
%s     |  |  \ \  %sOS:%s        %s  
%s  ___!  !__/ /  %sKERNEL:%s    %s 
%s / __, %s ,___/   %sUPTIME:%s    %s
%s/ / %s |  |       %sPACKAGES:%s  %s
%s\ \__!  |       %sSHELL:%s     %s
%s \______/       %sWM:%s        %s%s

`,
		c1,
		c1, font2, user, font1, font2, host,
		c1, font1, reset, osname,
		c1, font1, reset, kernel,
		c1, c2, font1, reset, uptime,
		c1, c2, font1, reset, getPackages(),
		c2, font1, reset, shell,
		c2, font1, reset, WM, reset)

}

