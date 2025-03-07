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

__all__ = ['FusionAuthEntityGrantArgs', 'FusionAuthEntityGrant']

@pulumi.input_type
class FusionAuthEntityGrantArgs:
    def __init__(__self__, *,
                 entity_id: pulumi.Input[str],
                 data: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipient_entity_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthEntityGrant resource.
        :param pulumi.Input[str] entity_id: The Id of the Entity to which access is granted.
        :param pulumi.Input[str] data: An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The set of permissions of this Grant.
        :param pulumi.Input[str] recipient_entity_id: The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        :param pulumi.Input[str] tenant_id: The unique Id of the tenant used to scope this API request.
        :param pulumi.Input[str] user_id: The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        pulumi.set(__self__, "entity_id", entity_id)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if recipient_entity_id is not None:
            pulumi.set(__self__, "recipient_entity_id", recipient_entity_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Input[str]:
        """
        The Id of the Entity to which access is granted.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of permissions of this Grant.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="recipientEntityId")
    def recipient_entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        """
        return pulumi.get(self, "recipient_entity_id")

    @recipient_entity_id.setter
    def recipient_entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recipient_entity_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique Id of the tenant used to scope this API request.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _FusionAuthEntityGrantState:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipient_entity_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthEntityGrant resources.
        :param pulumi.Input[str] data: An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        :param pulumi.Input[str] entity_id: The Id of the Entity to which access is granted.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The set of permissions of this Grant.
        :param pulumi.Input[str] recipient_entity_id: The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        :param pulumi.Input[str] tenant_id: The unique Id of the tenant used to scope this API request.
        :param pulumi.Input[str] user_id: The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if entity_id is not None:
            pulumi.set(__self__, "entity_id", entity_id)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if recipient_entity_id is not None:
            pulumi.set(__self__, "recipient_entity_id", recipient_entity_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of the Entity to which access is granted.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of permissions of this Grant.
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="recipientEntityId")
    def recipient_entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        """
        return pulumi.get(self, "recipient_entity_id")

    @recipient_entity_id.setter
    def recipient_entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recipient_entity_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        The unique Id of the tenant used to scope this API request.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class FusionAuthEntityGrant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipient_entity_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Entity Grant Resource

        Entities can have Grants. Grants are relationships between a target Entity and one of two other types:

        * A Recipient Entity
        * A User.

        Grants can have zero or more Permissions associated with them.

        [Entity Grant API](https://fusionauth.io/docs/v1/tech/apis/entity-management/grants)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        :param pulumi.Input[str] entity_id: The Id of the Entity to which access is granted.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The set of permissions of this Grant.
        :param pulumi.Input[str] recipient_entity_id: The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        :param pulumi.Input[str] tenant_id: The unique Id of the tenant used to scope this API request.
        :param pulumi.Input[str] user_id: The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthEntityGrantArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Entity Grant Resource

        Entities can have Grants. Grants are relationships between a target Entity and one of two other types:

        * A Recipient Entity
        * A User.

        Grants can have zero or more Permissions associated with them.

        [Entity Grant API](https://fusionauth.io/docs/v1/tech/apis/entity-management/grants)

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param FusionAuthEntityGrantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthEntityGrantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 recipient_entity_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthEntityGrantArgs.__new__(FusionAuthEntityGrantArgs)

            __props__.__dict__["data"] = data
            if entity_id is None and not opts.urn:
                raise TypeError("Missing required property 'entity_id'")
            __props__.__dict__["entity_id"] = entity_id
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["recipient_entity_id"] = recipient_entity_id
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["user_id"] = user_id
        super(FusionAuthEntityGrant, __self__).__init__(
            'fusionauth:index/fusionAuthEntityGrant:FusionAuthEntityGrant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[str]] = None,
            entity_id: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            recipient_entity_id: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'FusionAuthEntityGrant':
        """
        Get an existing FusionAuthEntityGrant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        :param pulumi.Input[str] entity_id: The Id of the Entity to which access is granted.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: The set of permissions of this Grant.
        :param pulumi.Input[str] recipient_entity_id: The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        :param pulumi.Input[str] tenant_id: The unique Id of the tenant used to scope this API request.
        :param pulumi.Input[str] user_id: The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthEntityGrantState.__new__(_FusionAuthEntityGrantState)

        __props__.__dict__["data"] = data
        __props__.__dict__["entity_id"] = entity_id
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["recipient_entity_id"] = recipient_entity_id
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["user_id"] = user_id
        return FusionAuthEntityGrant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[str]]:
        """
        An object that can hold any information about the Grant that should be persisted. Please review the limits on data field types as you plan for and build your custom data schema.  Must be a JSON string.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Output[str]:
        """
        The Id of the Entity to which access is granted.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of permissions of this Grant.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="recipientEntityId")
    def recipient_entity_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Entity Id for which access is granted. If `recipient_entity_id` is not provided, then the `user_id` will be required.
        """
        return pulumi.get(self, "recipient_entity_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The unique Id of the tenant used to scope this API request.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        The User Id for which access is granted. If `user_id` is not provided, then the `recipient_entity_id` will be required.
        """
        return pulumi.get(self, "user_id")

