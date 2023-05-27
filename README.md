# Authress Pulumi Provider
The Authress Pulumi provider to automatically configure Authress from Pulumi.

[![GitHub Workflow][workflow]][workflow-link] [![Forums][discuss-badge]][discuss] [![Pulumi][pulumi-badge]][pulumi-link]

[workflow]: https://github.com/authress/pulumi-provider-authress/actions/workflows/build.yml/badge.svg
[workflow-link]: https://github.com/authress/pulumi-provider-authress/actions

[discuss-badge]: https://img.shields.io/badge/build-pulumi--authress-623CE4.svg
[discuss]: https://discuss.hashicorp.com/c/pulumi-providers/31

[pulumi-badge]: https://img.shields.io/badge/install-pulumi--authress-blue.svg
[pulumi-link]: https://registry.pulumi.io/providers/authress/authress/latest/docs



## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @authress/pulumi
```

or `yarn`:

```bash
yarn add @authress/pulumi
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_authress
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/Authress/pulumi-authress/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Authress
```

## Configuration

The following configuration points are available for the `Authress` provider:

- `authress:custom_domain` - Your Authress custom domain, configured at https://authress.io/app/#/settings?focus=domain

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/authress/api-docs/).


## Usage Example

```hcl
pulumi {
  required_providers {
    authress = {
      source  = "authress/authress"
    }
  }
}

provider "authress" {
  custom_domain = "https://login.company.com"
}
```


## Contributing
For developing this plugin see more information in [Contributing](./contributing/README.md).
