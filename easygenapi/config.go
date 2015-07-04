package easygenapi

import (
	"flag"
	"fmt"
	"os"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "easygen" // os.Args[0]

// The Options struct defines the structure to holds the commandline values
type Options struct {
	HTML         bool   // treat the template file as html instead of text
	TemplateStr  string // template string (in text)
	TemplateFile string // .tmpl template file name (default: same as .yaml file)
	debug        int    // debugging level
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

// Opts holds the actual values from the command line paramters
var Opts Options

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {
	flag.BoolVar(&Opts.HTML, "html", false,
		"treat the template file as html instead of text")
	flag.StringVar(&Opts.TemplateStr, "ts", "",
		"template string (in text)")
	flag.StringVar(&Opts.TemplateFile, "tf", "",
		".tmpl template file name (default: same as .yaml file)")
	flag.IntVar(&Opts.debug, "debug", 0,
		"debugging level")
}

// The Usage function shows help on commandline usage
func Usage() {
	fmt.Fprintf(os.Stderr,
		"\nUsage:\n %s [flags] YamlFileName\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr,
		"\nYamlFileName: The name for the .yaml data and .tmpl template file\n\tOnly the name part, without extension. Can include the path as well.\n")
	os.Exit(0)
}