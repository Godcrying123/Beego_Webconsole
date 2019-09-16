package models

type MainTasks struct {
	ID             int    `form:"-"`
	TaskTitle      string `form:"task_name"`
	SubTasks       []EachTask
	MainTaskResult string
}

type EachTask struct {
	TaskID      int    `form:"-"`
	TaskSummary string `form:"task_summary"`
	TaskNode    string `form:"task_nodes"`
	TaskCommand string `form:"task_commands"`
	TaskResults string
}
