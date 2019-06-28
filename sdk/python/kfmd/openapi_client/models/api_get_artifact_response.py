# coding: utf-8

"""
    api/service.proto

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: version not set
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six


class ApiGetArtifactResponse(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'artifact': 'MlMetadataArtifact'
    }

    attribute_map = {
        'artifact': 'artifact'
    }

    def __init__(self, artifact=None):  # noqa: E501
        """ApiGetArtifactResponse - a model defined in OpenAPI"""  # noqa: E501

        self._artifact = None
        self.discriminator = None

        if artifact is not None:
            self.artifact = artifact

    @property
    def artifact(self):
        """Gets the artifact of this ApiGetArtifactResponse.  # noqa: E501


        :return: The artifact of this ApiGetArtifactResponse.  # noqa: E501
        :rtype: MlMetadataArtifact
        """
        return self._artifact

    @artifact.setter
    def artifact(self, artifact):
        """Sets the artifact of this ApiGetArtifactResponse.


        :param artifact: The artifact of this ApiGetArtifactResponse.  # noqa: E501
        :type: MlMetadataArtifact
        """

        self._artifact = artifact

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ApiGetArtifactResponse):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
