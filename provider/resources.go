// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authress

import (
	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"fmt"
	"path/filepath"
	"unicode"

	"github.com/Authress/pulumi-authress/provider/pkg/version"
	authressTerraformProvider "github.com/authress/terraform-provider-authress/src"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	authressProvider = "authress"
	// modules:
	mainMod = "index" // the authress module
)

// // preConfigureCallback is called before the providerConfigure function of the underlying provider.
// // It should validate that the provider can be configured, and provide actionable errors in the case
// // it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// // for example `stringValue(vars, "accessKey")`.
// func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
// 	return nil
// }

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() pf.ProviderInfo {
	info := tfbridge.ProviderInfo{
		Name: "authress",
		DisplayName: "Authress",
		Publisher: "Authress",
		// Deployed from the Authress Restrict UI
		LogoURL: "https://authress.io/app/img/logo.svg",
		// https://www.pulumi.com/docs/using-pulumi/pulumi-packages/how-to-author/#support-for-github-releases
		PluginDownloadURL: "github://api.github.com/Authress/pulumi-authress",
		Description:       "A Pulumi package for creating and managing Authress resources.",
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "authress", "category/cloud", "authorization", "authentication"},
		License:    "Apache-2.0",
		Homepage:   "https://authress.io",
		Repository: "https://github.com/Authress/pulumi-authress",
		GitHubOrg: "Authress",
		Version:      version.Version,

		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Resources: map[string]*tfbridge.ResourceInfo{
			"random_id":       {Tok: randomResource(mainMod, "RandomId")},
			"random_password": {Tok: randomResource(mainMod, "RandomPassword")},
			"random_pet":      {Tok: randomResource(mainMod, "RandomPet")},
			"random_shuffle":  {Tok: randomResource(mainMod, "RandomShuffle")},
			"random_integer":  {Tok: randomResource(mainMod, "RandomInteger")},
			"random_uuid":     {Tok: randomResource(mainMod, "RandomUuid")},

			"random_string": {
				Tok: randomResource(mainMod, "RandomString"),
				PreStateUpgradeHook: func(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
					// States for RandomString may be contaminated by
					// https://github.com/pulumi/pulumi-random/issues/258 bug where the state is
					// missing the version marker. Pretend that these states are at V1, which is the
					// best guess. V1->V2/V3 migrations seem idempotent, this is probably safe.
					if args.PriorStateSchemaVersion == 0 {
						return 1, args.PriorState, nil
					}
					return args.PriorStateSchemaVersion, args.PriorState, nil
				},
			},
		},

		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", authressProvider),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				authressProvider,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"authress": "Authress",
			},
		},
	}
	return pf.ProviderInfo{
		ProviderInfo: info,
		NewProvider:  authressTerraformProvider.New(),
	}
}

// randomMember manufactures a type token for the random package and the given module and type.
func randomMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(authressProvider + ":" + mod + ":" + mem)
}

// randomType manufactures a type token for the random package and the given module and type.
func randomType(mod string, typ string) tokens.Type {
	return tokens.Type(randomMember(mod, typ))
}

// randomResource manufactures a standard resource token given a module and resource name.  It automatically uses the
// random package and names the file by simply lower casing the resource's first character.
func randomResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return randomType(mod+"/"+fn, res)
}

// go:embed cmd/pulumi-resource-authress/bridge-metadata.json
var metadata []byte
