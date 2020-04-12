//+build mage

package main

import (
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