(ns clojure.examples.struct
  (:gen-class))

(defn main []
   (defstruct Employee :Name :ID)
   (def e (struct Employee "John" 1))
   (println "Employee: ", e)
)
(main)
