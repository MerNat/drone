package main

import "github.com/urfave/cli"

var registryDeleteCmd = cli.Command{
	Name:   "rm",
	Usage:  "remove a registry",
	Action: registryDelete,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "repository",
			Usage: "repository name (e.g. octocat/hello-world)",
		},
		cli.StringFlag{
			Name:  "hostname",
			Usage: "registry hostname",
			Value: "index.docker.io",
		},
	},
}

func registryDelete(c *cli.Context) error {
	var (
		hostname = c.String("hostname")
		reponame = c.String("repository")
	)
	if reponame == "" {
		reponame = c.Args().First()
	}
	owner, name, err := parseRepo(reponame)
	if err != nil {
		return err
	}
	client, err := newClient(c)
	if err != nil {
		return err
	}
	return client.RegistryDelete(owner, name, hostname)
}
