////////////////////////////////////////////////////////////////////////////
// Porgram: EasyGen
// Purpose: Easy to use universal code/text generator
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////
// Program start

package main

import (
	"bytes"
	"flag"
	"fmt"
	ht "html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	tt "text/template"

	"github.com/danverbraganza/varcaser/varcaser"
	"gopkg.in/yaml.v2"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const progname = "EasyGen" // os.Args[0]

// The Options structure holds the values for/from commandline
type Options struct {
	HTML         bool
	templateStr  string
	templateFile string
}

// common type for a *(text|html).Template value
type template interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	Name() string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var opts Options

// pre-config some varcaser transformers
var (
	ck2lc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.LowerCamelCase}
	ck2uc = varcaser.Caser{From: varcaser.KebabCase, To: varcaser.UpperCamelCase}
)

////////////////////////////////////////////////////////////////////////////
// Commandline definitions

func init() {
	flag.BoolVar(&opts.HTML, "html", false, "treat the template file as html instead of text")
	flag.StringVar(&opts.templateStr, "ts", "", "template string (in text)")
	flag.StringVar(&opts.templateFile, "tf", "", ".tmpl template file name (default: same as .yaml file)")
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage:\n %s [flags] YamlFileName\n\nFlags:\n\n",
		progname)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nYamlFileName: The name for the .yaml data and .tmpl template file\n\tOnly the name part, without extension. Can include the path as well.\n")
	os.Exit(0)
}

////////////////////////////////////////////////////////////////////////////
// Main

func main() {
	flag.Usage = usage
	flag.Parse()

	// One mandatory non-flag arguments
	if len(flag.Args()) < 1 {
		usage()
	}
	fileName := flag.Args()[0]

	fmt.Print(Generate(opts.HTML, fileName))
}

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Generate will produce output from the template according to driving data
func Generate(HTML bool, fileName string) string {
	source, err := ioutil.ReadFile(fileName + ".yaml")
	checkError(err)

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(source, &m)
	checkError(err)

	// template file name
	fileNameT := fileName
	if len(opts.templateFile) > 0 {
		fileNameT = opts.templateFile
	}

	t, err := parseFiles(HTML, fileNameT+".tmpl")
	checkError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, m)
	checkError(err)

	return buf.String()
}

// parseFiles, intialization. By Matt Harden @gmail.com
func parseFiles(HTML bool, filenames ...string) (template, error) {
	tname := filepath.Base(filenames[0])

	if HTML {
		// use html template
		t, err := ht.ParseFiles(filenames...)
		return t, err
	}

	// use text template
	funcMap := tt.FuncMap{
		"minus1": minus1,
	}

	if len(opts.templateStr) > 0 {
		t, err := tt.New("TT").Funcs(funcMap).Parse(opts.templateStr)
		return t, err
	}

	t, err := tt.New(tname).Funcs(funcMap).ParseFiles(filenames...)
	return t, err
}

// Exit if error occurs
func checkError(err error) {
	if err != nil {
		fmt.Printf("[%s] Fatal error - %v", progname, err.Error())
		os.Exit(1)
	}
}