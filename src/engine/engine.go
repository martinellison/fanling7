// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
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

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
