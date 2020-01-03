Chart
-

[doc](https://developers.google.com/chart/interactive/docs/)

````js
data.addColumn('string', 'Topping');
data.addColumn('number', 'Slices');
data.addColumn('date', 'Dates');
data.addColumn('datetime', 'Time');

const options = {
  legend: {position: 'none'},
  title: '',
  vAxis: {
    viewWindow: { max: 320, min: 0},
  }
};
````
