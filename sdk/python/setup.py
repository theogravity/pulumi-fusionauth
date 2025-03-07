# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "0.0.0"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "fusionauth Pulumi Package - Development Version"


setup(name='theogravity_pulumi_fusionauth',
      python_requires='>=3.9',
      version=VERSION,
      description="A Pulumi package for managing FusionAuth instances.",
      long_description=readme(),
      long_description_content_type='text/markdown',
      keywords='pulumi fusionauth category/infrastructure',
      url='https://github.com/theogravity/pulumi-fusionauth',
      project_urls={
          'Repository': 'https://github.com/theogravity/pulumi-fusionauth'
      },
      license='MIT',
      packages=find_packages(),
      package_data={
          'theogravity_pulumi_fusionauth': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'semver>=2.8.1',
          'typing-extensions>=4.11,<5; python_version < "3.11"'
      ],
      zip_safe=False)
