package models

type MainSteps struct {
	ID             int64
	StepTitle      string
	MainStepResult string
	SubSteps       []SubSteps
}

type SubSteps struct {
	StepID      int64
	StepName    string
	StepCommand []Command
	StepResults string
}
