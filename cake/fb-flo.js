var flo = require('fb-flo'),
    path = require('path'),
    fs = require('fs');

var sourceDirToWatch = './src'

var server = flo(
  sourceDirToWatch,
  {
    port: 8888,
    host: 'localhost',
    verbose: true,
    glob: [
       // All JS files in `sourceDirToWatch` and subdirectories
      '**/*.js',
       // All CSS files in `sourceDirToWatch` and subdirectories
      '**/*.css'
    ]
  },
  function resolver(filepath, callback) {
    // 1. Call into your compiler / bundler.
    // 2. Assuming that `bundle.js` is your output file, update `bundle.js`
    //    and `bundle.css` when a JS or CSS file changes.
    callback({
      resourceURL: 'bundle' + path.extname(filepath),
      // any string-ish value is acceptable. i.e. strings, Buffers etc.
      contents: fs.readFileSync(filepath)
    });
  }
);
