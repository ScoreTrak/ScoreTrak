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
    describe('TeamRead', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.TeamRead();
      });

      it('should create an instance of TeamRead', function() {
        // TODO: update the code to test TeamRead
        expect(instance).to.be.a(EntSchemaApi.TeamRead);
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

      it('should have the property index (base name: "index")', function() {
        // TODO: update the code to test the property index
        expect(instance).to.have.property('index');
        // expect(instance.index).to.be(expectedValueLiteral);
      });

    });
  });

}));
