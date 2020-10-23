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
    EntUserEdges,
    EntUserEdgesFromJSON,
    EntUserEdgesFromJSONTyped,
    EntUserEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUser
 */
export interface EntUser {
    /**
     * 
     * @type {EntUserEdges}
     * @memberof EntUser
     */
    edges?: EntUserEdges;
    /**
     * Email holds the value of the "email" field.
     * @type {string}
     * @memberof EntUser
     */
    email?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntUser
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntUser
     */
    name?: string;
    /**
     * Password holds the value of the "password" field.
     * @type {string}
     * @memberof EntUser
     */
    password?: string;
}

export function EntUserFromJSON(json: any): EntUser {
    return EntUserFromJSONTyped(json, false);
}

export function EntUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntUserEdgesFromJSON(json['edges']),
        'email': !exists(json, 'email') ? undefined : json['email'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'password': !exists(json, 'password') ? undefined : json['password'],
    };
}

export function EntUserToJSON(value?: EntUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntUserEdgesToJSON(value.edges),
        'email': value.email,
        'id': value.id,
        'name': value.name,
        'password': value.password,
    };
}


