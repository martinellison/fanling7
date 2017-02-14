// Copyright 2017 Martin Ellison. For GPL3 licence notice, see the end of this file.
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

// This file is part of Fanling7. Fanling7 is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version. Fanling7 is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details. You should have received a copy of the GNU General Public License along with Fanling7. If not, see <http://www.gnu.org/licenses/>.
