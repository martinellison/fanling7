#ifndef INTERFACE_H
#define INTERFACE_H
#include <memory>
#include <string>
#include <vector>
enum class Severity {
    ok, user, system
};
class Error {
public:
    virtual ~Error() {}
    virtual bool ok();
    virtual Severity severity();
    virtual std::string text();
};
class Engine {
public:
    virtual ~Engine();
    virtual bool pageExists(const std::string& ident); /* used by UI */
    virtual Error* applyAction(const std::string ident, const std::string actionName,const int actionNumber); /* used by UI and command line */
    virtual Error* createPage(const std::string newIdent,const std::string newType); /* used by UI and command line */
    virtual Error* exportPages(); /* used by command line */
    virtual void getInput(); /* used by UI and command line */
    virtual std::vector<std::string> getPageTypes(); /* used by UI */
    virtual std::vector<std::string> getPages(); /* used by UI */
    virtual void setConfig(const std::string& path); /* used by command line */
    virtual void setIndir(const std::string& dir); /* used by command line */
    virtual void setOutdir(const std::string& dir); /* used by command line */
    virtual void setMetadir(const std::string& dir); /* used by command line */
    virtual void setVerbose(const int verbosity); /* used by command line */
    virtual void init(); /* used by command line */
    virtual void readOptions(); /* used by command line */
    virtual std::string getPageOutURL(const std::string& ident); /* used by UI */
    virtual std::string identFromURL(const std::string& url); /* used by UI */
    virtual std::string getPageYAMLDetail(const std::string& ident); /* used by UI */
    virtual Error* setPageDetailAndProcess(const std::string& ident, const std::string& text); /* used by UI */
    virtual bool canEditPage(const std::string& ident); /* used by UI */
    virtual std::vector<std::string> actionsForPage(const std::string& ident); /* used by UI */
    virtual void dumpOptions(); /* used by command line */
};
class UserInterface {
public:
    virtual ~UserInterface();
    virtual void setEngine(Engine* engine); /* used by UIand command line */
    virtual void setVerbose(const int verbosity);
    virtual void start();
private:
    Engine* _engine;
    int _verbosity;
};

UserInterface* makeUserInterface();
#endif
