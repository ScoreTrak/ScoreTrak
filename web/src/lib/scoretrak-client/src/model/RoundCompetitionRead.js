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
 * The RoundCompetitionRead model module.
 * @module model/RoundCompetitionRead
 * @version 0.1.0
 */
export class RoundCompetitionRead {
  /**
   * Constructs a new <code>RoundCompetitionRead</code>.
   * @alias module:model/RoundCompetitionRead
   * @class
   * @param id {Number} 
   * @param createTime {Date} 
   * @param updateTime {Date} 
   * @param hidden {Boolean} 
   * @param pause {Boolean} 
   * @param name {String} 
   * @param displayName {String} 
   * @param roundDuration {Number} 
   */
  constructor(id, createTime, updateTime, hidden, pause, name, displayName, roundDuration) {
    this.id = id;
    this.createTime = createTime;
    this.updateTime = updateTime;
    this.hidden = hidden;
    this.pause = pause;
    this.name = name;
    this.displayName = displayName;
    this.roundDuration = roundDuration;
  }

  /**
   * Constructs a <code>RoundCompetitionRead</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/RoundCompetitionRead} obj Optional instance to populate.
   * @return {module:model/RoundCompetitionRead} The populated <code>RoundCompetitionRead</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new RoundCompetitionRead();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'Number');
      if (data.hasOwnProperty('create_time'))
        obj.createTime = ApiClient.convertToType(data['create_time'], 'Date');
      if (data.hasOwnProperty('update_time'))
        obj.updateTime = ApiClient.convertToType(data['update_time'], 'Date');
      if (data.hasOwnProperty('hidden'))
        obj.hidden = ApiClient.convertToType(data['hidden'], 'Boolean');
      if (data.hasOwnProperty('pause'))
        obj.pause = ApiClient.convertToType(data['pause'], 'Boolean');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('display_name'))
        obj.displayName = ApiClient.convertToType(data['display_name'], 'String');
      if (data.hasOwnProperty('round_duration'))
        obj.roundDuration = ApiClient.convertToType(data['round_duration'], 'Number');
      if (data.hasOwnProperty('to_be_started_at'))
        obj.toBeStartedAt = ApiClient.convertToType(data['to_be_started_at'], 'Date');
      if (data.hasOwnProperty('started_at'))
        obj.startedAt = ApiClient.convertToType(data['started_at'], 'Date');
      if (data.hasOwnProperty('finished_at'))
        obj.finishedAt = ApiClient.convertToType(data['finished_at'], 'Date');
    }
    return obj;
  }
}

/**
 * @member {Number} id
 */
RoundCompetitionRead.prototype.id = undefined;

/**
 * @member {Date} createTime
 */
RoundCompetitionRead.prototype.createTime = undefined;

/**
 * @member {Date} updateTime
 */
RoundCompetitionRead.prototype.updateTime = undefined;

/**
 * @member {Boolean} hidden
 */
RoundCompetitionRead.prototype.hidden = undefined;

/**
 * @member {Boolean} pause
 */
RoundCompetitionRead.prototype.pause = undefined;

/**
 * @member {String} name
 */
RoundCompetitionRead.prototype.name = undefined;

/**
 * @member {String} displayName
 */
RoundCompetitionRead.prototype.displayName = undefined;

/**
 * @member {Number} roundDuration
 */
RoundCompetitionRead.prototype.roundDuration = undefined;

/**
 * @member {Date} toBeStartedAt
 */
RoundCompetitionRead.prototype.toBeStartedAt = undefined;

/**
 * @member {Date} startedAt
 */
RoundCompetitionRead.prototype.startedAt = undefined;

/**
 * @member {Date} finishedAt
 */
RoundCompetitionRead.prototype.finishedAt = undefined;

