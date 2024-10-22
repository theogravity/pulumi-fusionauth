# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['FusionAuthEMailArgs', 'FusionAuthEMail']

@pulumi.input_type
class FusionAuthEMailArgs:
    def __init__(__self__, *,
                 default_html_template: pulumi.Input[str],
                 default_subject: pulumi.Input[str],
                 default_text_template: pulumi.Input[str],
                 default_from_name: Optional[pulumi.Input[str]] = None,
                 email_id: Optional[pulumi.Input[str]] = None,
                 from_email: Optional[pulumi.Input[str]] = None,
                 localized_from_names: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_html_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_subjects: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_text_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthEMail resource.
        :param pulumi.Input[str] default_html_template: The default HTML Email Template.
        :param pulumi.Input[str] default_subject: The default Subject used when sending emails.
        :param pulumi.Input[str] default_text_template: The default Text Email Template.
        :param pulumi.Input[str] default_from_name: The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[str] email_id: The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] from_email: The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[Mapping[str, Any]] localized_from_names: The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_html_templates: The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_subjects: The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_text_templates: The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[str] name: A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        pulumi.set(__self__, "default_html_template", default_html_template)
        pulumi.set(__self__, "default_subject", default_subject)
        pulumi.set(__self__, "default_text_template", default_text_template)
        if default_from_name is not None:
            pulumi.set(__self__, "default_from_name", default_from_name)
        if email_id is not None:
            pulumi.set(__self__, "email_id", email_id)
        if from_email is not None:
            pulumi.set(__self__, "from_email", from_email)
        if localized_from_names is not None:
            pulumi.set(__self__, "localized_from_names", localized_from_names)
        if localized_html_templates is not None:
            pulumi.set(__self__, "localized_html_templates", localized_html_templates)
        if localized_subjects is not None:
            pulumi.set(__self__, "localized_subjects", localized_subjects)
        if localized_text_templates is not None:
            pulumi.set(__self__, "localized_text_templates", localized_text_templates)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="defaultHtmlTemplate")
    def default_html_template(self) -> pulumi.Input[str]:
        """
        The default HTML Email Template.
        """
        return pulumi.get(self, "default_html_template")

    @default_html_template.setter
    def default_html_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_html_template", value)

    @property
    @pulumi.getter(name="defaultSubject")
    def default_subject(self) -> pulumi.Input[str]:
        """
        The default Subject used when sending emails.
        """
        return pulumi.get(self, "default_subject")

    @default_subject.setter
    def default_subject(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_subject", value)

    @property
    @pulumi.getter(name="defaultTextTemplate")
    def default_text_template(self) -> pulumi.Input[str]:
        """
        The default Text Email Template.
        """
        return pulumi.get(self, "default_text_template")

    @default_text_template.setter
    def default_text_template(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_text_template", value)

    @property
    @pulumi.getter(name="defaultFromName")
    def default_from_name(self) -> Optional[pulumi.Input[str]]:
        """
        The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        """
        return pulumi.get(self, "default_from_name")

    @default_from_name.setter
    def default_from_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_from_name", value)

    @property
    @pulumi.getter(name="emailId")
    def email_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "email_id")

    @email_id.setter
    def email_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_id", value)

    @property
    @pulumi.getter(name="fromEmail")
    def from_email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        """
        return pulumi.get(self, "from_email")

    @from_email.setter
    def from_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_email", value)

    @property
    @pulumi.getter(name="localizedFromNames")
    def localized_from_names(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_from_names")

    @localized_from_names.setter
    def localized_from_names(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_from_names", value)

    @property
    @pulumi.getter(name="localizedHtmlTemplates")
    def localized_html_templates(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_html_templates")

    @localized_html_templates.setter
    def localized_html_templates(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_html_templates", value)

    @property
    @pulumi.getter(name="localizedSubjects")
    def localized_subjects(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_subjects")

    @localized_subjects.setter
    def localized_subjects(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_subjects", value)

    @property
    @pulumi.getter(name="localizedTextTemplates")
    def localized_text_templates(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_text_templates")

    @localized_text_templates.setter
    def localized_text_templates(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_text_templates", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FusionAuthEMailState:
    def __init__(__self__, *,
                 default_from_name: Optional[pulumi.Input[str]] = None,
                 default_html_template: Optional[pulumi.Input[str]] = None,
                 default_subject: Optional[pulumi.Input[str]] = None,
                 default_text_template: Optional[pulumi.Input[str]] = None,
                 email_id: Optional[pulumi.Input[str]] = None,
                 from_email: Optional[pulumi.Input[str]] = None,
                 localized_from_names: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_html_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_subjects: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_text_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthEMail resources.
        :param pulumi.Input[str] default_from_name: The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[str] default_html_template: The default HTML Email Template.
        :param pulumi.Input[str] default_subject: The default Subject used when sending emails.
        :param pulumi.Input[str] default_text_template: The default Text Email Template.
        :param pulumi.Input[str] email_id: The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] from_email: The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[Mapping[str, Any]] localized_from_names: The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_html_templates: The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_subjects: The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_text_templates: The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[str] name: A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        if default_from_name is not None:
            pulumi.set(__self__, "default_from_name", default_from_name)
        if default_html_template is not None:
            pulumi.set(__self__, "default_html_template", default_html_template)
        if default_subject is not None:
            pulumi.set(__self__, "default_subject", default_subject)
        if default_text_template is not None:
            pulumi.set(__self__, "default_text_template", default_text_template)
        if email_id is not None:
            pulumi.set(__self__, "email_id", email_id)
        if from_email is not None:
            pulumi.set(__self__, "from_email", from_email)
        if localized_from_names is not None:
            pulumi.set(__self__, "localized_from_names", localized_from_names)
        if localized_html_templates is not None:
            pulumi.set(__self__, "localized_html_templates", localized_html_templates)
        if localized_subjects is not None:
            pulumi.set(__self__, "localized_subjects", localized_subjects)
        if localized_text_templates is not None:
            pulumi.set(__self__, "localized_text_templates", localized_text_templates)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="defaultFromName")
    def default_from_name(self) -> Optional[pulumi.Input[str]]:
        """
        The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        """
        return pulumi.get(self, "default_from_name")

    @default_from_name.setter
    def default_from_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_from_name", value)

    @property
    @pulumi.getter(name="defaultHtmlTemplate")
    def default_html_template(self) -> Optional[pulumi.Input[str]]:
        """
        The default HTML Email Template.
        """
        return pulumi.get(self, "default_html_template")

    @default_html_template.setter
    def default_html_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_html_template", value)

    @property
    @pulumi.getter(name="defaultSubject")
    def default_subject(self) -> Optional[pulumi.Input[str]]:
        """
        The default Subject used when sending emails.
        """
        return pulumi.get(self, "default_subject")

    @default_subject.setter
    def default_subject(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_subject", value)

    @property
    @pulumi.getter(name="defaultTextTemplate")
    def default_text_template(self) -> Optional[pulumi.Input[str]]:
        """
        The default Text Email Template.
        """
        return pulumi.get(self, "default_text_template")

    @default_text_template.setter
    def default_text_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_text_template", value)

    @property
    @pulumi.getter(name="emailId")
    def email_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "email_id")

    @email_id.setter
    def email_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_id", value)

    @property
    @pulumi.getter(name="fromEmail")
    def from_email(self) -> Optional[pulumi.Input[str]]:
        """
        The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        """
        return pulumi.get(self, "from_email")

    @from_email.setter
    def from_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "from_email", value)

    @property
    @pulumi.getter(name="localizedFromNames")
    def localized_from_names(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_from_names")

    @localized_from_names.setter
    def localized_from_names(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_from_names", value)

    @property
    @pulumi.getter(name="localizedHtmlTemplates")
    def localized_html_templates(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_html_templates")

    @localized_html_templates.setter
    def localized_html_templates(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_html_templates", value)

    @property
    @pulumi.getter(name="localizedSubjects")
    def localized_subjects(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_subjects")

    @localized_subjects.setter
    def localized_subjects(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_subjects", value)

    @property
    @pulumi.getter(name="localizedTextTemplates")
    def localized_text_templates(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_text_templates")

    @localized_text_templates.setter
    def localized_text_templates(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "localized_text_templates", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class FusionAuthEMail(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_from_name: Optional[pulumi.Input[str]] = None,
                 default_html_template: Optional[pulumi.Input[str]] = None,
                 default_subject: Optional[pulumi.Input[str]] = None,
                 default_text_template: Optional[pulumi.Input[str]] = None,
                 email_id: Optional[pulumi.Input[str]] = None,
                 from_email: Optional[pulumi.Input[str]] = None,
                 localized_from_names: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_html_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_subjects: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_text_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Email Resource

        This resource contains the APIs for managing Email Templates.

        [Emails API](https://fusionauth.io/docs/v1/tech/apis/emails)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        hello_world = fusionauth.FusionAuthEMail("helloWorld",
            default_from_name="Welcome Team",
            default_html_template=(lambda path: open(path).read())(f"{path['module']}/email_templates/HelloWorld.html.ftl"),
            default_subject="Hello",
            default_text_template=(lambda path: open(path).read())(f"{path['module']}/email_templates/HelloWorld.txt.ftl"),
            from_email="welcome@example.com.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_from_name: The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[str] default_html_template: The default HTML Email Template.
        :param pulumi.Input[str] default_subject: The default Subject used when sending emails.
        :param pulumi.Input[str] default_text_template: The default Text Email Template.
        :param pulumi.Input[str] email_id: The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] from_email: The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[Mapping[str, Any]] localized_from_names: The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_html_templates: The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_subjects: The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_text_templates: The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[str] name: A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FusionAuthEMailArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Email Resource

        This resource contains the APIs for managing Email Templates.

        [Emails API](https://fusionauth.io/docs/v1/tech/apis/emails)

        ## Example Usage

        ```python
        import pulumi
        import theogravity_pulumi_fusionauth as fusionauth

        hello_world = fusionauth.FusionAuthEMail("helloWorld",
            default_from_name="Welcome Team",
            default_html_template=(lambda path: open(path).read())(f"{path['module']}/email_templates/HelloWorld.html.ftl"),
            default_subject="Hello",
            default_text_template=(lambda path: open(path).read())(f"{path['module']}/email_templates/HelloWorld.txt.ftl"),
            from_email="welcome@example.com.com")
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthEMailArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthEMailArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_from_name: Optional[pulumi.Input[str]] = None,
                 default_html_template: Optional[pulumi.Input[str]] = None,
                 default_subject: Optional[pulumi.Input[str]] = None,
                 default_text_template: Optional[pulumi.Input[str]] = None,
                 email_id: Optional[pulumi.Input[str]] = None,
                 from_email: Optional[pulumi.Input[str]] = None,
                 localized_from_names: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_html_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_subjects: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 localized_text_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthEMailArgs.__new__(FusionAuthEMailArgs)

            __props__.__dict__["default_from_name"] = default_from_name
            if default_html_template is None and not opts.urn:
                raise TypeError("Missing required property 'default_html_template'")
            __props__.__dict__["default_html_template"] = default_html_template
            if default_subject is None and not opts.urn:
                raise TypeError("Missing required property 'default_subject'")
            __props__.__dict__["default_subject"] = default_subject
            if default_text_template is None and not opts.urn:
                raise TypeError("Missing required property 'default_text_template'")
            __props__.__dict__["default_text_template"] = default_text_template
            __props__.__dict__["email_id"] = email_id
            __props__.__dict__["from_email"] = from_email
            __props__.__dict__["localized_from_names"] = localized_from_names
            __props__.__dict__["localized_html_templates"] = localized_html_templates
            __props__.__dict__["localized_subjects"] = localized_subjects
            __props__.__dict__["localized_text_templates"] = localized_text_templates
            __props__.__dict__["name"] = name
        super(FusionAuthEMail, __self__).__init__(
            'fusionauth:index/fusionAuthEMail:FusionAuthEMail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_from_name: Optional[pulumi.Input[str]] = None,
            default_html_template: Optional[pulumi.Input[str]] = None,
            default_subject: Optional[pulumi.Input[str]] = None,
            default_text_template: Optional[pulumi.Input[str]] = None,
            email_id: Optional[pulumi.Input[str]] = None,
            from_email: Optional[pulumi.Input[str]] = None,
            localized_from_names: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            localized_html_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            localized_subjects: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            localized_text_templates: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'FusionAuthEMail':
        """
        Get an existing FusionAuthEMail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_from_name: The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[str] default_html_template: The default HTML Email Template.
        :param pulumi.Input[str] default_subject: The default Subject used when sending emails.
        :param pulumi.Input[str] default_text_template: The default Text Email Template.
        :param pulumi.Input[str] email_id: The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        :param pulumi.Input[str] from_email: The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        :param pulumi.Input[Mapping[str, Any]] localized_from_names: The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_html_templates: The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_subjects: The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        :param pulumi.Input[Mapping[str, Any]] localized_text_templates: The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        :param pulumi.Input[str] name: A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthEMailState.__new__(_FusionAuthEMailState)

        __props__.__dict__["default_from_name"] = default_from_name
        __props__.__dict__["default_html_template"] = default_html_template
        __props__.__dict__["default_subject"] = default_subject
        __props__.__dict__["default_text_template"] = default_text_template
        __props__.__dict__["email_id"] = email_id
        __props__.__dict__["from_email"] = from_email
        __props__.__dict__["localized_from_names"] = localized_from_names
        __props__.__dict__["localized_html_templates"] = localized_html_templates
        __props__.__dict__["localized_subjects"] = localized_subjects
        __props__.__dict__["localized_text_templates"] = localized_text_templates
        __props__.__dict__["name"] = name
        return FusionAuthEMail(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultFromName")
    def default_from_name(self) -> pulumi.Output[Optional[str]]:
        """
        The default From Name used when sending emails. If not provided, and a localized value cannot be determined, the default value for the tenant will be used. This is the display name part of the email address ( i.e. Jared Dunn <jared@piedpiper.com>).
        """
        return pulumi.get(self, "default_from_name")

    @property
    @pulumi.getter(name="defaultHtmlTemplate")
    def default_html_template(self) -> pulumi.Output[str]:
        """
        The default HTML Email Template.
        """
        return pulumi.get(self, "default_html_template")

    @property
    @pulumi.getter(name="defaultSubject")
    def default_subject(self) -> pulumi.Output[str]:
        """
        The default Subject used when sending emails.
        """
        return pulumi.get(self, "default_subject")

    @property
    @pulumi.getter(name="defaultTextTemplate")
    def default_text_template(self) -> pulumi.Output[str]:
        """
        The default Text Email Template.
        """
        return pulumi.get(self, "default_text_template")

    @property
    @pulumi.getter(name="emailId")
    def email_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Id to use for the new Email Template. If not specified a secure random UUID will be generated.
        """
        return pulumi.get(self, "email_id")

    @property
    @pulumi.getter(name="fromEmail")
    def from_email(self) -> pulumi.Output[Optional[str]]:
        """
        The email address that this email will be sent from. If not provided, the default value for the tenant will be used. This is the address part email address (i.e. Jared Dunn <jared@piedpiper.com>).
        """
        return pulumi.get(self, "from_email")

    @property
    @pulumi.getter(name="localizedFromNames")
    def localized_from_names(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The From Name used when sending emails to users who speak other languages. This overrides the default From Name based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_from_names")

    @property
    @pulumi.getter(name="localizedHtmlTemplates")
    def localized_html_templates(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The HTML Email Template used when sending emails to users who speak other languages. This overrides the default HTML Email Template based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_html_templates")

    @property
    @pulumi.getter(name="localizedSubjects")
    def localized_subjects(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The Subject used when sending emails to users who speak other languages. This overrides the default Subject based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_subjects")

    @property
    @pulumi.getter(name="localizedTextTemplates")
    def localized_text_templates(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The Text Email Template used when sending emails to users who speak other languages. This overrides the default Text Email Template based on the user’s list of preferred languages.
        """
        return pulumi.get(self, "localized_text_templates")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A descriptive name for the email template (i.e. "April 2016 Coupon Email")
        """
        return pulumi.get(self, "name")
