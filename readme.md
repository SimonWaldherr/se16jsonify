# se16jsonify

access SAP Database Tables and RFC-able Functions  
this is an example - howto use the [saprfc-package](https://github.com/SimonWaldherr/saprfc)  

## dependencies

* [simonwaldherr.de/go/golibs](https://simonwaldherr.de/go/golibs)
* [simonwaldherr.de/go/saprfc](https://simonwaldherr.de/go/saprfc)
* [simonwaldherr.de/go/gwv](https://simonwaldherr.de/go/gwv)

## install

1. download the package via ```go get -u -t simonwaldherr.de/go/se16jsonify```
2. edit the credentials file ```vi $GOPATH/src/simonwaldherr.de/go/se16jsonify/conn/sap_credentials.go```
3. run it ```go run $GOPATH/src/simonwaldherr.de/go/se16jsonify/se16jsonify.go```


## license

```
The MIT License (MIT)

Copyright (c) 2017 Simon Waldherr

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
