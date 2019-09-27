module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    uglify: {
      one_example_app: {
        files: {
          'dist/app.min.js': ['src/index.js']
        }
      }
    },

    htmlmin: {
      dist: {
        options: {
          removeComments: true,
          collapseWhitespace: true
        },
        files: {
          'dist/index.html': 'src/index.html',
        }
      }
    },

    cssmin: {
      options: {
        mergeIntoShorthands: false,
        roundingPrecision: -1
      },
      target: {
        files: {
          'dist/app.min.css': ['src/index.css']
        }
      }
    },

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
  });

  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-htmlmin');
  grunt.loadNpmTasks('grunt-contrib-cssmin');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['uglify', 'htmlmin', 'cssmin', 'watch']);
};