# Stablehand
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
RANCHER_SERVER=https://localhost
RANCHER_ACCESS_KEY=abcdefg
RANCHER_SECRET_KEY=abcdefg
```

```
# Deactivate hosts which are in the reconnecting state
docker run --rm --env-file stablehand.env sierrasoftworks/stablehand stablehand deactivate --agent-state=reconnecting

# Remove hosts which are in the inactive state
docker run --rm --env-file stablehand.env sierrasoftworks/stablehand stablehand remove --state=inactive
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
   --server value  The URL of your Rancher server [%RANCHER_SERVER%]
   --key value     The access key used to sign into your Rancher server [%RANCHER_ACCESS_KEY%]
   --secret value  The secret key used to sign into your Rancher server [%RANCHER_SECRET_KEY%]
   --help, -h      show help
   --version, -v   print the version
   
COPYRIGHT:
   Sierra Softworks © 2016
```
