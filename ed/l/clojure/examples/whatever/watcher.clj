(ns clojure.examples.watcher
  (:gen-class))

(defn main []
  (def x (atom 0))
  (add-watch x :watcher
    (fn [key atom old-state new-state]
      (println "The value of the atom has been changed")
      (println "old-state" old-state)
      (println "new-state" new-state)
    )
  )
  (reset! x 2)
)
(main)
