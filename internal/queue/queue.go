package queue
import "github.com/aniruddhrane/gotasker/internal/job"
type Queue struct{
	jobs chan job.Job
}

func NewQueue(size int)*Queue{
  q:= Queue{
	jobs:make(chan job.Job,size),
  }
  return &q
}

func (q*Queue) Enqueue(job job.Job){
    q.jobs <- job
	// It sends the job value into the jobs channel. If the channel has available buffer space, 
	// the send completes immediately. If the buffer is full, the sending goroutine blocks until space becomes available
}

func(q*Queue) Dequeue() job.Job{
	 value := <- q.jobs
    return value
	//the worker waits till the channel has something when it has something it can remove it
}