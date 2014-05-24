'use strict';

describe('Login', function () {
  var Login, component;

  beforeEach(function () {
    Login = require('../../../src/scripts/components/Login.jsx');
    component = Login();
  });

  it('should create a new instance of Login', function () {
    expect(component).toBeDefined();
  });
});
