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
    describe('ServicePropertiesList', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.ServicePropertiesList();
      });

      it('should create an instance of ServicePropertiesList', function() {
        // TODO: update the code to test ServicePropertiesList
        expect(instance).to.be.a(EntSchemaApi.ServicePropertiesList);
      });

      it('should have the property id (base name: "id")', function() {
        // TODO: update the code to test the property id
        expect(instance).to.have.property('id');
        // expect(instance.id).to.be(expectedValueLiteral);
      });

      it('should have the property log (base name: "log")', function() {
        // TODO: update the code to test the property log
        expect(instance).to.have.property('log');
        // expect(instance.log).to.be(expectedValueLiteral);
      });

      it('should have the property error (base name: "error")', function() {
        // TODO: update the code to test the property error
        expect(instance).to.have.property('error');
        // expect(instance.error).to.be(expectedValueLiteral);
      });

      it('should have the property passed (base name: "passed")', function() {
        // TODO: update the code to test the property passed
        expect(instance).to.have.property('passed');
        // expect(instance.passed).to.be(expectedValueLiteral);
      });

    });
  });

}));
