function returnTrue() {
  console.log('returnTrue');
  return true;
}

function returnFalse() {
  console.log('returnFalse');
  return false;
}

function main() {
  console.log('\ncheck1');
  if (returnTrue() && returnFalse()) {
    console.log('in1');
  }

  console.log('\ncheck2');
  if (returnFalse() && returnTrue()) {
    console.log('in2');
  }

  console.log('\ncheck3');
  if (returnTrue() && returnFalse() && returnTrue()) {
    console.log('in3');
  }
}

main();

/*

check1
returnTrue
returnFalse

check2
returnFalse

check3
returnTrue
returnFalse

*/
