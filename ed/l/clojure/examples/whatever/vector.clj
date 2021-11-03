(ns clojure.examples.vector
  (:gen-class))

(defn main []
  (def v (vector 1 2 3))
  (println "vector v:", v)
)
(main)
