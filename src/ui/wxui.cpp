// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
#include "wxui.h"
using namespace std;

//-- Fanling7Frame --
Fanling7Frame::Fanling7Frame(Engine* engine)
    : wxFrame(NULL, wxID_ANY, "Fanling7", wxDefaultPosition, wxDefaultSize,
              wxDEFAULT_FRAME_STYLE), _engine(engine) {
    if(verbosity>0)  cerr<<"creating user interface\n";
//_engine->getInput();
    wxMenu* menuFile = new wxMenu;
    menuFile->Append(wxID_EXIT);
    wxMenuBar* menuBar = new wxMenuBar;
    menuBar->Append(menuFile, "&File");
    SetMenuBar(menuBar);
    CreateStatusBar();
    SetStatusText("Fanling7 started.");
    makeControls();
    bindControls() ;
    SetSizerAndFit(_sizer);
}

void Fanling7Frame::makeControls() {
    _sizer = new wxFlexGridSizer(1, 1, 1);
    _controlSizer = new wxFlexGridSizer(4, 1, 1);
    _sizer->Add(_controlSizer);
    _sizer->SetFlexibleDirection(wxBOTH);
    _controlSizer->Add(new wxStaticText(this, wxID_ANY, "Page type:"));
    wxArrayString items;
    std::vector<std::string> pageTypes=_engine->getPageTypes();
    for(string& s : pageTypes) items.Add(s);
    items.Sort();
    wxChoice* typeListChoice = new wxChoice(this, IDTYPE, wxDefaultPosition, wxDefaultSize, items, wxCB_SORT);
    _controlSizer->Add(typeListChoice);
    _controlSizer->Add(new wxStaticText(this, wxID_ANY, "Page ident:"));
    //std::vector<std::string> pages=_engine->getPages();
    //items.Clear();
    //for(string& s : pages) items.Add(s);
    //items.Sort();
    //_identCombo = new wxComboBox(this, IDIDENT);
    //_identCombo->Append(items);
    //_controlSizer->Add(_identCombo);
    _controlSizer->Add(new wxStaticText(this, wxID_ANY, ""));
    wxButton* makePageButton = new wxButton(this, IDMAKEPAGE, "Make page");
    _controlSizer->Add(makePageButton);
    _controlSizer->Add(new wxStaticText(this, wxID_ANY, "Action:"));
    items.Clear();
    _actNameChoice = new wxChoice(this, IDACTNAME, wxDefaultPosition, wxDefaultSize, items, wxCB_SORT);
    _controlSizer->Add(_actNameChoice);
    _controlSizer->Add(new wxStaticText(this, wxID_ANY, "Number:"));
    _actNumSpin = new wxSpinCtrl(this,IDACTNUM, wxEmptyString, wxDefaultPosition, wxDefaultSize, wxSP_ARROW_KEYS, 1, 100, 1);
    _controlSizer->Add(_actNumSpin);
    _controlSizer->Add(new wxStaticText(this, wxID_ANY, ""));
    _actionButton = new wxButton(this, IDACTION, "Do action");
    _controlSizer->Add(_actionButton);
    _showEditCheck = new wxCheckBox(this, IDSHOWEDIT, "Edit?");
    _showEditCheck->SetValue(false);
    _controlSizer->Add(_showEditCheck);
    _saveEditButton = new wxButton(this, IDSAVEEDIT, "Save");
    _controlSizer->Add(_saveEditButton);
    _revertButton = new wxButton(this, IDREVERT, "Revert");
    _controlSizer->Add(_revertButton);
    _webView =wxWebView::New(this, IDWEBVIEW, "", wxDefaultPosition, wxSize(200,200), wxWebViewBackendDefault, 0);
    _sizer->Add(_webView, 1, wxEXPAND|wxALIGN_TOP|wxALIGN_LEFT);
    styleEditor();
    _sizer->Add(_textEd, 1, wxEXPAND|wxALIGN_TOP|wxALIGN_LEFT);
    _sizer->Show(_textEd,false);
    _sizer->AddGrowableCol(0);
    _sizer->AddGrowableRow(1);
    _sizer->AddGrowableRow(2);

}
void Fanling7Frame::bindControls() {

    Bind(wxEVT_MENU, [=](wxCommandEvent&) {
        Close(true);
    }, wxID_EXIT);
    Bind(wxEVT_CHOICE, [=](wxCommandEvent&event) {
        _chosenType = event.GetString();
    }, IDTYPE);
    Bind(wxEVT_TEXT, [=](wxCommandEvent&) {
        setPage(string(_identCombo->GetValue()));
    }, IDIDENT);
    Bind(wxEVT_COMBOBOX, [=](wxCommandEvent&event) {
        setPage(string(event.GetString()));
    }, IDIDENT);
    Bind(wxEVT_BUTTON, [=](wxCommandEvent&) {
        string newIdent = string(_identCombo->GetValue());
        if(newIdent=="" or _chosenType=="") {
            showError("no page to create or no page type.");
            return;
        }
        Result result;
        _engine-> createPage(newIdent,_chosenType, result);
        showResult(result);
        _identCombo->Append(newIdent);
        setPage(newIdent);
        SetStatusText("Page created.");
    }, IDMAKEPAGE);
    Bind(wxEVT_CHOICE, [=](wxCommandEvent&) {
        _actionName=_actNameChoice->GetStringSelection();
    }, IDACTNAME);
    Bind(wxEVT_SPINCTRL, [=](wxCommandEvent&) {
        _actionNumber=_actNumSpin->GetValue();
    }, IDACTNUM);
    Bind(wxEVT_BUTTON, [=](wxCommandEvent&) {
        Result result;
        _engine->getPage(_chosenIdent,  result);
        showResult(result);
        if(_chosenIdent=="" or result.severity!=Severity::okFound or _actionName=="") {
            showError("no page to show, no action, or page does not exist.");
            return;
        }
        PagePtr page =  result.page;
        page->applyAction(_actionName,_actionNumber, result);
        showResult(result);
        setPage(_chosenIdent, true, true);
        SetStatusText(_actionName+" done.");
    }, IDACTION);
    Bind(wxEVT_CHECKBOX, [=](wxCommandEvent&) {
        showWebEdit(_showEditCheck->GetValue());
    }, IDSHOWEDIT);
    Bind(wxEVT_WEBVIEW_NAVIGATING, [=](wxWebViewEvent&event) {
        setPage(_engine->identFromURL(string(event.GetURL())),false);
        event.Allow();
        if(verbosity>0)  cerr << "Navigating to "<<event.GetURL() <<"\n";
    }, IDWEBVIEW);
    Bind(wxEVT_WEBVIEW_ERROR , [=](wxWebViewEvent&event) {
        cerr << "Error in "<<event.GetURL() <<": "<<event.GetString()<<"\n";
    }, IDWEBVIEW);
    Bind(wxEVT_BUTTON, [=](wxCommandEvent&) {
        if(_chosenIdent=="") {
            showError("no page to save.");
            return;
        }
        Result result;
        _engine->getPage(_chosenIdent,  result);
        showResult(result);
        if(result.severity!=Severity::okFound and _chosenType != "") {
            _engine-> createPage(_chosenIdent,_chosenType,result);
            showResult(result);
        }
        savePage(_chosenIdent) ;
    }, IDSAVEEDIT);
    Bind(wxEVT_BUTTON, [=](wxCommandEvent&) {
        if(_chosenIdent=="") {
            showError("no page to revert.");
            return;
        }
        showWebEdit(_showEditCheck->GetValue());
    }, IDREVERT);
}
void Fanling7Frame::styleEditor() {
#pragma GCC diagnostic ignored "-Woverflow"
    _textEd=new wxStyledTextCtrl(this, IDEDIT, wxDefaultPosition, wxSize(200,200));
    _textEd->SetLexerLanguage("yaml");
    _textEd->SetMarginWidth(EDITMARGIN, 50);
    _textEd->StyleSetForeground(wxSTC_STYLE_LINENUMBER, wxColour(75, 75, 75));
    _textEd->StyleSetBackground(wxSTC_STYLE_LINENUMBER, wxColour(220, 220, 220));
    _textEd->SetMarginType(EDITMARGIN, wxSTC_MARGIN_NUMBER);
    _textEd->SetWrapMode(wxSTC_WRAP_WORD);
    _textEd->SetTabWidth(4);
    _textEd->SetProperty(wxT("indentation.smartindenttype"), wxT("simple"));
    _textEd->SetProperty(wxT("indentation.indentwidth"), wxT("4"));
    _textEd->SetProperty(wxT("indentation.tabwidth"), wxT("4"));
    _textEd->SetProperty("spell.mistake.indicator", "style:squigglelow");
    _textEd->SetTabIndents(true);
    _textEd->SetBackSpaceUnIndents(true);
    _textEd->SetUseTabs(false);
    _textEd->StyleSetForeground(wxSTC_YAML_COMMENT, wxColour(128, 128, 128));
    _textEd->StyleSetForeground(wxSTC_YAML_DEFAULT, wxColour(128, 256, 256));
    _textEd->StyleSetForeground(wxSTC_YAML_DOCUMENT, wxColour(128, 256, 128));
    _textEd->StyleSetForeground(wxSTC_YAML_ERROR, wxColour(256,0,0));
    _textEd->StyleSetForeground(wxSTC_YAML_IDENTIFIER, wxColour(128, 256, 128));
    _textEd->StyleSetForeground(wxSTC_YAML_KEYWORD, wxColour(128,0, 256));
    _textEd->StyleSetBold(wxSTC_YAML_KEYWORD, true);
    _textEd->StyleSetForeground(wxSTC_YAML_NUMBER, wxColour(256, 256, 128));
    _textEd->StyleSetForeground(wxSTC_YAML_REFERENCE, wxColour(128, 256, 256));
    _textEd->StyleSetForeground(wxSTC_YAML_TEXT , wxColour(256, 256, 256));
#pragma GCC diagnostic pop
}
void Fanling7Frame::setPage(const string ident,const bool web, const bool force) {
    string oldIdent = _chosenIdent;
    if(verbosity>0) cerr<<ident<<": getting page, previous "<<oldIdent<<"\n";
    Result result;
    _engine->getPage(ident,  result);
    showResult(result);
    if((!force and oldIdent==ident) or result.severity!=Severity::okFound) return;
    PagePtr page = result.page;
    _chosenIdent=ident;
    _identCombo->SetValue(_chosenIdent);
    loadEditor(oldIdent, _chosenIdent);
    const bool canEdit=page->canEdit();
    if(!canEdit) {
        showWebEdit(false);
        _showEditCheck->SetValue(false);
    }
    _showEditCheck->Enable(canEdit);
    wxArrayString actionsWx;
//TODO: enable/disable other controls
    if(verbosity>0)
        cerr<<"setting controls for "<<ident<<(canEdit?", can":", cannot")<<" edit\n";
    if(ident!="") {
        vector<string> actions = page->actions();
        for(string& s : actions) {
            actionsWx.Add(s);
            if(verbosity>0)cerr << s << ": action\n";
        }
    }
    _actNameChoice->Set(actionsWx);
    if(web)_webView->LoadURL(_engine->getPageOutURL(ident));
}
void Fanling7Frame::showWebEdit(const bool showEdit) {
    Result result;
    _engine->getPage(_chosenIdent,  result);
    showResult(result);
    if(_chosenIdent=="" or result.severity!=Severity::okFound) {
        showError("no page to show or page does not exist.");
        return;
    }
    PagePtr page = result.page;
    _sizer->Show(_webView, !showEdit);
    _sizer->Show(_textEd,showEdit);
    _controlSizer->Show(_saveEditButton,showEdit);
    _controlSizer->Show(_revertButton,showEdit);
    _sizer-> Layout();
    if(showEdit) _textEd->ChangeValue(page->getPageYAMLDetail());
    else _webView->LoadURL(_engine->getPageOutURL(_chosenIdent));
}
void Fanling7Frame::loadEditor(const string oldIdent, const string ident) {
    if(oldIdent!="" and _textEd->IsModified()) {
        int res = wxMessageBox("Edit text "+oldIdent+" has been changed, save?", "Save?", wxYES_NO|wxCANCEL|wxCENTRE);
        switch(res) {
        case wxNO:
            ;
        case wxCANCEL:
            return;
        default:
            savePage(oldIdent);
        }
    }
    Result result;
    _engine->getPage(ident,  result);
    showResult(result);
    if(ident != "" and result.severity==Severity::okFound)
        _textEd->ChangeValue(result.page->getPageYAMLDetail());
    _textEd->SetModified(false);
}
void Fanling7Frame::savePage(string ident) {
    string newValue=string(_textEd->GetValue());
    if(verbosity>0) cerr<<"saving page "<<ident<< " with "<<newValue<<"\n-----\n";
    Result result;
    _engine->getPage(ident,  result);
    showResult(result);
    result.page->setDetailAndProcess(newValue, result);
    showResult(result);
    _textEd->SetModified(false);
}
void Fanling7Frame::showIndex() {
    setPage("index");
}
void showError(const string& msg, const Severity severity) {
    std::string text;
    switch(severity) {
    default:
        text="User error";
        break;
    case  Severity::system:
        text="System error";
        break;
    case Severity::okFound:
    case Severity::notFound:
        return  ;
    }
    (void) wxMessageBox(msg, text, wxOK|wxCENTRE|wxICON_ERROR);
}
void showResult(const Result& result) {
    if(result.ok()) return;
    (void) wxMessageBox(result.text, result.severity==Severity::system?"System error":"User error", wxOK|wxCENTRE|wxICON_ERROR);
}

//-- Fanling7App --
bool Fanling7App::OnInit() {
    if(verbosity>0) cerr << "init user interface...\n";
    Fanling7Frame* frame = new Fanling7Frame(_engine);
    frame->verbosity=verbosity;
    frame->showIndex();
    frame->Show(true);
    return true;
}
wxIMPLEMENT_APP_NO_MAIN(Fanling7App);

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
