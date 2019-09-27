const del = require('del');
const gulp = require('gulp');
const htmlmin = require('gulp-htmlmin');
const less = require('gulp-less');
const cleanCSS = require('gulp-clean-css');
const rename = require('gulp-rename');
const uglify = require('gulp-uglify');
const concat = require('gulp-concat');
const livereload = require('gulp-livereload');

const paths = {
  html: {
    src: 'src/**/*.html',
    dest: 'dist/'
  },
  styles: {
    src: 'src/**/*.less',
    dest: 'dist/s/'
  },
  scripts: {
    src: 'src/**/*.js',
    dest: 'dist/js/'
  }
};

function clean() {
  return del(['dist']);
}

function html() {
  return gulp.src(paths.html.src)
    .pipe(htmlmin({collapseWhitespace: true}))
    .pipe(gulp.dest(paths.html.dest))
    .pipe(livereload());
}

function styles() {
  return gulp.src(paths.styles.src)
    .pipe(less())
    .pipe(cleanCSS())
    .pipe(rename({basename: 'app', suffix: '.min'}))
    .pipe(gulp.dest(paths.styles.dest))
    .pipe(livereload());
}

function scripts() {
  return gulp.src(paths.scripts.src, { sourcemaps: true })
    .pipe(uglify())
    .pipe(concat('app.min.js'))
    .pipe(gulp.dest(paths.scripts.dest))
    .pipe(livereload());
}

function watch() {
  livereload.listen();

  gulp.watch(paths.html.src, html);
  gulp.watch(paths.styles.src, styles);
  gulp.watch(paths.scripts.src, scripts);
}

exports.clean = clean;
exports.html = html;
exports.styles = styles;
exports.scripts = scripts;
exports.watch = watch;

const build = gulp.series(clean, gulp.parallel(html, styles, scripts));

gulp.task('default', build);
gulp.task('build', build);
