// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
#ifndef INTERFACE_H
#define INTERFACE_H
#include <memory>
#include <string>
#include <vector>
struct Result;
class Page ;
class Engine;
class Page {
public:
    virtual ~Page() {}
    virtual void applyAction(const std::string actionName,const int actionNumber, Result& result) {} /* used by UI and command line */
    virtual std::string getPageYAMLDetail() {} /* used by UI */
    virtual void setDetailAndProcess(const std::string& text, Result& result) {} /* used by UI */
    virtual bool canEdit() {
        return false;   /* used by UI */
    }
    virtual std::vector<std::string> actions() {} /* used by UI */
};
typedef Page* PagePtr;
enum class Severity {
    okFound, notFound, user, system
};
struct Result {
    bool ok() const {
        switch(severity) {
        default:
            return false;
        case Severity::okFound:
        case Severity::notFound:
            return true;
        }
    }
    bool found() const {
        return severity==Severity::okFound;
    }
    bool notFound() const {
        return severity==Severity::notFound;
    }
    Result() : severity(Severity::okFound), text(""), page(nullptr) {}
    void setError(const std::string message="",const Severity s=Severity::user) {
        severity=s;
        text=message;
    }
    void setFound(const PagePtr p, const bool found=true) {
        severity=found?Severity::okFound:Severity::notFound;
        page=p;
    }
    Severity severity;
    std::string text;
    PagePtr page;
//    std::vector<PagePtr> pages;
};
class Engine {
public:
    virtual ~Engine() {}
    //virtual bool pageExists(const std::string& ident) {
    //throw "bad call";   /* used by UI */
    //}
    virtual void getPage(const std::string& ident, Result& result) {
        throw "bad call";
    }
    virtual void createPage(const std::string newIdent, const std::string pageType, Result& result) {
        throw "bad call";   /* used by UI and command line */
    }
    virtual void exportPages(Result& result) {
        throw "bad call";   /* used by command line */
    }
    virtual void getInput() {
        throw "bad call";   /* used by UI and command line */
    }
    virtual std::vector<std::string> getPageTypes() {
        throw "bad call";   /* used by UI */
    }
    // virtual std::vector<std::string> getPages(); /* used by UI */
    virtual void setConfig(const std::string& path) {
        throw "bad call";   /* used by command line */
    }
    virtual void setIndir(const std::string& dir) {
        throw "bad call";   /* used by command line */
    }
    virtual void setOutdir(const std::string& dir) {
        throw "bad call";   /* used by command line */
    }
    virtual void setMetadir(const std::string& dir) {
        throw "bad call";   /* used by command line */
    }
    virtual void setVerbose(const int verbosity) {
        throw "bad call";   /* used by command line */
    }
    virtual void init() {
        throw "bad call";   /* used by command line */
    }
    virtual void readOptions() {
        throw "bad call";
    }/* used by command line*/
    virtual std::string getPageOutURL(const std::string& ident) {
        throw "bad call";   /* used by UI */
    }
    virtual std::string identFromURL(const std::string& url) {
        throw "bad call";   /* used by UI */
    }
    virtual void dumpOptions() {
        throw "bad call";   /* used by command line */
    }
};
class UserInterface {
public:
    ~UserInterface() {}
    void setEngine(Engine* engine) {
        _engine=engine;
    }
    void setVerbose(const int verbosity) {
        _verbosity=verbosity;
    }
    void start();
private:
    Engine* _engine;
    int _verbosity;
};

UserInterface* makeUserInterface();
#endif

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
