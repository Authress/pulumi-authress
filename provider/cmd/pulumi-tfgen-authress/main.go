package main

import (
	authressPulumi "github.com/Authress/pulumi-authress/provider"
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("authress", authressPulumi.Provider())
}
