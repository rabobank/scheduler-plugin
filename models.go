package main

import "time"

// SchedulerPlugin is the struct implementing the interface defined by the core CLI. It can be found at  "code.cloudfoundry.org/cli/plugin/plugin.go"
type SchedulerPlugin struct{}

type GenericRequestFitsAll struct {
	SpaceGUID      string `json:"spaceguid"`
	AppGUID        string `json:"appguid,omitempty"`
	Name           string `json:"name,omitempty"`
	CronExpression string `json:"cronexpression,omitempty"`
	ExpressionType string `json:"expressiontype,omitempty"`
	Command        string `json:"command,omitempty"`
	Url            string `json:"url,omitempty"`
	AuthHeader     string `json:"authheader,omitempty"`
	ScheduleGuid   string `json:"scheduleguid,omitempty"`
}

type JobListResponse struct {
	Jobs []Job
}

type Job struct {
	JobName string `json:"jobname"`
	AppName string `json:"appname"`
	Command string `json:"command"`
}

type CallListResponse struct {
	Calls []Call
}

type Call struct {
	CallName   string `json:"callname"`
	AppName    string `json:"appname"`
	Url        string `json:"url"`
	AuthHeader string `json:"authheader"`
}

type JobScheduleListResponse struct {
	JobSchedules []JobSchedule
}

type JobSchedule struct {
	AppGuid        string `json:"appguid"`
	Name           string `json:"name"`
	Command        string `json:"command"`
	CronExpression string `json:"cronexpression"`
	ScheduleGuid   string `json:"scheduleguid"`
}

type CallScheduleListResponse struct {
	CallSchedules []CallSchedule
}

type CallSchedule struct {
	AppName        string `json:"appname"`
	Name           string `json:"name"`
	Url            string `json:"Url"`
	CronExpression string `json:"cronexpression"`
	ScheduleGuid   string `json:"scheduleguid"`
}

type HistoryRequest struct {
	SpaceGUID string `json:"spaceguid"`
	Name      string `json:"name"`
}

type HistoryListResponse struct {
	Histories []History
}

type History struct {
	Guid               string
	ScheduledTime      time.Time
	ExecutionStartTime time.Time
	ExecutionEndTime   time.Time
	Message            string
	State              string
	ScheduleGuid       string
	TaskGuid           string
	CreatedAt          time.Time
}
