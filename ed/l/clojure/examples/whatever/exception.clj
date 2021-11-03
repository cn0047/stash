(ns clojure.examples.exception
  (:gen-class))

(defn main []
  (try
    (aget (int-array [1 2 3]) 5)
  (catch Exception e (
    println (str "caught exception: " (.toString e))
  ))
  (finally (
    println "final")
  ))
  (println "...")
)
(main)
