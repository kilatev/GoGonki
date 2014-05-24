/**
 * @jsx React.DOM
 */

'use strict';

var React = require('react/addons');
var ReactTransitionGroup = React.addons.TransitionGroup;

// CSS
require('../../styles/reset.css');
require('../../styles/main.css');

var imageURL = '../../images/yeoman.png';

var CakeApp = React.createClass({
  render: function() {
    return (
         <Locations>
        <Location path="/" handler={MainPage} />
        <Location path="/users/:username" handler={UserPage} />
      </Locations>
      <div className='main'>
        <ReactTransitionGroup transitionName="fade">
          <img src={imageURL} />
        </ReactTransitionGroup>
      </div>
    );
  }
});

React.renderComponent(<CakeApp />, document.getElementById('content')); // jshint ignore:line

module.exports = CakeApp;
