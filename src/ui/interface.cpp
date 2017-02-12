#include "interface.h"
#include "wxui.h"
#include <iostream>
#include <exception>
using namespace std;

//-- UserInterface --
void UserInterface::start() {
    try {
        cerr << "starting wx ui\n";
        Fanling7App* app = new Fanling7App();
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
