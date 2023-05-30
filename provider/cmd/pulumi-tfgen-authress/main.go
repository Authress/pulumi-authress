package main

import (
	"fmt"
	"github.com/Authress/pulumi-authress/provider/pkg/version"
	authressPulumi "github.com/Authress/pulumi-authress/provider"
	"github.com/pulumi/pulumi-terraform-bridge/pf/tfgen"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

func main() {
	fmt.Println("Authress TF Generator Version:\t", version.Version)

	// Validates the version is set correctly
	tfbridge.GetModuleMajorVersion(version.Version)

	// Modify the path to point to the new provider
	tfgen.Main("authress", authressPulumi.Provider())
}
