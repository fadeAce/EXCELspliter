# EXCELspliter
## **EN documentation** | [中文文档](./doc/中文文档.md)

EXCELspliter is a tool to make a split Excel documents easily. If you are interested in source code, you can build one and modify it for features add-on you want it to be, then read build from the source.

Typical scenario: you have a large Excel file with 1 million lines and you want it to split them into multiple files with 250 lines per file while retain the header.

EXCELspliter can do the job.

## Build from source
### build execution file:
```apple js
cd EXCELspliter
go get
go build
```
### run from source:
```apple js
ls
LICENSE  README.md  spliter.exe*  spliter.go  writer_test.go
```
Do not double-click/open the EXCELspliter.exe file directly. You should open Terminal or Command Prompt `(Win+X).` then `cd` or change directory to EXCELspliter and run the command from there.

Example:

```apple js
EXCELspliter.exe -p=./lzy.xlsx -head -l=10 -s=1
```

### Parameters:

_-head :_ (optional) Whether to keep first line header for this file when splitting.

_-p :_ path for excel file.

_-l :_ total rows you want to keep per excel file.

_-s :_ excel sheet index.

_-t :_ where the output files to be saved.

#### for more details with CLI params:

_-h :_ for help with command line params

#### get executable file:
[release](https://github.com/fadeAce/EXCELspliter/releases)

#### advanced usage & config file:
For advanced usage and versatile functions , you may get it from `-h` console print helper, which shows a series functions that should be configured in `yaml` file.