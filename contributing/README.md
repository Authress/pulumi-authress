## How to Contribute to the Authress Pulumi provider
The reference guide for building and publishing the package is available in the [Pulumi docs](https://www.pulumi.com/docs/using-pulumi/pulumi-packages/how-to-author/).

## Prerequisites

Ensure the following tools are installed and present in your `$PATH`:

- [Pulumi](https://www.pulumi.com/docs/install/)
- [`pulumictl`](https://github.com/pulumi/pulumictl/releases) - Download and copy into this repos top level directory
- [Go 1.17](https://golang.org/dl/) or 1.latest
- [NodeJS](https://nodejs.org/en/) 14.x.  We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations.
- [Yarn](https://yarnpkg.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Python](https://www.python.org/downloads/) (called as `python3`).  For recent versions of MacOS, the system-installed version is fine.
- [.NET](https://dotnet.microsoft.com/download)

## Development

### Install necessary Go packages:

```sh
cd provider
go mod tidy
```

### Regenerate TF conversion
Blocked on https://github.com/pulumi/pulumi-terraform-bridge/issues/956

```sh
export VERSION=v1.1.1
export PATH="/home/$USER/git/authress/sdk/pulumi/bin:/home/$USER/git/authress/sdk/pulumi/provider/cmd:$PATH"
make tfgen
make provider
make build_nodejs
```

### Run an example

```sh
make install_nodejs_sdk
cd examples/typescript-example
yarn
yarn link @pulumi/authress
```

1. Configure any necessary environment variables for authentication, e.g `$AUTHRESS_KEY`, in your local environment.
1. Ensure the program runs successfully via `pulumi up`.
1. Once the program completes successfully, verify the resource was created in the provider's UI.
1. Destroy any resources created by the program via `pulumi destroy`.