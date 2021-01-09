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
    EntAgency,
    EntAgencyFromJSON,
    EntAgencyFromJSONTyped,
    EntAgencyToJSON,
    EntPlace,
    EntPlaceFromJSON,
    EntPlaceFromJSONTyped,
    EntPlaceToJSON,
    EntStudent,
    EntStudentFromJSON,
    EntStudentFromJSONTyped,
    EntStudentToJSON,
    EntYear,
    EntYearFromJSON,
    EntYearFromJSONTyped,
    EntYearToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntActivityEdges
 */
export interface EntActivityEdges {
    /**
     * 
     * @type {EntAgency}
     * @memberof EntActivityEdges
     */
    actiAgen?: EntAgency;
    /**
     * 
     * @type {EntPlace}
     * @memberof EntActivityEdges
     */
    actiPlace?: EntPlace;
    /**
     * 
     * @type {EntStudent}
     * @memberof EntActivityEdges
     */
    actiStud?: EntStudent;
    /**
     * 
     * @type {EntYear}
     * @memberof EntActivityEdges
     */
    actiYear?: EntYear;
}

export function EntActivityEdgesFromJSON(json: any): EntActivityEdges {
    return EntActivityEdgesFromJSONTyped(json, false);
}

export function EntActivityEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntActivityEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'actiAgen': !exists(json, 'actiAgen') ? undefined : EntAgencyFromJSON(json['actiAgen']),
        'actiPlace': !exists(json, 'actiPlace') ? undefined : EntPlaceFromJSON(json['actiPlace']),
        'actiStud': !exists(json, 'actiStud') ? undefined : EntStudentFromJSON(json['actiStud']),
        'actiYear': !exists(json, 'actiYear') ? undefined : EntYearFromJSON(json['actiYear']),
    };
}

export function EntActivityEdgesToJSON(value?: EntActivityEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'actiAgen': EntAgencyToJSON(value.actiAgen),
        'actiPlace': EntPlaceToJSON(value.actiPlace),
        'actiStud': EntStudentToJSON(value.actiStud),
        'actiYear': EntYearToJSON(value.actiYear),
    };
}


