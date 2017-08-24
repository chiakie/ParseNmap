# ParseNmap

## Purpose
This tool is used to parse xml result from NMAP port scanning.<br>
We parse xml to golang struct, and filter out the machine with os Windows.<br>
Finally, we will get a csv list as output.

## How to Build?
`cd <gopath>/src/ParseNmap`<br>
`go build`

## How to use it?
Windows<br>
`ParseNmap.exe <input_file>.xml`<br>
Unix<br>
`./ParseNmap <input_file>.xml`

## What is the Output?
`<input_file>.csv`
