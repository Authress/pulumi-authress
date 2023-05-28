## Adding Mappings, Building the Provider and SDKs

In this section we will add the mappings that allow the interoperation between the Pulumi provider and the Terraform provider.  Terraform resources map to an identically named concept in Pulumi.  Terraform data sources map to plain old functions in your supported programming language of choice.  Pulumi also allows provider functions and resources to be grouped into _namespaces_ to improve the cohesion of a provider's code, thereby making it easier for developers to use.  If your provider has a large number of resources, consider using namespaces to improve usability.

The following instructions all pertain to `provider/resources.go`, in the section of the code where we construct a `tfbridge.ProviderInfo` object:

1. **Add resource mappings:** For each resource in the provider, add an entry in the `Resources` property of the `tfbridge.ProviderInfo`, e.g.:

    ```go
    // Most providers will have all resources (and data sources) in the main module.
    // Note the mapping from snake_case HCL naming conventions to UpperCamelCase Pulumi SDK naming conventions.
    // The name of the provider is omitted from the mapped name due to the presence of namespaces in all supported Pulumi languages.
    "authress_something":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Something")},
    "authress_something_else": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SomethingElse")},
    ```

1. **Add CSharpName (if necessary):** Dotnet does not allow for fields named the same as the enclosing type, which sometimes results in errors during the dotnet SDK build.
    If you see something like

    ```text
    error CS0542: 'ApiKey': member names cannot be the same as their enclosing type [/Users/guin/go/src/github.com/pulumi/pulumi-artifactory/sdk/dotnet/Pulumi.Artifactory.csproj]
    ```

    you'll want to give your Resource a CSharpName, which can have any value that makes sense:

    ```go
    "authress_something_dotnet": {
        Tok: tfbridge.MakeResource(mainPkg, mainMod, "SomethingDotnet"),
        Fields: map[string]*tfbridge.SchemaInfo{
            "something_dotnet": {
                CSharpName: "SpecialName",
            },
        },
    },
    ```

   [See the underlying terraform-bridge code here.](https://github.com/pulumi/pulumi-terraform-bridge/blob/master/pkg/tfbridge/info.go#L168)
1. **Add data source mappings:** For each data source in the provider, add an entry in the `DataSources` property of the `tfbridge.ProviderInfo`, e.g.:

    ```go
    // Note the 'get' prefix for data sources
    "authress_something":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSomething")},
    "authress_something_else": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSomethingElse")},
    ```


1. **Add provider configuration overrides (not typically needed):** Pulumi's Terraform bridge automatically detects configuration options for the upstream provider.  However, in rare cases these settings may need to be overridden, e.g. if we want to change an environment variable default from `API_KEY` to `FOO_API_KEY`.  Examples of common uses cases:

    ```go
    "additional_required_parameter": {},
    "additional_optional_string_parameter": {
        Default: &tfbridge.DefaultInfo{
            Value: "default_value",
        },
    "additional_optional_boolean_parameter": {
        Default: &tfbridge.DefaultInfo{
            Value: true,
        },
    // Renamed environment variables can be accounted for like so:
    "apikey": {
        Default: &tfbridge.DefaultInfo{
            EnvVars: []string{"FOO_API_KEY"},
        },
    ```

1. Build the provider binary and ensure there are no warnings about unmapped resources and no warnings about unmapped data sources:

    ```bash
    make provider
    ```

    You may see warnings about documentation and examples, including "unexpected code snippets".  These can be safely ignored for now.  Pulumi will add additional documentation on mapping docs in a future revision of this guide.

1. Build the SDKs in the various languages Pulumi supports:

    ```bash
    make build_sdks
    ```

1. Ensure the Golang SDK is a proper go module:

    ```bash
    cd sdk && go mod tidy && cd -
    ```

    This will pull in the correct dependencies in `sdk/go.mod` as well as setting the dependency tree in `sdk/go.sum`.

1. Finally, ensure the provider code conforms to Go standards:

    ```bash
    make lint_provider
    ```

    Fix any issues found by the linter.

**Note:** If you make revisions to code in `resources.go`, you must re-run the `make tfgen` target to regenerate the schema.
The `make tfgen` target will take the file `schema.json` and serialize it to a byte array so that it can be included in the build output.
(This is a holdover from Go 1.16, which does not have the ability to directly embed text files. We are working on removing the need for this step.)

## Sample Program

1. Create a minimal program for the provider, i.e. one that creates the smallest-footprint resource.  Place this code in `index.ts`.
1. Configure any necessary environment variables for authentication, e.g `$FOO_USERNAME`, `$FOO_TOKEN`, in your local environment.
1. Ensure the program runs successfully via `pulumi up`.
1. Once the program completes successfully, verify the resource was created in the provider's UI.
1. Destroy any resources created by the program via `pulumi destroy`.

Optionally, you may create additional examples for SDKs in other languages supported by Pulumi:

1. Python:

    ```bash
    mkdir examples/my-example/py
    cd examples/my-example/py
    pulumi new python
    # (Go through the prompts with the default values)
    source venv/bin/activate # use the virtual Python env that Pulumi sets up for you
    pip install pulumi_authress
    ```


## Add End-to-end Testing

We can run integration tests on our examples using the `*_test.go` files in the `examples/` folder.

1. Add code to `examples_nodejs_test.go` to call the example you created, e.g.:

    ```go
    // Swap out MyExample and "my-example" below with the name of your integration test.
    func TestAccMyExampleTs(t *testing.T) {
        test := getJSBaseOptions(t).
            With(integration.ProgramTestOptions{
                Dir: filepath.Join(getCwd(t), "my-example", "ts"),
            })
        integration.ProgramTest(t, &test)
    }
    ```

1. Add a similar function for each example that you want to run in an integration test.  For examples written in other languages, create similar files for `examples_${LANGUAGE}_test.go`.

1. You can run these tests locally via Make:

    ```bash
    make test
    ```

    You can also run each test file separately via test tags:

    ```bash
    cd examples && go test -v -tags=nodejs
    ```

## Configuring CI with GitHub Actions

### Third-party providers

1. Follow the instructions laid out in the [deployment templates](./deployment-templates/README-DEPLOYMENT.md).

### Pulumi Internal

In this section, we'll add the necessary configuration to work with GitHub Actions for Pulumi's standard CI/CD workflows for providers.

1. Generate GitHub workflows per [the instructions in the ci-mgmt repository](https://github.com/pulumi/ci-mgmt/) and copy to `.github/` in this repository.

1. Ensure that any required secrets are present as repository-level [secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) in GitHub.  These will be used by the integration tests during the CI/CD process.

1. Repository settings: Toggle `Allow auto-merge` on in your provider repo to automate GitHub Actions workflow updates.

## Final Steps

1. Ensure all required configurations (API keys, etc.) are documented in README-PROVIDER.md.

1. Replace this file with the README for the provider and push your changes:

    ```bash
    mv README-PROVIDER.md README.md
    ```

1. If publishing the npm package fails during the "Publish SDKs" Action, perform the following steps:
    1. Go to [NPM Packages](https://www.npmjs.com/) and sign in as pulumi-bot.
    1. Click on the bot's profile pic and navigate to "Packages".
    1. On the left, under "Organizations, click on the Pulumi organization.
    1. On the last page of the listed packages, you should see the new package.
    1. Under "Settings", set the Package Status to "public".

Now you are ready to use the provider, cut releases, and have some well-deserved :ice_cream:!

## Building the Provider Locally

There are 2 ways the provider can be built locally:

`make provider` will use the current operating system and architecture to create a binary that can be used on your PATH.

To build the provider for another set of operating systems / architectures, the project uses [goreleaser](https://goreleaser.com/).
Goreleaser, a CLI tool, that allows a user to build a matrix of binaries.

Create a `.goreleaser.yml` file in the root of your project:

```yaml

archives:
- id: archive
  name_template: "{{ .Binary }}-{{ .Tag }}-{{ .Os }}-{{ .Arch }}"
before:
  hooks:
  - make tfgen
builds:
- binary: pulumi-resource-authress
  dir: provider
  goarch:
  - amd64
  - arm64
  goos:
  - darwin
  - windows
  - linux
  ignore: []
  ldflags:
  - -X github.com/pulumi/pulumi-authress/provider/pkg/version.Version={{.Tag}}
  main: ./cmd/pulumi-resource-authress/
  sort: asc
  use: git
release:
  disable: false
snapshot:
  name_template: "{{ .Tag }}-SNAPSHOT"
```

To build the provider for the combination of architectures and operating systems, you can run the following CLI command:

```bash
goreleaser build --rm-dist --skip-validate
```

That will ensure that a list of binaries are available to use:

```bash

▶ tree dist
dist
├── CHANGELOG.md
├── artifacts.json
├── config.yaml
├── metadata.json
├── pulumi-authress_darwin_amd64_v1
│   └── pulumi-resource-authress
├── pulumi-authress_darwin_arm64
│   └── pulumi-resource-authress
├── pulumi-authress_linux_amd64_v1
│   └── pulumi-resource-authress
├── pulumi-authress_linux_arm64
│   └── pulumi-resource-authress
├── pulumi-authress_windows_amd64_v1
│   └── pulumi-resource-authress.exe
└── pulumi-authress_windows_arm64
    └── pulumi-resource-authress.exe
```

Any of the provider binaries can be used to target the correct machine architecture


## And finally publish
https://www.pulumi.com/docs/using-pulumi/pulumi-packages/how-to-author/#support-for-github-releases