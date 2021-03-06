/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
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
    EntActivityEdges,
    EntActivityEdgesFromJSON,
    EntActivityEdgesFromJSONTyped,
    EntActivityEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntActivity
 */
export interface EntActivity {
    /**
     * ACTIVITYNAME holds the value of the "ACTIVITYNAME" field.
     * @type {string}
     * @memberof EntActivity
     */
    aCTIVITYNAME?: string;
    /**
     * Added holds the value of the "added" field.
     * @type {string}
     * @memberof EntActivity
     */
    added?: string;
    /**
     * 
     * @type {EntActivityEdges}
     * @memberof EntActivity
     */
    edges?: EntActivityEdges;
    /**
     * Hours holds the value of the "hours" field.
     * @type {number}
     * @memberof EntActivity
     */
    hours?: number;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntActivity
     */
    id?: number;
}

export function EntActivityFromJSON(json: any): EntActivity {
    return EntActivityFromJSONTyped(json, false);
}

export function EntActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntActivity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'aCTIVITYNAME': !exists(json, 'ACTIVITYNAME') ? undefined : json['ACTIVITYNAME'],
        'added': !exists(json, 'added') ? undefined : json['added'],
        'edges': !exists(json, 'edges') ? undefined : EntActivityEdgesFromJSON(json['edges']),
        'hours': !exists(json, 'hours') ? undefined : json['hours'],
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntActivityToJSON(value?: EntActivity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ACTIVITYNAME': value.aCTIVITYNAME,
        'added': value.added,
        'edges': EntActivityEdgesToJSON(value.edges),
        'hours': value.hours,
        'id': value.id,
    };
}


