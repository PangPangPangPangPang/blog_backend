# golang的包管理入门

> 记录一下golang的官方包管理工具[dep](https://github.com/golang/dep)的使用。

[date] 2019-08-16 14:32:00
[tag] Golang 技术

## 安装
虽然提供了各个平台包管理工具安装方式，但是还是推荐*go get*一把梭

```sh
    go get -u github.com/golang/dep/cmd/dep
```

## 初始化
不管是新仓库，还是以前没有用过包管理工具的仓库，配置好路径之后，直接跑下面的命令就可以进行初始化

```sh
    dep init
```

稍等之后，上面的命令会生成*Gopkg.toml*,*Gopkg.lock*,*vender/*三个文件/文件夹，这就是包管理的所有内容了。

## Tips

* *dep check* will quickly report any ways in which your project is out of sync.
* *dep ensure -update* is the preferred way to update dependencies, though it's less effective for projects that don't publish semver releases.
* *dep ensure -add* is usually the easiest way to introduce new dependencies, though you can also just add new import statements then run dep ensure.
* If you ever make a manual change in *Gopkg.toml*, it's best to run *dep ensure* to make sure everything's in sync.
* *dep ensure* is almost never the wrong thing to run; if you're not sure what's going on, running it will bring you back to safety ("the nearest lilypad"), or fail informatively.

## Other

[Gopkg.toml](https://golang.github.io/dep/docs/Gopkg.toml.html)

[Gopkg.lock](https://golang.github.io/dep/docs/Gopkg.lock.html)
