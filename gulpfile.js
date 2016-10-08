'use strict';

const gulp = require('gulp');
const sass = require('gulp-sass');
const autoprefixer = require('gulp-autoprefixer');
const cssnano = require('gulp-cssnano');
const sourcemaps = require('gulp-sourcemaps');

const dir = {
    src: 'resources',
    dist: 'public'
};

const options = {
    sass: { outputStyle: 'nested'},
};

gulp.task('sass', () => {
    return gulp.src(`${dir.src}/scss/*.scss`)
        .pipe(sourcemaps.init())
        .pipe(sass(options.sass).on('error', sass.logError))
        .pipe(autoprefixer())
        .pipe(cssnano())
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest(`${dir.dist}/css`));
});

gulp.task('watch', ['sass'], () => {
    gulp.watch(`${dir.src}/scss/**/*.scss`, ['sass']);
});

gulp.task('default', ['watch']);
