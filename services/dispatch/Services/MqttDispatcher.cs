using System.Text.Json;
using DispatchAgent.Models;
using MQTTnet;
using MQTTnet.Client;

namespace DispatchAgent.Services;

public class MqttDispatcher
{
    private IMqttClient _mqttClient;
    private MqttClientOptions _options;

    public MqttDispatcher(string brokerUrl, string clientId)
    {
        var factory = new MqttFactory();
        _mqttClient = factory.CreateMqttClient();
        
        _options = new MqttClientOptionsBuilder()
            .WithClientId(clientId)
            .WithTcpServer(brokerUrl)
            .Build();
    }

    public async Task ConnectAsync()
    {
        try
        {
            await _mqttClient.ConnectAsync(_options, CancellationToken.None);
            Console.WriteLine("Connected to Cloud MQTT Broker.");
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Failed to connect to broker: {ex.Message}");
            throw;
        }
    }

    public async Task BroadcastStrategyAsync(DispatchStrategy strategy)
    {
        if (!_mqttClient.IsConnected)
        {
            Console.WriteLine("MQTT Client not connected. Cannot broadcast.");
            return;
        }

        string payload = JsonSerializer.Serialize(strategy);
        
        var message = new MqttApplicationMessageBuilder()
            .WithTopic("c2/strategy/dispatch")
            .WithPayload(payload)
            .WithQualityOfServiceLevel(MQTTnet.Protocol.MqttQualityOfServiceLevel.AtLeastOnce)
            .Build();

        await _mqttClient.PublishAsync(message, CancellationToken.None);
        Console.WriteLine($"Broadcasted Strategy [{strategy.IntentId}] with Target Throughput: {strategy.TargetThroughput}");
    }
}
