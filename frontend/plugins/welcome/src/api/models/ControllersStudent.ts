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
 * @interface ControllersStudent
 */
export interface ControllersStudent {
    /**
     * 
     * @type {string}
     * @memberof ControllersStudent
     */
    addr?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersStudent
     */
    degree?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersStudent
     */
    email?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersStudent
     */
    fname?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersStudent
     */
    lname?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersStudent
     */
    province?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersStudent
     */
    school?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersStudent
     */
    sex?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersStudent
     */
    tel?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersStudent
     */
    title?: number;
}

export function ControllersStudentFromJSON(json: any): ControllersStudent {
    return ControllersStudentFromJSONTyped(json, false);
}

export function ControllersStudentFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersStudent {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addr': !exists(json, 'addr') ? undefined : json['addr'],
        'degree': !exists(json, 'degree') ? undefined : json['degree'],
        'email': !exists(json, 'email') ? undefined : json['email'],
        'fname': !exists(json, 'fname') ? undefined : json['fname'],
        'lname': !exists(json, 'lname') ? undefined : json['lname'],
        'province': !exists(json, 'province') ? undefined : json['province'],
        'school': !exists(json, 'school') ? undefined : json['school'],
        'sex': !exists(json, 'sex') ? undefined : json['sex'],
        'tel': !exists(json, 'tel') ? undefined : json['tel'],
        'title': !exists(json, 'title') ? undefined : json['title'],
    };
}

export function ControllersStudentToJSON(value?: ControllersStudent | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addr': value.addr,
        'degree': value.degree,
        'email': value.email,
        'fname': value.fname,
        'lname': value.lname,
        'province': value.province,
        'school': value.school,
        'sex': value.sex,
        'tel': value.tel,
        'title': value.title,
    };
}

