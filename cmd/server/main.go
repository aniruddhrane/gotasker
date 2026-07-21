package main
import(
 "time"
 "github.com/aniruddhrane/gotasker/internal/queue"
  "github.com/aniruddhrane/gotasker/internal/worker"
  "github.com/aniruddhrane/gotasker/internal/job"

 )
 func main(){
   q :=queue.NewQueue(10)
   w1 :=worker.Worker{
	 ID:1,
	 Queue: q,
   }
   w2 :=worker.Worker{
	 ID:2,
	 Queue:q,
   }
   w3 :=worker.Worker{
	 ID:3,
	 Queue: q,
   }

   j1:=job.Job{
	  ID :1,
	  Type:"email",
	  Payload:"request for job",
   }
   j2:=job.Job{
	ID :2,
	Type:"promeheus_log",
	Payload:"check the health of the system",
   }
   

   go w1.Start()
   go w2.Start()
   go w3.Start()
   q.Enqueue(j1)
   q.Enqueue(j2)
   time.Sleep(5 * time.Second)
   
 }