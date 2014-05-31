/** @jsx React.DOM */

var Register = React.createClass({
  render: function() {
    return (
      <form action="">
        <input type="text" name="email" refs="email"></input>
        <input type="text" name="password" refs="password"></input>
        <button> Submit</button>
      </form>
    );
  }
});
