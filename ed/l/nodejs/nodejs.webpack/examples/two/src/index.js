console.log('Init btn.');

document.getElementById('btn').addEventListener('click', function (e) {
  e.target.innerHTML = parseInt(e.target.innerHTML) + 1;
});
