package main

import (
	"os"
	"time"
	"flag"
	"fmt"
	"go/build"
	"path/filepath"
	"text/template"
	"math/rand"
)

const (
	Black       = "0;30"
	DarkGray    = "1;30"
	Blue        = "0;34"
	LightBlue   = "1;34"
	Green       = "0;32"
	LightGreen  = "1;32"
	Cyan        = "0;36"
	LightCyan   = "1;36"
	Red         = "0;31"
	LightRed    = "1;31"
	Purple      = "0;35"
	LightPurple = "1;35"
	Brown       = "0;33"
	Yellow      = "1;33"
	LightGray   = "0;37"
	White       = "1;37"
)

var templateDir string

func init() {
	//Get template dir
	pkgInfo, err := build.Import("github.com/smyrman/goenv", "", build.FindOnly)
	if err != nil {
		panic("Could not locate the package's source code directory. " +
		"I need the templates that are located there! The error was: " + err.Error())
	}
	templateDir = filepath.Join(pkgInfo.Dir, "templates")

	// Define usage string
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s <path to new environment>\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Generate a seed
	rand.Seed(int64(time.Now().Unix()))
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "%s: This commnad takes exactly one argument!\n", os.Args[0])
		flag.Usage()
		os.Exit(1)
	}

	// Get path and name for the environmet
	envdir := flag.Arg(0)
	envname := filepath.Base(envdir)

	// Get a pseudo-random color
	colorSet := []string{Green, Cyan, Red, Purple, Brown}
	color := colorSet[rand.Intn(len(colorSet))]

	// Read template
	t, err := template.ParseFiles(filepath.Join(templateDir, "sourceme.bash"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: Could not read the 'sourceme.bash' template: %s\n", os.Args[0], err.Error())
		os.Exit(2)
	}

	// Define template data
	tdata := make(map[string]string, 2)
	tdata["Name"] = envname
	tdata["Color"] = color

	// Try to create the environment, only pre-create the src directory.
	srcdir := filepath.Join(envdir, "src")
	if err := os.MkdirAll(srcdir, os.ModeDir|0755); err != nil {
		fmt.Fprintf(os.Stderr, "%s: Could not create directory '%s': %s\n", os.Args[0], srcdir, err.Error())
		os.Exit(2)
	}

	// Write to sourcme.bash file
	fname := filepath.Join(envdir, "sourceme.bash")
	fp, err := os.Create(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: Could not create '%s': %s\n", os.Args[0], fname, err.Error())
		os.Exit(2)
	}
	err = t.Execute(fp, tdata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: Could not create '%s': %s\n", os.Args[0], fname, err.Error())
		os.Exit(2)
	}
	return
}
