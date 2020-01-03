const sum = require('./../src/sum');

describe('Test suite', () => {
  test('Adds 1 + 2 to equal 3', () => {
    expect(sum(1, 2)).toBe(3);
  });

  test('Object', () => {
    const actual = {
      some: 'payload'
    };

    expect(typeof actual).toBe('object');
  });
});
