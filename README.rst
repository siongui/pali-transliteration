====================
Pāli Transliteration
====================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/pali-transliteration?status.svg
   :target: https://godoc.org/github.com/siongui/pali-transliteration

.. image:: https://travis-ci.org/siongui/pali-transliteration.svg?branch=master
    :target: https://travis-ci.org/siongui/pali-transliteration

.. image:: https://goreportcard.com/badge/github.com/siongui/pali-transliteration
   :target: https://goreportcard.com/report/github.com/siongui/pali-transliteration

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/pali-transliteration/blob/master/UNLICENSE


Convert `thai script`_/Devanagari/Sinhalese/Burmese/... to/from
romanized `Pāli`_.

Development Environment: `Ubuntu 20.04`_ and Go_.


Rules for Romanized Pāli to Thai (Traditional Chinese)
++++++++++++++++++++++++++++++++++++++++++++++++++++++

**泰文字母-羅馬字母 巴利文轉換規則**

(泰文鍵盤可以使用： `Thai Keyboard - แป้นพิมพ์ไทย`_ ，或者是在電腦裡安裝泰文輸入法)

子音對照
========

.. list-table:: 子音對照
   :header-rows: 1

   * - 泰文
     - 羅馬拼寫
   * - ก
     - k
   * - ข
     - kh
   * - ค
     - g
   * - ฆ
     - gh
   * - ง
     - ṅ
   * - จ
     - c
   * - ฉ
     - ch
   * - ช
     - j
   * - ฌ
     - jh
   * - ญ
     - ñ
   * - ฏ
     - ṭ
   * - ฐ
     - ṭh
   * - ฑ
     - ḍ
   * - ฒ
     - ḍh
   * - ณ
     - ṇ
   * - ต
     - t
   * - ถ
     - th
   * - ท
     - d
   * - ธ
     - dh
   * - น
     - n
   * - ป
     - p
   * - ผ
     - ph
   * - พ
     - b
   * - ภ
     - bh
   * - ม
     - m
   * - ย
     - y
   * - ร
     - r
   * - ล
     - l
   * - ฬ
     - ḷ
   * - ว
     - v
   * - ส
     - s
   * - ห
     - h
   * - อ
     - (空子音, 見 [02]_ )

子母音組合
==========

- 泰文如天城體、緬文、柬埔寨文、寮文為元音附標文字 (Abugida)，是以子音為主體，而母音是加在子音上、下、左、右。
- 以下先列出巴利文母音的部分：

.. list-table:: 母音對照
   :header-rows: 1

   * - 泰文
     - 羅馬拼寫
   * - (見 [01]_ )
     - a
   * -  ิ
     - i
   * -  ุ
     - u
   * - - า
     - ā
   * -  ี
     - ī
   * -  ู
     - ū
   * - เ -
     - e
   * - โ -
     - o

.. [01] 泰文拼法中，單個子音符號的出現，即預設短母音a。例如 กจ 即表示 kaca。

- 其他的母音中有出現 - 或是圈圈的部分，就是要塞上子音的，例如：

.. list-table:: 母音+子音
   :header-rows: 1

   * - 泰文
     - 羅馬拼寫
   * - ก
     - ka
   * - กิ
     - ki
   * - กุ
     - ku
   * - กา
     - kā
   * - กี
     - kī
   * - กู
     - kū
   * - เก
     - ke
   * - โก
     - ko

- 在泰文鍵盤上，除了-e, -o之外，都是先打子音，再輸入母音。只有 -e, -o 是先打母音，才打子音。

.. [02] 空子音 อ

        由於泰文的拼讀法無論如何都需要子音，像是 iti 的開頭 i- ，在羅馬拼法是沒有子音只有母音的，在泰文就必須塞進一個空子音，所以例如

        * iti 記為 อิติ
        * eva  記為 เอว
        * api 記為 อปิ （อ若單獨出現即表示無子音開頭音節的 a ）
        * āji 記為 อาชิ
        * upa 記為 อุป


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `ภาษาบาลี - วิกิพีเดีย <https://th.wikipedia.org/wiki/%E0%B8%A0%E0%B8%B2%E0%B8%A9%E0%B8%B2%E0%B8%9A%E0%B8%B2%E0%B8%A5%E0%B8%B5>`_

.. [2] `romanized pali`_

.. [3] `佛學數位圖書館暨博物館 ::: 語言教學 <http://buddhism.lib.ntu.edu.tw/lesson/>`_

.. [4] | `translit_overlay.js <https://github.com/yuttadhammo/digitalpalireader/blob/master/ThunDPR/content/js/translit_overlay.js>`_
       | `String.prototype.charAt() - JavaScript | MDN <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/charAt>`_

.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.5.3: https://golang.org/dl/
.. _Pāli: https://en.wikipedia.org/wiki/Pali
.. _romanized pali: https://www.google.com/search?q=romanized+pali
.. _thai script: https://www.google.com/search?q=thai+script
.. _Thai Keyboard - แป้นพิมพ์ไทย: https://www.branah.com/thai
.. _UNLICENSE: https://unlicense.org/
