package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

var List cli.Command = buildMapHostsFunction(cli.Command{
	Name:        "list",
	Description: "Get a list of the hosts on your Rancher server",
	Flags:       []cli.Flag{},
}, func(c *filterContext) {

}, func(c *mapContext) error {
	fmt.Printf("#%s %s - %s\n", c.Host.Id, c.Host.Hostname, c.Host.State)
	return nil
})
