var fetch = require('node-fetch');


function fetchData(url) {
  return fetch(url).then(resp => resp.json()).then(data => data);
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

async function retry(fn, n, ...args) {
  for (let i = 0; i < n; i++) {
    try {
      console.log('retry #', i);
      return await fn(...args);
    } catch {}
    await sleep(2000 * i);
  }

  throw new Error(`failed after ${n} retries`);
}

async function retry2(n, ctx, fn, ...args) {
  for (let i = 0; i < n; i++) {
    try {
      console.log('retry #', i);
      return await fn.call(ctx, ...args);
    } catch {}
    await sleep(5000 * i);
  }

  throw new Error(`failed after ${n} retries`);
}

async function main() {
  let r = await retry(fetchData, 5, 'xhttps://api.icndb.com/jokes/random');
  console.log(r);
}

main();
