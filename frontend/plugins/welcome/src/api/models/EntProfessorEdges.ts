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
    EntFaculty,
    EntFacultyFromJSON,
    EntFacultyFromJSONTyped,
    EntFacultyToJSON,
    EntPrefix,
    EntPrefixFromJSON,
    EntPrefixFromJSONTyped,
    EntPrefixToJSON,
    EntProfessorship,
    EntProfessorshipFromJSON,
    EntProfessorshipFromJSONTyped,
    EntProfessorshipToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProfessorEdges
 */
export interface EntProfessorEdges {
    /**
     * 
     * @type {EntFaculty}
     * @memberof EntProfessorEdges
     */
    profFacu?: EntFaculty;
    /**
     * 
     * @type {EntPrefix}
     * @memberof EntProfessorEdges
     */
    profPre?: EntPrefix;
    /**
     * 
     * @type {EntProfessorship}
     * @memberof EntProfessorEdges
     */
    profPros?: EntProfessorship;
}

export function EntProfessorEdgesFromJSON(json: any): EntProfessorEdges {
    return EntProfessorEdgesFromJSONTyped(json, false);
}

export function EntProfessorEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProfessorEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'profFacu': !exists(json, 'ProfFacu') ? undefined : EntFacultyFromJSON(json['ProfFacu']),
        'profPre': !exists(json, 'ProfPre') ? undefined : EntPrefixFromJSON(json['ProfPre']),
        'profPros': !exists(json, 'ProfPros') ? undefined : EntProfessorshipFromJSON(json['ProfPros']),
    };
}

export function EntProfessorEdgesToJSON(value?: EntProfessorEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'profFacu': EntFacultyToJSON(value.profFacu),
        'profPre': EntPrefixToJSON(value.profPre),
        'profPros': EntProfessorshipToJSON(value.profPros),
    };
}


