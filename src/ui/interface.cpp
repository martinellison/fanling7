#include "interface.h"
#include "wxui.h"
#include <iostream>
#include <exception>
using namespace std;
// Error --

bool Error::ok() {
    return severity()==Severity::ok;
}
Severity Error::severity() {
    return Severity::system;
}
std::string Error::text() {
    return "bad call text";
}

//-- Engine (dummy stubs) --
Engine::~Engine() {}
bool Engine::pageExists(const std::string& ident) {
    throw "bad call";
}
Error* Engine::applyAction(std::string ident, std::string actionName , int actionNumber) {
    throw "bad call";
}
Error* Engine::createPage(std::string newIdent, std::string newType) {
    throw "bad call";
}
Error* Engine::exportPages() {
    throw "bad call";
}
void Engine::getInput() {
    throw "bad call";
}
std::vector<std::string> Engine::getPageTypes() {
    throw "bad call";
}
std::vector<std::string> Engine::getPages() {
    throw "bad call";
}
void Engine::setConfig(const std::string& path) {
    throw "bad call";
}
void Engine::setIndir(const std::string& dir) {
    throw "bad call";
}
void Engine::setOutdir(const std::string& dir) {
    throw "bad call";
}
void Engine::setMetadir(const std::string& dir) {
    throw "bad call";
}
void Engine::setVerbose(const int verbosity) {
    throw "bad call";
}
void Engine::init() {
    throw "bad call";
}
void Engine::readOptions() {
    throw "bad call";
}
string Engine::getPageOutURL(const std::string& ident) {
    throw "bad call getPageOutURL";
}
std::string Engine::identFromURL(const std::string& url) {
    throw "bad call identFromURL";
}
std::string Engine::getPageYAMLDetail(const std::string& ident) {
    throw "bad call";
}
Error* Engine::setPageDetailAndProcess(const std::string& ident, const std::string& text) {
    throw "bad call";
}
bool Engine::canEditPage(const std::string& ident) {
    throw "bad call";
}
std::vector<std::string> Engine::actionsForPage(const std::string& ident) {
    throw "bad call";
}
void Engine::dumpOptions() {
    cerr << "should not dump like this!";
}


//-- UserInterface --
UserInterface::~UserInterface() {}
void UserInterface::setEngine(Engine* engine) {
    _engine=engine;
}
void UserInterface::setVerbose(const int verbosity) {
    _verbosity=verbosity;
}
void UserInterface::start() {
    try {
        cerr << "starting wx ui\n";
        Fanling6App* app = new Fanling6App();
        app->setEngine(_engine);
        app->verbosity=_verbosity;
        wxApp::SetInstance(app);
        int argCount = 0;
        char* argv[0];
        (void)wxEntry(argCount, argv);
        cerr << "wx ui terminated\n";
    } catch(exception& ex) {
        cerr << "exception: " << ex.what() << "\n";
    } catch(...) {
        cerr << "unknown exception!\n";
    }
}

UserInterface* makeUserInterface() {
    return new UserInterface;
}
