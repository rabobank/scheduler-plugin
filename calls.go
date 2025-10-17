package main

import (
	"bytes"
	"code.cloudfoundry.org/cli/cf/terminal"
	"encoding/json"
	"fmt"
	cfclient "github.com/cloudfoundry/go-cfclient/v3/client"
	"io"
	"net/http"
	"net/url"
	"os"
)

func createCall(args []string) {
	if len(args) != 3 && len(args) != 5 {
		fmt.Printf("Incorrect Usage: the required arguments are `APP_NAME` and `CALL_NAME` and `URL`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", CreateCallHelpText, CreateCallUsage)
		os.Exit(1)
	}
	app, err := cfClient.Applications.Single(ctx, &cfclient.AppListOptions{ListOptions: &cfclient.ListOptions{}, Names: cfclient.Filter{Values: []string{args[0]}}, SpaceGUIDs: cfclient.Filter{Values: []string{currentSpace.Guid}}})
	if err != nil {
		fmt.Printf("app lookup for %s in space %s returned error: %s\n", args[0], currentSpace.Name, err)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{AppGUID: app.GUID, SpaceGUID: currentSpace.Guid, Name: args[1], Url: args[2], AuthHeader: FlagAuthHeader})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/calls", *serviceInstance.DashboardURL))
	httpRequest := http.Request{Method: http.MethodPost, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to create call, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func runCall(args []string) {
	if len(args) != 1 {
		fmt.Printf("Incorrect Usage: the required arguments are `CALL_NAME`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", RunCallHelpText, RunCallUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/calls", *serviceInstance.DashboardURL))
	httpRequest := http.Request{Method: http.MethodPut, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to run call, reponse code %d, response: %s\n", resp.StatusCode, body)
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}

func calls(args []string) {
	if len(args) != 0 {
		fmt.Printf("Incorrect Usage: there should be no arguments to this command`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", ListCallsHelpText, ListCallsUsage)
		os.Exit(1)
	}
	request := GenericRequestFitsAll{SpaceGUID: currentSpace.Guid}
	requestBody, _ := json.Marshal(request)
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/calls", *serviceInstance.DashboardURL))
	httpRequest := http.Request{Method: http.MethodGet, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	fmt.Printf("Getting scheduled calls for org %s / space %s as %s\n\n", terminal.AdvisoryColor(currentOrg.Name), terminal.AdvisoryColor(currentSpace.Name), terminal.AdvisoryColor(currentUser))
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("response (%d) from scheduler service: %s", resp.StatusCode, body)))
		os.Exit(1)
	}
	jsonResponse := CallListResponse{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed to parse response: %s", err)))
	}
	table := terminal.NewTable([]string{"Call Name", "App Name", "Url", "Auth Header"})
	for _, call := range jsonResponse.Calls {
		table.Add(call.CallName, call.AppName, call.Url, call.AuthHeader)
	}
	_ = table.PrintTo(os.Stdout)
}

func deleteCall(args []string) {
	if len(args) != 1 {
		fmt.Printf("Incorrect Usage: the required arguments are `JOB_NAME`\n\nNAME:\n   %s\n\nUSAGE:\n   %s\n", DeleteCallHelpText, DeleteCallUsage)
		os.Exit(1)
	}
	requestBody, _ := json.Marshal(GenericRequestFitsAll{SpaceGUID: currentSpace.Guid, Name: args[0]})
	requestUrl, _ := url.Parse(fmt.Sprintf("%s/api/calls", *serviceInstance.DashboardURL))
	httpRequest := http.Request{Method: http.MethodDelete, URL: requestUrl, Header: requestHeader, Body: io.NopCloser(bytes.NewReader(requestBody))}
	resp, err := httpClient.Do(&httpRequest)
	if err != nil {
		fmt.Println(terminal.FailureColor(fmt.Sprintf("failed response from scheduler service: %s", err)))
		os.Exit(1)
	}
	if resp != nil {
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("failed to delete call, reponse code %d, response: %s\n", resp.StatusCode, body)
			if FlagForce {
				os.Exit(0)
			}
			os.Exit(1)
		} else {
			body, _ := io.ReadAll(resp.Body)
			fmt.Printf("%s\n", body)
			fmt.Println(terminal.SuccessColor("OK"))
		}
	}
}
