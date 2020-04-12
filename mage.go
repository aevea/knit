//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/outillage/magefiles"
)

func Install() error {
	dependencies := []string{
		"github.com/outillage/oto-tools",
		"github.com/pacedotdev/oto",
	}
	return magefiles.Install(dependencies)
}

func Test() error {
	return magefiles.Test()
}

func GoModTidy() error {
	return magefiles.GoModTidy()
}

type Build mg.Namespace

func (Build) Server() error {
	var args []string

	args = append(args, "-template", "./templates/oto/server.go.plush")

	args = append(args, "-out", "./api/generated/oto.gen.go")

	args = append(args, "-ignore", "Ignorer")

	args = append(args, "-pkg", "generated")

	args = append(args, "./api/definitions")

	return sh.RunV("oto", args...)
}

func PublishClient() error {
	err := sh.RunV("oto-tools", "generate",
		"--package-name", "@outillage/merge-master",
		"--oto-template", "./templates/oto/client.js.plush",
		"--oto-definitions", "./api/definitions")

	if err != nil {
		return err
	}

	token := os.Getenv("GITHUB_TOKEN")

	return sh.RunV("oto-tools", "publish-npm", "--token", token, "--registry", "github", "--owner", "outillage")
}
