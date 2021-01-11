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
    EntPrefixEdges,
    EntPrefixEdgesFromJSON,
    EntPrefixEdgesFromJSONTyped,
    EntPrefixEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPrefix
 */
export interface EntPrefix {
    /**
     * 
     * @type {EntPrefixEdges}
     * @memberof EntPrefix
     */
    edges?: EntPrefixEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPrefix
     */
    id?: number;
    /**
     * Prefix holds the value of the "prefix" field.
     * @type {string}
     * @memberof EntPrefix
     */
    prefix?: string;
}

export function EntPrefixFromJSON(json: any): EntPrefix {
    return EntPrefixFromJSONTyped(json, false);
}

export function EntPrefixFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPrefix {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntPrefixEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'prefix': !exists(json, 'prefix') ? undefined : json['prefix'],
    };
}

export function EntPrefixToJSON(value?: EntPrefix | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntPrefixEdgesToJSON(value.edges),
        'id': value.id,
        'prefix': value.prefix,
    };
}


