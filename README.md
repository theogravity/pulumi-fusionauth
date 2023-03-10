# FusionAuth Resource Provider

The FusionAuth Resource Provider lets you manage [FusionAuth](http://fusionauth.io/) resources.

This is bridged using the [Terraform FusionAuth](https://github.com/gpsinsight/terraform-provider-fusionauth) package.

You can look at the [Terraform FusionAuth docs by gpsinsight](https://registry.terraform.io/providers/gpsinsight/fusionauth/latest/docs)
to get an idea of what is available in this package.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install pulumi-fusionauth
```

or `yarn`:

```bash
yarn add pulumi-fusionauth
```

### Python

To use from Python, install using `pip`:

```bash
pip install theogravity-pulumi-fusionauth
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package theogravity.Fusionauth
```

## Configuration

The following configuration points are available for the `fusionauth` provider:

- `fusionauth:host` (environment: `FUSION_AUTH_HOST_URL`) - the URL to the FusionAuth instance with the trailing slash omitted (ex: `https://instance.fusionauth.io`)
- `fusionauth:api_key` (environment: `FUSION_AUTH_API_KEY`) - the API key for `fusionauth`

## Publishing new versions

- Update `provider/go.mod` with the appropriate `github.com/gpsinsight/terraform-provider-fusionauth` version
- Run `make tfgen` in the root of this directory
- Check the warnings for any new mappings that need to be added to `provider/resources.go` and run `make tfgen` again after
corrections are made
- Commit all the generated files and push
- Make sure the CI passes green.
- Create a release by making a git tag in the format of `major.minor.patch` and push that tag to origin via `git push origin <tag>`
