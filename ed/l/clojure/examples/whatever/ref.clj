(ns clojure.examples.ref
  (:gen-class))

(defn main []
  (def r (ref 1 :validator pos?))
  (println "ref r:", @r) ;; 1
)
(main)
