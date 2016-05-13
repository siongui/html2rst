=================================
HTML_ to reStructuredText_ in Go_
=================================

Development Environment: `Ubuntu 16.04 LTS`_ and `Go 1.6.2`_.

Goal: Convert HTML_ to reStructuredText_ in Golang_. (not fully implemented)


Install
+++++++

.. code-block:: go

  $ go get -u github.com/siongui/html2rst


Usage
+++++

See `example <usage/example.go>`_:

.. code-block:: go

  package main

  import (
  	"github.com/siongui/html2rst"
  	"net/http"
  )

  func main() {
  	//resp, err := http.Get("http://nanda.online-dhamma.net/")
  	//resp, err := http.Get("https://siongui.github.io/zh/2016/03/14/pillow-useful-items-for-me-notes/")
  	resp, err := http.Get("http://myweb.ncku.edu.tw/~lsn46/Tipitaka/Sutta/Khuddaka/Khuddaka-patha/Khuddaka-patha.html")
  	if err != nil {
  		panic(err)
  	}
  	defer resp.Body.Close()

  	print(html2rst.HtmlToRst(resp.Body))
  }


Deploy on `Google App Engine Go`_
+++++++++++++++++++++++++++++++++

See `gae <gae/>`_ directory.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `golang.org/x/net/html - GoDoc <https://godoc.org/golang.org/x/net/html>`_

.. [2] `[Golang] HTML to reStructuredText <https://siongui.github.io/2016/05/12/go-html-to-rst/>`_


.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Ubuntu 16.04 LTS: http://releases.ubuntu.com/16.04/
.. _Go 1.6.2: https://golang.org/dl/
.. _reStructuredText: http://docutils.sourceforge.net/rst.html
.. _HTML: https://www.google.com/search?q=HTML
.. _Google App Engine Go: https://cloud.google.com/appengine/docs/go/
.. _UNLICENSE: http://unlicense.org/
