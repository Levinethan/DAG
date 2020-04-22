package main

import "sync"

func RunAsync(job *Job)  {
	wg := &sync.WaitGroup{} //等待
	wg.Add(len(job.tasks))  //等待的任务数量
	 for _ , task := range job.tasks{
	 	go func(task func()) {
	 		task()  //并发执行
	 		wg.Done() //批量执行任务
		}(task)
	 }
	 wg.Wait()
	if job.onComplete != nil{
		job.onComplete()  //完成的时候调用  控制异步执行
	}
}