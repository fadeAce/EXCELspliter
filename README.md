### **EN documentation** | [中文文档](./doc/中文文档.md)

EXECELspliter is a tool for excelees to make a split of document easily .
if you are interested in source code , you can build one and modify it for
features add-on you want it to be , then read build from source underlying. 

 <Br> typical scenario : you have a '.xlsx' file with 1 million lines and you want it to be split to ones limited to 250 lines , and each sheet into a single file , moreover , each one owns the same header line whitch usually been a type definition ,then , choose EXCELspliter is a simple solution.  <Br/>

#### build from source:
###### build execution file:
```apple js
cd EXCELspliter
go get
go build
```
###### run from source:
```apple js
ls
LICENSE  README.md  spliter.exe*  spliter.go  writer_test.go
```
there is an execution file spliter.exe , remember do not run it directly , you should open console (win + X) or cmd.exe
then cd to EXCELspliter and make a command like underlying

an example for you to tap'in CLI :

```apple js
spliter.exe -p=./lzy.xlsx -head -l=10 -s=1
```

#### params to control splitting:

_-head :_ (optional) Whether to keep head for this file when processing splitting.

_-p :_ path for given xlsx file to split from.

_-l :_ length for a splitting.

_-s :_ what sheet you want to split at.

_-t :_ where the output files targeting at.

#### for more details with CLI params:

_-h :_ for help with command line params

#### get executable file:
[release](https://github.com/fadeAce/EXCELspliter/releases)

#### advanced usage & config file:
for advanced usage and versatile functions , you may get it from `-h` console print helper , which shows a series functions that should be configured in `yaml` file.