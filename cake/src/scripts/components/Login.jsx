/**
 * @jsx React.DOM
 */

'use strict';

var React = require('react/addons');
require('../../styles/Login.css');

var Login = React.createClass({
  render: function () {
    return (
        <div>
          <form>
            <input type="text"></input>
            <input type="text"></input>
            <input type="button"></input>
          </form>
        </div>
      );
  }
});

module.exports = Login;
