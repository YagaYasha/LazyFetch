cmake_minimum_required(VERSION 3.2)

project(judenfetch
        LANGUAGES CXX
        DESCRIPTUION "Fetch for Linux made by an Israeli jew on C++"
)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
add_link_options("-static" "-static-libgcc" "-static-libstdc++")

add_subdirectory(src)

install(TARGETS ${PROJECT_NAME} RUNTIME DESTINATION bin)

set(CPACK_PACKAGE_NAME "judenfetch")

SET(CPACK_GENERATOR "TGZ")
SET(CPACK_INCLUDE_TOPLEVEL_DIRECTORY 1)

include(CPack)
