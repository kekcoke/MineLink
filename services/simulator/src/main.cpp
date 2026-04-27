#include "operator_agent.h"
#include <vector>
#include <iostream>
#include <thread>
#include <csignal>
#include <cstdlib>
#include <string>

bool g_running = true;

void signal_handler(int signal) {
    std::cout << "\nShutdown signal received. Stopping agents..." << std::endl;
    g_running = false;
}

int main() {
    std::signal(SIGINT, signal_handler);
    std::signal(SIGTERM, signal_handler);

    std::string brokerUrl = "tcp://localhost:1883";
    
    // Support dynamic fleet sizing for stress testing
    int fleetSize = 5; // Default size
    if (const char* env_p = std::getenv("FLEET_SIZE")) {
        fleetSize = std::atoi(env_p);
        if (fleetSize <= 0) fleetSize = 5;
    }

    std::vector<std::string> truckIds;
    for (int i = 1; i <= fleetSize; ++i) {
        std::string id = "TRK-" + std::string(3 - std::to_string(i).length(), '0') + std::to_string(i);
        truckIds.push_back(id);
    }
    
    std::vector<std::unique_ptr<minelink::OperatorAgent>> agents;

    std::cout << "--- MineLink Operator Simulator Starting ---" << std::endl;
    std::cout << "Provisioning " << truckIds.size() << " digital workers (Internal Threads)..." << std::endl;

    for (const auto& id : truckIds) {
        auto agent = std::make_unique<minelink::OperatorAgent>(id, brokerUrl);
        agent->start();
        agents.push_back(std::move(agent));
    }

    std::cout << "All agents active. Press Ctrl+C to terminate simulation." << std::endl;

    while (g_running) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    for (auto& agent : agents) {
        agent->stop();
    }

    std::cout << "Simulation terminated gracefully." << std::endl;
    return 0;
}

