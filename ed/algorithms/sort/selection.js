function selectionSort(a) {
  for (let i = 0; i < a.length; i++) {
    for (let j = i + 1; j < a.length; j++) {
      if (a[i] > a[j]) {
        let t = a[i];
        a[i] = a[j];
        a[j] = t;
      }
    }
  }
  return a;
}

console.log(selectionSort([5, 3, 1, 2, 4]));
console.log(selectionSort([2, 4, 6, 8, 3]));
console.log(selectionSort([2, 4, 6, 8, 9]));
console.log(selectionSort([2, 4, 6, 8, 0]));
console.log(selectionSort(['f', 's', 'w', 'e', 'l', 'x', 'v', 'd', 'a']));
