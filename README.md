# Stablehand [![Build Status](https://drone.sierrasoftworks.com/api/badges/SierraSoftworks/stablehand/status.svg)](https://drone.sierrasoftworks.com/SierraSoftworks/stablehand)
**A cleanup tool for Rancher servers**

Stablehand is a tool designed to keep your Rancher server in a nice, clean, maintanable
state under production conditions (read: when an autoscaling group is destroying your hosts).

It is intended to be run on a `cron` schedule to deactivate, remove and purge hosts which meet
certain conditions. You would, for example, run it with the following `crontab`.

```crontab
# Deactivate hosts which are in the reconnecting state every 15 minutes
*/15 * * * * stablehand deactivate --agent-state=reconnecting

# Remove hosts which are in the inactive state on the hour
0 * * * * stablehand remove --state=inactive
```

## Docker

Stablehand is also available as a Docker image at `sierrasoftworks/stablehand`, allowing you to run this as a scheduled Docker container on your Rancher infrastructure for added-meta. When doing so, it may be easiest to use an environment file to hold your
config.

```
CATTLE_URL=https://localhost
CATTLE_ACCESS_KEY=abcdefg
CATTLE_SECRET_KEY=abcdefg
```

```
# Deactivate hosts which are in the reconnecting state
docker run --rm --env-file stablehand.env sierrasoftworks/stablehand stablehand deactivate --agent-state=reconnecting

# Remove hosts which are in the inactive state
docker run --rm --env-file stablehand.env sierrasoftworks/stablehand stablehand remove --state=inactive
```

### On Rancher
If you're running Stablehand on Rancher, you'll want to set the following labels to have Rancher automatically populate the `CATTLE_URL`,
`CATTLE_ACCESS_KEY` and `CATTLE_SECRET_KEY` environment variables.

```yml
labels:
    io.rancher.container.create_agent: 'true'
    io.rancher.container.agent.role: environment
```

A typical example for the `docker-compose.yml` on Rancher would look like the following, it makes use of [rancher-cron](https://github.com/SocialEngine/rancher-cron) to provide scheduled tasks and will run the cleanup
tasks once every 15m to 1h. You can easily change this by modifying the corresponding `com.socialengine.rancher-cron.schedule`
label on each service.

- Use `scale: 1` for each of them in `rancher-compose.yml`

```yml
cron:
  labels:
    io.rancher.container.create_agent: 'true'
    io.rancher.scheduler.affinity:host_label: role=worker
    io.rancher.container.agent.role: environment
  image: socialengine/rancher-cron:0.2.0
stablehand-deactivate-reconnecting:
  labels:
    io.rancher.container.pull_image: always
    io.rancher.container.start_once: 'true'
    io.rancher.container.create_agent: 'true'
    io.rancher.container.agent.role: environment
    com.socialengine.rancher-cron.schedule: '@every 15m'
  entrypoint:
  - /stablehand
  command:
  - deactivate
  - --agent-state=reconnecting
  image: sierrasoftworks/stablehand
stablehand-remove-inactive:
  labels:
    io.rancher.container.pull_image: always
    io.rancher.container.start_once: 'true'
    io.rancher.container.create_agent: 'true'
    io.rancher.container.agent.role: environment
    com.socialengine.rancher-cron.schedule: '@every 1h'
  entrypoint:
  - /stablehand
  command:
  - remove
  - --state=inactive
  image: sierrasoftworks/stablehand
```

## `man stablehand`

```
NAME:
   Stablehand - A tool to help you keep your Rancher server clean in production
USAGE:
   debug [global options] command [command options] [arguments...]
   
VERSION:
   1.0.0-dev.000000
   
AUTHOR(S):
   Benjamin Pannell <admin@sierrasoftworks.com> 
   
COMMANDS:
     list        
     deactivate  [HOST_ID]
     remove      [HOST_ID]
     purge       [HOST_ID]
     help, h     Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --server value  The URL of your Rancher server [%CATTLE_URL%]
   --key value     The access key used to sign into your Rancher server [%CATTLE_ACCESS_KEY%]
   --secret value  The secret key used to sign into your Rancher server [%CATTLE_SECRET_KEY%]
   --help, -h      show help
   --version, -v   print the version
   
COPYRIGHT:
   Sierra Softworks © 2016
```
