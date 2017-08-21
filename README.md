# ParseNmap

## Purpose
用來解析nmap掃描的結果xml檔<br>
可過濾出os為windows的主機ip, 並產出CSV檔

## How to Build?
`cd <gopath>/src/ParseNmap`<br>
`go build`

## How to use it?
`parseNmap <input_file>.xml`

## What is the Output?
`<input_file>.csv`
