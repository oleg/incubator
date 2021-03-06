(ns maze
  (:require [clojure.string :as str]))

;; (def maze [[{} {} {} {}]
;;            [{} {:open #{:east}} {:open #{:west}} {}]
;;            [{} {} {:open #{:south}} {}]
;;            [{} {} {:open #{:north}} {}]])


(defn empty-maze [rows columns]
  {:rows rows
   :columns columns
   :cells (vec (map (fn [x1] (vec (map (fn [y1] {:x x1 :y y1})
                                       (range 0 columns))))
                    (range rows)))})


(defn generate [rows columns]
  (let [m (empty-maze rows columns)]
    m))


(defn render-cell-top [c]
  (str (if (:north (:open c)) "   +" "---+")))

(defn render-cell-middle [c]
  (str (if (:east (:open c)) "    " "   |")))

(defn render-cell-bottom [c]
  (str (if (:south (:open c)) "   +" "---+")))

(defn render-east-cell-top [c]
  "+")

(defn render-east-cell-middle [c]
  (str (if (:west (:open c)) " " "|")))

(defn render-east-cell-bottom [c]
  "+")

(defn make-row-renderer [f1 f2]
  (fn [[f :as all]]
    (str/join (apply str (f1 f) (map f2 all)))))

(defn render-row [r]
  (let [render-row-middle (make-row-renderer render-east-cell-middle render-cell-middle)
        render-row-bottom (make-row-renderer render-east-cell-bottom render-cell-bottom)]
    (str/join "\n" [(render-row-middle r)
                    (render-row-bottom r)])))

(defn render [[f :as all]]
  (let [render-row-top (make-row-renderer render-east-cell-top render-cell-top)]
    (str/join "\n" (cons (render-row-top f)
                         (map render-row all)))))

(defn run [opts]
  (let [{rows :rows columns :columns} opts]
    (println (render (:cells (generate rows columns))))))


(defn cell-at [m r c]
  ((m r) c))

