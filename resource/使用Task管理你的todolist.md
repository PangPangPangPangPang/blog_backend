# 使用Taskwarrior管理你的todolist
[date] 2019-03-20 22:04:11
[tag] 工具 Todolist

> 写在前面：其实自己原来并没有一个记录**Todolist**的习惯，真是挺抱歉没有及早的养成一个好习惯。最近领导任命我为组长~~狗腿子~~有意让我去负责一些管理的事情，其实我很早就表态过不希望做这些杂七杂八的事情，尽管给我安排需求就可以了，然而领导还是让我分配了不喜欢的工作,怨念！

既然作为一个管理者，必须合理的安排自己的事情，这时候就需要一个**Todolist**来帮助你规划事项。经过一系列选型（装逼为主），最终选用了这个终端工具：**Taskwarrior**

当然了本文主要是用来安利~~备忘~~的，如果你是一个自身Geek，那么根据你的系统安装后即可关闭本篇文章。
```sh
yay -S task
man task
```
好啦，正式进入正文！

## 简介
```sh
# 基础命令构成
task <filter> <command> [ <mods> | <args> ]
```
* **task**：顾名思义，就是主命令啦！
* **filter**：顾名思义，就是filter啦！开个玩笑，这个位置可以加一些限定条件。
```sh
task project:Home list # 限定工作区
task project:Home +weekend garden list  # +weekend为限定标签， graden为模糊匹配
task project:Home due.before:today # due.before:today为今天之前
task 28 # 28为task的ID
```

* **mods**: 顾名思义（能不能少用这个词），指定task的属性。
```sh
task <filter> <command> project:Home
task <filter> <command> +weekend +garden due:tomorrow
task <filter> <command> Description/annotation text
task <filter> <command> /from/to/     <- replace first match
task <filter> <command> /from/to/g    <- replace all matches
```

* **command**：**command**就太多了，其中的**read subcommands**建议自己用**man**去看，我这里记录几个常用的命令。
```sh
# Read subcommands
task <filter> # 展示相关task
task <filter> active/all/blocked/completed/newest
task commands # 速查
task canlendar

# Write subcommands
task add <mods> # 添加task
task <filter> start <mods> # 开始task
task <filter> stop <mods> # 结束task
task <filter> annotate <mods> # 为task添加注释
task <filter> denotate <mods> # 为task删除注释
task <filter> append <mods> # 为task补充信息
task <filter> delete <mods> # 删除指定task
task <filter> done <mods> # 标注指定task完成
task <filter> edit # 用编辑器编辑指定task
task <filter> modify <mods> # 用编辑指定task
```
* **attributes**：task的属性。
```sh
+<tag> # 添加标签
project:<project-name> # 工作区名称
priority:H|M|L # 优先级
due:<due-date> # 到期日期
depends:<id1,id2> # 指定依赖的task（需要在指定task后开始）
recur:<frequency> # day/month 等指定循环频率，用于设置周期性任务
```

* **attributes modifiers**：属性的修饰符
```sh
before (synonyms under, below)
after (synonyms over, above)
none
any
is (synonym equals)
isnt (synonym not)
has (synonym contains)
hasnt
startswith (synonym left)
endswith (synonym right)
word
noword

# demo
task due.before:eom priority.not:L list
```

* **date**：特别强调一下due的描述
```sh
due:2019-03-21
due:now
due:today
due:yesterday
due:tomorrow
due:23rd
due:3wks
due:1day
due:9hrs
```

## 同步

1. 实用[inthe.am](https://inthe.am/)生成3个密钥，并离线保存好。
2. 在**~/.taskrc**文件中添加如下内容(当然了这些设置都可以在[inthe.am](https://inthe.am/configure))中找到。
```sh
taskd.certificate=/path/to/private.certificate.pem
taskd.key=/path/to/private.key.pem
taskd.ca=/path/to/ca.cert.pem
taskd.server=taskwarrior.inthe.am:53589
taskd.credentials=<your credentials>
taskd.trust=ignore hostname
```
3. 执行**task sync init**进行初始化。
4. 正常执行添加task等操作。
5. 执行**task sync**将本地修改同步到云端。（从云端将内容同步到本地也是用这个命令）

