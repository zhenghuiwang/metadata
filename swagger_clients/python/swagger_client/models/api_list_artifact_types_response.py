# coding: utf-8

"""
    api/service.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from swagger_client.models.api_artifact_type import ApiArtifactType  # noqa: F401,E501


class ApiListArtifactTypesResponse(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'artifact_types': 'list[ApiArtifactType]'
    }

    attribute_map = {
        'artifact_types': 'artifact_types'
    }

    def __init__(self, artifact_types=None):  # noqa: E501
        """ApiListArtifactTypesResponse - a model defined in Swagger"""  # noqa: E501

        self._artifact_types = None
        self.discriminator = None

        if artifact_types is not None:
            self.artifact_types = artifact_types

    @property
    def artifact_types(self):
        """Gets the artifact_types of this ApiListArtifactTypesResponse.  # noqa: E501


        :return: The artifact_types of this ApiListArtifactTypesResponse.  # noqa: E501
        :rtype: list[ApiArtifactType]
        """
        return self._artifact_types

    @artifact_types.setter
    def artifact_types(self, artifact_types):
        """Sets the artifact_types of this ApiListArtifactTypesResponse.


        :param artifact_types: The artifact_types of this ApiListArtifactTypesResponse.  # noqa: E501
        :type: list[ApiArtifactType]
        """

        self._artifact_types = artifact_types

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
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
        if issubclass(ApiListArtifactTypesResponse, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, ApiListArtifactTypesResponse):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
