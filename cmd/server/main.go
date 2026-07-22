package main

import (
	"sync"

	"github.com/aniruddhrane/gotasker/internal/job"
	"github.com/aniruddhrane/gotasker/internal/queue"
	"github.com/aniruddhrane/gotasker/internal/worker"
)
 func main(){
  var wg sync.WaitGroup

   q :=queue.NewQueue(10)
   w1 :=worker.Worker{
	 ID:1,
	 Queue: q,
   Wg:&wg,
   }
   w2 :=worker.Worker{
	 ID:2,
	 Queue:q,
   Wg:&wg,
   }
   w3 :=worker.Worker{
	 ID:3,
	 Queue: q,
   Wg:&wg,
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
   
   wg.Add(1)
   go w1.Start()
   wg.Add(1)
   go w2.Start()
   wg.Add(1)
   go w3.Start()

   q.Enqueue(j1)
   q.Enqueue(j2)
   q.Close()
   wg.Wait()   
 }