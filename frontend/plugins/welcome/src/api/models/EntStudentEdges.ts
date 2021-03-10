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
     * @type {EntProvince}
     * @memberof EntStudentEdges
     */
    studDist?: EntProvince;
    /**
     * 
     * @type {EntGender}
     * @memberof EntStudentEdges
     */
    studGend?: EntGender;
    /**
     * 
     * @type {EntProvince}
     * @memberof EntStudentEdges
     */
    studPost?: EntProvince;
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
    /**
     * 
     * @type {EntProvince}
     * @memberof EntStudentEdges
     */
    studSubd?: EntProvince;
}

export function EntStudentEdgesFromJSON(json: any): EntStudentEdges {
    return EntStudentEdgesFromJSONTyped(json, false);
}

export function EntStudentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStudentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'studActi': !exists(json, 'StudActi') ? undefined : ((json['StudActi'] as Array<any>).map(EntActivityFromJSON)),
        'studDegr': !exists(json, 'StudDegr') ? undefined : EntDegreeFromJSON(json['StudDegr']),
        'studDist': !exists(json, 'StudDist') ? undefined : EntProvinceFromJSON(json['StudDist']),
        'studGend': !exists(json, 'StudGend') ? undefined : EntGenderFromJSON(json['StudGend']),
        'studPost': !exists(json, 'StudPost') ? undefined : EntProvinceFromJSON(json['StudPost']),
        'studPref': !exists(json, 'StudPref') ? undefined : EntPrefixFromJSON(json['StudPref']),
        'studProv': !exists(json, 'StudProv') ? undefined : EntProvinceFromJSON(json['StudProv']),
        'studResu': !exists(json, 'StudResu') ? undefined : ((json['StudResu'] as Array<any>).map(EntResultsFromJSON)),
        'studSubd': !exists(json, 'StudSubd') ? undefined : EntProvinceFromJSON(json['StudSubd']),
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
        'studDist': EntProvinceToJSON(value.studDist),
        'studGend': EntGenderToJSON(value.studGend),
        'studPost': EntProvinceToJSON(value.studPost),
        'studPref': EntPrefixToJSON(value.studPref),
        'studProv': EntProvinceToJSON(value.studProv),
        'studResu': value.studResu === undefined ? undefined : ((value.studResu as Array<any>).map(EntResultsToJSON)),
        'studSubd': EntProvinceToJSON(value.studSubd),
    };
}


