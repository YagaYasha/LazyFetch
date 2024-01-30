#pragma once

#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <cstdlib>



// substring the last directory element
void dirscan(std::string directory);

bool isWanted(const std::string& line, std::string field);

std::vector<std::string> explode (const std::string& str, const char& ch);

std::string parse(std::string field, std::string file);

std::string extractLine(std::string field, std::string file, char delimeter);

namespace status 
{
std::string user = std::getenv("LOGNAME");
std::ifstream hostname("/etc/hostname");
std::string distro = extractLine("PRETTY_NAME", "/etc/os-release", '=');
std::string kernel = extractLine("Linux version", "proc/version", '-');
std::string shell = std::getenv("SHELL");
std::string deskManager = std::getenv("DESKTOP_SESSION");
std::string terminal = std::getenv("TERM");
std::string local = std::getenv("LANG");
} 

// namespace status
// User, HostName, Distro, Kernel, Ram, Shell
