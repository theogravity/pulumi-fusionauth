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
    'GetUserGroupMembershipResult',
    'AwaitableGetUserGroupMembershipResult',
    'get_user_group_membership',
    'get_user_group_membership_output',
]

@pulumi.output_type
class GetUserGroupMembershipResult:
    """
    A collection of values returned by getUserGroupMembership.
    """
    def __init__(__self__, data=None, group_id=None, id=None, membership_id=None, user_id=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if membership_id and not isinstance(membership_id, str):
            raise TypeError("Expected argument 'membership_id' to be a str")
        pulumi.set(__self__, "membership_id", membership_id)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def data(self) -> Mapping[str, str]:
        """
        (Optional) An object that can hold any information about the User for this membership that should be persisted.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="membershipId")
    def membership_id(self) -> str:
        """
        (Optional) The Id of the User Group Membership. If not provided, a random UUID will be generated.
        """
        return pulumi.get(self, "membership_id")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")


class AwaitableGetUserGroupMembershipResult(GetUserGroupMembershipResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserGroupMembershipResult(
            data=self.data,
            group_id=self.group_id,
            id=self.id,
            membership_id=self.membership_id,
            user_id=self.user_id)


def get_user_group_membership(group_id: Optional[str] = None,
                              user_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserGroupMembershipResult:
    """
    ## # User Group Membership Resource

    [User Group Membership API](https://fusionauth.io/docs/apis/groups#request-5)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    this = fusionauth.get_user_group_membership(group_id=fusionauth_group["this"]["id"],
        user_id=fusionauth_user["this"]["id"])
    ```


    :param str group_id: The Id of the Group of this membership.
    :param str user_id: "The Id of the User of this membership.
    """
    __args__ = dict()
    __args__['groupId'] = group_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getUserGroupMembership:getUserGroupMembership', __args__, opts=opts, typ=GetUserGroupMembershipResult).value

    return AwaitableGetUserGroupMembershipResult(
        data=pulumi.get(__ret__, 'data'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        membership_id=pulumi.get(__ret__, 'membership_id'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_user_group_membership_output(group_id: Optional[pulumi.Input[str]] = None,
                                     user_id: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetUserGroupMembershipResult]:
    """
    ## # User Group Membership Resource

    [User Group Membership API](https://fusionauth.io/docs/apis/groups#request-5)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    this = fusionauth.get_user_group_membership(group_id=fusionauth_group["this"]["id"],
        user_id=fusionauth_user["this"]["id"])
    ```


    :param str group_id: The Id of the Group of this membership.
    :param str user_id: "The Id of the User of this membership.
    """
    __args__ = dict()
    __args__['groupId'] = group_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('fusionauth:index/getUserGroupMembership:getUserGroupMembership', __args__, opts=opts, typ=GetUserGroupMembershipResult)
    return __ret__.apply(lambda __response__: GetUserGroupMembershipResult(
        data=pulumi.get(__response__, 'data'),
        group_id=pulumi.get(__response__, 'group_id'),
        id=pulumi.get(__response__, 'id'),
        membership_id=pulumi.get(__response__, 'membership_id'),
        user_id=pulumi.get(__response__, 'user_id')))
