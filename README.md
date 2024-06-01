## Distributed System Simulation of Chandy-Lamport Algorithm

This Go program simulates a distributed system implementing the Chandy-Lamport algorithm for taking a global snapshot. The program features four processes that communicate through message passing, allowing for asynchronous communication. It captures the state of each process and the communication channels between them to form a consistent global snapshot.

### Features

- **Processes and Channels:** Four processes, each with dedicated input and output channels for message passing.
- **Message Passing:** Processes send and receive messages asynchronously to simulate real-world distributed systems.
- **Chandy-Lamport Algorithm:** Implements the Chandy-Lamport algorithm using marker messages to record the global state of the system.
- **State Recording:** Each process records its state and the state of its incoming channels when it receives a marker message.
- **Communication Simulation:** Processes send various messages to each other with simulated delays, demonstrating asynchronous communication and snapshot recording.

### Usage

1. **Compile and Run the Program:**
    ```sh
    go run main.go
    ```

2. **Output:**
    The program outputs the state of each process and their channels after sending and receiving messages, followed by the global snapshot captured by the Chandy-Lamport algorithm.

### Main Components

- **Process Struct:** Defines the structure of a process, including its ID, state, channels, and mechanisms for recording state.
- **Message Struct:** Defines the structure of a message, including its content and sender ID.
- **ChannelState Struct:** Records the state of a communication channel, including its ID and messages.
- **NewProcess Function:** Initializes a new process with specified channels and states.
- **RecordState Method:** Records the state of a process and its channels.
- **SendMarker Method:** Sends marker messages to other processes to initiate state recording.
- **ReceiveMarker Method:** Handles the reception of marker messages and triggers state recording.
- **SendMessage Method:** Sends a message to another process and records it in the channel state.
- **ReceiveMessage Method:** Handles the reception of messages and records them in the channel state.
- **PrintSnapshot Function:** Prints the global snapshot of all processes and their channel states.

### Example Output

```plaintext
Process 3 sent message 'Hello' to process 0
Process 3 sent message 'Bonjour' to process 1
Process 2 sent message 'Hello' to process 3  
Process 2 sent message 'Bonjour' to process 0
Process 1 sent message 'Hello' to process 2  
Process 1 sent message 'Bonjour' to process 3
Process 0 sent message 'Hello' to process 1  
Process 0 sent message 'Bonjour' to process 2
Process 3 sent message 'Hola' to process 0
Process 3 sent message 'Ciao' to process 1
Process 1 sent message 'Hola' to process 2
Process 1 sent message 'Ciao' to process 3
Process 2 sent message 'Hola' to process 3
Process 2 sent message 'Ciao' to process 0
Process 0 sent message 'Hola' to process 1
Process 0 sent message 'Ciao' to process 2
Process 0 sent message 'wow' to process 1
Process 0 sent message 'nice' to process 2
Process 3 sent message 'wow' to process 0
Process 3 sent message 'nice' to process 1
Process 1 sent message 'wow' to process 2
Process 1 sent message 'nice' to process 3
Process 2 sent message 'wow' to process 3
Process 2 sent message 'nice' to process 0
Process 3 sent message 'rishith' to process 0
Process 3 sent message 'hi' to process 1
Process 3 sent message 'naice' to process 0
Process 3 sent message 'da' to process 1
Process 2 sent message 'rishith' to process 3
Process 2 sent message 'hi' to process 0
Process 2 sent message 'naice' to process 3
Process 2 sent message 'da' to process 0
Process 0 sent message 'rishith' to process 1
Process 0 sent message 'hi' to process 2
Process 0 sent message 'naice' to process 1
Process 0 sent message 'da' to process 2
Process 1 sent message 'rishith' to process 2
Process 1 sent message 'hi' to process 3
Process 1 sent message 'naice' to process 2
Process 1 sent message 'da' to process 3
Process 0 sent marker to process 0
Process 0 sent marker to process 1
Process 0 sent marker to process 2
Process 0 sent marker to process 3
Global Snapshot:
Process 0 State: Initial
    Channel 0 State: []
    Channel 1 State: [{Hello 0} {Hola 0} {wow 0} {rishith 0} {naice 0}]
    Channel 2 State: [{Bonjour 0} {Ciao 0} {nice 0} {hi 0} {da 0}]
    Channel 3 State: []
Process 1 State: Initial
    Channel 0 State: []
    Channel 1 State: []
    Channel 2 State: [{Hello 1} {Hola 1} {wow 1} {rishith 1} {naice 1}]
    Channel 3 State: [{Bonjour 1} {Ciao 1} {nice 1} {hi 1} {da 1}]
Process 2 State: Initial
    Channel 0 State: [{Bonjour 2} {Ciao 2} {nice 2} {hi 2} {da 2}]
    Channel 1 State: []
    Channel 2 State: []
    Channel 3 State: [{Hello 2} {Hola 2} {wow 2} {rishith 2} {naice 2}]
Process 3 State: Initial
    Channel 0 State: [{Hello 3} {Hola 3} {wow 3} {rishith 3} {naice 3}]
    Channel 1 State: [{Bonjour 3} {Ciao 3} {nice 3} {hi 3} {da 3}]
    Channel 2 State: []
    Channel 3 State: []
```
