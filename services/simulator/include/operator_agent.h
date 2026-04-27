#pragma once

#include <string>
#include <chrono>
#include <mqtt/async_client.h>
#include <nlohmann/json.hpp>

namespace minelink {

struct VehicleMetrics {
    double fuelLevel = 100.0;
    double engineTemp = 85.0;
    double tirePressure = 102.0;
};

class OperatorAgent {
public:
    OperatorAgent(const std::string& vehicleId, const std::string& brokerUrl);
    ~OperatorAgent();

    void start();
    void stop();

private:
    void telemetryLoop();
    void simulatePhysics();
    void publishTelemetry();

    std::string vehicleId_;
    std::string brokerUrl_;
    VehicleMetrics metrics_;
    std::string currentTask_ = "HAUL_ORE";
    
    bool running_ = false;
    mqtt::async_client client_;
};

} // namespace minelink
