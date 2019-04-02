# health-check

This is an application that periodically checks the health of various websites

## Start

You can start the whole stack using docker-compose  
`docker-compose up`  
This starts the web server, cron job to perform regular health checks and mongodb server

After this head on to [webpage](http://localhost:9000/web/)

## Note

The cron job is set to run every 5 minutes. You may change it in [here](https://github.com/abhiche/health-check/blob/master/internal/prober/cronjobs)

## Todo

Improve unit test coverage  
Add integration tests
