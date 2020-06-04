// Cypher



// NODES
CREATE (n:Person {name:'James Bond', code:'007', active:true}) RETURN n;
CREATE (n:Person {name:'Moneypenny', code:'mp', active:true}) RETURN n;
CREATE (n:Person {name:'Felix Leiter', code:'felix'});
CREATE (n:Person {name:'008', code:'008'}) RETURN n;
CREATE (n:Person {name:'Q', code: 'q'});
CREATE (n:Person {name:'M', code: 'm'});
CREATE (n:Organization {name:'MI6'});
CREATE (n:Organization {name:'CIA'});
CREATE (n:Organization {name:'test'});
CREATE (n:Country {name:'UK'});
CREATE (n:Country {name:'USA'});

// get all
MATCH (n:Person) RETURN n;
MATCH (n:Organization) RETURN n;

// get
MATCH (n:Person {code:'007'}) RETURN n;
MATCH (n:Organization {name:'test'}) RETURN n;
MATCH (n:Person) WHERE n.active = true RETURN n;

// get 2
MATCH (p:Person {code:'007'})
MATCH (o:Organization {name:'MI6'})
RETURN p, o;

// update
MATCH (n:Organization {name:'test'}) SET n.name = 'TEST_ORG' RETURN n.name;
MATCH (n:Organization {name:'TEST_ORG'}) RETURN n;

// delete
MATCH (n:Person {name:'James Bond'}) DELETE n;
MATCH (n:Organization {name:'test'}) DELETE n;
MATCH (n:Organization {name:'TEST_ORG'}) DELETE n;

// select all
MATCH (n) RETURN n LIMIT 100;
MATCH (n:Organization) RETURN n;



// RELATIONSHIPS
// bond familiar with
MATCH (a:Person),(b:Person) WHERE a.code = '007' AND b.code = 'mp'    CREATE (a)-[r:FAMILIAR]->(b) RETURN type(r);
MATCH (a:Person),(b:Person) WHERE a.code = '007' AND b.code = 'felix' CREATE (a)-[r:FAMILIAR]->(b);
MATCH (a:Person),(b:Person) WHERE a.code = '007' AND b.code = '008'   CREATE (a)-[r:FAMILIAR]->(b);
MATCH (a:Person),(b:Person) WHERE a.code = '007' AND b.code = 'q'     CREATE (a)-[r:FAMILIAR]->(b);
MATCH (a:Person),(b:Person) WHERE a.code = '007' AND b.code = 'm'     CREATE (a)-[r:FAMILIAR]->(b);
// works at
MATCH (a:Person),(b:Organization) WHERE a.code = '007'   AND b.name = 'MI6' CREATE (a)-[r:WORKS_AT]->(b);
MATCH (a:Person),(b:Organization) WHERE a.code = '008'   AND b.name = 'MI6' CREATE (a)-[r:WORKS_AT]->(b);
MATCH (a:Person),(b:Organization) WHERE a.code = 'mp'    AND b.name = 'MI6' CREATE (a)-[r:WORKS_AT]->(b);
MATCH (a:Person),(b:Organization) WHERE a.code = 'q'     AND b.name = 'MI6' CREATE (a)-[r:WORKS_AT]->(b);
MATCH (a:Person),(b:Organization) WHERE a.code = 'm'     AND b.name = 'MI6' CREATE (a)-[r:WORKS_AT]->(b);
MATCH (a:Person),(b:Organization) WHERE a.code = 'felix' AND b.name = 'CIA' CREATE (a)-[r:WORKS_AT]->(b);
// country of organization
MATCH (o:Organization {name: 'MI6'}),(c:Country {name:'UK'}) CREATE (o)-[r:COUNTRY]->(c) RETURN o, c;
MATCH (o:Organization {name: 'CIA'}),(c:Country {name:'USA'}) CREATE (o)-[r:COUNTRY]->(c) RETURN o, c;
// wisited
MATCH (p:Person {code: '007'}),(c:Country {name:'USA'}) CREATE (p)-[r:WISITED]->(c) RETURN p, c;

// node's relationships
MATCH (:Person {code: '007'})-[r]-() RETURN r;
MATCH (:Organization {name: 'MI6'})-[r]-() RETURN r;

// bond familiar with
MATCH (b:Person {code: '007'})-[:FAMILIAR]->(n) RETURN b, n.name ORDER BY n.name;

// bond works at
MATCH (b:Person {code: '007'})-[:WORKS_AT]->(n) RETURN b, n.name ORDER BY n.name;

// country of MI6
MATCH (:Organization {name: 'MI6'})-[:COUNTRY]->(c) RETURN c;

// country for persons
MATCH (p:Person)-[:WORKS_AT]->(o)-[:COUNTRY]->(c) RETURN p, o, c;
MATCH (p:Person)-[]->()-[]->(c:Country) RETURN p, c;



//
MATCH (a:Person)-[*2]->(b:Country) RETURN a, b;



// AGGREGATION
MATCH (n:Person) RETURN COUNT(n) as count;
MATCH (n:Person) RETURN DISTINCT LABELS(n), COUNT(*);
MATCH (n:Person) RETURN DISTINCT n.code;



// ADMIN
call db.schema();
