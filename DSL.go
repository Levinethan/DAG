package main

type PipeLineResult struct {
	dag *DAG  //有向无环图
}

type PipeLineDSL struct {
	dag *DAG
}

func (result *PipeLineResult)Then() *PipeLineDSL  {
	return  &PipeLineDSL{result.dag} //开启一个任务序列
}

func (result *PipeLineResult )OnComplete (do func())*PipeLineResult {
	job := result.dag.PreJob()
	if job!=nil{
		job.onComplete = do
	}
	return result  //设置完成以后执行
}

func (dsl *PipeLineDSL)Spawn(tasks ...func()) *SpanResult  {
	dsl.dag.Spawn(tasks...) //并发执行任务
	return  &SpanResult{dsl.dag}
}

type SpanResult struct {
	dag *DAG
}

func (result *SpanResult)Join ()*SpawnDSL  {
	return &SpawnDSL{result.dag}
}

type SpawnDSL struct {
	dag *DAG
}
func (result *SpanResult )OnComplete (do func())*SpanResult {
	job := result.dag.PreJob()
	if job!=nil{
		job.onComplete = do
	}
	return result  //设置完成以后执行
}

func (dsl *SpawnDSL)PipeLine(tasks...func()) *PipeLineResult  {

	dsl.dag.PipeLine(tasks...)
	return &PipeLineResult{dsl.dag}
}