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
    describe('ServiceUpdate', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.ServiceUpdate();
      });

      it('should create an instance of ServiceUpdate', function() {
        // TODO: update the code to test ServiceUpdate
        expect(instance).to.be.a(EntSchemaApi.ServiceUpdate);
      });

      it('should have the property id (base name: "id")', function() {
        // TODO: update the code to test the property id
        expect(instance).to.have.property('id');
        // expect(instance.id).to.be(expectedValueLiteral);
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

      it('should have the property weight (base name: "weight")', function() {
        // TODO: update the code to test the property weight
        expect(instance).to.have.property('weight');
        // expect(instance.weight).to.be(expectedValueLiteral);
      });

      it('should have the property pointBoost (base name: "point_boost")', function() {
        // TODO: update the code to test the property pointBoost
        expect(instance).to.have.property('pointBoost');
        // expect(instance.pointBoost).to.be(expectedValueLiteral);
      });

      it('should have the property roundUnits (base name: "round_units")', function() {
        // TODO: update the code to test the property roundUnits
        expect(instance).to.have.property('roundUnits');
        // expect(instance.roundUnits).to.be(expectedValueLiteral);
      });

      it('should have the property reoundDelay (base name: "reound_delay")', function() {
        // TODO: update the code to test the property reoundDelay
        expect(instance).to.have.property('reoundDelay');
        // expect(instance.reoundDelay).to.be(expectedValueLiteral);
      });

    });
  });

}));
