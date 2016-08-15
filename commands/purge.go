package commands

import (
	"log"

	"github.com/urfave/cli"
)

var Purge cli.Command = buildMapHostsFunction(cli.Command{
	Name:        "purge",
	Description: "Purges a specific host or all hosts which match a supplied rule.",
	Usage:       "[HOST_ID]",
	Flags:       []cli.Flag{},
}, func(c *filterContext) {
	if c.CLI.NArg() > 0 {
		c.Filters["hostId"] = c.CLI.Args().Get(0)
	}
}, func(c *mapContext) error {
	log.Printf("Purging host #%s %s: ", c.Host.Id, c.Host.Hostname)
	_, err := c.Rancher.Host.ActionPurge(c.Host)
	if err != nil {
		log.Printf("Failed [%s]\n", err)
	} else {
		log.Printf("Success\n")
	}

	return nil
})
