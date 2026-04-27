using DispatchAgent.Models;
using DispatchAgent.Services;

Console.WriteLine("Starting Dispatch Optimization Agent...");

// Ensure we don't crash if the broker isn't up immediately in dev
string brokerUrl = Environment.GetEnvironmentVariable("MQTT_BROKER_URL") ?? "localhost";
string clientId = "cloud-dispatch-01";

var dispatcher = new MqttDispatcher(brokerUrl, clientId);

try
{
    // Await connection
    await dispatcher.ConnectAsync();
}
catch (Exception)
{
    Console.WriteLine("Could not connect to broker on startup. Ensure broker is running.");
    // In a real environment, we would implement Polly retry policies here
}

// Simulation Loop: Evaluate fleet telemetry and broadcast intents
Console.WriteLine("Beginning Strategic Optimization Loop...");
while (true)
{
    // Generate a simulated strategic intent
    var strategy = new DispatchStrategy
    {
        IntentId = Guid.NewGuid().ToString(),
        TargetThroughput = new Random().Next(100, 200) + 0.5, // e.g., 150.5 tons/hr
        PriorityZones = new List<string> { "zone-alpha", "zone-beta" },
        Timestamp = DateTime.UtcNow.ToString("O")
    };

    try
    {
        await dispatcher.BroadcastStrategyAsync(strategy);
    }
    catch (Exception ex)
    {
         Console.WriteLine($"Error broadcasting strategy: {ex.Message}");
    }
    
    // Wait before the next strategic evaluation (e.g., every 5 minutes in prod, 10 seconds for testing)
    await Task.Delay(TimeSpan.FromSeconds(10));
}
