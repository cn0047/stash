// @see: https://owasp.org/www-community/attacks/CSV_Injection

function sanitizeForCSV(val) {
  let result = (val || '').toString();
  result = result.replace('"', '""');
  // result = result.replace("'", "''");
  result = /^[=\-+@\t\r]/.test(result) ? "'" + result : result;

  return '"' + result + '"';
}

const r = [
  sanitizeForCSV('=7*7'),
  sanitizeForCSV('-4-1'),
  sanitizeForCSV('+2+3'),
  sanitizeForCSV('@next'),
  sanitizeForCSV('\ttab1'),
  sanitizeForCSV('0x09tab2'),
  sanitizeForCSV('\r1ret'),
  sanitizeForCSV('0x0D2ret'),
  sanitizeForCSV('=1+2\';=1+2'),
  sanitizeForCSV("=1+2';=1+2"),
  sanitizeForCSV('=1+2";=1+2'), // "'=1+2"";=1+2"
  sanitizeForCSV('=1+2\'" ;,=1+2'), // "'=1+2'"" ;,=1+2"
];
console.log(r);
