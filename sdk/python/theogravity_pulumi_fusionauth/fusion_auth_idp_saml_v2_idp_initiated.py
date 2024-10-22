# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FusionAuthIdpSamlV2IdpInitiatedArgs', 'FusionAuthIdpSamlV2IdpInitiated']

@pulumi.input_type
class FusionAuthIdpSamlV2IdpInitiatedArgs:
    def __init__(__self__, *,
                 issuer: pulumi.Input[str],
                 key_id: pulumi.Input[str],
                 application_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 email_claim: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp_id: Optional[pulumi.Input[str]] = None,
                 lambda_reconcile_id: Optional[pulumi.Input[str]] = None,
                 linking_strategy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]] = None,
                 use_name_for_email: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a FusionAuthIdpSamlV2IdpInitiated resource.
        :param pulumi.Input[str] issuer: The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
               this value was required to be a URL.
        :param pulumi.Input[str] key_id: The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
               identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
               component).
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]] application_configurations: The configuration for each Application that the identity provider is enabled for.
        :param pulumi.Input[bool] debug: Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
               an Event Log will be created.
        :param pulumi.Input[str] email_claim: The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
               identity the user. If this is not set, the `use_name_for_email` flag must be true.
        :param pulumi.Input[bool] enabled: Determines if this provider is enabled. If it is false then it will be disabled globally.
        :param pulumi.Input[str] idp_id: The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] lambda_reconcile_id: The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        :param pulumi.Input[str] linking_strategy: The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        :param pulumi.Input[str] name: The name of this SAML v2 identity provider. This is only used for display purposes.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]] tenant_configurations: The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        :param pulumi.Input[bool] use_name_for_email: Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
               processing. If this is false, then the `email_claim` property must be set.
        """
        pulumi.set(__self__, "issuer", issuer)
        pulumi.set(__self__, "key_id", key_id)
        if application_configurations is not None:
            pulumi.set(__self__, "application_configurations", application_configurations)
        if debug is not None:
            pulumi.set(__self__, "debug", debug)
        if email_claim is not None:
            pulumi.set(__self__, "email_claim", email_claim)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if idp_id is not None:
            pulumi.set(__self__, "idp_id", idp_id)
        if lambda_reconcile_id is not None:
            pulumi.set(__self__, "lambda_reconcile_id", lambda_reconcile_id)
        if linking_strategy is not None:
            pulumi.set(__self__, "linking_strategy", linking_strategy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tenant_configurations is not None:
            pulumi.set(__self__, "tenant_configurations", tenant_configurations)
        if use_name_for_email is not None:
            pulumi.set(__self__, "use_name_for_email", use_name_for_email)

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Input[str]:
        """
        The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
        this value was required to be a URL.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Input[str]:
        """
        The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
        identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
        component).
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="applicationConfigurations")
    def application_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]:
        """
        The configuration for each Application that the identity provider is enabled for.
        """
        return pulumi.get(self, "application_configurations")

    @application_configurations.setter
    def application_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]):
        pulumi.set(self, "application_configurations", value)

    @property
    @pulumi.getter
    def debug(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
        an Event Log will be created.
        """
        return pulumi.get(self, "debug")

    @debug.setter
    def debug(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "debug", value)

    @property
    @pulumi.getter(name="emailClaim")
    def email_claim(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
        identity the user. If this is not set, the `use_name_for_email` flag must be true.
        """
        return pulumi.get(self, "email_claim")

    @email_claim.setter
    def email_claim(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_claim", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if this provider is enabled. If it is false then it will be disabled globally.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "idp_id")

    @idp_id.setter
    def idp_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_id", value)

    @property
    @pulumi.getter(name="lambdaReconcileId")
    def lambda_reconcile_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        """
        return pulumi.get(self, "lambda_reconcile_id")

    @lambda_reconcile_id.setter
    def lambda_reconcile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_reconcile_id", value)

    @property
    @pulumi.getter(name="linkingStrategy")
    def linking_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        """
        return pulumi.get(self, "linking_strategy")

    @linking_strategy.setter
    def linking_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "linking_strategy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this SAML v2 identity provider. This is only used for display purposes.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tenantConfigurations")
    def tenant_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]:
        """
        The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        """
        return pulumi.get(self, "tenant_configurations")

    @tenant_configurations.setter
    def tenant_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]):
        pulumi.set(self, "tenant_configurations", value)

    @property
    @pulumi.getter(name="useNameForEmail")
    def use_name_for_email(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
        processing. If this is false, then the `email_claim` property must be set.
        """
        return pulumi.get(self, "use_name_for_email")

    @use_name_for_email.setter
    def use_name_for_email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_name_for_email", value)


@pulumi.input_type
class _FusionAuthIdpSamlV2IdpInitiatedState:
    def __init__(__self__, *,
                 application_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 email_claim: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 lambda_reconcile_id: Optional[pulumi.Input[str]] = None,
                 linking_strategy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_configurations: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]] = None,
                 use_name_for_email: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering FusionAuthIdpSamlV2IdpInitiated resources.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]] application_configurations: The configuration for each Application that the identity provider is enabled for.
        :param pulumi.Input[bool] debug: Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
               an Event Log will be created.
        :param pulumi.Input[str] email_claim: The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
               identity the user. If this is not set, the `use_name_for_email` flag must be true.
        :param pulumi.Input[bool] enabled: Determines if this provider is enabled. If it is false then it will be disabled globally.
        :param pulumi.Input[str] idp_id: The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] issuer: The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
               this value was required to be a URL.
        :param pulumi.Input[str] key_id: The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
               identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
               component).
        :param pulumi.Input[str] lambda_reconcile_id: The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        :param pulumi.Input[str] linking_strategy: The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        :param pulumi.Input[str] name: The name of this SAML v2 identity provider. This is only used for display purposes.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]] tenant_configurations: The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        :param pulumi.Input[bool] use_name_for_email: Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
               processing. If this is false, then the `email_claim` property must be set.
        """
        if application_configurations is not None:
            pulumi.set(__self__, "application_configurations", application_configurations)
        if debug is not None:
            pulumi.set(__self__, "debug", debug)
        if email_claim is not None:
            pulumi.set(__self__, "email_claim", email_claim)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if idp_id is not None:
            pulumi.set(__self__, "idp_id", idp_id)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if lambda_reconcile_id is not None:
            pulumi.set(__self__, "lambda_reconcile_id", lambda_reconcile_id)
        if linking_strategy is not None:
            pulumi.set(__self__, "linking_strategy", linking_strategy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tenant_configurations is not None:
            pulumi.set(__self__, "tenant_configurations", tenant_configurations)
        if use_name_for_email is not None:
            pulumi.set(__self__, "use_name_for_email", use_name_for_email)

    @property
    @pulumi.getter(name="applicationConfigurations")
    def application_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]:
        """
        The configuration for each Application that the identity provider is enabled for.
        """
        return pulumi.get(self, "application_configurations")

    @application_configurations.setter
    def application_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]):
        pulumi.set(self, "application_configurations", value)

    @property
    @pulumi.getter
    def debug(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
        an Event Log will be created.
        """
        return pulumi.get(self, "debug")

    @debug.setter
    def debug(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "debug", value)

    @property
    @pulumi.getter(name="emailClaim")
    def email_claim(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
        identity the user. If this is not set, the `use_name_for_email` flag must be true.
        """
        return pulumi.get(self, "email_claim")

    @email_claim.setter
    def email_claim(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_claim", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if this provider is enabled. If it is false then it will be disabled globally.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "idp_id")

    @idp_id.setter
    def idp_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_id", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
        this value was required to be a URL.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
        identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
        component).
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="lambdaReconcileId")
    def lambda_reconcile_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        """
        return pulumi.get(self, "lambda_reconcile_id")

    @lambda_reconcile_id.setter
    def lambda_reconcile_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_reconcile_id", value)

    @property
    @pulumi.getter(name="linkingStrategy")
    def linking_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        """
        return pulumi.get(self, "linking_strategy")

    @linking_strategy.setter
    def linking_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "linking_strategy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of this SAML v2 identity provider. This is only used for display purposes.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="tenantConfigurations")
    def tenant_configurations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]:
        """
        The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        """
        return pulumi.get(self, "tenant_configurations")

    @tenant_configurations.setter
    def tenant_configurations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]):
        pulumi.set(self, "tenant_configurations", value)

    @property
    @pulumi.getter(name="useNameForEmail")
    def use_name_for_email(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
        processing. If this is false, then the `email_claim` property must be set.
        """
        return pulumi.get(self, "use_name_for_email")

    @use_name_for_email.setter
    def use_name_for_email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_name_for_email", value)


class FusionAuthIdpSamlV2IdpInitiated(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 email_claim: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 lambda_reconcile_id: Optional[pulumi.Input[str]] = None,
                 linking_strategy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]] = None,
                 use_name_for_email: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a FusionAuthIdpSamlV2IdpInitiated resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]] application_configurations: The configuration for each Application that the identity provider is enabled for.
        :param pulumi.Input[bool] debug: Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
               an Event Log will be created.
        :param pulumi.Input[str] email_claim: The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
               identity the user. If this is not set, the `use_name_for_email` flag must be true.
        :param pulumi.Input[bool] enabled: Determines if this provider is enabled. If it is false then it will be disabled globally.
        :param pulumi.Input[str] idp_id: The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] issuer: The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
               this value was required to be a URL.
        :param pulumi.Input[str] key_id: The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
               identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
               component).
        :param pulumi.Input[str] lambda_reconcile_id: The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        :param pulumi.Input[str] linking_strategy: The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        :param pulumi.Input[str] name: The name of this SAML v2 identity provider. This is only used for display purposes.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]] tenant_configurations: The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        :param pulumi.Input[bool] use_name_for_email: Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
               processing. If this is false, then the `email_claim` property must be set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthIdpSamlV2IdpInitiatedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FusionAuthIdpSamlV2IdpInitiated resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FusionAuthIdpSamlV2IdpInitiatedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthIdpSamlV2IdpInitiatedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 email_claim: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp_id: Optional[pulumi.Input[str]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 lambda_reconcile_id: Optional[pulumi.Input[str]] = None,
                 linking_strategy: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tenant_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]] = None,
                 use_name_for_email: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthIdpSamlV2IdpInitiatedArgs.__new__(FusionAuthIdpSamlV2IdpInitiatedArgs)

            __props__.__dict__["application_configurations"] = application_configurations
            __props__.__dict__["debug"] = debug
            __props__.__dict__["email_claim"] = email_claim
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["idp_id"] = idp_id
            if issuer is None and not opts.urn:
                raise TypeError("Missing required property 'issuer'")
            __props__.__dict__["issuer"] = issuer
            if key_id is None and not opts.urn:
                raise TypeError("Missing required property 'key_id'")
            __props__.__dict__["key_id"] = key_id
            __props__.__dict__["lambda_reconcile_id"] = lambda_reconcile_id
            __props__.__dict__["linking_strategy"] = linking_strategy
            __props__.__dict__["name"] = name
            __props__.__dict__["tenant_configurations"] = tenant_configurations
            __props__.__dict__["use_name_for_email"] = use_name_for_email
        super(FusionAuthIdpSamlV2IdpInitiated, __self__).__init__(
            'fusionauth:index/fusionAuthIdpSamlV2IdpInitiated:FusionAuthIdpSamlV2IdpInitiated',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]]] = None,
            debug: Optional[pulumi.Input[bool]] = None,
            email_claim: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            idp_id: Optional[pulumi.Input[str]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            lambda_reconcile_id: Optional[pulumi.Input[str]] = None,
            linking_strategy: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tenant_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]]] = None,
            use_name_for_email: Optional[pulumi.Input[bool]] = None) -> 'FusionAuthIdpSamlV2IdpInitiated':
        """
        Get an existing FusionAuthIdpSamlV2IdpInitiated resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedApplicationConfigurationArgs']]]] application_configurations: The configuration for each Application that the identity provider is enabled for.
        :param pulumi.Input[bool] debug: Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
               an Event Log will be created.
        :param pulumi.Input[str] email_claim: The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
               identity the user. If this is not set, the `use_name_for_email` flag must be true.
        :param pulumi.Input[bool] enabled: Determines if this provider is enabled. If it is false then it will be disabled globally.
        :param pulumi.Input[str] idp_id: The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] issuer: The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
               this value was required to be a URL.
        :param pulumi.Input[str] key_id: The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
               identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
               component).
        :param pulumi.Input[str] lambda_reconcile_id: The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        :param pulumi.Input[str] linking_strategy: The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        :param pulumi.Input[str] name: The name of this SAML v2 identity provider. This is only used for display purposes.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthIdpSamlV2IdpInitiatedTenantConfigurationArgs']]]] tenant_configurations: The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        :param pulumi.Input[bool] use_name_for_email: Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
               processing. If this is false, then the `email_claim` property must be set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthIdpSamlV2IdpInitiatedState.__new__(_FusionAuthIdpSamlV2IdpInitiatedState)

        __props__.__dict__["application_configurations"] = application_configurations
        __props__.__dict__["debug"] = debug
        __props__.__dict__["email_claim"] = email_claim
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["idp_id"] = idp_id
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["lambda_reconcile_id"] = lambda_reconcile_id
        __props__.__dict__["linking_strategy"] = linking_strategy
        __props__.__dict__["name"] = name
        __props__.__dict__["tenant_configurations"] = tenant_configurations
        __props__.__dict__["use_name_for_email"] = use_name_for_email
        return FusionAuthIdpSamlV2IdpInitiated(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationConfigurations")
    def application_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.FusionAuthIdpSamlV2IdpInitiatedApplicationConfiguration']]]:
        """
        The configuration for each Application that the identity provider is enabled for.
        """
        return pulumi.get(self, "application_configurations")

    @property
    @pulumi.getter
    def debug(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if debug is enabled for this provider. When enabled, each time this provider is invoked to reconcile a login
        an Event Log will be created.
        """
        return pulumi.get(self, "debug")

    @property
    @pulumi.getter(name="emailClaim")
    def email_claim(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the email claim (Attribute in the Assertion element) in the SAML response that FusionAuth uses to uniquely
        identity the user. If this is not set, the `use_name_for_email` flag must be true.
        """
        return pulumi.get(self, "email_claim")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if this provider is enabled. If it is false then it will be disabled globally.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="idpId")
    def idp_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID to use for the new identity provider. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "idp_id")

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[str]:
        """
        The EntityId (unique identifier) of the SAML v2 identity provider. This value should be provided to you. Prior to 1.27.1
        this value was required to be a URL.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        The id of the key stored in Key Master that is used to verify the SAML response sent back to FusionAuth from the
        identity provider. This key must be a verification only key or certificate (meaning that it only has a public key
        component).
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="lambdaReconcileId")
    def lambda_reconcile_id(self) -> pulumi.Output[Optional[str]]:
        """
        The id of a SAML reconcile lambda that is applied when the identity provider sends back a successful SAML response.
        """
        return pulumi.get(self, "lambda_reconcile_id")

    @property
    @pulumi.getter(name="linkingStrategy")
    def linking_strategy(self) -> pulumi.Output[str]:
        """
        The linking strategy to use when creating the link between the {idp_display_name} Identity Provider and the user.
        """
        return pulumi.get(self, "linking_strategy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of this SAML v2 identity provider. This is only used for display purposes.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="tenantConfigurations")
    def tenant_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.FusionAuthIdpSamlV2IdpInitiatedTenantConfiguration']]]:
        """
        The configuration for each Tenant that limits the number of links a user may have for a particular identity provider.
        """
        return pulumi.get(self, "tenant_configurations")

    @property
    @pulumi.getter(name="useNameForEmail")
    def use_name_for_email(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not FusionAuth will use the NameID element value as the email address of the user for reconciliation
        processing. If this is false, then the `email_claim` property must be set.
        """
        return pulumi.get(self, "use_name_for_email")

