package worker
import "github.com/aniruddhrane/gotasker/internal/queue"
import "fmt"
type Worker struct{
	ID int
	Queue *queue.Queue
}
// worker's behavior with a Start() method that runs forever, 
// dequeues jobs, and processes them

func (w*Worker) Start(){
	for{
       job := w.Queue.Dequeue()
       fmt.Printf("Worker %d processing job %d \n",w.ID,job.ID)
	}
}