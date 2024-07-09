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

__all__ = ['FusionAuthApiKeyArgs', 'FusionAuthApiKey']

@pulumi.input_type
class FusionAuthApiKeyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_access_control_list_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 permissions_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthApiKey resource.
        :param pulumi.Input[str] description: Description of the key.
        :param pulumi.Input[str] ip_access_control_list_id: The Id of the IP Access Control List limiting access to this API key.
        :param pulumi.Input[str] key: API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        :param pulumi.Input[str] key_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]] permissions_endpoints: The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        :param pulumi.Input[str] tenant_id: The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_access_control_list_id is not None:
            pulumi.set(__self__, "ip_access_control_list_id", ip_access_control_list_id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if permissions_endpoints is not None:
            pulumi.set(__self__, "permissions_endpoints", permissions_endpoints)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the key.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAccessControlListId")
    def ip_access_control_list_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of the IP Access Control List limiting access to this API key.
        """
        return pulumi.get(self, "ip_access_control_list_id")

    @ip_access_control_list_id.setter
    def ip_access_control_list_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_access_control_list_id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new Form. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="permissionsEndpoints")
    def permissions_endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]]]:
        """
        The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        """
        return pulumi.get(self, "permissions_endpoints")

    @permissions_endpoints.setter
    def permissions_endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]]]):
        pulumi.set(self, "permissions_endpoints", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _FusionAuthApiKeyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_access_control_list_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 permissions_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthApiKey resources.
        :param pulumi.Input[str] description: Description of the key.
        :param pulumi.Input[str] ip_access_control_list_id: The Id of the IP Access Control List limiting access to this API key.
        :param pulumi.Input[str] key: API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        :param pulumi.Input[str] key_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]] permissions_endpoints: The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        :param pulumi.Input[str] tenant_id: The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_access_control_list_id is not None:
            pulumi.set(__self__, "ip_access_control_list_id", ip_access_control_list_id)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if permissions_endpoints is not None:
            pulumi.set(__self__, "permissions_endpoints", permissions_endpoints)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the key.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAccessControlListId")
    def ip_access_control_list_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of the IP Access Control List limiting access to this API key.
        """
        return pulumi.get(self, "ip_access_control_list_id")

    @ip_access_control_list_id.setter
    def ip_access_control_list_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_access_control_list_id", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new Form. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="permissionsEndpoints")
    def permissions_endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]]]:
        """
        The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        """
        return pulumi.get(self, "permissions_endpoints")

    @permissions_endpoints.setter
    def permissions_endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthApiKeyPermissionsEndpointArgs']]]]):
        pulumi.set(self, "permissions_endpoints", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


class FusionAuthApiKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_access_control_list_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 permissions_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthApiKeyPermissionsEndpointArgs']]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # API Key

        The FusionAuth APIs are primarily secured using API keys. This API can only be accessed using an API key that is set as a keyManager. In order to retrieve, update or delete an API key, an API key with equal or greater permissions must be used. A "tenant-scoped" API key can retrieve, create, update or delete an API key for the same tenant. This page describes APIs that are used to manage API keys.

        [API Key](https://fusionauth.io/docs/v1/tech/apis/api-keys/)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi-fusionauth as fusionauth

        example = fusionauth.FusionAuthApiKey("example",
            description="my super secret key",
            key="super-secret-key",
            permissions_endpoints=[fusionauth.FusionAuthApiKeyPermissionsEndpointArgs(
                delete=True,
                endpoint="/api/application",
                get=True,
                patch=True,
                post=True,
                put=True,
            )],
            tenant_id="94f751c5-4883-4684-a817-6b106778edec")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the key.
        :param pulumi.Input[str] ip_access_control_list_id: The Id of the IP Access Control List limiting access to this API key.
        :param pulumi.Input[str] key: API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        :param pulumi.Input[str] key_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthApiKeyPermissionsEndpointArgs']]]] permissions_endpoints: The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        :param pulumi.Input[str] tenant_id: The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FusionAuthApiKeyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # API Key

        The FusionAuth APIs are primarily secured using API keys. This API can only be accessed using an API key that is set as a keyManager. In order to retrieve, update or delete an API key, an API key with equal or greater permissions must be used. A "tenant-scoped" API key can retrieve, create, update or delete an API key for the same tenant. This page describes APIs that are used to manage API keys.

        [API Key](https://fusionauth.io/docs/v1/tech/apis/api-keys/)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi-fusionauth as fusionauth

        example = fusionauth.FusionAuthApiKey("example",
            description="my super secret key",
            key="super-secret-key",
            permissions_endpoints=[fusionauth.FusionAuthApiKeyPermissionsEndpointArgs(
                delete=True,
                endpoint="/api/application",
                get=True,
                patch=True,
                post=True,
                put=True,
            )],
            tenant_id="94f751c5-4883-4684-a817-6b106778edec")
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthApiKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthApiKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_access_control_list_id: Optional[pulumi.Input[str]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 permissions_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthApiKeyPermissionsEndpointArgs']]]]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthApiKeyArgs.__new__(FusionAuthApiKeyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ip_access_control_list_id"] = ip_access_control_list_id
            __props__.__dict__["key"] = None if key is None else pulumi.Output.secret(key)
            __props__.__dict__["key_id"] = key_id
            __props__.__dict__["permissions_endpoints"] = permissions_endpoints
            __props__.__dict__["tenant_id"] = tenant_id
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["key"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(FusionAuthApiKey, __self__).__init__(
            'fusionauth:index/fusionAuthApiKey:FusionAuthApiKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_access_control_list_id: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            permissions_endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthApiKeyPermissionsEndpointArgs']]]]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'FusionAuthApiKey':
        """
        Get an existing FusionAuthApiKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the key.
        :param pulumi.Input[str] ip_access_control_list_id: The Id of the IP Access Control List limiting access to this API key.
        :param pulumi.Input[str] key: API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        :param pulumi.Input[str] key_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FusionAuthApiKeyPermissionsEndpointArgs']]]] permissions_endpoints: The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        :param pulumi.Input[str] tenant_id: The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthApiKeyState.__new__(_FusionAuthApiKeyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["ip_access_control_list_id"] = ip_access_control_list_id
        __props__.__dict__["key"] = key
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["permissions_endpoints"] = permissions_endpoints
        __props__.__dict__["tenant_id"] = tenant_id
        return FusionAuthApiKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the key.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipAccessControlListId")
    def ip_access_control_list_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id of the IP Access Control List limiting access to this API key.
        """
        return pulumi.get(self, "ip_access_control_list_id")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        API key string. When you create an API key the key is defaulted to a secure random value but the API key is simply a string, so you may call it super-secret-key if you’d like. However a long and random value makes a good API key in that it is unique and difficult to guess.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id to use for the new Form. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="permissionsEndpoints")
    def permissions_endpoints(self) -> pulumi.Output[Optional[Sequence['outputs.FusionAuthApiKeyPermissionsEndpoint']]]:
        """
        The unique Id of the private key downloaded from Apple and imported into Key Master that will be used to sign the client secret.
        """
        return pulumi.get(self, "permissions_endpoints")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        The unique Id of the Tenant. This value is required if the key is meant to be tenant scoped. Tenant scoped keys can only be used to access users and other tenant scoped objects for the specified tenant. This value is read-only once the key is created.
        """
        return pulumi.get(self, "tenant_id")

