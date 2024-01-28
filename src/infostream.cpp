#include <iostream>
#include <string>
#include <cstring>
#include <fstream>
#include "infostream.h"


void dirscan(std::string envdir)
{


    std::string lastDir = "";

    size_t length = envdir.size();

    for ( auto i = length; i >= 0; i--)
    {

        if(envdir[i] == '/')
        {

            lastDir = envdir.substr(i+1, length);
            std::cout << lastDir << std::endl;
            break;
        }

    }
}

std::string logName = std::getenv("LOGNAME");
std::string shell = std::getenv("SHELL");
std::string desktopManager = std::getenv("GDMSESSION");






