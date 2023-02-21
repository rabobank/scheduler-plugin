package main

import (
	"bytes"
	"code.cloudfoundry.org/cli/cf/terminal"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func createJobSchedule(args []string) {
	if len(args) != 2 {
		fmt.Printf("Incorrect Usage: the required arguments are `JOB_NAME` and `CRON_EXPRESSION`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", CreateJobScheduleHelpText, CreateJobScheduleUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0], CronExpression: args[1]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobschedules", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodPost, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to create job schedule, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func jobSchedules(args []string) {
	if len(args) != 0 {
		fmt.Printf("Incorrect Usage: there should be no arguments to this command`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", ListJobSchedulesHelpText, ListJobSchedulesUsage)
		os.Exit(1)
	}
	request := GenericRequestFitsAll{SpaceGUID: currentSpace.Guid}
	requestBody, _ := json.Marshal(request)
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobschedules", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodGet, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	fmt.Printf("Getting job schedules for org %s / space %s as %s\n\n", terminal.AdvisoryColor(currentOrg.Name), terminal.AdvisoryColor(currentSpace.Name), terminal.AdvisoryColor(currentUser))
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to list job schedules: %s", err)))
		os.Exit(1)
	}
	body, _ := io.ReadAll(resp.Body)
	jsonResponse := JobScheduleListResponse{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to parse response: %s", err)))
	}
	table := terminal.NewTable([]string{"Job Name", "Command", "Cron Expression", "Schedule Guid"})
	for _, jobSchedule := range jsonResponse.JobSchedules {
		table.Add(jobSchedule.Name, jobSchedule.Command, jobSchedule.CronExpression, jobSchedule.ScheduleGuid)
	}
	_ = table.PrintTo(os.Stdout)
}

func deleteJobSchedule(args []string) {
	if len(args) != 2 {
		fmt.Printf("Incorrect Usage: the required arguments are `JOB_NAME` and `SCHEDULE_GUID`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", DeleteJobScheduleHelpText, DeleteJobScheduleUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0], ScheduleGuid: args[1]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobschedules", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodDelete, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to delete job schedule, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func createCallSchedule(args []string) {
	if len(args) != 2 {
		fmt.Printf("Incorrect Usage: the required arguments are `CALL_NAME` and `CRON_EXPRESSION`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", CreateCallScheduleHelpText, CreateCallScheduleUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0], CronExpression: args[1]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/callschedules", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodPost, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to create call schedule, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func callSchedules(args []string) {
	if len(args) != 0 {
		fmt.Printf("Incorrect Usage: there should be no arguments to this command`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", ListCallSchedulesHelpText, ListCallSchedulesUsage)
		os.Exit(1)
	}
	request := GenericRequestFitsAll{SpaceGUID: currentSpace.Guid}
	requestBody, _ := json.Marshal(request)
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/callschedules", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodGet, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	fmt.Printf("Getting call schedules for org %s / space %s as %s\n\n", terminal.AdvisoryColor(currentOrg.Name), terminal.AdvisoryColor(currentSpace.Name), terminal.AdvisoryColor(currentUser))
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to list call schedules: %s", err)))
		os.Exit(1)
	}
	body, _ := io.ReadAll(resp.Body)
	jsonResponse := CallScheduleListResponse{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to parse response: %s", err)))
	}
	table := terminal.NewTable([]string{"Call Name", "Command", "Cron Expression", "Schedule Guid"})
	for _, callSchedule := range jsonResponse.CallSchedules {
		table.Add(callSchedule.Name, callSchedule.Url, callSchedule.CronExpression, callSchedule.ScheduleGuid)
	}
	_ = table.PrintTo(os.Stdout)
}

func deleteCallSchedule(args []string) {
	if len(args) != 2 {
		fmt.Printf("Incorrect Usage: the required arguments are `CALL_NAME` and `SCHEDULE_GUID`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", DeleteCallScheduleHelpText, DeleteCallScheduleUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0], ScheduleGuid: args[1]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/callschedules", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodDelete, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to delete call schedule, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}
