package main

//根据任务性质 决定顺序或异步执行

func run(job *Job)  {
	if job.squential{
		RunSync(job)
	}else {
		RunAsync(job)
	}
}