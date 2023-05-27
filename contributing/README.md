## How to Contribute to the Authress Pulumi provider

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
export $PATH="~/git/authress/sdk/pulumi:$PATH"
cd provider && go mod tidy
```

### Regenerate TF conversion
Blocked on https://github.com/pulumi/pulumi-terraform-bridge/issues/956

```sh
make tfgen
```