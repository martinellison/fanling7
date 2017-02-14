// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
#ifndef WXUI_H
#define WXUI_H
#include <wx/wx.h>
#include <wx/app.h>
#include <wx/webview.h>
#include <wx/spinctrl.h>
#include <wx/stc/stc.h>
#include <iostream>
#include "interface.h"
using namespace std;
void showResult(const Result& result) ;
void showError(const string& msg, const Severity severity=Severity::user);
enum  EditStyles {EDITMARGIN};
//-- Fanling7Frame --
class Fanling7Frame: public wxFrame {
public:
    Fanling7Frame(Engine* engine);
    void showIndex();
    int verbosity=0;
private:
    void OnExit(wxCommandEvent& event);
    void makeControls();
    void bindControls();
    Engine* _engine;
    string _chosenType = "";
    string _chosenIdent = "";
    string _actionName = "";
    int _actionNumber = 0;
    wxFlexGridSizer* _sizer;
    wxFlexGridSizer* _controlSizer;
    wxSpinCtrl* _actNumSpin;
    wxButton* _actionButton ;
    wxChoice* _actNameChoice;
    wxCheckBox* _showEditCheck;
    wxComboBox* _identCombo;
    wxWebView* _webView;
    wxStyledTextCtrl* _textEd;
    wxButton* _saveEditButton;
    wxButton* _revertButton;
    void setPage(const string ident,const bool web=true,  const bool force=false);
    void loadEditor(const string oldIdent, const string ident);
    void styleEditor();
    void showWebEdit(const bool showWeb);
    void savePage(string ident);
};
enum {IDDOIT = 1, IDTYPE, IDIDENT, IDMAKEPAGE, IDACTIDENT, IDACTION, IDACTNAME, IDACTNUM, IDSHOWEDIT, IDWEBVIEW, IDSAVEEDIT, IDREVERT, IDEDIT} WindoIds;

//-- Fanling7App --
class Fanling7App : public wxApp {
public:
    Fanling7App() {
    }
    void setEngine(Engine* engine) {
        _engine=engine;
    }
    virtual bool OnInit() override;
    int verbosity=0;
private:
    Engine* _engine;
};
#endif

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
