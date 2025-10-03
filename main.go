package main

import (
	"code.cloudfoundry.org/cli/cf/i18n"
	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/plugin"
	plugin_models "code.cloudfoundry.org/cli/plugin/models"
	"code.cloudfoundry.org/cli/util/configv3"
	"context"
	"flag"
	"fmt"
	cfclient "github.com/cloudfoundry/go-cfclient/v3/client"
	cfconfig "github.com/cloudfoundry/go-cfclient/v3/config"
	"github.com/cloudfoundry/go-cfclient/v3/resource"
	"github.com/rabobank/scheduler-plugin/version"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	HttpTimeoutDefault = 5

	CreateJobHelpText = "creates a schedulable job"
	CreateJobUsage    = "create-job APP_NAME JOB_NAME COMMAND [MEMORY_IN_MB] [DISK_IN_MB]"
	ListJobsHelpText  = "lists all schedulable jobs"
	ListJobsUsage     = "jobs"
	DeleteJobHelpText = "deletes a schedulable job"
	DeleteJobUsage    = "delete-job [--force] JOB_NAME"
	RunJobHelpText    = "runs a job"
	RunJobUsage       = "run-job JOB_NAME"

	CreateCallHelpText = "creates a schedulable call"
	CreateCallUsage    = "create-call [--auth-header AUTH-HEADER] APP_NAME CALL_NAME URL"
	ListCallsHelpText  = "lists all schedulable calls"
	ListCallsUsage     = "calls"
	DeleteCallHelpText = "deletes a schedulable call"
	DeleteCallUsage    = "delete-call [--force] CALL_NAME"
	RunCallHelpText    = "runs a call"
	RunCallUsage       = "run-call CALL_NAME"

	CreateJobScheduleHelpText = "schedules a job"
	CreateJobScheduleUsage    = "schedule-job JOB_NAME CRON_EXPRESSION"
	ListJobSchedulesHelpText  = "lists all job schedules"
	ListJobSchedulesUsage     = "job-schedules"
	DeleteJobScheduleHelpText = "deletes a schedule for a job"
	DeleteJobScheduleUsage    = "delete-job-schedule JOB_NAME SCHEDULE_GUID"

	CreateCallScheduleHelpText = "schedules a call"
	CreateCallScheduleUsage    = "schedule-call CALL_NAME CRON_EXPRESSION"
	ListCallSchedulesHelpText  = "lists all call schedules"
	ListCallSchedulesUsage     = "call-schedules"
	DeleteCallScheduleHelpText = "deletes a schedule for a call"
	DeleteCallScheduleUsage    = "delete-call-schedule CALL_NAME SCHEDULE_GUID"

	ListJobHistoriesHelpText  = "lists the history for a scheduled job"
	ListJobHistoriesUsage     = "job-history JOB_NAME"
	ListCallHistoriesHelpText = "lists the history for a scheduled call"
	ListCallHistoriesUsage    = "call-history CALL_NAME"
)

var (
	accessToken     string
	cfClient        *cfclient.Client
	ctx             = context.Background()
	currentOrg      plugin_models.Organization
	currentSpace    plugin_models.Space
	currentUser     string
	serviceInstance resource.ServiceInstance
	requestHeader   http.Header
	httpClient      http.Client
	FlagForce       bool
	FlagAuthHeader  string
	FlagTimeout     int
	FlagMemoryInMB  int
	FlagDiskInMB    int
)

// Run must be implemented by any plugin because it is part of the plugin interface defined by the core CLI.
//
// Run(....) is the entry point when the core CLI is invoking a command defined by a plugin.
// The first parameter, plugin.CliConnection, is a struct that can be used to invoke cli commands. The second parameter, args, is a slice of strings.
// args[0] will be the name of the command, and will be followed by any additional arguments a cli user typed in.
//
// Any error handling should be handled with the plugin itself (this means printing user facing errors).
// The CLI will exit 0 if the plugin exits 0 and will exit 1 should the plugin exits nonzero.
func (c *SchedulerPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	pluginFlagSet := flag.NewFlagSet("scheduler-plugin", flag.ExitOnError)
	pluginFlagSet.BoolVar(&FlagForce, "force", false, "exit with rc=0, even if the command fails (only applicable for deletes)")
	pluginFlagSet.StringVar(&FlagAuthHeader, "auth-header", "", "the authorization header to use on the http call (only applicable for create-call)")
	pluginFlagSet.IntVar(&FlagTimeout, "timeout", HttpTimeoutDefault, "the timeout (in secs) to use on http calls")
	pluginFlagSet.IntVar(&FlagMemoryInMB, "memory_in_mb", 0, "memory in MB to use for the job (only applicable for create-job)")
	pluginFlagSet.IntVar(&FlagDiskInMB, "disk_in_mb", 0, "disk in MB to use for the job (only applicable for create-job)")
	if err := pluginFlagSet.Parse(args[1:]); err != nil {
		fmt.Printf("failed to parse arguments: %s", err)
	}

	httpClient = http.Client{Timeout: time.Duration(FlagTimeout) * time.Second}
	if args[0] != "install-plugin" && args[0] != "CLI-MESSAGE-UNINSTALL" {
		preCheck(cliConnection)
		requestHeader = map[string][]string{"Content-Type": {"application/json"}, "Authorization": {accessToken}}
	}

	switch args[0] {
	case "create-job":
		createJob(pluginFlagSet.Args())
	case "jobs":
		jobs(pluginFlagSet.Args())
	case "delete-job":
		deleteJob(pluginFlagSet.Args())
	case "create-call":
		createCall(pluginFlagSet.Args())
	case "calls":
		calls(pluginFlagSet.Args())
	case "delete-call":
		deleteCall(pluginFlagSet.Args())
	case "schedule-job":
		createJobSchedule(pluginFlagSet.Args())
	case "job-schedules":
		jobSchedules(pluginFlagSet.Args())
	case "delete-job-schedule":
		deleteJobSchedule(pluginFlagSet.Args())
	case "schedule-call":
		createCallSchedule(pluginFlagSet.Args())
	case "call-schedules":
		callSchedules(pluginFlagSet.Args())
	case "delete-call-schedule":
		deleteCallSchedule(pluginFlagSet.Args())
	case "job-history":
		jobHistories(pluginFlagSet.Args())
	case "call-history":
		callHistories(pluginFlagSet.Args())
	case "run-call":
		runCall(pluginFlagSet.Args())
	case "run-job":
		runJob(pluginFlagSet.Args())
	}
}

// GetMetadata returns a PluginMetadata struct. The first field, Name, determines the name of the plugin which should generally be without spaces.
// If there are spaces in the name a user will need to properly quote the name during uninstall otherwise the name will be treated as separate arguments.
// The second value is a slice of Command structs. Our slice only contains one Command Struct, but could contain any number of them.
// The first field Name defines the command `cf basic-plugin-command` once installed into the CLI.
// The second field, HelpText, is used by the core CLI to display help information to the user in the core commands `cf help`, `cf`, or `cf -h`.
func (c *SchedulerPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:          "Scheduler",
		Version:       plugin.VersionType{Major: version.GetMajorVersion(), Minor: version.GetMinorVersion(), Build: version.GetPatchVersion()},
		MinCliVersion: plugin.VersionType{Major: 6, Minor: 7, Build: 0},
		Commands: []plugin.Command{
			{Name: "create-job", HelpText: CreateJobHelpText, UsageDetails: plugin.Usage{Usage: CreateJobUsage}},
			{Name: "jobs", HelpText: ListJobsHelpText, UsageDetails: plugin.Usage{Usage: ListJobsUsage}},
			{Name: "delete-job", HelpText: DeleteJobHelpText, UsageDetails: plugin.Usage{Usage: DeleteJobUsage}},
			{Name: "create-call", HelpText: CreateCallHelpText, UsageDetails: plugin.Usage{Usage: CreateCallUsage}},
			{Name: "calls", HelpText: ListCallsHelpText, UsageDetails: plugin.Usage{Usage: ListCallsUsage}},
			{Name: "delete-call", HelpText: DeleteCallHelpText, UsageDetails: plugin.Usage{Usage: DeleteCallUsage}},
			{Name: "schedule-job", HelpText: CreateJobScheduleHelpText, UsageDetails: plugin.Usage{Usage: CreateJobScheduleUsage}},
			{Name: "job-schedules", HelpText: ListJobSchedulesHelpText, UsageDetails: plugin.Usage{Usage: ListJobSchedulesUsage}},
			{Name: "delete-job-schedule", HelpText: DeleteJobScheduleHelpText, UsageDetails: plugin.Usage{Usage: DeleteJobScheduleUsage}},
			{Name: "schedule-call", HelpText: CreateCallScheduleHelpText, UsageDetails: plugin.Usage{Usage: CreateCallScheduleUsage}},
			{Name: "call-schedules", HelpText: ListCallSchedulesHelpText, UsageDetails: plugin.Usage{Usage: ListCallSchedulesUsage}},
			{Name: "delete-call-schedule", HelpText: DeleteCallScheduleHelpText, UsageDetails: plugin.Usage{Usage: DeleteCallScheduleUsage}},
			{Name: "job-history", HelpText: ListJobHistoriesHelpText, UsageDetails: plugin.Usage{Usage: ListJobHistoriesUsage}},
			{Name: "call-history", HelpText: ListCallHistoriesHelpText, UsageDetails: plugin.Usage{Usage: ListCallHistoriesUsage}},
			{Name: "run-call", HelpText: RunCallHelpText, UsageDetails: plugin.Usage{Usage: RunCallUsage}},
			{Name: "run-job", HelpText: RunJobHelpText, UsageDetails: plugin.Usage{Usage: RunJobUsage}},
		},
	}
}

// preCheck Does all common validations, like being logged in, and having a targeted org and space, and if there is an instance of the scheduler-service.
func preCheck(cliConnection plugin.CliConnection) {
	config, _ := configv3.LoadConfig()
	i18n.T = i18n.Init(config)
	var schedulerService resource.ServiceInstance
	loggedIn, err := cliConnection.IsLoggedIn()
	if err != nil || !loggedIn {
		fmt.Println(terminal.NotLoggedInText())
		os.Exit(1)
	}
	currentUser, _ = cliConnection.Username()
	hasOrg, err := cliConnection.HasOrganization()
	if err != nil || !hasOrg {
		fmt.Println(terminal.FailureColor("please target your org/space first"))
		os.Exit(1)
	}
	org, _ := cliConnection.GetCurrentOrg()
	currentOrg = org
	hasSpace, err := cliConnection.HasSpace()
	if err != nil || !hasSpace {
		fmt.Println(terminal.FailureColor("please target your space first"))
		os.Exit(1)
	}
	space, _ := cliConnection.GetCurrentSpace()
	currentSpace = space
	if accessToken, err = cliConnection.AccessToken(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	CF_HOME := os.Getenv("CF_HOME")
	if CF_HOME == "" {
		CF_HOME = os.Getenv("HOME")
	}

	if cfConfig, err := cfconfig.NewFromCFHomeDir(CF_HOME); err != nil {
		log.Fatalf("failed to create new config: %s", err)
	} else {
		if cfClient, err = cfclient.New(cfConfig); err != nil {
			log.Fatalf("failed to create new cf client: %s", err)
		} else {
			if servicePlans, err := cfClient.ServicePlans.ListAll(ctx, &cfclient.ServicePlanListOptions{ServiceOfferingNames: cfclient.Filter{Values: []string{"scheduler"}}}); err != nil {
				log.Fatalf("failed to get service plans for scheduler service: %s", err)
			} else if len(servicePlans) == 0 {
				fmt.Println(terminal.FailureColor("no service plan found for scheduler service"))
				os.Exit(1)
			} else {
				schedulerServices, err := cfClient.ServiceInstances.ListAll(ctx, &cfclient.ServiceInstanceListOptions{
					ListOptions:      &cfclient.ListOptions{},
					Type:             "managed",
					SpaceGUIDs:       cfclient.Filter{Values: []string{currentSpace.Guid}},
					ServicePlanGUIDs: cfclient.Filter{Values: []string{servicePlans[0].GUID}},
				})
				if len(schedulerServices) == 0 {
					fmt.Println(terminal.FailureColor("no scheduler service instance found, please create a scheduler service instance first"))
				}
				schedulerService = *schedulerServices[0]
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
			serviceInstance = schedulerService
		}
	}
}

// Unlike most Go programs, the `Main()` function will not be used to run all of the commands provided in your plugin.
// Main will be used to initialize the plugin process, as well as any dependencies you might require for your plugin.
func main() {
	// Any initialization for your plugin can be handled here
	//
	// Note: to run the plugin.Start method, we pass in a pointer to the struct implementing the interface defined at "code.cloudfoundry.org/cli/plugin/plugin.go"
	//
	// Note: The plugin's main() method is invoked at install time to collect metadata. The plugin will exit 0 and the Run([]string) method will not be invoked.
	plugin.Start(new(SchedulerPlugin))
	// Plugin code should be written in the Run([]string) method, ensuring the plugin environment is bootstrapped.
}
