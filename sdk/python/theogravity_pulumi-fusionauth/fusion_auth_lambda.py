# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FusionAuthLambdaArgs', 'FusionAuthLambda']

@pulumi.input_type
class FusionAuthLambdaArgs:
    def __init__(__self__, *,
                 body: pulumi.Input[str],
                 type: pulumi.Input[str],
                 debug: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lambda_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthLambda resource.
        :param pulumi.Input[str] body: The lambda function body, a JavaScript function.
        :param pulumi.Input[str] type: The lambda type. The possible values are:
               - `JWTPopulate`
               - `OpenIDReconcile`
               - `SAMLv2Reconcile`
               - `SAMLv2Populate`
               - `AppleReconcile`
               - `ExternalJWTReconcile`
               - `FacebookReconcile`
               - `GoogleReconcile`
               - `HYPRReconcile`
               - `TwitterReconcile`
               - `LDAPConnectorReconcile`
               - `LinkedInReconcile`
               - `EpicGamesReconcile`
               - `NintendoReconcile`
               - `SonyPSNReconcile`
               - `SteamReconcile`
               - `TwitchReconcile`
               - `XboxReconcile`
               - `SelfServiceRegistrationValidation`
               - `ClientCredentialsJWTPopulate`
        :param pulumi.Input[bool] debug: Whether or not debug event logging is enabled for this Lambda.
        :param pulumi.Input[bool] enabled: Whether or not this Lambda is enabled.
        :param pulumi.Input[str] engine_type: The JavaScript execution engine for the lambda.
        :param pulumi.Input[str] lambda_id: The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The name of the lambda.
        """
        pulumi.set(__self__, "body", body)
        pulumi.set(__self__, "type", type)
        if debug is not None:
            pulumi.set(__self__, "debug", debug)
        if enabled is not None:
            warnings.warn("""Not currently used and may be removed in a future version.""", DeprecationWarning)
            pulumi.log.warn("""enabled is deprecated: Not currently used and may be removed in a future version.""")
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if lambda_id is not None:
            pulumi.set(__self__, "lambda_id", lambda_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def body(self) -> pulumi.Input[str]:
        """
        The lambda function body, a JavaScript function.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: pulumi.Input[str]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The lambda type. The possible values are:
        - `JWTPopulate`
        - `OpenIDReconcile`
        - `SAMLv2Reconcile`
        - `SAMLv2Populate`
        - `AppleReconcile`
        - `ExternalJWTReconcile`
        - `FacebookReconcile`
        - `GoogleReconcile`
        - `HYPRReconcile`
        - `TwitterReconcile`
        - `LDAPConnectorReconcile`
        - `LinkedInReconcile`
        - `EpicGamesReconcile`
        - `NintendoReconcile`
        - `SonyPSNReconcile`
        - `SteamReconcile`
        - `TwitchReconcile`
        - `XboxReconcile`
        - `SelfServiceRegistrationValidation`
        - `ClientCredentialsJWTPopulate`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def debug(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not debug event logging is enabled for this Lambda.
        """
        return pulumi.get(self, "debug")

    @debug.setter
    def debug(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "debug", value)

    @property
    @pulumi.getter
    @_utilities.deprecated("""Not currently used and may be removed in a future version.""")
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not this Lambda is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The JavaScript execution engine for the lambda.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter(name="lambdaId")
    def lambda_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "lambda_id")

    @lambda_id.setter
    def lambda_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the lambda.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FusionAuthLambdaState:
    def __init__(__self__, *,
                 body: Optional[pulumi.Input[str]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lambda_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthLambda resources.
        :param pulumi.Input[str] body: The lambda function body, a JavaScript function.
        :param pulumi.Input[bool] debug: Whether or not debug event logging is enabled for this Lambda.
        :param pulumi.Input[bool] enabled: Whether or not this Lambda is enabled.
        :param pulumi.Input[str] engine_type: The JavaScript execution engine for the lambda.
        :param pulumi.Input[str] lambda_id: The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The name of the lambda.
        :param pulumi.Input[str] type: The lambda type. The possible values are:
               - `JWTPopulate`
               - `OpenIDReconcile`
               - `SAMLv2Reconcile`
               - `SAMLv2Populate`
               - `AppleReconcile`
               - `ExternalJWTReconcile`
               - `FacebookReconcile`
               - `GoogleReconcile`
               - `HYPRReconcile`
               - `TwitterReconcile`
               - `LDAPConnectorReconcile`
               - `LinkedInReconcile`
               - `EpicGamesReconcile`
               - `NintendoReconcile`
               - `SonyPSNReconcile`
               - `SteamReconcile`
               - `TwitchReconcile`
               - `XboxReconcile`
               - `SelfServiceRegistrationValidation`
               - `ClientCredentialsJWTPopulate`
        """
        if body is not None:
            pulumi.set(__self__, "body", body)
        if debug is not None:
            pulumi.set(__self__, "debug", debug)
        if enabled is not None:
            warnings.warn("""Not currently used and may be removed in a future version.""", DeprecationWarning)
            pulumi.log.warn("""enabled is deprecated: Not currently used and may be removed in a future version.""")
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if engine_type is not None:
            pulumi.set(__self__, "engine_type", engine_type)
        if lambda_id is not None:
            pulumi.set(__self__, "lambda_id", lambda_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def body(self) -> Optional[pulumi.Input[str]]:
        """
        The lambda function body, a JavaScript function.
        """
        return pulumi.get(self, "body")

    @body.setter
    def body(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "body", value)

    @property
    @pulumi.getter
    def debug(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not debug event logging is enabled for this Lambda.
        """
        return pulumi.get(self, "debug")

    @debug.setter
    def debug(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "debug", value)

    @property
    @pulumi.getter
    @_utilities.deprecated("""Not currently used and may be removed in a future version.""")
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not this Lambda is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> Optional[pulumi.Input[str]]:
        """
        The JavaScript execution engine for the lambda.
        """
        return pulumi.get(self, "engine_type")

    @engine_type.setter
    def engine_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_type", value)

    @property
    @pulumi.getter(name="lambdaId")
    def lambda_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "lambda_id")

    @lambda_id.setter
    def lambda_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the lambda.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The lambda type. The possible values are:
        - `JWTPopulate`
        - `OpenIDReconcile`
        - `SAMLv2Reconcile`
        - `SAMLv2Populate`
        - `AppleReconcile`
        - `ExternalJWTReconcile`
        - `FacebookReconcile`
        - `GoogleReconcile`
        - `HYPRReconcile`
        - `TwitterReconcile`
        - `LDAPConnectorReconcile`
        - `LinkedInReconcile`
        - `EpicGamesReconcile`
        - `NintendoReconcile`
        - `SonyPSNReconcile`
        - `SteamReconcile`
        - `TwitchReconcile`
        - `XboxReconcile`
        - `SelfServiceRegistrationValidation`
        - `ClientCredentialsJWTPopulate`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FusionAuthLambda(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lambda_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Lambda Resource

        Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.

        [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi-fusionauth as fusionauth

        preferred__username = fusionauth.FusionAuthLambda("preferred Username",
            body=\"\"\"// Using the user and registration parameters add additional values to the jwt object.
        function populate(jwt, user, registration) {
          jwt.preferred_username = registration.username;
        }
          
        \"\"\",
            enabled=True,
            type="JWTPopulate")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: The lambda function body, a JavaScript function.
        :param pulumi.Input[bool] debug: Whether or not debug event logging is enabled for this Lambda.
        :param pulumi.Input[bool] enabled: Whether or not this Lambda is enabled.
        :param pulumi.Input[str] engine_type: The JavaScript execution engine for the lambda.
        :param pulumi.Input[str] lambda_id: The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The name of the lambda.
        :param pulumi.Input[str] type: The lambda type. The possible values are:
               - `JWTPopulate`
               - `OpenIDReconcile`
               - `SAMLv2Reconcile`
               - `SAMLv2Populate`
               - `AppleReconcile`
               - `ExternalJWTReconcile`
               - `FacebookReconcile`
               - `GoogleReconcile`
               - `HYPRReconcile`
               - `TwitterReconcile`
               - `LDAPConnectorReconcile`
               - `LinkedInReconcile`
               - `EpicGamesReconcile`
               - `NintendoReconcile`
               - `SonyPSNReconcile`
               - `SteamReconcile`
               - `TwitchReconcile`
               - `XboxReconcile`
               - `SelfServiceRegistrationValidation`
               - `ClientCredentialsJWTPopulate`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthLambdaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Lambda Resource

        Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.

        [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi-fusionauth as fusionauth

        preferred__username = fusionauth.FusionAuthLambda("preferred Username",
            body=\"\"\"// Using the user and registration parameters add additional values to the jwt object.
        function populate(jwt, user, registration) {
          jwt.preferred_username = registration.username;
        }
          
        \"\"\",
            enabled=True,
            type="JWTPopulate")
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthLambdaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthLambdaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 body: Optional[pulumi.Input[str]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 lambda_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthLambdaArgs.__new__(FusionAuthLambdaArgs)

            if body is None and not opts.urn:
                raise TypeError("Missing required property 'body'")
            __props__.__dict__["body"] = body
            __props__.__dict__["debug"] = debug
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["engine_type"] = engine_type
            __props__.__dict__["lambda_id"] = lambda_id
            __props__.__dict__["name"] = name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(FusionAuthLambda, __self__).__init__(
            'fusionauth:index/fusionAuthLambda:FusionAuthLambda',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            body: Optional[pulumi.Input[str]] = None,
            debug: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            engine_type: Optional[pulumi.Input[str]] = None,
            lambda_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FusionAuthLambda':
        """
        Get an existing FusionAuthLambda resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] body: The lambda function body, a JavaScript function.
        :param pulumi.Input[bool] debug: Whether or not debug event logging is enabled for this Lambda.
        :param pulumi.Input[bool] enabled: Whether or not this Lambda is enabled.
        :param pulumi.Input[str] engine_type: The JavaScript execution engine for the lambda.
        :param pulumi.Input[str] lambda_id: The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The name of the lambda.
        :param pulumi.Input[str] type: The lambda type. The possible values are:
               - `JWTPopulate`
               - `OpenIDReconcile`
               - `SAMLv2Reconcile`
               - `SAMLv2Populate`
               - `AppleReconcile`
               - `ExternalJWTReconcile`
               - `FacebookReconcile`
               - `GoogleReconcile`
               - `HYPRReconcile`
               - `TwitterReconcile`
               - `LDAPConnectorReconcile`
               - `LinkedInReconcile`
               - `EpicGamesReconcile`
               - `NintendoReconcile`
               - `SonyPSNReconcile`
               - `SteamReconcile`
               - `TwitchReconcile`
               - `XboxReconcile`
               - `SelfServiceRegistrationValidation`
               - `ClientCredentialsJWTPopulate`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthLambdaState.__new__(_FusionAuthLambdaState)

        __props__.__dict__["body"] = body
        __props__.__dict__["debug"] = debug
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["engine_type"] = engine_type
        __props__.__dict__["lambda_id"] = lambda_id
        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        return FusionAuthLambda(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def body(self) -> pulumi.Output[str]:
        """
        The lambda function body, a JavaScript function.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def debug(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not debug event logging is enabled for this Lambda.
        """
        return pulumi.get(self, "debug")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Not currently used and may be removed in a future version.""")
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not this Lambda is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> pulumi.Output[Optional[str]]:
        """
        The JavaScript execution engine for the lambda.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter(name="lambdaId")
    def lambda_id(self) -> pulumi.Output[str]:
        """
        The Id to use for the new lambda. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "lambda_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the lambda.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The lambda type. The possible values are:
        - `JWTPopulate`
        - `OpenIDReconcile`
        - `SAMLv2Reconcile`
        - `SAMLv2Populate`
        - `AppleReconcile`
        - `ExternalJWTReconcile`
        - `FacebookReconcile`
        - `GoogleReconcile`
        - `HYPRReconcile`
        - `TwitterReconcile`
        - `LDAPConnectorReconcile`
        - `LinkedInReconcile`
        - `EpicGamesReconcile`
        - `NintendoReconcile`
        - `SonyPSNReconcile`
        - `SteamReconcile`
        - `TwitchReconcile`
        - `XboxReconcile`
        - `SelfServiceRegistrationValidation`
        - `ClientCredentialsJWTPopulate`
        """
        return pulumi.get(self, "type")

