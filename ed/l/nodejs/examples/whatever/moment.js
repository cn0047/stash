const moment = require('moment');


const f = (lineItem, wli) => {
  let liStart = moment.utc(lineItem.ship_start, ['YYYY/MM/DD', 'MM/DD/YYYY']);
  let liEnd = moment.utc(lineItem.ship_end, ['YYYY/MM/DD', 'MM/DD/YYYY']);

  if (typeof lineItem.ship_start === 'number' && typeof lineItem.ship_end === 'number') {
    liStart = moment.utc(lineItem.ship_start);
    liEnd = moment.utc(lineItem.ship_end);
  }

  const matchWindow = ({ ship_start, ship_end }) => {
      return liStart.isSameOrAfter(ship_start, 'day') && liEnd.isSameOrBefore(ship_end, 'day');
  }

  return matchWindow(wli);
}

let lineItem, wli;

lineItem = {ship_start: '2021-03-26T00:00:00.000Z', ship_end: '2021-04-09T00:00:00.000Z'};
wli      = {ship_start: '2021-03-26T00:00:00Z',     ship_end: '2021-04-09T00:00:00Z'};
console.log('case 1:', f(lineItem, wli));
wli      = {ship_start: '2021-04-20T00:00:00Z',     ship_end: '2021-04-24T00:00:00Z'};
console.log('case 2:', f(lineItem, wli));

lineItem = {};
wli = {ship_start: '2022-01-03T00:00:00.000Z', ship_end: '2022-01-31T00:00:00.000Z'};
console.log('case 3:', f(lineItem, wli));

lineItem = {ship_start: '2022-01-06T00:00:00.000Z', ship_end: '2022-01-31T00:00:00.000Z'};
wli = {ship_start: '2022-01-06T00:00:00.000Z', ship_end: '2022-01-31T00:00:00.000Z'};
console.log('case 4:', f(lineItem, wli));

lineItem = {ship_start: 1639008000000, ship_end: 1640995199999};
wli = {ship_start: '2021-12-09T00:00:00.000Z', ship_end: '2021-12-31T00:00:00.000Z'};
console.log('case 5:', f(lineItem, wli));
