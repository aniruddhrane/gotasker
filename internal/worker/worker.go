package worker

import (
	"fmt"
	"sync"

	"github.com/aniruddhrane/gotasker/internal/queue"
)
type Worker struct{
	ID int
	Queue *queue.Queue
	Wg *sync.WaitGroup
}
// worker's behavior with a Start() method that runs forever, 
// dequeues jobs, and processes them
//#update we removed dequeue() because the range can search for available job
//and also process it through receving channel

func (w*Worker) Start(){
	defer w.Wg.Done()
	for job:=range w.Queue.Jobs(){
       fmt.Printf("Worker %d processing job %d \n",w.ID,job.ID)
	}
}

