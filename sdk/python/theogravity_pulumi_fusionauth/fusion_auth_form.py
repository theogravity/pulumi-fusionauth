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
from . import outputs
from ._inputs import *

__all__ = ['FusionAuthFormArgs', 'FusionAuthForm']

@pulumi.input_type
class FusionAuthFormArgs:
    def __init__(__self__, *,
                 steps: pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]],
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 form_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthForm resource.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]] steps: An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the Form Field that should be persisted.
        :param pulumi.Input[str] form_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The unique name of the Form Field.
        :param pulumi.Input[str] type: The type of form being created, a form type cannot be changed after the form has been created.
        """
        pulumi.set(__self__, "steps", steps)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if form_id is not None:
            pulumi.set(__self__, "form_id", form_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def steps(self) -> pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]]:
        """
        An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        """
        return pulumi.get(self, "steps")

    @steps.setter
    def steps(self, value: pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]]):
        pulumi.set(self, "steps", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object that can hold any information about the Form Field that should be persisted.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="formId")
    def form_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new Form. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "form_id")

    @form_id.setter
    def form_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "form_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the Form Field.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of form being created, a form type cannot be changed after the form has been created.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _FusionAuthFormState:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 form_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 steps: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthForm resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the Form Field that should be persisted.
        :param pulumi.Input[str] form_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The unique name of the Form Field.
        :param pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]] steps: An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        :param pulumi.Input[str] type: The type of form being created, a form type cannot be changed after the form has been created.
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if form_id is not None:
            pulumi.set(__self__, "form_id", form_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if steps is not None:
            pulumi.set(__self__, "steps", steps)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        An object that can hold any information about the Form Field that should be persisted.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="formId")
    def form_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new Form. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "form_id")

    @form_id.setter
    def form_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "form_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the Form Field.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def steps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]]]:
        """
        An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        """
        return pulumi.get(self, "steps")

    @steps.setter
    def steps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FusionAuthFormStepArgs']]]]):
        pulumi.set(self, "steps", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of form being created, a form type cannot be changed after the form has been created.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class FusionAuthForm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 form_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 steps: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FusionAuthFormStepArgs', 'FusionAuthFormStepArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Form Resource

        A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.

        [Form API](https://fusionauth.io/docs/v1/tech/apis/forms/)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        form = fusionauth.FusionAuthForm("form",
            data={
                "description": "This form customizes the registration experience.",
            },
            steps=[
                {
                    "fields": ["91909721-7d4f-b110-8f21-cfdee2a1edb8"],
                },
                {
                    "fields": [
                        "8ed89a31-c325-3156-72ed-6e89183af917",
                        "a977cfd4-a9ed-c4cf-650f-f4539268ac38",
                    ],
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the Form Field that should be persisted.
        :param pulumi.Input[str] form_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The unique name of the Form Field.
        :param pulumi.Input[Sequence[pulumi.Input[Union['FusionAuthFormStepArgs', 'FusionAuthFormStepArgsDict']]]] steps: An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        :param pulumi.Input[str] type: The type of form being created, a form type cannot be changed after the form has been created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthFormArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Form Resource

        A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.

        [Form API](https://fusionauth.io/docs/v1/tech/apis/forms/)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        form = fusionauth.FusionAuthForm("form",
            data={
                "description": "This form customizes the registration experience.",
            },
            steps=[
                {
                    "fields": ["91909721-7d4f-b110-8f21-cfdee2a1edb8"],
                },
                {
                    "fields": [
                        "8ed89a31-c325-3156-72ed-6e89183af917",
                        "a977cfd4-a9ed-c4cf-650f-f4539268ac38",
                    ],
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthFormArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthFormArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 form_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 steps: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FusionAuthFormStepArgs', 'FusionAuthFormStepArgsDict']]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthFormArgs.__new__(FusionAuthFormArgs)

            __props__.__dict__["data"] = data
            __props__.__dict__["form_id"] = form_id
            __props__.__dict__["name"] = name
            if steps is None and not opts.urn:
                raise TypeError("Missing required property 'steps'")
            __props__.__dict__["steps"] = steps
            __props__.__dict__["type"] = type
        super(FusionAuthForm, __self__).__init__(
            'fusionauth:index/fusionAuthForm:FusionAuthForm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            form_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            steps: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FusionAuthFormStepArgs', 'FusionAuthFormStepArgsDict']]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'FusionAuthForm':
        """
        Get an existing FusionAuthForm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] data: An object that can hold any information about the Form Field that should be persisted.
        :param pulumi.Input[str] form_id: The Id to use for the new Form. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] name: The unique name of the Form Field.
        :param pulumi.Input[Sequence[pulumi.Input[Union['FusionAuthFormStepArgs', 'FusionAuthFormStepArgsDict']]]] steps: An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        :param pulumi.Input[str] type: The type of form being created, a form type cannot be changed after the form has been created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthFormState.__new__(_FusionAuthFormState)

        __props__.__dict__["data"] = data
        __props__.__dict__["form_id"] = form_id
        __props__.__dict__["name"] = name
        __props__.__dict__["steps"] = steps
        __props__.__dict__["type"] = type
        return FusionAuthForm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An object that can hold any information about the Form Field that should be persisted.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="formId")
    def form_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id to use for the new Form. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "form_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the Form Field.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def steps(self) -> pulumi.Output[Sequence['outputs.FusionAuthFormStep']]:
        """
        An ordered list of objects containing one or more Form Fields. A Form must have at least one step defined.
        """
        return pulumi.get(self, "steps")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of form being created, a form type cannot be changed after the form has been created.
        """
        return pulumi.get(self, "type")

