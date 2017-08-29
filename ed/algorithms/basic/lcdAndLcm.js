/**
 * Largest (Greatest) Common Divider.
 */
function lcd(a, b) {
  if (b == 0) return a;
  return gcd(b, a % b);
}

console.log(gcd(4, 6)); // 2
console.log(gcd(15, 25)); // 5

/**
 * Least Common Multiple.
 */
function lcm(a, b) {
  return a / lcd(a, b) * b;
}
