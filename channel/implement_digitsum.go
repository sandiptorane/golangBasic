//this program is og worker pool
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct{
	id int
	random_num int
}

type Result struct{
	jobs Job
	digitsum int
}

var jobs = make(chan Job,10)
var results = make(chan Result,10)

func digits(num int) int{   //do  the sum of all digits
	sum :=0
	no :=num
	for no!=0{
		digit := no%10
		sum += digit
		no = no/10
	}
	time.Sleep(2*time.Second)
	return sum
}

func  allocate(noOfjobs int){
	for i := 0;i< noOfjobs;i++{
		randomnum := rand.Intn(999)
		job := Job{i,randomnum}
		jobs <- job //add to  jobs channel
	}
	close(jobs)
}

func worker(wg *sync.WaitGroup){
	for job := range jobs{
		output := Result{job,digits(job.random_num)}
		results<-output
	}
	wg.Done()
}

func createWorkerPool(noOfWorker int){
	var wg sync.WaitGroup
	for i:=0;i<noOfWorker;i++{
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func result(done chan bool){
	for result := range results{
		fmt.Printf("id = %d , random number = %d, sum of digits = %d\n",result.jobs.id,result.jobs.random_num,result.digitsum)
	}
	done <-true
}

func main(){
	starttime := time.Now()
	noOfJobs :=100
	go allocate(noOfJobs)
	done := make(chan bool)
	go  result(done)
	noOfworker := 10
	go createWorkerPool(noOfworker)
	<-done
	endtime :=time.Now()
	diff := endtime.Sub(starttime)
	fmt.Println("total time taken =" ,diff)
}

