# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetLambdaResult',
    'AwaitableGetLambdaResult',
    'get_lambda',
    'get_lambda_output',
]

@pulumi.output_type
class GetLambdaResult:
    """
    A collection of values returned by getLambda.
    """
    def __init__(__self__, body=None, debug=None, id=None, name=None, type=None):
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if debug and not isinstance(debug, bool):
            raise TypeError("Expected argument 'debug' to be a bool")
        pulumi.set(__self__, "debug", debug)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def body(self) -> str:
        """
        The lambda function body, a JavaScript function.
        """
        return pulumi.get(self, "body")

    @property
    @pulumi.getter
    def debug(self) -> bool:
        """
        Whether or not debug event logging is enabled for this Lambda.
        """
        return pulumi.get(self, "debug")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetLambdaResult(GetLambdaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLambdaResult(
            body=self.body,
            debug=self.debug,
            id=self.id,
            name=self.name,
            type=self.type)


def get_lambda(id: Optional[str] = None,
               name: Optional[str] = None,
               type: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLambdaResult:
    """
    ## # Lambda Resource

    Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.

    [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default_google_reconcile = fusionauth.get_lambda(name="Default Google Reconcile provided by FusionAuth",
        type="GoogleReconcile")
    ```


    :param str id: The ID of the Lambda. At least one of `id` or `name` must be specified.
    :param str name: The name of the Lambda. At least one of `id` or `name` must be specified.
    :param str type: The Lambda type. The possible values are:
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getLambda:getLambda', __args__, opts=opts, typ=GetLambdaResult).value

    return AwaitableGetLambdaResult(
        body=__ret__.body,
        debug=__ret__.debug,
        id=__ret__.id,
        name=__ret__.name,
        type=__ret__.type)


@_utilities.lift_output_func(get_lambda)
def get_lambda_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      type: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLambdaResult]:
    """
    ## # Lambda Resource

    Lambdas are user defined JavaScript functions that may be executed at runtime to perform various functions. Lambdas may be used to customize the claims returned in a JWT, reconcile a SAML v2 response or an OpenID Connect response when using these external identity providers.

    [Lambdas API](https://fusionauth.io/docs/v1/tech/apis/lambdas)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default_google_reconcile = fusionauth.get_lambda(name="Default Google Reconcile provided by FusionAuth",
        type="GoogleReconcile")
    ```


    :param str id: The ID of the Lambda. At least one of `id` or `name` must be specified.
    :param str name: The name of the Lambda. At least one of `id` or `name` must be specified.
    :param str type: The Lambda type. The possible values are:
    """
    ...
