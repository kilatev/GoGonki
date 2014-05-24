'use strict';

describe('Register', function () {
  var Register, component;

  beforeEach(function () {
    Register = require('../../../src/scripts/components/Register.jsx');
    component = Register();
  });

  it('should create a new instance of Register', function () {
    expect(component).toBeDefined();
  });
});
