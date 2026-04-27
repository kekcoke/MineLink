#include "operator_agent.h"
#include <iostream>
#include <thread>
#include <iomanip>
#include <sstream>

namespace minelink {

OperatorAgent::OperatorAgent(const std::string& vehicleId, const std::string& brokerUrl)
    : vehicleId_(vehicleId), brokerUrl_(brokerUrl), client_(brokerUrl, "operator-" + vehicleId) {}

OperatorAgent::~OperatorAgent() {
    stop();
}

void OperatorAgent::start() {
    try {
        mqtt::connect_options connOpts;
        connOpts.set_keep_alive_interval(20);
        connOpts.set_clean_session(true);

        client_.connect(connOpts)->wait();
        std::cout << "Operator [" << vehicleId_ << "] connected to Edge Broker." << std::endl;

        running_ = true;
        std::thread([this]() { telemetryLoop(); }).detach();
    } catch (const mqtt::exception& exc) {
        std::cerr << "MQTT Connection failed for [" << vehicleId_ << "]: " << exc.what() << std::endl;
    }
}

void OperatorAgent::stop() {
    running_ = false;
    if (client_.is_connected()) {
        client_.disconnect()->wait();
    }
}

void OperatorAgent::telemetryLoop() {
    while (running_) {
        simulatePhysics();
        publishTelemetry();
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }
}

void OperatorAgent::simulatePhysics() {
    // Simple resource depletion logic
    metrics_.fuelLevel -= 0.05;
    if (metrics_.fuelLevel < 0) metrics_.fuelLevel = 0;
    
    metrics_.engineTemp += 0.1;
    if (metrics_.engineTemp > 110) metrics_.engineTemp = 110;

    metrics_.tirePressure -= 0.01;
}

void OperatorAgent::publishTelemetry() {
    using nlohmann::json;

    auto now = std::chrono::system_clock::now();
    auto in_time_t = std::chrono::system_clock::to_time_t(now);
    
    std::stringstream ss;
    ss << std::put_time(std::gmtime(&in_time_t), "%Y-%m-%dT%H:%M:%SZ");

    json payload = {
        {"vehicleId", vehicleId_},
        {"timestamp", ss.str()},
        {"metrics", {
            {"fuelLevel", metrics_.fuelLevel},
            {"engineTemp", metrics_.engineTemp},
            {"tirePressure", metrics_.tirePressure}
        }},
        {"currentTask", currentTask_}
    };

    std::string topic = "telemetry/truck/" + vehicleId_ + "/status";
    client_.publish(topic, payload.dump(), 1, false);
}

} // namespace minelink
