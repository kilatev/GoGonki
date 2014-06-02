/** @jsx React.DOM */

var Register = React.createClass({
  render: function() {
    return (
      <form action="/create-room">
        <input type="text" name="roomname" refs="roomname"></input>
        <input type="text" name="roomsize" refs="roomsize"></input>
        <input type="text" name="privacy" refs="privacy"></input>
        <button> Submit</button>
      </form>
    );
  }
});
