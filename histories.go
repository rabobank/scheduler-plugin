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
	"time"
)

func jobHistories(args []string) {
	if len(args) != 1 {
		fmt.Printf("Incorrect Usage: the required arguments are `JOB_NAME`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", ListJobHistoriesHelpText, ListJobHistoriesUsage)
		os.Exit(1)
	}
	histories(args, "job")
}

func callHistories(args []string) {
	if len(args) != 1 {
		fmt.Printf("Incorrect Usage: the required arguments are `CALL_NAME`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", ListCallHistoriesHelpText, ListCallHistoriesUsage)
		os.Exit(1)
	}
	histories(args, "call")
}

func histories(args []string, jobOrCall string) {
	request := HistoryRequest{SpaceGUID: currentSpace.Guid, Name: args[0]}
	requestBody, _ := json.Marshal(request)
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/%shistories", serviceInstance.DashboardURL, jobOrCall))
	httpRequest := http.Request{Method: http.MethodGet, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	fmt.Printf("Getting job/call history for org %s / space %s as %s\n\n", terminal.AdvisoryColor(currentOrg.Name), terminal.AdvisoryColor(currentSpace.Name), terminal.AdvisoryColor(currentUser))
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	body, _ := io.ReadAll(resp.Body)
	jsonResponse := HistoryListResponse{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to parse response: %s", err)))
	}
	table := terminal.NewTable([]string{"Execution GUID", "Execution State", "Scheduled Time", "Execution Start Time", "Execution End Time", "Exit Message"})
	for _, hist := range jsonResponse.Histories {
		table.Add(hist.Guid, hist.State, hist.ScheduledTime.Format(time.RFC3339), hist.ExecutionStartTime.Format(time.RFC3339), hist.ExecutionEndTime.Format(time.RFC3339), hist.Message)
	}
	_ = table.PrintTo(os.Stdout)
}
