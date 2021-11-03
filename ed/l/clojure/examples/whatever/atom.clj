(ns clojure.examples.atom
  (:gen-class))

(defn main []
  (def a (atom 1))
  (println "atom a:", a)
  ;; atom a: #object[clojure.lang.Atom 0x210386e0 {:status :ready, :val 1}]
  (println "atom a:", @a) ;; 1
)
(main)
