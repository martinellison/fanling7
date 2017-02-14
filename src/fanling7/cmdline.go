// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
package main

import (
	"engine"
	"fmt"
	"global"
	"log"
	"strings"
	"ui"

	"github.com/lanior/panic"
	"github.com/y0ssar1an/q"
	"gopkg.in/alecthomas/kingpin.v2"
)

// An `operation` describes something that the program can do.
type operation int

const (
	exportOp operation = iota
	createOp
	actionOp
	showIdentsOp
	showTypesOp
	uiOp
)

// `Operation`
var Operation operation

// `ident`
var ident, newType, actionName string

// `actionNumber`
var actionNumber int

// `verbosity`
var verbosity int

// `getOptions` processes the command line Options.
func getOptions(engine ui.Engine) (stop bool) {
	configOption := kingpin.Flag("config", "config file location").Short('c').String()
	indirOption := kingpin.Flag("indir", "directory for input").Short('i').String()
	outdirOption := kingpin.Flag("outdir", "directory for output").Short('o').String()
	metadirOption := kingpin.Flag("metadir", "directory for templates etc").Short('m').String()
	exportOperation := kingpin.Flag("export", "export all pages as HTML").Bool()
	createOperation := kingpin.Flag("create", "create a new page").Bool()
	identOption := kingpin.Flag("ident", "page identOption for create").String()
	pageTypeOption := kingpin.Flag("type", "page type for create").Short('t').String()
	actionOption := kingpin.Flag("action", "action type (page type dependent)").Short('a').String()
	actnum := kingpin.Flag("num", "extra numeric parameter (for some actions)").Short('n').Int()
	verboseOption := kingpin.Flag("verbose", "verbosity level").Short('v').Int()
	types := kingpin.Flag("types", "list all page types").Bool()
	idents := kingpin.Flag("idents", "list all idents").Bool()
	kingpin.Parse()
	if *exportOperation {
		Operation = exportOp
	} else if *createOperation {
		Operation = createOp
	} else if *types {
		Operation = showTypesOp
	} else if *idents {
		Operation = showIdentsOp
	} else if *actionOption != "" {
		Operation = actionOp
	} else {
		Operation = uiOp
	}
	engine.SetConfig(*configOption)
	engine.ReadOptions()
	if *indirOption != "" {
		id := *indirOption
		engine.SetIndir(id)
		q.Q("in dir option set", id)
	}
	if *metadirOption != "" {
		md := *metadirOption
		engine.SetMetadir(md)
		q.Q("metadir option set", md)
	}
	if *verboseOption > 0 {
		verbosity = *verboseOption
		engine.SetVerbose(verbosity)
	}
	switch Operation {
	case exportOp:
		if *outdirOption != "" {
			engine.SetOutdir(*outdirOption)
		}
		if *outdirOption == "" {
			log.Panic("no output directory")
		}
	case createOp:
		ident = *identOption
		newType = *pageTypeOption
		if ident == "" || newType == "" {
			_, err := fmt.Printf("must specify --ident option and --type")
			panic.PanicIf(err)
			return true
		}
	case actionOp:
		ident = *identOption
		actionName = *actionOption
		if ident == "" || actionName == "" {
			_, err := fmt.Printf("must specify --ident option ")
			panic.PanicIf(err)
			return true
		}
		actionNumber = *actnum
	default:
	}
	return false
}

// `main` is the main program.
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic trapped in main: %v\n", r)
			global.TellPanic()
			q.Q(r)
		}
	}()
	var theEngine ui.Engine = engine.MakeEngine()
	if getOptions(theEngine) {
		return
	}
	theEngine.Init()
	theEngine.DumpOptions()
	switch Operation {
	case exportOp:
		theEngine.ExportPages()
	case createOp:
		theEngine.CreatePage(ident, newType)
	case actionOp:
		theEngine.ApplyAction(ident, actionName, actionNumber)
	case showIdentsOp:
		theEngine.GetInput()
		fmt.Print(strings.Join(global.SortStrings(global.StringVectorToList(theEngine.GetPages())), " "))
	case showTypesOp:
		fmt.Print(strings.Join(global.SortStrings(global.StringVectorToList(theEngine.GetPageTypes())), " "))
	case uiOp:
		theEngine.ExportPages()
		var ui ui.UserInterface = ui.MakeUserInterface()
		ui.SetEngine(theEngine)
		ui.SetVerbose(verbosity)
		ui.Start()
	default:

	}
}

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
