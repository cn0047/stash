const b = require('./bridge');

const objects = [
  {objectID: "example1", tags: [], text: 'Theory: Leonardo DiCaprio in Romeo + Juliet Is the Next Big Style Icon'},
  {objectID: "example2", tags: [], text: 'Leo Turns 40 Today! A Look Back at the Star&#039;s Best Roles'},
  {objectID: "example3", tags: [], text: 'Meghan Markle Completely Changed Her Duchess Style With This Bold Look'},
  {objectID: "example4", tags: [], text: 'The 8 Zara Items Everyone Will Buy This Month'},
  {objectID: "example5", tags: [], text: 'Victoria Beckham Just Explained How She Chose Her Royal Wedding Outfit'},
];
b.index.addObjects(objects, (err, content) => console.log(err, content));
