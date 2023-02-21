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