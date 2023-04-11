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
    describe('HostsIdBody', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.HostsIdBody();
      });

      it('should create an instance of HostsIdBody', function() {
        // TODO: update the code to test HostsIdBody
        expect(instance).to.be.a(EntSchemaApi.HostsIdBody);
      });

      it('should have the property address (base name: "address")', function() {
        // TODO: update the code to test the property address
        expect(instance).to.have.property('address');
        // expect(instance.address).to.be(expectedValueLiteral);
      });

      it('should have the property addressListRange (base name: "address_list_range")', function() {
        // TODO: update the code to test the property addressListRange
        expect(instance).to.have.property('addressListRange');
        // expect(instance.addressListRange).to.be(expectedValueLiteral);
      });

      it('should have the property services (base name: "services")', function() {
        // TODO: update the code to test the property services
        expect(instance).to.have.property('services');
        // expect(instance.services).to.be(expectedValueLiteral);
      });

      it('should have the property team (base name: "team")', function() {
        // TODO: update the code to test the property team
        expect(instance).to.have.property('team');
        // expect(instance.team).to.be(expectedValueLiteral);
      });

    });
  });

}));
