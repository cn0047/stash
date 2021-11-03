(ns clojure.examples.function
  (:gen-class))

(defn x [msg]
  (println "func x param:", msg) ;; func x param: p1
)

(defn y [msg]
  (str "[y]", msg) ;; return
)

(defn main []
  (x "p1")
  (def v (y "p2"))
  (println "v:", v) ;; v: [y]p2
)
(main)
