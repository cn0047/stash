'use strict';

var fetch = require('node-fetch');

function fetchQuote() {
  return fetch('https://api.icndb.com/jokes/random').then(function( resp ){
    return resp.json();
  }).then(function( data ){
    return data.value.joke;
  });
}

async function sayJoke() {
  try {
    let result = await fetchQuote();
    console.log({'Joke:': result});
  } catch(err) {
    console.error(err);
  }
}

sayJoke();
