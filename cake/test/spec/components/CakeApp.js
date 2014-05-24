'use strict';

describe('Main', function () {
  var CakeApp, component;

  beforeEach(function () {
    var container = document.createElement('div');
    container.id = 'content';
    document.body.appendChild(container);

    CakeApp = require('../../../src/scripts/components/CakeApp.jsx');
    component = CakeApp();
  });

  it('should create a new instance of CakeApp', function () {
    expect(component).toBeDefined();
  });
});
