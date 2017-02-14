package engine

import (
	"pagestore"
	"ui"
)

type Engine struct {
}

func NewEngine() ui.Engine { return ui.NewDirectorEngine(&Engine{}) }

func (e *Engine) GetPage(arg2 string, arg3 ui.Result)                 {}
func (e *Engine) CreatePage(arg2 string, arg3 string, arg4 ui.Result) {}
func (e *Engine) ExportPages(arg2 ui.Result)                          {}
func (e *Engine) GetInput()                                           {}
func (e *Engine) GetPageTypes() ui.StringVector                       { return ui.NewStringVector() }
func (e *Engine) SetConfig(arg2 string)                               {}
func (e *Engine) SetIndir(arg2 string)                                {}
func (e *Engine) SetOutdir(arg2 string)                               {}
func (e *Engine) SetMetadir(arg2 string)                              {}
func (e *Engine) SetVerbose(arg2 int)                                 {}
func (e *Engine) Init()                                               {}
func (e *Engine) ReadOptions()                                        {}
func (e *Engine) GetPageOutURL(arg2 string) string                    { return "" }
func (e *Engine) IdentFromURL(arg2 string) string                     { return "" }
func (e *Engine) DumpOptions()                                        {}

type Page interface {
	pagestore.Freezable
	ui.Page
}
