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
    EntDegree,
    EntDegreeFromJSON,
    EntDegreeFromJSONTyped,
    EntDegreeToJSON,
    EntFaculty,
    EntFacultyFromJSON,
    EntFacultyFromJSONTyped,
    EntFacultyToJSON,
    EntInstitution,
    EntInstitutionFromJSON,
    EntInstitutionFromJSONTyped,
    EntInstitutionToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCourseEdges
 */
export interface EntCourseEdges {
    /**
     * 
     * @type {EntDegree}
     * @memberof EntCourseEdges
     */
    courDegr?: EntDegree;
    /**
     * 
     * @type {EntFaculty}
     * @memberof EntCourseEdges
     */
    courFacu?: EntFaculty;
    /**
     * 
     * @type {EntInstitution}
     * @memberof EntCourseEdges
     */
    courInst?: EntInstitution;
}

export function EntCourseEdgesFromJSON(json: any): EntCourseEdges {
    return EntCourseEdgesFromJSONTyped(json, false);
}

export function EntCourseEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCourseEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'courDegr': !exists(json, 'courDegr') ? undefined : EntDegreeFromJSON(json['courDegr']),
        'courFacu': !exists(json, 'courFacu') ? undefined : EntFacultyFromJSON(json['courFacu']),
        'courInst': !exists(json, 'courInst') ? undefined : EntInstitutionFromJSON(json['courInst']),
    };
}

export function EntCourseEdgesToJSON(value?: EntCourseEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'courDegr': EntDegreeToJSON(value.courDegr),
        'courFacu': EntFacultyToJSON(value.courFacu),
        'courInst': EntInstitutionToJSON(value.courInst),
    };
}

