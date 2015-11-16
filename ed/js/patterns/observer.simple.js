<?php

var Subject = function () {
    this.listeners = [];
    this.attach = function (l) {
        this.listeners.push(l);
    }
    this.notify = function (e) {
        for (i in this.listeners) {
            this.listeners[i].notify(e);
        }
    }
}
var Tv = function () {
    this.notify = function (e) {
        console.log('Breaking news: %s', e);
    }
}
var Press = function () {
    this.notify = function (e) {
        console.log('Fresh press: %s', e);
    }
}

var s = new Subject();
s.attach(new Tv());
s.attach(new Press());
s.notify('Ukraine win!');

/*
Breaking news: Ukraine win!
Fresh press: Ukraine win!
*/
