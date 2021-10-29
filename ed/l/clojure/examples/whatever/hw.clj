(ns clojure.examples.hello
  (:gen-class))

;; This is "Hellow World" program.
(defn hello-world []
  (def msg "Hello World")
  (println msg)
)
(hello-world)
