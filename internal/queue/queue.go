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
//dequee removed after implementing receiving channel

func(q*Queue)Close(){
   close(q.jobs)
}

//implement a receiving channel so woker can see that channel is emply
//receiving channel so because they have limited access to our main channel
//and dont disturb it
func(q*Queue) Jobs() <-chan job.Job{
	return q.jobs
}