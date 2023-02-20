package main

import (
	"bytes"
	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/plugin"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func createJob(cliConnection plugin.CliConnection, args []string) {
	if len(args) != 3 {
		fmt.Printf("Incorrect Usage: the required arguments are `APP_NAME` and `JOB_NAME` and `COMMAND`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", CreateJobHelpText, CreateJobUsage)
		os.Exit(1)
	}
	app, err := cliConnection.GetApp(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{AppGUID: app.Guid, SpaceGUID: currentSpace.Guid, Name: args[1], Command: args[2]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobs", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodPost, URL: requestUrl, Header: requestHeader, Body: ioutil.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusCreated {
			body, _ := io.readAll(resp.Body)
			fmt.Printf("failed to create job, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.readAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func runJob(args []string) {
	if len(args) != 1 {
		fmt.Printf("Incorrect Usage: the required arguments are `JOB_NAME`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", RunJobHelpText, RunJobUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobs", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodPut, URL: requestUrl, Header: requestHeader, Body: ioutil.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, _ := io.readAll(resp.Body)
			fmt.Printf("failed to run job, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.readAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func jobs(args []string) {
	if len(args) != 0 {
		fmt.Printf("Incorrect Usage: there should be no arguments to this command`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", ListJobsHelpText, ListJobsUsage)
		os.Exit(1)
	}
	request := GenericRequestFitsAll{SpaceGUID: currentSpace.Guid}
	requestBody, _ := json.Marshal(request)
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobs", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodGet, URL: requestUrl, Header: requestHeader, Body: ioutil.NopCloser(bytes.NewReader(requestBody))}
	fmt.Printf("Getting scheduled jobs for org %s / space %s as %s\n\n", terminal.AdvisoryColor(currentOrg.Name), terminal.AdvisoryColor(currentSpace.Name), terminal.AdvisoryColor(currentUser))
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to list job: %s", err)))
		os.Exit(1)
	}
	body, _ := io.readAll(resp.Body)
	jsonResponse := JobListResponse{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to parse response: %s", err)))
	}
	table := terminal.NewTable([]string{"Job Name", "App Name", "Command"})
	for _, job := range jsonResponse.Jobs {
		table.Add(job.JobName, job.AppName, job.Command)
	}
	_ = table.PrintTo(os.Stdout)
}

func deleteJob(args []string) {
	if len(args) != 1 {
		fmt.Printf("Incorrect Usage: the required arguments are `JOB_NAME`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", DeleteJobHelpText, DeleteJobUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/jobs", serviceInstance.DashboardUrl))
	httpRequest := http.Request{Method: http.MethodDelete, URL: requestUrl, Header: requestHeader, Body: ioutil.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, _ := io.readAll(resp.Body)
			fmt.Printf("failed to delete job, reponse code %d, response: %s\n", resp.StatusCode, body)
			if FlagForce {
				os.Exit(0)
			}
			os.Exit(1)
		} else {
			body, _ := io.readAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}
