#include "wxui.h"
using namespace std;

//-- Fanling6Frame --
Fanling6Frame::Fanling6Frame(Engine* engine)
    : wxFrame(NULL, wxID_ANY, "Fanling6", wxDefaultPosition, wxDefaultSize,
              wxDEFAULT_FRAME_STYLE), _engine(engine) {
    if(verbosity>0)  cerr<<"creating user interface\n";
//_engine->getInput();
    wxMenu* menuFile = new wxMenu;
    menuFile->Append(wxID_EXIT);
    wxMenuBar* menuBar = new wxMenuBar;
    menuBar->Append(menuFile, "&File");
    SetMenuBar(menuBar);
    CreateStatusBar();
    SetStatusText("Fanling6 started.");
    makeControls();
    bindControls() ;
    SetSizerAndFit(_sizer);
}

void Fanling6Frame::makeControls() {
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
    std::vector<std::string> pages=_engine->getPages();
    items.Clear();
    for(string& s : pages) items.Add(s);
    items.Sort();
    _identCombo = new wxComboBox(this, IDIDENT);
    _identCombo->Append(items);
    _controlSizer->Add(_identCombo);
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
void Fanling6Frame::bindControls() {

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
        showError(_engine-> createPage(newIdent,_chosenType));
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
        if(_chosenIdent=="" or !_engine->pageExists(_chosenIdent)or _actionName=="") {
            showError("no page to show, no action, or page does not exist.");
            return;
        }
        showError(_engine->applyAction(_chosenIdent, _actionName,_actionNumber));
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
        if(!_engine->pageExists(_chosenIdent) and _chosenType != "")
            showError(_engine-> createPage(_chosenIdent,_chosenType));
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
void Fanling6Frame::styleEditor() {
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
void Fanling6Frame::setPage(const string ident,const bool web, const bool force) {
    string oldIdent = _chosenIdent;
    if(verbosity>0) cerr<<ident<<": getting page, previous "<<oldIdent<<"\n";
    if((!force and oldIdent==ident) or !_engine->pageExists(ident)) return;
    _chosenIdent=ident;
    _identCombo->SetValue(_chosenIdent);
    loadEditor(oldIdent, _chosenIdent);
    const bool canEdit= _engine->canEditPage(ident);
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
        vector<string> actions = _engine->actionsForPage(ident);
        for(string& s : actions) {
            actionsWx.Add(s);
            if(verbosity>0)cerr << s << ": action\n";
        }
    }
    _actNameChoice->Set(actionsWx);
    if(web)_webView->LoadURL(_engine->getPageOutURL(ident));
}
void Fanling6Frame::showWebEdit(const bool showEdit) {
    if(_chosenIdent=="" or !_engine->pageExists(_chosenIdent)) {
        showError("no page to show or page does not exist.");
        return;
    }
    _sizer->Show(_webView, !showEdit);
    _sizer->Show(_textEd,showEdit);
    _controlSizer->Show(_saveEditButton,showEdit);
    _controlSizer->Show(_revertButton,showEdit);
    _sizer-> Layout();
    if(showEdit) _textEd->ChangeValue(_engine->getPageYAMLDetail(_chosenIdent));
    else _webView->LoadURL(_engine->getPageOutURL(_chosenIdent));
}
void Fanling6Frame::loadEditor(const string oldIdent, const string ident) {
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
    if(ident != "" and _engine->pageExists(ident))
        _textEd->ChangeValue(_engine->getPageYAMLDetail(ident));
    _textEd->SetModified(false);
}
void Fanling6Frame::savePage(string ident) {
    string newValue=string(_textEd->GetValue());
    if(verbosity>0) cerr<<"saving page "<<ident<< " with "<<newValue<<"\n-----\n";
    showError(_engine->setPageDetailAndProcess(ident,newValue));
    _textEd->SetModified(false);
}
void Fanling6Frame::showIndex() {
    setPage("index");
}
void showError(Error* err) {
    if(err->ok()) return;
    (void) wxMessageBox(err->text(), err->severity()==Severity::system?"System error":"User error", wxOK|wxCENTRE|wxICON_ERROR);
}
void showError(const string& msg, const Severity severity) {
    if(severity==Severity::ok) return;
    (void) wxMessageBox(msg, severity==Severity::system?"System error":"User error", wxOK|wxCENTRE|wxICON_ERROR);
}

//-- Fanling6App --
bool Fanling6App::OnInit() {
    if(verbosity>0) cerr << "init user interface...\n";
    Fanling6Frame* frame = new Fanling6Frame(_engine);
    frame->verbosity=verbosity;
    frame->showIndex();
    frame->Show(true);
    return true;
}
wxIMPLEMENT_APP_NO_MAIN(Fanling6App);