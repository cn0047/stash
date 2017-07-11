inst_section = {
  "trumpet"   => "brass",
  "violin"    => "string"
}

h = { "dog" => "canine", "cat" => "feline", "donkey" => "asinine" }
h.length → 3
h["dog"] → "canine"
h["cow"] = "bovine"
h[12] = "dodecine"
h["cat"] = 99
# h → {"cow"=>"bovine", "cat"=>99, 12=>"dodecine", "donkey"=>"asinine", "dog"=>"canine"}
