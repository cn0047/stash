grunt
-

````sh
npm install grunt-init # project scaffolding
npm install grunt-cli

node_modules/.bin/grunt

node_modules/.bin/grunt taskName -v
node_modules/.bin/grunt watch
````

[Live reload](https://github.com/gruntjs/grunt-contrib-watch#optionslivereload)

````
# in config
watch: {
  scripts: {
    files: ['src/*.js', 'src/*.html', 'src/*.css'],
    tasks: ['uglify', 'htmlmin', 'cssmin'],
    options: {
      spawn: false,
      livereload: true,
    },
  },
},

# in html file
<script src="//localhost:35729/livereload.js"></script>
````
