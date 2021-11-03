(ns clojure.examples.set
  (:gen-class))

(defn main []
  (def s (set '(1 2 3)))
  (println "set s:", s) ;; set s: #{1 3 2}
)
(main)
