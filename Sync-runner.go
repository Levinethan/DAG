package main


//顺序

func RunSync(job *Job)  {
	for _ , task := range job.tasks{
		task() //顺序调用
	}

	if job.onComplete != nil{
		job.onComplete()  //完成的时候调用  控制异步执行
	}

	//同步执行
}