(ns clojure.examples.hw
  (:gen-class))

;; This is "Hellow World" program.
(defn hello-world []
  (def msg "Hello World")
  (println msg)
)
(hello-world)
