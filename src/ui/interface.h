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
    virtual void applyAction(const std::string actionName,const int actionNumber, Result& result); /* used by UI and command line */
    virtual std::string getPageYAMLDetail(); /* used by UI */
    virtual void setDetailAndProcess(const std::string& text, Result& result); /* used by UI */
    virtual bool canEdit(); /* used by UI */
    virtual std::vector<std::string> actions(); /* used by UI */
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
