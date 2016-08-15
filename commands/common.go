package commands

import (
	"github.com/SierraSoftworks/stablehand/config"
	"github.com/rancher/go-rancher/client"
	"github.com/urfave/cli"
)

func getClient() (*client.RancherClient, error) {
	return client.NewRancherClient(&client.ClientOpts{
		Url:       config.Rancher.Server,
		AccessKey: config.Rancher.AccessKey,
		SecretKey: config.Rancher.SecretKey,
	})
}

type filterContext struct {
	CLI     *cli.Context
	Filters map[string]interface{}
}

type mapContext struct {
	CLI     *cli.Context
	Host    *client.Host
	Rancher *client.RancherClient
}

func buildMapHostsFunction(cmd cli.Command, filter func(c *filterContext), apply func(c *mapContext) error) cli.Command {
	cmd.Flags = append(
		cmd.Flags,
		cli.StringFlag{
			Name:  "state",
			Usage: "All hosts which are currently in the specified state.\nOne of active|reconnecting|deactivating|inactive|purging|purged|removing|removed",
		})

	cmd.Action = func(c *cli.Context) error {
		cl, err := getClient()
		if err != nil {
			return err
		}

		listOpts := client.ListOpts{
			Filters: map[string]interface{}{},
		}

		if c.IsSet("state") {
			listOpts.Filters["state"] = c.String("state")
		}

		filter(&filterContext{
			CLI:     c,
			Filters: listOpts.Filters,
		})

		hosts, err := cl.Host.List(&listOpts)
		if err != nil {
			return err
		}

		for _, host := range hosts.Data {
			err := apply(&mapContext{c, &host, cl})
			if err != nil {
				return err
			}
		}

		return nil
	}

	return cmd
}
