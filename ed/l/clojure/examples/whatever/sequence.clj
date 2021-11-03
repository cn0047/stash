(ns clojure.examples.seq
  (:gen-class))

(defn main []
  (def s (set '(1 2 3)))
  (println "seq s:", s) ;; seq s: #{1 3 2}
)
(main)
