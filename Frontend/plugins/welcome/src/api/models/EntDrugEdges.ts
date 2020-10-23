/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRegisterstore,
    EntRegisterstoreFromJSON,
    EntRegisterstoreFromJSONTyped,
    EntRegisterstoreToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDrugEdges
 */
export interface EntDrugEdges {
    /**
     * Drugs holds the value of the drugs edge.
     * @type {Array<EntRegisterstore>}
     * @memberof EntDrugEdges
     */
    drugs?: Array<EntRegisterstore>;
}

export function EntDrugEdgesFromJSON(json: any): EntDrugEdges {
    return EntDrugEdgesFromJSONTyped(json, false);
}

export function EntDrugEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDrugEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'drugs': !exists(json, 'drugs') ? undefined : ((json['drugs'] as Array<any>).map(EntRegisterstoreFromJSON)),
    };
}

export function EntDrugEdgesToJSON(value?: EntDrugEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'drugs': value.drugs === undefined ? undefined : ((value.drugs as Array<any>).map(EntRegisterstoreToJSON)),
    };
}


