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
    EntProfessorEdges,
    EntProfessorEdgesFromJSON,
    EntProfessorEdgesFromJSONTyped,
    EntProfessorEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProfessor
 */
export interface EntProfessor {
    /**
     * 
     * @type {EntProfessorEdges}
     * @memberof EntProfessor
     */
    edges?: EntProfessorEdges;
    /**
     * Email holds the value of the "email" field.
     * @type {string}
     * @memberof EntProfessor
     */
    email?: string;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntProfessor
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntProfessor
     */
    name?: string;
    /**
     * Tel holds the value of the "tel" field.
     * @type {string}
     * @memberof EntProfessor
     */
    tel?: string;
}

export function EntProfessorFromJSON(json: any): EntProfessor {
    return EntProfessorFromJSONTyped(json, false);
}

export function EntProfessorFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProfessor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntProfessorEdgesFromJSON(json['edges']),
        'email': !exists(json, 'email') ? undefined : json['email'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'tel': !exists(json, 'tel') ? undefined : json['tel'],
    };
}

export function EntProfessorToJSON(value?: EntProfessor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntProfessorEdgesToJSON(value.edges),
        'email': value.email,
        'id': value.id,
        'name': value.name,
        'tel': value.tel,
    };
}

