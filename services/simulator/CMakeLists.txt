cmake_minimum_required(VERSION 3.14)
project(OperatorAgent CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find Paho MQTT C and C++ libraries
find_package(PahoMqttCpp REQUIRED)
find_package(PahoMqttC REQUIRED)

# Include nlohmann/json for payload serialization
# In a real environment, you would use FetchContent or a package manager (vcpkg/conan)
# For this scaffold, we assume it is installed globally.
find_package(nlohmann_json REQUIRED)

add_executable(operator_simulator
    src/main.cpp
    src/operator_agent.cpp
)

target_include_directories(operator_simulator PRIVATE include)

target_link_libraries(operator_simulator PRIVATE
    PahoMqttCpp::paho-mqttpp3
    nlohmann_json::nlohmann_json
    pthread
)
