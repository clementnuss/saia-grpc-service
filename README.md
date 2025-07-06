# SAIA PCD gRPC Service

A gRPC service for communicating with SAIA PCD controllers, providing remote access to inputs, outputs, flags, counters, timers, and registers.

## Overview

This service implements a gRPC interface for SAIA PCD (Programmable Control Device) controllers, allowing remote monitoring and control operations. The service is built using Protocol Buffers and provides methods for:

- Reading inputs, outputs, flags, counters, timers, and registers
- Writing flags and registers
- Real-time communication with PCD controllers

## Features

- **Read Operations**: Access input states, output states, flag values, counter values, timer values, and register contents
- **Write Operations**: Modify flag states and register values
- **gRPC Interface**: High-performance, language-agnostic communication protocol
- **Protocol Buffers**: Efficient serialization and strong typing

## Prerequisites

- Python 3.7+
- Protocol Buffers compiler (`protoc`)
- Buf CLI tool for Protocol Buffer management

## Setup

1. **Create and activate a virtual environment:**
   ```bash
   python3 -m venv ~/saia/
   source ~/saia/bin/activate
   ```

2. **Configure pip for Buf's Python package index:**
   ```bash
   python -m pip config set --site global.extra-index-url https://buf.build/gen/python
   ```

3. **Install dependencies:**
   ```bash
   pip install -r requirements.txt
   ```

4. **Generate gRPC code from Protocol Buffers:**
   ```bash
   buf generate
   ```

5. **Configure SAIA server connection:**
   
   Copy the `.env` file and update the SAIA server settings:
   ```bash
   cp .env .env.local
   ```
   
   Edit `.env.local` with your SAIA server details:
   ```
   SAIA_SERVER_ADDRESS="192.168.1.100"
   SAIA_SERVER_STATION="0"
   ```

## Usage

### Starting the gRPC Server

```bash
python saia_grpc_service.py
```

The server will start on port 50051 and connect to the configured SAIA PCD controller.

### Client Examples

The generated gRPC client code can be used to interact with the service:

```python
import grpc
from gen import saia_pb2, saia_pb2_grpc

# Create a gRPC channel
channel = grpc.insecure_channel('localhost:50051')
stub = saia_pb2_grpc.SaiaPcdServiceStub(channel)

# Read a register as integer
request = saia_pb2.ReadRegisterRequest(
    address=100, 
    data_type=saia_pb2.REGISTER_DATA_TYPE_INT
)
response = stub.ReadRegister(request)
print(f"Register 100 value: {response.int_value}")

# Read a flag
request = saia_pb2.ReadFlagRequest(address=1)
response = stub.ReadFlag(request)
print(f"Flag 1 state: {response.value}")

# Write a register
request = saia_pb2.WriteRegisterRequest(
    address=100, 
    data_type=saia_pb2.REGISTER_DATA_TYPE_INT,
    int_value=42
)
response = stub.WriteRegister(request)

# Write a flag
request = saia_pb2.WriteFlagRequest(address=1, value=True)
response = stub.WriteFlag(request)

# Read counter and timer values
counter_request = saia_pb2.ReadCounterRequest(address=5)
counter_response = stub.ReadCounter(counter_request)
print(f"Counter 5 value: {counter_response.value}")

timer_request = saia_pb2.ReadTimerRequest(address=10)
timer_response = stub.ReadTimer(timer_request)
print(f"Timer 10 value: {timer_response.value}")
```

## Project Structure

```
.
├── proto/saia.proto          # Protocol Buffer definitions
├── gen/                      # Generated gRPC code
│   ├── saia_pb2.py          # Protocol Buffer messages
│   ├── saia_pb2.pyi         # Type hints
│   └── saia_pb2_grpc.py     # gRPC service definitions
├── saia_grpc_service.py     # Main gRPC service implementation
├── buf.yaml                 # Buf configuration
├── buf.gen.yaml            # Buf generation configuration
├── requirements.txt        # Python dependencies
└── .env                    # Environment configuration template
```

## API Reference

### Read Operations

- `ReadInput(address)` → `bool` - Read digital input state
- `ReadOutput(address)` → `bool` - Read digital output state  
- `ReadFlag(address)` → `bool` - Read flag value
- `ReadCounter(address)` → `int32` - Read counter value
- `ReadTimer(address)` → `int32` - Read timer value
- `ReadRegister(address, data_type)` → `int32|float` - Read register as integer or float

### Write Operations

- `WriteFlag(address, value)` - Write boolean value to flag
- `WriteRegister(address, data_type, value)` - Write integer or float value to register

### Data Types

- `REGISTER_DATA_TYPE_INT` - Read/write register as 32-bit signed integer
- `REGISTER_DATA_TYPE_FLOAT` - Read/write register as 32-bit float

## Development

### Regenerating gRPC Code

After modifying the Protocol Buffer definitions in `proto/saia.proto`, regenerate the gRPC code:

```bash
buf generate
```

### Adding New Operations

1. Update `proto/saia.proto` with new service methods and message types
2. Run `buf generate` to regenerate Python code
3. Implement the new methods in `SaiaPcdServicer` class in `saia_grpc_service.py`

## Error Handling

The service returns gRPC status codes for different error conditions:

- `INTERNAL` - Unable to communicate with SAIA controller or read/write values
- `INVALID_ARGUMENT` - Missing required fields or invalid data types

## Contributing

1. Make changes to the Protocol Buffer definitions in `proto/saia.proto`
2. Regenerate code with `buf generate`
3. Update the service implementation in `saia_grpc_service.py`
4. Test your changes thoroughly

## License

[Add your license information here]
