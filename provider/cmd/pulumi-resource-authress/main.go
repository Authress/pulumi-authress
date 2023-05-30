//go:generate go run ./generate.go

package main

import (
	"fmt"
	"context"
	_ "embed"

	"github.com/Authress/pulumi-authress/provider/pkg/version"
	authressPulumi "github.com/Authress/pulumi-authress/provider"
	tfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
)

//go:embed schema-embed.json
var pulumiSchema []byte

//go:embed bridge-metadata.json
var bridgeMetadata []byte

func main() {
	meta := tfbridge.ProviderMetadata{
		PackageSchema:  pulumiSchema,
		BridgeMetadata: bridgeMetadata,
	}

	fmt.Println("Authress Package Version:\t", version.Version)

	tfbridge.Main(context.Background(), "authress", authressPulumi.Provider(), meta)
}
