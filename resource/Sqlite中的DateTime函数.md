# SQLite 中的 Date/Time 函数

> 水一篇

[date] 2019-09-05 16:47:33
[tag] 技术 SQLite

### 1.总览

SQLite 总共有如下 5 个函数:

1. date(timestring, modifier, modifier, ...) YYYY-MM-DD
2. time(timestring, modifier, modifier, ...) HH:MM:SS
3. datetime(timestring, modifier, modifier, ...) YYYY-MM-DD HH:MM:SS
4. julianday(timestring, modifier, modifier, ...) Julian day
5. strftime(format, timestring, modifier, modifier, ...)

其中前四个函数返回固定的 format，最后一个**strftime**可以指定 format(这个 format 遵循 standard C library， 并且额外增加两个类型如下)

- %J [Julian day number](http://en.wikipedia.org/wiki/Julian_day)
- %f fractional seconds: SS.SSS

### 2.timestring

timestring 作为第一个参数，需要满足如下的格式:

1. YYYY-MM-DD
2. YYYY-MM-DD HH:MM
3. YYYY-MM-DD HH:MM:SS
4. YYYY-MM-DD HH:MM:SS.SSS
5. YYYY-MM-DD**T**HH:MM 其中**T**用来拆分 DD/HH 的字符
6. YYYY-MM-DD**T**HH:MM:SS
7. YYYY-MM-DD**T**HH:MM:SS.SSS
8. HH:MM 没指定日期则默认**2000-01-01**
9. HH:MM:SS
10. HH:MM:SS.SSS
11. now 当前时间
12. DDDDDDDDDD Julian day 或者 Unix Time

#### 时区

SQLite 内部使用**UTC**作为标准，或者称作"Zulu"。以下都是等价用法:

- 2013-10-07 08:23:19.120
- 2013-10-07T08:23:19.120Z
- 2013-10-07 04:23:19.120-04:00

### 3.modifiers

追加在 timestring 后的参数可以进一步更改时间

1. NNN days "+1 day(s)" "-5 day(s)"
2. NNN hours
3. NNN minutes
4. NNN.NNNN seconds
5. NNN months 注意这个增减月份会转换成当前月份的天数来进行增加，所以尽量少用这个...
6. NNN years
7. start of month
8. start of year
9. start of day
10. weekday N
11. unixepoch 这个参数只在 timestring 为时间戳的情况下才能生效，而且必须追加
12. localtime 如果 timestring 是 utc 时间，则 localtime 参数会将 utc 时间转换成本地时间(不是很建议用这个，这个首先是依赖本机的 time zone，而且不能根据请求来源来做国际化)
13. utc 如果 timestring 是本地时间，则 utc 参数会将本地时间转换成 utc 时间(听起来有点绕，但是如果你的数据库里面存的已经是 utc 的时间戳，那么就不需要再用这个参数了)

### 示例

```sql
SELECT datetime(1092941466, 'unixepoch'); -- format时间戳，重要！
SELECT datetime(1092941466, 'unixepoch', 'localtime'); -- 获取本地时区的时间
SELECT strftime('%s','now'); -- 获取当前时间戳
SELECT date('now','start of month','+1 month','-1 day'); -- 获取本月的最后一天
```

### 需要注意的一些坑

1. 不要依赖数据库来处理时区的问题
2. 月份的计算是按照当前月的天数来计算的
3. 以后碰到在追加...
