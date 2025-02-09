# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['FusionAuthApplicationOAuthScopeArgs', 'FusionAuthApplicationOAuthScope']

@pulumi.input_type
class FusionAuthApplicationOAuthScopeArgs:
    def __init__(__self__, *,
                 application_id: pulumi.Input[str],
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 default_consent_detail: Optional[pulumi.Input[str]] = None,
                 default_consent_message: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthApplicationOAuthScope resource.
        :param pulumi.Input[str] application_id: ID of the application that this role is for.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the OAuth Scope that should be persisted.
        :param pulumi.Input[str] default_consent_detail: "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] default_consent_message: The default message to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] description: A description of the OAuth Scope. This is used for display purposes only.
        :param pulumi.Input[str] name: The name of the Role.
        :param pulumi.Input[bool] required: Determines if the OAuth Scope is required when requested in an OAuth workflow.
        :param pulumi.Input[str] scope_id: The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        pulumi.set(__self__, "application_id", application_id)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if default_consent_detail is not None:
            pulumi.set(__self__, "default_consent_detail", default_consent_detail)
        if default_consent_message is not None:
            pulumi.set(__self__, "default_consent_message", default_consent_message)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Input[str]:
        """
        ID of the application that this role is for.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object that can hold any information about the OAuth Scope that should be persisted.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="defaultConsentDetail")
    def default_consent_detail(self) -> Optional[pulumi.Input[str]]:
        """
        "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_detail")

    @default_consent_detail.setter
    def default_consent_detail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_consent_detail", value)

    @property
    @pulumi.getter(name="defaultConsentMessage")
    def default_consent_message(self) -> Optional[pulumi.Input[str]]:
        """
        The default message to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_message")

    @default_consent_message.setter
    def default_consent_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_consent_message", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the OAuth Scope. This is used for display purposes only.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the OAuth Scope is required when requested in an OAuth workflow.
        """
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)


@pulumi.input_type
class _FusionAuthApplicationOAuthScopeState:
    def __init__(__self__, *,
                 application_id: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 default_consent_detail: Optional[pulumi.Input[str]] = None,
                 default_consent_message: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthApplicationOAuthScope resources.
        :param pulumi.Input[str] application_id: ID of the application that this role is for.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the OAuth Scope that should be persisted.
        :param pulumi.Input[str] default_consent_detail: "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] default_consent_message: The default message to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] description: A description of the OAuth Scope. This is used for display purposes only.
        :param pulumi.Input[str] name: The name of the Role.
        :param pulumi.Input[bool] required: Determines if the OAuth Scope is required when requested in an OAuth workflow.
        :param pulumi.Input[str] scope_id: The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        if application_id is not None:
            pulumi.set(__self__, "application_id", application_id)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if default_consent_detail is not None:
            pulumi.set(__self__, "default_consent_detail", default_consent_detail)
        if default_consent_message is not None:
            pulumi.set(__self__, "default_consent_message", default_consent_message)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if scope_id is not None:
            pulumi.set(__self__, "scope_id", scope_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the application that this role is for.
        """
        return pulumi.get(self, "application_id")

    @application_id.setter
    def application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_id", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object that can hold any information about the OAuth Scope that should be persisted.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="defaultConsentDetail")
    def default_consent_detail(self) -> Optional[pulumi.Input[str]]:
        """
        "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_detail")

    @default_consent_detail.setter
    def default_consent_detail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_consent_detail", value)

    @property
    @pulumi.getter(name="defaultConsentMessage")
    def default_consent_message(self) -> Optional[pulumi.Input[str]]:
        """
        The default message to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_message")

    @default_consent_message.setter
    def default_consent_message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_consent_message", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the OAuth Scope. This is used for display purposes only.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines if the OAuth Scope is required when requested in an OAuth workflow.
        """
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "scope_id")

    @scope_id.setter
    def scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope_id", value)


class FusionAuthApplicationOAuthScope(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 default_consent_detail: Optional[pulumi.Input[str]] = None,
                 default_consent_message: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Application OAuth Scope Resource

        The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.

        [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        this = fusionauth.FusionAuthApplicationOAuthScope("this",
            application_id=fusionauth_application["this"]["id"],
            data={
                "addedBy": "Tom",
                "addedOn": "2025-02-05",
            },
            default_consent_detail="This will provide the requesting application read-only access to your data",
            default_consent_message="View your data",
            description="Provides an application read-only access to a user's data",
            required=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: ID of the application that this role is for.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the OAuth Scope that should be persisted.
        :param pulumi.Input[str] default_consent_detail: "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] default_consent_message: The default message to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] description: A description of the OAuth Scope. This is used for display purposes only.
        :param pulumi.Input[str] name: The name of the Role.
        :param pulumi.Input[bool] required: Determines if the OAuth Scope is required when requested in an OAuth workflow.
        :param pulumi.Input[str] scope_id: The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthApplicationOAuthScopeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Application OAuth Scope Resource

        The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.

        [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        this = fusionauth.FusionAuthApplicationOAuthScope("this",
            application_id=fusionauth_application["this"]["id"],
            data={
                "addedBy": "Tom",
                "addedOn": "2025-02-05",
            },
            default_consent_detail="This will provide the requesting application read-only access to your data",
            default_consent_message="View your data",
            description="Provides an application read-only access to a user's data",
            required=True)
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthApplicationOAuthScopeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthApplicationOAuthScopeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 default_consent_detail: Optional[pulumi.Input[str]] = None,
                 default_consent_message: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 scope_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthApplicationOAuthScopeArgs.__new__(FusionAuthApplicationOAuthScopeArgs)

            if application_id is None and not opts.urn:
                raise TypeError("Missing required property 'application_id'")
            __props__.__dict__["application_id"] = application_id
            __props__.__dict__["data"] = data
            __props__.__dict__["default_consent_detail"] = default_consent_detail
            __props__.__dict__["default_consent_message"] = default_consent_message
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["required"] = required
            __props__.__dict__["scope_id"] = scope_id
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="fusionauth:index/applicationOauthScope:ApplicationOauthScope")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(FusionAuthApplicationOAuthScope, __self__).__init__(
            'fusionauth:index/fusionAuthApplicationOAuthScope:FusionAuthApplicationOAuthScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            default_consent_detail: Optional[pulumi.Input[str]] = None,
            default_consent_message: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            required: Optional[pulumi.Input[bool]] = None,
            scope_id: Optional[pulumi.Input[str]] = None) -> 'FusionAuthApplicationOAuthScope':
        """
        Get an existing FusionAuthApplicationOAuthScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: ID of the application that this role is for.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the OAuth Scope that should be persisted.
        :param pulumi.Input[str] default_consent_detail: "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] default_consent_message: The default message to display on the OAuth consent screen if one cannot be found in the theme.
        :param pulumi.Input[str] description: A description of the OAuth Scope. This is used for display purposes only.
        :param pulumi.Input[str] name: The name of the Role.
        :param pulumi.Input[bool] required: Determines if the OAuth Scope is required when requested in an OAuth workflow.
        :param pulumi.Input[str] scope_id: The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthApplicationOAuthScopeState.__new__(_FusionAuthApplicationOAuthScopeState)

        __props__.__dict__["application_id"] = application_id
        __props__.__dict__["data"] = data
        __props__.__dict__["default_consent_detail"] = default_consent_detail
        __props__.__dict__["default_consent_message"] = default_consent_message
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["required"] = required
        __props__.__dict__["scope_id"] = scope_id
        return FusionAuthApplicationOAuthScope(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        ID of the application that this role is for.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An object that can hold any information about the OAuth Scope that should be persisted.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="defaultConsentDetail")
    def default_consent_detail(self) -> pulumi.Output[Optional[str]]:
        """
        "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_detail")

    @property
    @pulumi.getter(name="defaultConsentMessage")
    def default_consent_message(self) -> pulumi.Output[Optional[str]]:
        """
        The default message to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_message")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the OAuth Scope. This is used for display purposes only.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def required(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines if the OAuth Scope is required when requested in an OAuth workflow.
        """
        return pulumi.get(self, "required")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> pulumi.Output[str]:
        """
        The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "scope_id")

