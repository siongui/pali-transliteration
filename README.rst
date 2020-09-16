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
     - (空子音, 見 [註2]_ )

子母音組合
==========

- 泰文如天城體、緬文、柬埔寨文、寮文為元音附標文字 (Abugida)，是以子音為主體，而母音是加在子音上、下、左、右。
- 以下先列出巴利文母音的部分：

.. list-table:: 母音對照
   :header-rows: 1

   * - 泰文
     - 羅馬拼寫
   * - (見 [註1]_ )
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
