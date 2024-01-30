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

bool isWanted(const std::string &line, std::string field)
{
    return (line.find(field) != std::string::npos);
}

std::vector<std::string> explode(const std::string& str, const char& ch)
{

    std::string next;
    std::vector<std::string> result;

    for (std::string::const_iterator it = str.begin(); it != str.end(); it++)
    {

        if (*it == ch)
        {

            if (!next.empty())
            {
                result.push_back(next);
                next.clear();
            }
        }
        else 
        {
            next += *it;
        }
    }

    if (!next.empty())
    {
        result.push_back(next);
    }
    return result;
}


std::string parse(std::string field, std::string file)
{

    std::ifstream thefile(file);
    std::string output, line, line_pre_array;

    while (getline(thefile, line))
    {
        if (isWanted(line, field))
        {
            line_pre_array = line;
        }
    }
    thefile.close();
    return line_pre_array;
}

