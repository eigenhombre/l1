;; Return the factorial of `n`:
(def fact
  (fn [n]
    (cond (= 0 n) 1
          :else (*' n (fact (- n 1))))))

(println (fact 100))
