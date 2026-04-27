import pytest
import yaml
import json
from jsonschema import validate, ValidationError

# Load the AsyncAPI specification
def load_contract():
    with open('contracts/asyncapi.yaml', 'r') as f:
        return yaml.safe_load(f)

CONTRACT = load_contract()

def get_message_schema(message_name):
    return CONTRACT['components']['messages'][message_name]['payload']

def test_telemetry_status_schema():
    schema = get_message_schema('TelemetryStatus')
    valid_payload = {
        "vehicleId": "TRK-001",
        "timestamp": "2026-04-26T10:00:00Z",
        "metrics": {
            "fuelLevel": 85.5,
            "engineTemp": 92.1,
            "tirePressure": 105.0
        },
        "currentTask": "HAUL_ORE"
    }
    # Should not raise ValidationError
    validate(instance=valid_payload, schema=schema)

    invalid_payload = {
        "vehicleId": "TRK-001",
        "metrics": { "fuelLevel": -10 } # Invalid fuel level
    }
    with pytest.raises(ValidationError):
        validate(instance=invalid_payload, schema=schema)

def test_tactical_assignment_schema():
    schema = get_message_schema('TacticalAssignment')
    valid_payload = {
        "commandId": "CMD-123",
        "supervisorId": "SUP-01",
        "action": "SCALE_UP",
        "parameters": {
            "workerCount": 5
        }
    }
    validate(instance=valid_payload, schema=schema)

    invalid_payload = {
        "action": "INVALID_ACTION"
    }
    with pytest.raises(ValidationError):
        validate(instance=invalid_payload, schema=schema)

def test_operator_command_schema():
    schema = get_message_schema('OperatorCommand')
    valid_payload = {
        "commandId": "CMD-999",
        "operatorId": "OP-77",
        "action": "ASSIGN_ROUTE",
        "targetDestination": {
            "lat": -22.95,
            "lon": 14.31
        }
    }
    validate(instance=valid_payload, schema=schema)
