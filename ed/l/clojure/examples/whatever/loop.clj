(ns clojure.examples.loop
  (:gen-class))

(defn main []
  (print "\n[doseq]: ")
  (doseq [v [0 1 2]]
    (print v, " ")
  ) ;; 0 1 2

  (print "\n[loop]: ")
  (loop [x 5]
    (when (> x 0)
      (print x, " ")
      (recur (- x 1))
    )
  ) ;; 5  4  3  2  1
)
(main)
