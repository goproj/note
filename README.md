note
====

just a todo note

## install

```
mkdir $GOPATH/src/github.com/goproj
cd $GOPATH/src/github.com/goproj
git clone git@github.com:goproj/note.git
go get -v github.com/goproj/note/...
cd note
go build
./note
```

## 开发约定

后端接口通常都是返回一个json字符串，有三个字段：

- err 错误码，大于0就是有错误，0表示成功
- msg 错误消息，如果err>0，msg就会有值，否则就不会有msg字段
- data 具体要返回的核心内容，如果err>0，就不会有data字段

所以前端可能要先判断err字段，再做具体处理

## 备忘

- prepare方法中不能每次都读DB，需要加cache，先跑通基本流程，之后统一加cache