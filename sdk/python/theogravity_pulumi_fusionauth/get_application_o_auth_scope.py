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

__all__ = [
    'GetApplicationOAuthScopeResult',
    'AwaitableGetApplicationOAuthScopeResult',
    'get_application_o_auth_scope',
    'get_application_o_auth_scope_output',
]

@pulumi.output_type
class GetApplicationOAuthScopeResult:
    """
    A collection of values returned by getApplicationOAuthScope.
    """
    def __init__(__self__, application_id=None, data=None, default_consent_detail=None, default_consent_message=None, description=None, id=None, name=None, required=None, scope_id=None):
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        pulumi.set(__self__, "application_id", application_id)
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if default_consent_detail and not isinstance(default_consent_detail, str):
            raise TypeError("Expected argument 'default_consent_detail' to be a str")
        pulumi.set(__self__, "default_consent_detail", default_consent_detail)
        if default_consent_message and not isinstance(default_consent_message, str):
            raise TypeError("Expected argument 'default_consent_message' to be a str")
        pulumi.set(__self__, "default_consent_message", default_consent_message)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if required and not isinstance(required, bool):
            raise TypeError("Expected argument 'required' to be a bool")
        pulumi.set(__self__, "required", required)
        if scope_id and not isinstance(scope_id, str):
            raise TypeError("Expected argument 'scope_id' to be a str")
        pulumi.set(__self__, "scope_id", scope_id)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> str:
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def data(self) -> Mapping[str, str]:
        """
        (Optional) An object that can hold any information about the OAuth Scope that should be persisted.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="defaultConsentDetail")
    def default_consent_detail(self) -> str:
        """
        (Optional) "The default detail to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_detail")

    @property
    @pulumi.getter(name="defaultConsentMessage")
    def default_consent_message(self) -> str:
        """
        (Optional) The default message to display on the OAuth consent screen if one cannot be found in the theme.
        """
        return pulumi.get(self, "default_consent_message")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (Optional) A description of the OAuth Scope. This is used for display purposes only.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def required(self) -> bool:
        """
        (Optional) Determines if the OAuth Scope is required when requested in an OAuth workflow.
        """
        return pulumi.get(self, "required")

    @property
    @pulumi.getter(name="scopeId")
    def scope_id(self) -> str:
        """
        (Optional) The Id to use for the new OAuth Scope. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "scope_id")


class AwaitableGetApplicationOAuthScopeResult(GetApplicationOAuthScopeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationOAuthScopeResult(
            application_id=self.application_id,
            data=self.data,
            default_consent_detail=self.default_consent_detail,
            default_consent_message=self.default_consent_message,
            description=self.description,
            id=self.id,
            name=self.name,
            required=self.required,
            scope_id=self.scope_id)


def get_application_o_auth_scope(application_id: Optional[str] = None,
                                 name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationOAuthScopeResult:
    """
    ## # Application OAuth Scope Resource

    The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.

    [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    this = fusionauth.get_application_o_auth_scope(application_id=data["fusionauth_application"]["this"]["id"],
        name="data:read")
    ```


    :param str application_id: ID of the application that this role is for.
    :param str name: The name of the Role.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getApplicationOAuthScope:getApplicationOAuthScope', __args__, opts=opts, typ=GetApplicationOAuthScopeResult).value

    return AwaitableGetApplicationOAuthScopeResult(
        application_id=pulumi.get(__ret__, 'application_id'),
        data=pulumi.get(__ret__, 'data'),
        default_consent_detail=pulumi.get(__ret__, 'default_consent_detail'),
        default_consent_message=pulumi.get(__ret__, 'default_consent_message'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        required=pulumi.get(__ret__, 'required'),
        scope_id=pulumi.get(__ret__, 'scope_id'))
def get_application_o_auth_scope_output(application_id: Optional[pulumi.Input[str]] = None,
                                        name: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetApplicationOAuthScopeResult]:
    """
    ## # Application OAuth Scope Resource

    The Application OAuth Scope resource allows you to define the scopes that an application can request when using OAuth.

    [Application OAuth Scope API](https://fusionauth.io/docs/apis/scopes)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    this = fusionauth.get_application_o_auth_scope(application_id=data["fusionauth_application"]["this"]["id"],
        name="data:read")
    ```


    :param str application_id: ID of the application that this role is for.
    :param str name: The name of the Role.
    """
    __args__ = dict()
    __args__['applicationId'] = application_id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('fusionauth:index/getApplicationOAuthScope:getApplicationOAuthScope', __args__, opts=opts, typ=GetApplicationOAuthScopeResult)
    return __ret__.apply(lambda __response__: GetApplicationOAuthScopeResult(
        application_id=pulumi.get(__response__, 'application_id'),
        data=pulumi.get(__response__, 'data'),
        default_consent_detail=pulumi.get(__response__, 'default_consent_detail'),
        default_consent_message=pulumi.get(__response__, 'default_consent_message'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        required=pulumi.get(__response__, 'required'),
        scope_id=pulumi.get(__response__, 'scope_id')))
