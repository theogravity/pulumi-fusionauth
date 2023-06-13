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

__all__ = [
    'GetFormFieldResult',
    'AwaitableGetFormFieldResult',
    'get_form_field',
    'get_form_field_output',
]

@pulumi.output_type
class GetFormFieldResult:
    """
    A collection of values returned by getFormField.
    """
    def __init__(__self__, confirm=None, consent_id=None, control=None, data=None, description=None, form_field_id=None, id=None, key=None, name=None, options=None, required=None, type=None, validator=None):
        if confirm and not isinstance(confirm, bool):
            raise TypeError("Expected argument 'confirm' to be a bool")
        pulumi.set(__self__, "confirm", confirm)
        if consent_id and not isinstance(consent_id, str):
            raise TypeError("Expected argument 'consent_id' to be a str")
        pulumi.set(__self__, "consent_id", consent_id)
        if control and not isinstance(control, str):
            raise TypeError("Expected argument 'control' to be a str")
        pulumi.set(__self__, "control", control)
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if form_field_id and not isinstance(form_field_id, str):
            raise TypeError("Expected argument 'form_field_id' to be a str")
        pulumi.set(__self__, "form_field_id", form_field_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if options and not isinstance(options, list):
            raise TypeError("Expected argument 'options' to be a list")
        pulumi.set(__self__, "options", options)
        if required and not isinstance(required, bool):
            raise TypeError("Expected argument 'required' to be a bool")
        pulumi.set(__self__, "required", required)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if validator and not isinstance(validator, dict):
            raise TypeError("Expected argument 'validator' to be a dict")
        pulumi.set(__self__, "validator", validator)

    @property
    @pulumi.getter
    def confirm(self) -> Optional[bool]:
        """
        Determines if the user input should be confirmed by requiring the value to be entered twice.
        - consent_id
        - control
        """
        return pulumi.get(self, "confirm")

    @property
    @pulumi.getter(name="consentId")
    def consent_id(self) -> Optional[str]:
        return pulumi.get(self, "consent_id")

    @property
    @pulumi.getter
    def control(self) -> str:
        return pulumi.get(self, "control")

    @property
    @pulumi.getter
    def data(self) -> Optional[Mapping[str, Any]]:
        """
        An object that can hold any information about the Form Field that should be persisted.
        - description
        - key
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="formFieldId")
    def form_field_id(self) -> str:
        return pulumi.get(self, "form_field_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The unique name of the Form Field.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> Optional[Sequence[str]]:
        """
        A list of options that are applied to checkbox, radio, or select controls.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def required(self) -> Optional[bool]:
        """
        Determines if a value is required to complete the form.
        """
        return pulumi.get(self, "required")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The form field type. The possible values are:
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def validator(self) -> 'outputs.GetFormFieldValidatorResult':
        return pulumi.get(self, "validator")


class AwaitableGetFormFieldResult(GetFormFieldResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFormFieldResult(
            confirm=self.confirm,
            consent_id=self.consent_id,
            control=self.control,
            data=self.data,
            description=self.description,
            form_field_id=self.form_field_id,
            id=self.id,
            key=self.key,
            name=self.name,
            options=self.options,
            required=self.required,
            type=self.type,
            validator=self.validator)


def get_form_field(confirm: Optional[bool] = None,
                   consent_id: Optional[str] = None,
                   control: Optional[str] = None,
                   data: Optional[Mapping[str, Any]] = None,
                   description: Optional[str] = None,
                   form_field_id: Optional[str] = None,
                   key: Optional[str] = None,
                   name: Optional[str] = None,
                   options: Optional[Sequence[str]] = None,
                   required: Optional[bool] = None,
                   type: Optional[str] = None,
                   validator: Optional[pulumi.InputType['GetFormFieldValidatorArgs']] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFormFieldResult:
    """
    ## # Form Field Resource

    A FusionAuth Form Field is an object that can be customized to receive input within a FusionAuth [Form](https://fusionauth.io/docs/v1/tech/apis/forms).

    [Form Field API](https://fusionauth.io/docs/v1/tech/apis/form-fields)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default = fusionauth.get_form_field(name="Email")
    ```


    :param bool confirm: Determines if the user input should be confirmed by requiring the value to be entered twice.
           - consent_id
           - control
    :param Mapping[str, Any] data: An object that can hold any information about the Form Field that should be persisted.
           - description
           - key
    :param str form_field_id: The unique id of the Form Field. Either `form_field_id` or `name` must be specified.
    :param str name: The name of the Form field. Either `form_field_id` or `name` must be specified.
    :param Sequence[str] options: A list of options that are applied to checkbox, radio, or select controls.
    :param bool required: Determines if a value is required to complete the form.
    :param str type: The form field type. The possible values are:
    """
    __args__ = dict()
    __args__['confirm'] = confirm
    __args__['consentId'] = consent_id
    __args__['control'] = control
    __args__['data'] = data
    __args__['description'] = description
    __args__['formFieldId'] = form_field_id
    __args__['key'] = key
    __args__['name'] = name
    __args__['options'] = options
    __args__['required'] = required
    __args__['type'] = type
    __args__['validator'] = validator
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getFormField:getFormField', __args__, opts=opts, typ=GetFormFieldResult).value

    return AwaitableGetFormFieldResult(
        confirm=__ret__.confirm,
        consent_id=__ret__.consent_id,
        control=__ret__.control,
        data=__ret__.data,
        description=__ret__.description,
        form_field_id=__ret__.form_field_id,
        id=__ret__.id,
        key=__ret__.key,
        name=__ret__.name,
        options=__ret__.options,
        required=__ret__.required,
        type=__ret__.type,
        validator=__ret__.validator)


@_utilities.lift_output_func(get_form_field)
def get_form_field_output(confirm: Optional[pulumi.Input[Optional[bool]]] = None,
                          consent_id: Optional[pulumi.Input[Optional[str]]] = None,
                          control: Optional[pulumi.Input[Optional[str]]] = None,
                          data: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                          description: Optional[pulumi.Input[Optional[str]]] = None,
                          form_field_id: Optional[pulumi.Input[Optional[str]]] = None,
                          key: Optional[pulumi.Input[Optional[str]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          options: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          required: Optional[pulumi.Input[Optional[bool]]] = None,
                          type: Optional[pulumi.Input[Optional[str]]] = None,
                          validator: Optional[pulumi.Input[Optional[pulumi.InputType['GetFormFieldValidatorArgs']]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFormFieldResult]:
    """
    ## # Form Field Resource

    A FusionAuth Form Field is an object that can be customized to receive input within a FusionAuth [Form](https://fusionauth.io/docs/v1/tech/apis/forms).

    [Form Field API](https://fusionauth.io/docs/v1/tech/apis/form-fields)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default = fusionauth.get_form_field(name="Email")
    ```


    :param bool confirm: Determines if the user input should be confirmed by requiring the value to be entered twice.
           - consent_id
           - control
    :param Mapping[str, Any] data: An object that can hold any information about the Form Field that should be persisted.
           - description
           - key
    :param str form_field_id: The unique id of the Form Field. Either `form_field_id` or `name` must be specified.
    :param str name: The name of the Form field. Either `form_field_id` or `name` must be specified.
    :param Sequence[str] options: A list of options that are applied to checkbox, radio, or select controls.
    :param bool required: Determines if a value is required to complete the form.
    :param str type: The form field type. The possible values are:
    """
    ...
