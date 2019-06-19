/**
 * api/service.proto
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: version not set
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { MlMetadataExecutionState } from './mlMetadataExecutionState';
import { MlMetadataValue } from './mlMetadataValue';

export class MlMetadataExecution {
    /**
    * The id of the execution.
    */
    'id'?: string;
    /**
    * The id of an ExecutionType. The ExecutionType must be specified and cannot be changed.
    */
    'typeId'?: string;
    'lastKnownState'?: MlMetadataExecutionState;
    /**
    * Properties of the Execution. Properties must be specified in the ExecutionType.
    */
    'properties'?: { [key: string]: MlMetadataValue; };
    /**
    * User provided custom properties which are not defined by its type.
    */
    'customProperties'?: { [key: string]: MlMetadataValue; };

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "typeId",
            "baseName": "type_id",
            "type": "string"
        },
        {
            "name": "lastKnownState",
            "baseName": "last_known_state",
            "type": "MlMetadataExecutionState"
        },
        {
            "name": "properties",
            "baseName": "properties",
            "type": "{ [key: string]: MlMetadataValue; }"
        },
        {
            "name": "customProperties",
            "baseName": "custom_properties",
            "type": "{ [key: string]: MlMetadataValue; }"
        }    ];

    static getAttributeTypeMap() {
        return MlMetadataExecution.attributeTypeMap;
    }
}

