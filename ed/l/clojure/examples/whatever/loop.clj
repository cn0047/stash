(ns clojure.examples.loop
  (:gen-class))

(defn main []
  (print "\n[loop]: ")
  (loop [x 5]
    (when (> x 0)
      (print x, "")
      (recur (- x 1))
    )
  ) ;; 5  4  3  2  1
  
  (print "\n[while]: ")
  (def x (atom 1))
  (while ( < @x 5 ) (do
    (print @x, "")
    (swap! x inc)
  ))

  (print "\n[doseq]: ")
  (doseq [v [0 1 2]]
    (print v, "")
  ) ;; 0 1 2

  (print "\n[dotimes]: ")
  (dotimes [n 5]
    (print n, "")
  ) ;; 0 1 2 3 4
)
(main)
