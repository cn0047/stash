System = function(name, hp){
  this.name = name;
  this.hp = [];
  this.setHP = function(hp){
     this.hp[0] = hp;
     this.hp[1] = hp;
  }
  this.setHP(hp);
}
Weapon = function(name, hp){
    System.apply(this, arguments);
}

console.log(new Weapon("Gun", 10));
console.log(new System("Hangar", 10));

/*
Weapon {name: "Gun", hp: Array[2], setHP: function}
System {name: "Hangar", hp: Array[2], setHP: function}
*/
