#pragma once

#include <fstream>
#include <iostream>
#include <string>
#include <vector>


// substring the last directory element
void dirscan(std::string directory);

bool isWanted(const std::string& line, std::string field);

std::vector<std::string> explode (const std::string& str, const char& ch);

std::string parse(std::string field, std::string file);



namespace status 
{
std::string user = std::getenv("LOGNAME");
std::string hostname;
std::string distro;
std::string kernel;
std::string shell = std::getenv("SHELL");
std::string deskManager = std::getenv("DESKTOP_SESSION");
std::string terminal = std::getenv("TERM");
std::string local = std::getenv("LANG");
} 

// namespace status
// User, HostName, Distro, Kernel, Ram, Shell
