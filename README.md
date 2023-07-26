# Authress Pulumi Provider
The Authress Pulumi provider to automatically configure Authress from Pulumi.

[![GitHub Workflow][workflow]][workflow-link] [![Forums][discuss-badge]][discuss] [![Pulumi][pulumi-badge]][pulumi-link]

[workflow]: https://github.com/authress/pulumi-authress/actions/workflows/build.yml/badge.svg
[workflow-link]: https://github.com/authress/pulumi-authress/actions

[discuss-badge]: https://img.shields.io/badge/Community-Authress-fbaf0b.svg
[discuss]: https://authress.io/community

[pulumi-badge]: https://img.shields.io/badge/Install-pulumi--authress-blue.svg
[pulumi-link]: https://www.pulumi.com/registry/packages/authress/



## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/authress
```

or `yarn`:

```bash
yarn add @pulumi/authress
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_authress
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/Authress/pulumi-authress/sdk/go
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Authress
```

## Configuration

The following configuration points are available for the `Authress` provider:

- `authress:accessKey` - The access key for the Authress API. Should be [configured by your CI/CD](https://authress.io/knowledge-base/docs/category/cicd) automatically. Or it can be overridden directly here. Do not commit this plaintext value to your source code.
- `authress:customDomain` - Your Authress custom domain. [Configure a custom domain for your Authress account](https://authress.io/app/#/settings?focus=domain) or use the [provided domain](https://authress.io/app/#/api?route=overview).

## Reference Examples

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/authress/api-docs/).



## Contributing
For developing this plugin see more information in [Contributing](./contributing/README.md).
