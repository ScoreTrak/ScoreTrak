/*
 * Ent Schema API
 * This is an auto generated API description made out of an Ent schema definition
 *
 * OpenAPI spec version: 0.1.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 3.0.33
 *
 * Do not edit the class manually.
 *
 */
import {ApiClient} from '../ApiClient';

/**
 * The ServiceChecksList model module.
 * @module model/ServiceChecksList
 * @version 0.1.0
 */
export class ServiceChecksList {
  /**
   * Constructs a new <code>ServiceChecksList</code>.
   * @alias module:model/ServiceChecksList
   * @class
   * @param id {Number} 
   * @param createTime {Date} 
   * @param updateTime {Date} 
   * @param pause {Boolean} 
   * @param hidden {Boolean} 
   * @param competitionId {Number} 
   * @param log {String} 
   * @param error {String} 
   * @param passed {Boolean} 
   */
  constructor(id, createTime, updateTime, pause, hidden, competitionId, log, error, passed) {
    this.id = id;
    this.createTime = createTime;
    this.updateTime = updateTime;
    this.pause = pause;
    this.hidden = hidden;
    this.competitionId = competitionId;
    this.log = log;
    this.error = error;
    this.passed = passed;
  }

  /**
   * Constructs a <code>ServiceChecksList</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ServiceChecksList} obj Optional instance to populate.
   * @return {module:model/ServiceChecksList} The populated <code>ServiceChecksList</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new ServiceChecksList();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'Number');
      if (data.hasOwnProperty('create_time'))
        obj.createTime = ApiClient.convertToType(data['create_time'], 'Date');
      if (data.hasOwnProperty('update_time'))
        obj.updateTime = ApiClient.convertToType(data['update_time'], 'Date');
      if (data.hasOwnProperty('pause'))
        obj.pause = ApiClient.convertToType(data['pause'], 'Boolean');
      if (data.hasOwnProperty('hidden'))
        obj.hidden = ApiClient.convertToType(data['hidden'], 'Boolean');
      if (data.hasOwnProperty('competition_id'))
        obj.competitionId = ApiClient.convertToType(data['competition_id'], 'Number');
      if (data.hasOwnProperty('log'))
        obj.log = ApiClient.convertToType(data['log'], 'String');
      if (data.hasOwnProperty('error'))
        obj.error = ApiClient.convertToType(data['error'], 'String');
      if (data.hasOwnProperty('passed'))
        obj.passed = ApiClient.convertToType(data['passed'], 'Boolean');
    }
    return obj;
  }
}

/**
 * @member {Number} id
 */
ServiceChecksList.prototype.id = undefined;

/**
 * @member {Date} createTime
 */
ServiceChecksList.prototype.createTime = undefined;

/**
 * @member {Date} updateTime
 */
ServiceChecksList.prototype.updateTime = undefined;

/**
 * @member {Boolean} pause
 */
ServiceChecksList.prototype.pause = undefined;

/**
 * @member {Boolean} hidden
 */
ServiceChecksList.prototype.hidden = undefined;

/**
 * @member {Number} competitionId
 */
ServiceChecksList.prototype.competitionId = undefined;

/**
 * @member {String} log
 */
ServiceChecksList.prototype.log = undefined;

/**
 * @member {String} error
 */
ServiceChecksList.prototype.error = undefined;

/**
 * @member {Boolean} passed
 */
ServiceChecksList.prototype.passed = undefined;

