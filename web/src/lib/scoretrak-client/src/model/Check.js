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
import {Competition} from './Competition';
import {Round} from './Round';
import {Service} from './Service';

/**
 * The Check model module.
 * @module model/Check
 * @version 0.1.0
 */
export class Check {
  /**
   * Constructs a new <code>Check</code>.
   * @alias module:model/Check
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
   * @param competition {module:model/Competition} 
   * @param rounds {module:model/Round} 
   * @param services {module:model/Service} 
   */
  constructor(id, createTime, updateTime, pause, hidden, competitionId, log, error, passed, competition, rounds, services) {
    this.id = id;
    this.createTime = createTime;
    this.updateTime = updateTime;
    this.pause = pause;
    this.hidden = hidden;
    this.competitionId = competitionId;
    this.log = log;
    this.error = error;
    this.passed = passed;
    this.competition = competition;
    this.rounds = rounds;
    this.services = services;
  }

  /**
   * Constructs a <code>Check</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Check} obj Optional instance to populate.
   * @return {module:model/Check} The populated <code>Check</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new Check();
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
      if (data.hasOwnProperty('competition'))
        obj.competition = Competition.constructFromObject(data['competition']);
      if (data.hasOwnProperty('rounds'))
        obj.rounds = Round.constructFromObject(data['rounds']);
      if (data.hasOwnProperty('services'))
        obj.services = Service.constructFromObject(data['services']);
    }
    return obj;
  }
}

/**
 * @member {Number} id
 */
Check.prototype.id = undefined;

/**
 * @member {Date} createTime
 */
Check.prototype.createTime = undefined;

/**
 * @member {Date} updateTime
 */
Check.prototype.updateTime = undefined;

/**
 * @member {Boolean} pause
 */
Check.prototype.pause = undefined;

/**
 * @member {Boolean} hidden
 */
Check.prototype.hidden = undefined;

/**
 * @member {Number} competitionId
 */
Check.prototype.competitionId = undefined;

/**
 * @member {String} log
 */
Check.prototype.log = undefined;

/**
 * @member {String} error
 */
Check.prototype.error = undefined;

/**
 * @member {Boolean} passed
 */
Check.prototype.passed = undefined;

/**
 * @member {module:model/Competition} competition
 */
Check.prototype.competition = undefined;

/**
 * @member {module:model/Round} rounds
 */
Check.prototype.rounds = undefined;

/**
 * @member {module:model/Service} services
 */
Check.prototype.services = undefined;

