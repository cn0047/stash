(ns clojure.examples.agent
  (:gen-class))

(defn main []
  (def counter (agent 0))
  (println "counter:", @counter)

  (send counter + 1)
  (await counter)
  (println "counter:", @counter)
)
(main)
