person1 = "Tim"
person2 = person1
person1[0] = "J"
puts person1
puts person2
# Jim
# Jim

person1 = "Tim"
person2 = person1.dup
person1[0] = "J"
puts person1
puts person2
# Jim
# Tim

person1.freeze # prevent modifications
person1[0] = "K"
# RuntimeError: can't modify frozen String
