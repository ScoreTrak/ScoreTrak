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
    describe('RoundCreate', function() {
      beforeEach(function() {
        instance = new EntSchemaApi.RoundCreate();
      });

      it('should create an instance of RoundCreate', function() {
        // TODO: update the code to test RoundCreate
        expect(instance).to.be.a(EntSchemaApi.RoundCreate);
      });

      it('should have the property id (base name: "id")', function() {
        // TODO: update the code to test the property id
        expect(instance).to.have.property('id');
        // expect(instance.id).to.be(expectedValueLiteral);
      });

      it('should have the property roundNumber (base name: "round_number")', function() {
        // TODO: update the code to test the property roundNumber
        expect(instance).to.have.property('roundNumber');
        // expect(instance.roundNumber).to.be(expectedValueLiteral);
      });

      it('should have the property note (base name: "note")', function() {
        // TODO: update the code to test the property note
        expect(instance).to.have.property('note');
        // expect(instance.note).to.be(expectedValueLiteral);
      });

      it('should have the property err (base name: "err")', function() {
        // TODO: update the code to test the property err
        expect(instance).to.have.property('err');
        // expect(instance.err).to.be(expectedValueLiteral);
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
