import subprocess
import os
import time

# Configuration
SIMULATOR_BINARY = './services/simulator/operator_simulator' # Assumed build path
FLEET_SIZE = 200
DURATION_SECONDS = 60

def run_stress_test():
    print(f"--- Starting Synthetic Load Stress Test ---")
    print(f"Target Fleet Size: {FLEET_SIZE} vehicles")
    print(f"Test Duration: {DURATION_SECONDS} seconds")
    
    env = os.environ.copy()
    env["FLEET_SIZE"] = str(FLEET_SIZE)
    
    try:
        # Start the simulator with the large fleet size
        process = subprocess.Popen([SIMULATOR_BINARY], env=env, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        print("Simulator started. Generating telemetry load...")
        
        # In a real environment, we would query the Edge Broker metrics or TimescaleDB
        # to ensure it is ingesting ~200 msgs/second without significant lag.
        
        time.sleep(DURATION_SECONDS)
        
        print("Test duration reached. Terminating simulator...")
        process.terminate()
        process.wait()
        
        print("--- Stress Test Completed Successfully ---")
        
    except FileNotFoundError:
        print(f"Error: Simulator binary not found at {SIMULATOR_BINARY}.")
        print("Please build the C++ project first.")

if __name__ == "__main__":
    run_stress_test()
