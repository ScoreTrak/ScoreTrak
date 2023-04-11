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
 * The RoundUpdate model module.
 * @module model/RoundUpdate
 * @version 0.1.0
 */
export class RoundUpdate {
  /**
   * Constructs a new <code>RoundUpdate</code>.
   * @alias module:model/RoundUpdate
   * @class
   * @param id {Number} 
   * @param createTime {Date} 
   * @param updateTime {Date} 
   * @param competitionId {Number} 
   * @param roundNumber {Number} 
   * @param note {String} 
   * @param err {String} 
   * @param startedAt {Date} 
   * @param finishedAt {Date} 
   */
  constructor(id, createTime, updateTime, competitionId, roundNumber, note, err, startedAt, finishedAt) {
    this.id = id;
    this.createTime = createTime;
    this.updateTime = updateTime;
    this.competitionId = competitionId;
    this.roundNumber = roundNumber;
    this.note = note;
    this.err = err;
    this.startedAt = startedAt;
    this.finishedAt = finishedAt;
  }

  /**
   * Constructs a <code>RoundUpdate</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/RoundUpdate} obj Optional instance to populate.
   * @return {module:model/RoundUpdate} The populated <code>RoundUpdate</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new RoundUpdate();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'Number');
      if (data.hasOwnProperty('create_time'))
        obj.createTime = ApiClient.convertToType(data['create_time'], 'Date');
      if (data.hasOwnProperty('update_time'))
        obj.updateTime = ApiClient.convertToType(data['update_time'], 'Date');
      if (data.hasOwnProperty('competition_id'))
        obj.competitionId = ApiClient.convertToType(data['competition_id'], 'Number');
      if (data.hasOwnProperty('round_number'))
        obj.roundNumber = ApiClient.convertToType(data['round_number'], 'Number');
      if (data.hasOwnProperty('note'))
        obj.note = ApiClient.convertToType(data['note'], 'String');
      if (data.hasOwnProperty('err'))
        obj.err = ApiClient.convertToType(data['err'], 'String');
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
RoundUpdate.prototype.id = undefined;

/**
 * @member {Date} createTime
 */
RoundUpdate.prototype.createTime = undefined;

/**
 * @member {Date} updateTime
 */
RoundUpdate.prototype.updateTime = undefined;

/**
 * @member {Number} competitionId
 */
RoundUpdate.prototype.competitionId = undefined;

/**
 * @member {Number} roundNumber
 */
RoundUpdate.prototype.roundNumber = undefined;

/**
 * @member {String} note
 */
RoundUpdate.prototype.note = undefined;

/**
 * @member {String} err
 */
RoundUpdate.prototype.err = undefined;

/**
 * @member {Date} startedAt
 */
RoundUpdate.prototype.startedAt = undefined;

/**
 * @member {Date} finishedAt
 */
RoundUpdate.prototype.finishedAt = undefined;

