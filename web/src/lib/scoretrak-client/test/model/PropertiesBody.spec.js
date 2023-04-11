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
    describe('PropertiesBody', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.PropertiesBody();
      });

      it('should create an instance of PropertiesBody', function() {
        // TODO: update the code to test PropertiesBody
        expect(instance).to.be.a(EntSchemaApi.PropertiesBody);
      });

      it('should have the property prop (base name: "prop")', function() {
        // TODO: update the code to test the property prop
        expect(instance).to.have.property('prop');
        // expect(instance.prop).to.be(expectedValueLiteral);
      });

    });
  });

}));
