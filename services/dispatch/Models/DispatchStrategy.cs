using System.Text.Json.Serialization;

namespace DispatchAgent.Models;

public class DispatchStrategy
{
    [JsonPropertyName("intentId")]
    public string IntentId { get; set; } = string.Empty;

    [JsonPropertyName("targetThroughput")]
    public double TargetThroughput { get; set; }

    [JsonPropertyName("priorityZones")]
    public List<string> PriorityZones { get; set; } = new();

    [JsonPropertyName("timestamp")]
    public string Timestamp { get; set; } = string.Empty;
}
