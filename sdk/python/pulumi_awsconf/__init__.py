# coding=utf-8
# *** WARNING: this file was generated by pulumi-gen-awsconf. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .configurer import *
from .provider import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "awsconf",
  "mod": "index",
  "fqn": "pulumi_awsconf",
  "classes": {
   "awsconf:index:Configurer": "Configurer"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "awsconf",
  "token": "pulumi:providers:awsconf",
  "fqn": "pulumi_awsconf",
  "class": "Provider"
 }
]
"""
)
