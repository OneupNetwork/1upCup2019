## 1up Cup 2019
![](https://img.shields.io/badge/date-2019--07--26-0385B1.svg)

### Problems

#### [Chocobo - Tanya and Stairways](https://codeforces.com/problemset/problem/1005/A)
宣告一個陣列紀錄每層樓梯數 steps，當坦雅喊 1 且不是第一層時，把上層的樓梯數丟進 steps，最後輸出 steps 長度跟 steps 內容

---

#### [Ifrit - Distances to Zero](https://codeforces.com/problemset/problem/803/B)
測資有極限值，輸入加一下 buffer size，且要一行一行輸入再轉數字陣列避免輸入花太多時間，最接近的 0 可能在後面，所以先用 pos 陣列記錄所有 0 的位置，再跑迴圈看元素距離左邊的 0 比較近還右邊的 0 比較近，這題只需要知道某元素跟元素值為 0 的元素的距離（array index 差）跟元素值沒關係

---

#### [Moogle - Stages](https://codeforces.com/problemset/problem/1011/A)
取最輕，所以先排序，從小的開始取，跟前一個是相鄰字母就跳過，取到 n 個為止，取不到 n 個代表無法建造

---

#### [Odin - Almost Arithmetic Progression](https://codeforces.com/problemset/problem/978/D)
等差數列，需要用到 ![](https://wikimedia.org/api/rest_v1/media/math/render/svg/10838b5a34501cb20961b2646292f88ba2daa82f) 這公式，輸入也有極限值，加大 buffer size，字串轉數字陣列避免 TLE，用上述公式取陣列頭尾為 n m 可以求出各種變化（9種）的差，差有小數點就不可能，剩下可能的跑迴圈，一樣用上述公式驗算每個元素是否在允許變化範圍內，順便紀錄變化數，最後輸出最小的變化數

---

#### [Shiva - Obtaining the String](https://codeforces.com/problemset/problem/1015/B)
先排序 s t，排序後不同就不可能，接著從 s 跟 t 第一個差異的字母開始換，順便記錄過程，最後輸出結果

---

#### [Titan - Three Parts of the Array](https://codeforces.com/problemset/problem/1006/C)
測資一樣有極限值，輸入與上述處理方式相同，題目只要求輸出最大 sum1，又 sum1 == sum3，所以從頭尾開始加即可，這樣的時間複雜度會小於 O(n) 不然正常方式一定 TLE，另外需注意加起來的值會大於 32 位元整數最大值