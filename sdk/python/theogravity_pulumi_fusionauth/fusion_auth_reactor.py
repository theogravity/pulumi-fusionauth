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

__all__ = ['FusionAuthReactorArgs', 'FusionAuthReactor']

@pulumi.input_type
class FusionAuthReactorArgs:
    def __init__(__self__, *,
                 license_id: pulumi.Input[str],
                 license: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthReactor resource.
        :param pulumi.Input[str] license_id: The license Id to activate.
        :param pulumi.Input[str] license: The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        """
        pulumi.set(__self__, "license_id", license_id)
        if license is not None:
            pulumi.set(__self__, "license", license)

    @property
    @pulumi.getter(name="licenseId")
    def license_id(self) -> pulumi.Input[str]:
        """
        The license Id to activate.
        """
        return pulumi.get(self, "license_id")

    @license_id.setter
    def license_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "license_id", value)

    @property
    @pulumi.getter
    def license(self) -> Optional[pulumi.Input[str]]:
        """
        The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        """
        return pulumi.get(self, "license")

    @license.setter
    def license(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license", value)


@pulumi.input_type
class _FusionAuthReactorState:
    def __init__(__self__, *,
                 license: Optional[pulumi.Input[str]] = None,
                 license_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthReactor resources.
        :param pulumi.Input[str] license: The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        :param pulumi.Input[str] license_id: The license Id to activate.
        """
        if license is not None:
            pulumi.set(__self__, "license", license)
        if license_id is not None:
            pulumi.set(__self__, "license_id", license_id)

    @property
    @pulumi.getter
    def license(self) -> Optional[pulumi.Input[str]]:
        """
        The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        """
        return pulumi.get(self, "license")

    @license.setter
    def license(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license", value)

    @property
    @pulumi.getter(name="licenseId")
    def license_id(self) -> Optional[pulumi.Input[str]]:
        """
        The license Id to activate.
        """
        return pulumi.get(self, "license_id")

    @license_id.setter
    def license_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "license_id", value)


class FusionAuthReactor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 license: Optional[pulumi.Input[str]] = None,
                 license_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Reactor Resource

        The Reactor is FusionAuth’s license system. Reactor is used to activate features based upon your licensing tier.

        [Reactor API](https://fusionauth.io/docs/v1/tech/apis/reactor)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        reactor = fusionauth.FusionAuthReactor("reactor",
            license="abc",
            license_id="xyz")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] license: The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        :param pulumi.Input[str] license_id: The license Id to activate.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthReactorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Reactor Resource

        The Reactor is FusionAuth’s license system. Reactor is used to activate features based upon your licensing tier.

        [Reactor API](https://fusionauth.io/docs/v1/tech/apis/reactor)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        reactor = fusionauth.FusionAuthReactor("reactor",
            license="abc",
            license_id="xyz")
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthReactorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthReactorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 license: Optional[pulumi.Input[str]] = None,
                 license_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthReactorArgs.__new__(FusionAuthReactorArgs)

            __props__.__dict__["license"] = None if license is None else pulumi.Output.secret(license)
            if license_id is None and not opts.urn:
                raise TypeError("Missing required property 'license_id'")
            __props__.__dict__["license_id"] = None if license_id is None else pulumi.Output.secret(license_id)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["license", "licenseId"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(FusionAuthReactor, __self__).__init__(
            'fusionauth:index/fusionAuthReactor:FusionAuthReactor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            license: Optional[pulumi.Input[str]] = None,
            license_id: Optional[pulumi.Input[str]] = None) -> 'FusionAuthReactor':
        """
        Get an existing FusionAuthReactor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] license: The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        :param pulumi.Input[str] license_id: The license Id to activate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthReactorState.__new__(_FusionAuthReactorState)

        __props__.__dict__["license"] = license
        __props__.__dict__["license_id"] = license_id
        return FusionAuthReactor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def license(self) -> pulumi.Output[Optional[str]]:
        """
        The Base64 encoded license value. This value is necessary in an air gapped configuration where outbound network access is not available.
        """
        return pulumi.get(self, "license")

    @property
    @pulumi.getter(name="licenseId")
    def license_id(self) -> pulumi.Output[str]:
        """
        The license Id to activate.
        """
        return pulumi.get(self, "license_id")

