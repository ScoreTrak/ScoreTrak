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
 * The PropertyCreate model module.
 * @module model/PropertyCreate
 * @version 0.1.0
 */
export class PropertyCreate {
  /**
   * Constructs a new <code>PropertyCreate</code>.
   * @alias module:model/PropertyCreate
   * @class
   * @param id {Number} 
   * @param createTime {Date} 
   * @param updateTime {Date} 
   * @param competitionId {Number} 
   * @param teamId {Number} 
   * @param key {String} 
   * @param value {String} 
   * @param status {module:model/PropertyCreate.StatusEnum} 
   */
  constructor(id, createTime, updateTime, competitionId, teamId, key, value, status) {
    this.id = id;
    this.createTime = createTime;
    this.updateTime = updateTime;
    this.competitionId = competitionId;
    this.teamId = teamId;
    this.key = key;
    this.value = value;
    this.status = status;
  }

  /**
   * Constructs a <code>PropertyCreate</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/PropertyCreate} obj Optional instance to populate.
   * @return {module:model/PropertyCreate} The populated <code>PropertyCreate</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new PropertyCreate();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'Number');
      if (data.hasOwnProperty('create_time'))
        obj.createTime = ApiClient.convertToType(data['create_time'], 'Date');
      if (data.hasOwnProperty('update_time'))
        obj.updateTime = ApiClient.convertToType(data['update_time'], 'Date');
      if (data.hasOwnProperty('competition_id'))
        obj.competitionId = ApiClient.convertToType(data['competition_id'], 'Number');
      if (data.hasOwnProperty('team_id'))
        obj.teamId = ApiClient.convertToType(data['team_id'], 'Number');
      if (data.hasOwnProperty('key'))
        obj.key = ApiClient.convertToType(data['key'], 'String');
      if (data.hasOwnProperty('value'))
        obj.value = ApiClient.convertToType(data['value'], 'String');
      if (data.hasOwnProperty('status'))
        obj.status = ApiClient.convertToType(data['status'], 'String');
    }
    return obj;
  }
}

/**
 * @member {Number} id
 */
PropertyCreate.prototype.id = undefined;

/**
 * @member {Date} createTime
 */
PropertyCreate.prototype.createTime = undefined;

/**
 * @member {Date} updateTime
 */
PropertyCreate.prototype.updateTime = undefined;

/**
 * @member {Number} competitionId
 */
PropertyCreate.prototype.competitionId = undefined;

/**
 * @member {Number} teamId
 */
PropertyCreate.prototype.teamId = undefined;

/**
 * @member {String} key
 */
PropertyCreate.prototype.key = undefined;

/**
 * @member {String} value
 */
PropertyCreate.prototype.value = undefined;

/**
 * Allowed values for the <code>status</code> property.
 * @enum {String}
 * @readonly
 */
PropertyCreate.StatusEnum = {
  /**
   * value: "view"
   * @const
   */
  view: "view",

  /**
   * value: "edit"
   * @const
   */
  edit: "edit",

  /**
   * value: "hide"
   * @const
   */
  hide: "hide"
};
/**
 * @member {module:model/PropertyCreate.StatusEnum} status
 * @default 'view'
 */
PropertyCreate.prototype.status = 'view';

