package main

import (
	"fmt"
	"time"

	"github.com/KRVPerera/DistributedJobRaft/model"
)

// CreateSocket creates a socket for communication in the distributed system
func CreateSocket() {
	// Implement socket creation logic here
}

// SubmitJob submits a job to the job queue
func SubmitJob(job model.Job) {
	// Implement job submission logic here
}

// ProcessJobs processes the job queue periodically
func ProcessJobs() {
	// Implement job processing logic here
}

func main() {
	// Create a socket for communication
	// Create a new node
	// node := &model.Node{}

	// CreateSocket()

	// // Create a socket for communication
	// socket := &model.Socket{}

	// Start the node's message processing routine
	// go node.ProcessMessages()

	// // Start the node's job processing routine
	// go node.ProcessJobs()

	// // Start the timed task
	// go node.TimedTask()

	// Start the REST API server for job submission
	// Implement REST API server logic here

	// Keep the main goroutine running
	for {
		fmt.Println("Hello world")
		time.Sleep(time.Second)
	}

	// Start a goroutine to process jobs periodically
	// go func() {
	// 	for {
	// 		ProcessJobs()
	// 		time.Sleep(time.Second) // Adjust the time interval as per your requirements
	// 	}
	// }()

	// // Start the main loop to handle messages
	// for {
	// 	// Receive a message from the distributed system
	// 	msg := ReceiveMessage()

	// 	// Process the received message
	// 	ProcessMessage(msg)
	// }

}
