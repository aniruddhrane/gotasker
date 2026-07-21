package queue
import "github.com/aniruddhrane/gotasker/internal/job"
type Queue struct{
	jobs chan job.Job
}
