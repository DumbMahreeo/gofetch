package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func do(cmd string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		panic(err)
	}

	return strings.Replace(string(out), "\n", "", 1)

}

func main() {

	user := do("whoami")
	host := do("hostname")
	osname := "Void"                 // might add other systems in the future
	kernel := do("uname", "-sr")
	uptime := do("uptime", "-p")[3:]
	packages := strings.Count(do("xbps-query", "-l"), "\n")
	shell := do("basename", os.Getenv("SHELL"))
	WM := os.Getenv("WM")

	// __--COLORS--__
	const reset = "\033[0m"

	// ample/coffee theme
	const c1 = "\033[0m\033[1m\033[38;2;192;103;9m" // orange
	const c2 = "\033[0m\033[1m\033[38;2;226;78;78m" // red
	const font1 = "\033[0m\033[1m\033[38;2;226;78;78m"
	const font2 = "\033[0m\033[1m\033[38;2;192;103;9m"
	// ample/coffee theme

	// --__COLORS__--

	/*

	*/

	fmt.Printf(
`      %s_______      %s%s%s@%s%s                 
   %s_ \______ -     %sOS:%s        %s  
  %s| \  %s___%s  \ |    %sKERNEL:%s    %s 
  %s| | %s/   \%s | |    %sUPTIME:%s    %s
  %s| | %s\___/%s | |    %sPACKAGES:%s  %d
  %s| \______ \_|    %sSHELL:%s     %s
   %s-_______\       %sWM:%s        %s%s         

`,
		c1, font2, user, font1, font2, host,
		c1, font1, reset, osname,
		c1, c2, c1, font1, reset, kernel,
		c1, c2, c1, font1, reset, uptime,
		c1, c2, c1, font1, reset, packages,
		c1, font1, reset, shell,
		c1, font1, reset, WM, reset)

}

