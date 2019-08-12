package models

type MainSteps struct {
	ID             int    `form:"-"`
	StepTitle      string `form:"main_step"`
	SubSteps       []SubSteps
	MainStepResult string
}

type SubSteps struct {
	StepID      int    `form:"-"`
	StepName    string `form:"step_name"`
	StepSummary string `form:"step_summary"`
	StepCommand string `form:"step_command"`
	StepResults string
}
