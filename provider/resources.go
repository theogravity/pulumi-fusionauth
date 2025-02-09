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

package fusionauth

import (
	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	// Replace this provider with the provider you are bridging.
	fusionauth "github.com/FusionAuth/terraform-provider-fusionauth/fusionauth"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "fusionauth"
	// modules:
	mainMod = "index" // the fusionauth module
)

//go:embed cmd/pulumi-resource-fusionauth/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(fusionauth.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "fusionauth",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "FusionAuth",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Theo Gravity",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://avatars.githubusercontent.com/u/41974756?s=200&v=4",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/theogravity/pulumi-fusionauth/releases/download/v${VERSION}",
		Description:       "A Pulumi package for managing FusionAuth instances.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "fusionauth", "category/infrastructure"},
		License:    "MIT",
		Homepage:   "https://github.com/theogravity/pulumi-fusionauth",
		Repository: "https://github.com/theogravity/pulumi-fusionauth",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "FusionAuth",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FUSION_AUTH_API_KEY"},
				},
			},
			"host": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"FUSION_AUTH_HOST_URL"},
				},
			},
		},
		// List based on https://registry.terraform.io/providers/fusionauth/fusionauth/latest/docs
		Resources: map[string]*tfbridge.ResourceInfo{
			"fusionauth_api_key":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthApiKey")},
			"fusionauth_application":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthApplication")},
			"fusionauth_application_oauth_scope": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthApplicationOAuthScope")},
			"fusionauth_application_role":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthApplicationRole")},
			"fusionauth_email":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthEMail")},
			"fusionauth_entity":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthEntity")},
			"fusionauth_entity_grant":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthEntityGrant")},
			"fusionauth_entity_type":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthEntityType")},
			"fusionauth_entity_type_permission":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthEntityTypePermission")},
			"fusionauth_form":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthForm")},
			"fusionauth_form_field":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthFormField")},
			"fusionauth_generic_connector":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthGenericConnector")},
			"fusionauth_group":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthGroup")},
			"fusionauth_idp_apple":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpApple")},
			"fusionauth_idp_external_jwt":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpExternalJwt")},
			"fusionauth_idp_facebook":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpFacebook")},
			"fusionauth_idp_google":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpGoogle")},
			"fusionauth_idp_linkedin":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpLinkedIn")},
			"fusionauth_idp_open_id_connect":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpOpenIdConnect")},
			"fusionauth_idp_saml_v2":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpSamlv2")},
			// This misspelling has been here for a really long time, not sure if it should / can be corrected
			"fusionauth_idp_saml_v2_idp_initated": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpSamlV2IdpInitiated")},
			"fusionauth_idp_sony_psn":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpPsn")},
			"fusionauth_idp_steam":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpSteam")},
			"fusionauth_idp_twitch":               {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpTwitch")},
			"fusionauth_idp_xbox":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthIdpXBox")},
			"fusionauth_imported_key":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthImportedKey")},
			"fusionauth_key":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthKey")},
			"fusionauth_lambda":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthLambda")},
			"fusionauth_reactor":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthReactor")},
			"fusionauth_registration":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthRegistration")},
			"fusionauth_system_configuration":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthSystemConfiguration")},
			"fusionauth_tenant":                   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthTenant")},
			"fusionauth_theme":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthTheme")},
			"fusionauth_user":                     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthUser")},
			"fusionauth_user_action":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthUserAction")},
			"fusionauth_user_group_membership":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthUserGroupMembership")},
			"fusionauth_webhook":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "FusionAuthWebhook")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"fusionauth_application":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplication")},
			"fusionauth_application_oauth_scope": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplicationOauthScope")},
			"fusionauth_application_role":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getApplicationRole")},
			"fusionauth_email":                   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEMail")},
			"fusionauth_form":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getForm")},
			"fusionauth_form_field":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getFormField")},
			"fusionauth_idp":                     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getIdp")},
			"fusionauth_lambda":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLambda")},
			"fusionauth_tenant":                  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTenant")},
			"fusionauth_user":                    {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"fusionauth_user_group_membership":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUserGroupMembership")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "pulumi-fusionauth",
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
			PackageName: "theogravity_pulumi_fusionauth",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/theogravity/pulumi-%[1]s/sdk/", mainPkg),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "theogravity",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("fusionauth_", mainMod,
		tokens.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
