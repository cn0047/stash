(ns clojure.examples.meta
  (:gen-class))

(defn main []
  (def numbers (with-meta [1 2 3] {:prop "values"}))
  (println "numbers: ", numbers, "; meta: ", (meta numbers))
  ;; numbers:  [1 2 3] ; meta:  {:prop values}
)
(main)
