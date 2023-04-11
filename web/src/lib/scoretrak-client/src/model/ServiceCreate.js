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
 * The ServiceCreate model module.
 * @module model/ServiceCreate
 * @version 0.1.0
 */
export class ServiceCreate {
  /**
   * Constructs a new <code>ServiceCreate</code>.
   * @alias module:model/ServiceCreate
   * @class
   * @param id {Number} 
   * @param createTime {Date} 
   * @param updateTime {Date} 
   * @param pause {Boolean} 
   * @param hidden {Boolean} 
   * @param competitionId {Number} 
   * @param teamId {Number} 
   * @param name {String} 
   * @param displayName {String} 
   * @param weight {Number} 
   * @param pointBoost {Number} 
   * @param roundUnits {Number} 
   * @param roundDelay {Number} 
   */
  constructor(id, createTime, updateTime, pause, hidden, competitionId, teamId, name, displayName, weight, pointBoost, roundUnits, roundDelay) {
    this.id = id;
    this.createTime = createTime;
    this.updateTime = updateTime;
    this.pause = pause;
    this.hidden = hidden;
    this.competitionId = competitionId;
    this.teamId = teamId;
    this.name = name;
    this.displayName = displayName;
    this.weight = weight;
    this.pointBoost = pointBoost;
    this.roundUnits = roundUnits;
    this.roundDelay = roundDelay;
  }

  /**
   * Constructs a <code>ServiceCreate</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ServiceCreate} obj Optional instance to populate.
   * @return {module:model/ServiceCreate} The populated <code>ServiceCreate</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new ServiceCreate();
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
      if (data.hasOwnProperty('team_id'))
        obj.teamId = ApiClient.convertToType(data['team_id'], 'Number');
      if (data.hasOwnProperty('name'))
        obj.name = ApiClient.convertToType(data['name'], 'String');
      if (data.hasOwnProperty('display_name'))
        obj.displayName = ApiClient.convertToType(data['display_name'], 'String');
      if (data.hasOwnProperty('weight'))
        obj.weight = ApiClient.convertToType(data['weight'], 'Number');
      if (data.hasOwnProperty('point_boost'))
        obj.pointBoost = ApiClient.convertToType(data['point_boost'], 'Number');
      if (data.hasOwnProperty('round_units'))
        obj.roundUnits = ApiClient.convertToType(data['round_units'], 'Number');
      if (data.hasOwnProperty('round_delay'))
        obj.roundDelay = ApiClient.convertToType(data['round_delay'], 'Number');
    }
    return obj;
  }
}

/**
 * @member {Number} id
 */
ServiceCreate.prototype.id = undefined;

/**
 * @member {Date} createTime
 */
ServiceCreate.prototype.createTime = undefined;

/**
 * @member {Date} updateTime
 */
ServiceCreate.prototype.updateTime = undefined;

/**
 * @member {Boolean} pause
 */
ServiceCreate.prototype.pause = undefined;

/**
 * @member {Boolean} hidden
 */
ServiceCreate.prototype.hidden = undefined;

/**
 * @member {Number} competitionId
 */
ServiceCreate.prototype.competitionId = undefined;

/**
 * @member {Number} teamId
 */
ServiceCreate.prototype.teamId = undefined;

/**
 * @member {String} name
 */
ServiceCreate.prototype.name = undefined;

/**
 * @member {String} displayName
 */
ServiceCreate.prototype.displayName = undefined;

/**
 * @member {Number} weight
 */
ServiceCreate.prototype.weight = undefined;

/**
 * @member {Number} pointBoost
 */
ServiceCreate.prototype.pointBoost = undefined;

/**
 * @member {Number} roundUnits
 */
ServiceCreate.prototype.roundUnits = undefined;

/**
 * @member {Number} roundDelay
 */
ServiceCreate.prototype.roundDelay = undefined;

