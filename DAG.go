package main

type DAG struct {
	jobs []*Job  //定义一个有向无环图的结构
}

func NewDag() *DAG  {
	return &DAG{make([]*Job,0)}  //分配内存

}

func (dag *DAG) PreJob () *Job  {
	jobsCount := len(dag.jobs)  //处理任务数量
	if jobsCount == 0{
		return  nil
	}

	return dag.jobs[jobsCount-1]  //返回最后一个任务

}
//运行任务

func (dag *DAG)Run()  {
	for _ , job := range dag.jobs{
		run(job)
	}
}

//控制顺序执行
func (dag *DAG)PipeLine(tasks...func()) *PipeLineResult  {
	job := &Job{
		tasks:      make([]func(),len(tasks)),
		squential:  true,
	}

	for i,task := range tasks {
		job.tasks[i] = task //设定任务
	}

	dag.jobs = append(dag.jobs,job)

	//任务叠加
	return &PipeLineResult{dag}


}


//控制异步执行
func (dag *DAG)Spawn(tasks ...func()) *SpanResult  {
	job := &Job{
		tasks:      make([]func(),len(tasks)),
		squential:  false,
	}

	for i,task := range tasks {
		job.tasks[i] = task //设定任务
	}

	dag.jobs = append(dag.jobs,job)

	//任务叠加
	return &SpanResult{dag}

}


