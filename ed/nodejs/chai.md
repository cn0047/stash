chai
-
4.0.2

````js
expect(arthur).to.not.be.a.model('person');
expect([1, 2]).to.be.an('array').that.does.not.include(3);
expect({a: 1}).to.deep.equal({a: 1});
expect({a: {b: ['x', 'y']}}).to.have.nested.property('a.b[1]');
expect({a: 1}).to.have.own.property('a');
expect('foo').to.be.a('string');
````

````js
````
