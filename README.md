### Scheduler cf cli plugin

This plugin should be used in concert with the [Scheduler Service Broker](https://github.com/rabobank/scheduler-service-broker).  
It allows for running and scheduling calls and jobs on Cloud Foundry

Install the plugins as usual with _cf install-plugin <plugin binary>_

The _cf plugins_ command will show which subcommands are available: 


<pre>
Scheduler                1.0.14    call-history                            lists the history for a scheduled call  
Scheduler                1.0.14    call-schedules                          lists all call schedules  
Scheduler                1.0.14    calls                                   lists all schedulable calls  
Scheduler                1.0.14    create-call                             creates a schedulable call  
Scheduler                1.0.14    create-job                              creates a schedulable job  
Scheduler                1.0.14    delete-call                             deletes a schedulable call  
Scheduler                1.0.14    delete-call-schedule                    deletes a schedule for a call  
Scheduler                1.0.14    delete-job                              deletes a schedulable job  
Scheduler                1.0.14    delete-job-schedule                     deletes a schedule for a job  
Scheduler                1.0.14    job-history                             lists the history for a scheduled job  
Scheduler                1.0.14    job-schedules                           lists all job schedules  
Scheduler                1.0.14    jobs                                    lists all schedulable jobs  
Scheduler                1.0.14    run-call                                runs a call  
Scheduler                1.0.14    run-job                                 runs a job  
Scheduler                1.0.14    schedule-call                           schedules a call  
Scheduler                1.0.14    schedule-job                            schedules a job  
</pre>

In addition there are a few flags you can use:

* --force - Can be used with _delete-call_ and _delete-job_, the exit code will be 0 regardless the delete result.
* --auth-header - Can be used with the _create-call_, the authorization header to use on the http call.
* --timeout - Can be used on all subcommands, the timeout in seconds that is used for all interactions with the scheduler service broker.
* --memory_in_mb - Can be used with _create-job_, the amount of memory in MB to use for the job.
* --disk_in_mb - Can be used with _create-job_, the amount of disk in MB to use for the job.

The flags should be used just after the subcommand, for example:
````
cf create-job --memory_in_mb=42 --disk_in_mb 67 myapp job4711 "cat /proc/cpuinfo" 
````
