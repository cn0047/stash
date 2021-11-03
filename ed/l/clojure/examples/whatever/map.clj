(ns clojure.examples.map
  (:gen-class))

(defn main []
  (def m (hash-map "a" "1" "b" "2"))
  (println "map m:", m) ;; map m: {a 1, b 2}
)
(main)
