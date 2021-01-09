/* tslint:disable */
/* eslint-disable */
/**
 * SUT SE Example API
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
    EntActivity,
    EntActivityFromJSON,
    EntActivityFromJSONTyped,
    EntActivityToJSON,
    EntDegree,
    EntDegreeFromJSON,
    EntDegreeFromJSONTyped,
    EntDegreeToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntPrefix,
    EntPrefixFromJSON,
    EntPrefixFromJSONTyped,
    EntPrefixToJSON,
    EntProvince,
    EntProvinceFromJSON,
    EntProvinceFromJSONTyped,
    EntProvinceToJSON,
    EntResults,
    EntResultsFromJSON,
    EntResultsFromJSONTyped,
    EntResultsToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStudentEdges
 */
export interface EntStudentEdges {
    /**
     * StudActi holds the value of the stud_acti edge.
     * @type {Array<EntActivity>}
     * @memberof EntStudentEdges
     */
    studActi?: Array<EntActivity>;
    /**
     * 
     * @type {EntDegree}
     * @memberof EntStudentEdges
     */
    studDegr?: EntDegree;
    /**
     * 
     * @type {EntGender}
     * @memberof EntStudentEdges
     */
    studGend?: EntGender;
    /**
     * 
     * @type {EntPrefix}
     * @memberof EntStudentEdges
     */
    studPref?: EntPrefix;
    /**
     * 
     * @type {EntProvince}
     * @memberof EntStudentEdges
     */
    studProv?: EntProvince;
    /**
     * StudResu holds the value of the stud_resu edge.
     * @type {Array<EntResults>}
     * @memberof EntStudentEdges
     */
    studResu?: Array<EntResults>;
}

export function EntStudentEdgesFromJSON(json: any): EntStudentEdges {
    return EntStudentEdgesFromJSONTyped(json, false);
}

export function EntStudentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStudentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'studActi': !exists(json, 'studActi') ? undefined : ((json['studActi'] as Array<any>).map(EntActivityFromJSON)),
        'studDegr': !exists(json, 'studDegr') ? undefined : EntDegreeFromJSON(json['studDegr']),
        'studGend': !exists(json, 'studGend') ? undefined : EntGenderFromJSON(json['studGend']),
        'studPref': !exists(json, 'studPref') ? undefined : EntPrefixFromJSON(json['studPref']),
        'studProv': !exists(json, 'studProv') ? undefined : EntProvinceFromJSON(json['studProv']),
        'studResu': !exists(json, 'studResu') ? undefined : ((json['studResu'] as Array<any>).map(EntResultsFromJSON)),
    };
}

export function EntStudentEdgesToJSON(value?: EntStudentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'studActi': value.studActi === undefined ? undefined : ((value.studActi as Array<any>).map(EntActivityToJSON)),
        'studDegr': EntDegreeToJSON(value.studDegr),
        'studGend': EntGenderToJSON(value.studGend),
        'studPref': EntPrefixToJSON(value.studPref),
        'studProv': EntProvinceToJSON(value.studProv),
        'studResu': value.studResu === undefined ? undefined : ((value.studResu as Array<any>).map(EntResultsToJSON)),
    };
}


