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
(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.EntSchemaApi);
  }
}(this, function(expect, EntSchemaApi) {
  'use strict';

  var instance;

  describe('(package)', function() {
    describe('ServiceCompetitionRead', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.ServiceCompetitionRead();
      });

      it('should create an instance of ServiceCompetitionRead', function() {
        // TODO: update the code to test ServiceCompetitionRead
        expect(instance).to.be.a(EntSchemaApi.ServiceCompetitionRead);
      });

      it('should have the property id (base name: "id")', function() {
        // TODO: update the code to test the property id
        expect(instance).to.have.property('id');
        // expect(instance.id).to.be(expectedValueLiteral);
      });

      it('should have the property createTime (base name: "create_time")', function() {
        // TODO: update the code to test the property createTime
        expect(instance).to.have.property('createTime');
        // expect(instance.createTime).to.be(expectedValueLiteral);
      });

      it('should have the property updateTime (base name: "update_time")', function() {
        // TODO: update the code to test the property updateTime
        expect(instance).to.have.property('updateTime');
        // expect(instance.updateTime).to.be(expectedValueLiteral);
      });

      it('should have the property hidden (base name: "hidden")', function() {
        // TODO: update the code to test the property hidden
        expect(instance).to.have.property('hidden');
        // expect(instance.hidden).to.be(expectedValueLiteral);
      });

      it('should have the property pause (base name: "pause")', function() {
        // TODO: update the code to test the property pause
        expect(instance).to.have.property('pause');
        // expect(instance.pause).to.be(expectedValueLiteral);
      });

      it('should have the property name (base name: "name")', function() {
        // TODO: update the code to test the property name
        expect(instance).to.have.property('name');
        // expect(instance.name).to.be(expectedValueLiteral);
      });

      it('should have the property displayName (base name: "display_name")', function() {
        // TODO: update the code to test the property displayName
        expect(instance).to.have.property('displayName');
        // expect(instance.displayName).to.be(expectedValueLiteral);
      });

      it('should have the property roundDuration (base name: "round_duration")', function() {
        // TODO: update the code to test the property roundDuration
        expect(instance).to.have.property('roundDuration');
        // expect(instance.roundDuration).to.be(expectedValueLiteral);
      });

      it('should have the property toBeStartedAt (base name: "to_be_started_at")', function() {
        // TODO: update the code to test the property toBeStartedAt
        expect(instance).to.have.property('toBeStartedAt');
        // expect(instance.toBeStartedAt).to.be(expectedValueLiteral);
      });

      it('should have the property startedAt (base name: "started_at")', function() {
        // TODO: update the code to test the property startedAt
        expect(instance).to.have.property('startedAt');
        // expect(instance.startedAt).to.be(expectedValueLiteral);
      });

      it('should have the property finishedAt (base name: "finished_at")', function() {
        // TODO: update the code to test the property finishedAt
        expect(instance).to.have.property('finishedAt');
        // expect(instance.finishedAt).to.be(expectedValueLiteral);
      });

    });
  });

}));
