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


/**
* A value in properties.
*/
export class MlMetadataValue {
    'intValue'?: string;
    'doubleValue'?: number;
    'stringValue'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "intValue",
            "baseName": "int_value",
            "type": "string"
        },
        {
            "name": "doubleValue",
            "baseName": "double_value",
            "type": "number"
        },
        {
            "name": "stringValue",
            "baseName": "string_value",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MlMetadataValue.attributeTypeMap;
    }
}

