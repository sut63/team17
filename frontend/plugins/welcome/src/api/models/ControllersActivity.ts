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
/**
 * 
 * @export
 * @interface ControllersActivity
 */
export interface ControllersActivity {
    /**
     * 
     * @type {string}
     * @memberof ControllersActivity
     */
    activityname?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersActivity
     */
    added?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersActivity
     */
    agency?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersActivity
     */
    hours?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersActivity
     */
    place?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersActivity
     */
    student?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersActivity
     */
    term?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersActivity
     */
    year?: number;
}

export function ControllersActivityFromJSON(json: any): ControllersActivity {
    return ControllersActivityFromJSONTyped(json, false);
}

export function ControllersActivityFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersActivity {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activityname': !exists(json, 'activityname') ? undefined : json['activityname'],
        'added': !exists(json, 'added') ? undefined : json['added'],
        'agency': !exists(json, 'agency') ? undefined : json['agency'],
        'hours': !exists(json, 'hours') ? undefined : json['hours'],
        'place': !exists(json, 'place') ? undefined : json['place'],
        'student': !exists(json, 'student') ? undefined : json['student'],
        'term': !exists(json, 'term') ? undefined : json['term'],
        'year': !exists(json, 'year') ? undefined : json['year'],
    };
}

export function ControllersActivityToJSON(value?: ControllersActivity | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activityname': value.activityname,
        'added': value.added,
        'agency': value.agency,
        'hours': value.hours,
        'place': value.place,
        'student': value.student,
        'term': value.term,
        'year': value.year,
    };
}


