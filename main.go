package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Content string
	Sender  int
}

type ChannelState struct {
	ChannelID int
	Messages  []Message
}

type Process struct {
	ID            int
	State         string
	InChannels    []chan Message
	OutChannels   []chan Message
	MarkerSent    []bool
	ReceivedFrom  []bool
	ChannelStates []ChannelState
	Mutex         sync.Mutex
}

func NewProcess(id int, numProcesses int, bufferSize int) *Process {
	process := &Process{
		ID:            id,
		State:         "Initial",
		InChannels:    make([]chan Message, numProcesses),
		OutChannels:   make([]chan Message, numProcesses),
		MarkerSent:    make([]bool, numProcesses),
		ReceivedFrom:  make([]bool, numProcesses),
		ChannelStates: make([]ChannelState, numProcesses),
	}

	for i := range process.InChannels {
		process.InChannels[i] = make(chan Message, bufferSize)
	}

	for i := range process.OutChannels {
		process.OutChannels[i] = make(chan Message, bufferSize)
	}

	return process
}

func (p *Process) RecordState() {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	p.State = fmt.Sprintf("State at %v", time.Now().Format("15:04:05"))
	fmt.Printf("Process %v recorded state: %v\n", p.ID, p.State)
	// Record the state of each channel
	for i := range p.ChannelStates {
		p.ChannelStates[i].Messages = append(p.ChannelStates[i].Messages, Message{Content: p.State, Sender: p.ID})
	}
}

func (p *Process) SendMarker() {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	for i := range p.OutChannels {
		if !p.MarkerSent[i] {
			p.MarkerSent[i] = true
			p.OutChannels[i] <- Message{Content: "Marker", Sender: p.ID}
			fmt.Printf("Process %v sent marker to process %v\n", p.ID, i)
		}
	}
}

func (p *Process) ReceiveMarker(senderID int) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	if !p.ReceivedFrom[senderID] {
		p.ReceivedFrom[senderID] = true
		fmt.Printf("Process %v received marker from process %v\n", p.ID, senderID)
		p.RecordState()
		p.SendMarker()
	}
}

func (p *Process) SendMessage(msg Message, receiverID int) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	p.ChannelStates[receiverID].Messages = append(p.ChannelStates[receiverID].Messages, msg)
	p.OutChannels[receiverID] <- msg
	fmt.Printf("Process %v sent message '%v' to process %v\n", p.ID, msg.Content, receiverID)
}

func (p *Process) ReceiveMessage(msg Message) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	fmt.Printf("Process %v received message '%v' from process %v\n", p.ID, msg.Content, msg.Sender)
	p.ChannelStates[msg.Sender].Messages = append(p.ChannelStates[msg.Sender].Messages, msg)
}

func PrintSnapshot(processes []*Process) {
	fmt.Println("Global Snapshot:")
	for _, p := range processes {
		fmt.Printf("Process %v State: %v\n", p.ID, p.State)
		for i, chState := range p.ChannelStates {
			fmt.Printf("    Channel %v State: %v\n", i, chState.Messages)
		}
	}
}




func main() {
	numProcesses := 4
	bufferSize := 100
	processes := make([]*Process, numProcesses)

	for i := 0; i < numProcesses; i++ {
		processes[i] = NewProcess(i, numProcesses, bufferSize)
	}

	// Simulate communication
	for _, p := range processes {
		go func(proc *Process) {
			time.Sleep(1 * time.Second) // Delay to ensure p1 sends the first marker
			proc.SendMessage(Message{Content: "Hello", Sender: proc.ID}, (proc.ID+1)%numProcesses)
			proc.SendMessage(Message{Content: "Bonjour", Sender: proc.ID}, (proc.ID+2)%numProcesses)
			time.Sleep(1 * time.Second)
			proc.SendMessage(Message{Content: "Hola", Sender: proc.ID}, (proc.ID+1)%numProcesses)
			proc.SendMessage(Message{Content: "Ciao", Sender: proc.ID}, (proc.ID+2)%numProcesses)

			time.Sleep(1 * time.Second) // Delay to ensure p1 sends the first marker
			proc.SendMessage(Message{Content: "wow", Sender: proc.ID}, (proc.ID+1)%numProcesses)
			proc.SendMessage(Message{Content: "nice", Sender: proc.ID}, (proc.ID+2)%numProcesses)

			time.Sleep(1 * time.Second) // Delay to ensure p1 sends the first marker
			proc.SendMessage(Message{Content: "rishith", Sender: proc.ID}, (proc.ID+1)%numProcesses)
			proc.SendMessage(Message{Content: "hi", Sender: proc.ID}, (proc.ID+2)%numProcesses)
			proc.SendMessage(Message{Content: "naice", Sender: proc.ID}, (proc.ID+1)%numProcesses)
			proc.SendMessage(Message{Content: "da", Sender: proc.ID}, (proc.ID+2)%numProcesses)
		}(p)
	}

	// Wait for message handling and marker sending to complete
	time.Sleep(5 * time.Second)

	// Send markers after message handling is complete
	for _, p := range processes {
		p.SendMarker()
		break
	}

	// Wait for simulation to finish
	time.Sleep(1 * time.Second)

	// Print global snapshot
	PrintSnapshot(processes)

	
}
