add_executable(${PROJECT_NAME} judenfetch.cpp
                               infostream.h 
                               infostream.cpp)

set_target_properties(${PROJECT_NAME} PROPERTIES RUNTIME_OUTPUT_DIRECTORY "${CMAKE_BINARY_DIR}")
