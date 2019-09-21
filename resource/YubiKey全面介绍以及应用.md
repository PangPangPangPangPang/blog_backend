# YubiKey 简单介绍以及应用

[date] 2019-04-02 21:15:00
[tag] Yubikey 折腾

> 写在前面，本文基于[Yubikey](https://wiki.archlinux.org/index.php/YubiKey)润色整理而来。争取简单易懂的介绍一下 Yubikey 是个什么东东，以及如何正确使用，再以及如何有效~~吃灰~~利用。

## Yubikey 是个什么鬼东东

![img4](http://ww1.sinaimg.cn/large/6ccb17ably1g1rronz4jsj20u00wemzl.jpg)
Yubikey 是一个**USB Security token**。它有以下能力：

1. 生成一次性密钥（OTP），基于**AES**或者**OATH-HOTP**
2. 生成一个 63 个字符长度的静态密码
3. Challenge/Response 认证
4. **U2F**认证，也就是有些产品的二次验证，比如 github、google 的帐号登录都支持**U2F**（注意有些分享里面也把 U2F 称为 FIDO，两者等价，改名了而已）
5. 智能卡模式（CCID）， 通过存储签名/加密/认证证书，从而可以实现 ssh、邮件的签名/加密、git commit 等能力

### 理解 Yubikey

Yubikey 利用一个小 LED 灯与一个按钮与用户进行交互。
Yubikey 可以模拟一个 usb 键盘当你点击按钮的时候输出一个**OPT**或者静态密码。
当然了，Yubikey 可以在你需要的时候切换成智能卡模式。

### Yubikey 的输入

- 利用 USB 接口的调用 API
- 短按 Yubikey 按钮
- 长按 Yubikey 按钮

### Yubikey 的输出

- 模拟键盘输出
  - 静态密码
  - **OTP**
- USB 的 API 回调
  - **Challenge-Response**请求的 response
  - **U2F**请求的 response
  - **CCID**智能卡的请求的 response

### LED 灯的状态

- 常亮：请求按压（通常在 U2F 模式请求二次验证通过时出现）
- 低频闪烁：激活/设置中/已可用
- 高频闪烁：出现错误/设置驱动中

## Yubikey 的应用

这部分应该是大家比较关心的部分了。总结一下我配置的一些内容

### OTP

这部分比较简单了，属于 Yubikey 开箱即用的东西。这部分内容建议安装[yubikey-manager](https://www.archlinux.org/packages/?name=yubikey-manager)以及[yubikey-manager-qt](https://www.archlinux.org/packages/?name=yubikey-manager-qt)直接利用 GUI 来进行配置。
总结起来就是 OTP 模式会提供两个 slot，并且每个 slot 可以提供 4 种能力支持。我的话在 slot1 配置的是静态密码，slot2 配置的是**yubico OTP**。
![pic1](http://ww1.sinaimg.cn/large/6ccb17ably1g1rpxzmxb1j20pt0c9dgo.jpg)
![pic2](http://ww1.sinaimg.cn/large/6ccb17ably1g1rpxixik6j20pn0byjsc.jpg)

### U2F

这部分也是属于开箱即用的功能，但是需要各位到支持的网站上自行配置。我也只是配置了**google**跟**github**的**U2F**，不过对于谷歌全家桶重度用户来说已经解决掉了绝大多数我的数据的安全问题。
至于你想知道你用到的服务是否支持**U2F**可以在[网站](https://twofactorauth.org/)查询。

### Smart Card

这一部分支持的能力有**PIV**/**OpenPGP**/**OATH**， 配置起来比较麻烦的部分。我介绍几个比较常用的。

#### OpenPGP

具体的定义大家可以在[wiki](https://zh.wikipedia.org/wiki/PGP)上具体查阅。
粗鲁的总结一下的话，就是一款很屌的用于加密/签名的开源规范。**GPG**是遵循**OpenPGP**规范的实现（名字很绕，谨记于心）。

那么 Yubikey 在这里面担任了一个什么角色呢？对了，就是密钥载体。
Yubikey 可以将**GPG**生成的密钥导入到自己体内（大误）。这样就可以随身携带啦。当有文本需要加密的时候，插入 Yubikey，就可以里利用体内的私钥进行加密。这个安全性以及便利性，想想就让人开心。

至于具体的配置，我帖一下 Yubico 的[官方文档](https://support.yubico.com/support/solutions/articles/15000006420-using-your-yubikey-with-openpgp_)，介绍的很详细，但是一定要有耐心的一步一步照着执行。（什么？你说的就像我看不懂文档一样？不是的！是**GPG**这个工具真的很难用！！不小心真的容易搞错！！！哦，记得，私钥一定记得要离线备份哦！！）

照着文档操作完之后，你就已经把**PGP**的几把密钥导入到 Yubikey 中了。接下来就是如何使用了。

如果说你的应用程序支持**PGP**（比如某某 mail）的话，在 GUI 界面应该直接就有选项可以加密，插上 Yubikey 进行加密就好了。
如果说你想给一些文本进行加密的话，可以运行如下命令行：

```sh
gpg --sign -a
# 输入完成后，输入回车，输入<C-d>即可得到加密后的内容
```

效果如图：
![img3](http://ww1.sinaimg.cn/large/6ccb17ably1g1rqxzvkrvj20k90c8wjb.jpg)

其他的**GPG**的使用技巧，大家可以翻一翻 wiki，或者在终端运行*gpg --help*来具体看一下。

#### PIV

> Starting from the fourth generation devices, the Yubikeys contain a PIV (Personal Identity Verification) application on the chip. PIV is a US government standard (FIPS 201) that specifies how a token using RSA or ECC (Elliptic Curve Cryptography) is used for personal electronic identification. The distinguishing characteristic of a PIV token is that it is built to protect private keys and operate on them on-chip. A private key never leaves the token after it has
> been installed on it. Optionally, the private key can even be generated on-chip with the aid of an on-chip random number generator. If generated on-chip, the private key is never handled outside of the chip, and there is no way to recover it from the token. When using the PIV mechanism, the Yubikey functions as a CCID device.

贴下 arch 文档上的介绍，主要是用于个人认证。
**PIV**可以方便的在之前提过的[yubikey-manager-qt](https://www.archlinux.org/packages/?name=yubikey-manager-qt)中进行配置。
下面说一下如何配置**PIV**进行 SSH 登录

1. 插入 Yubikey 并验证是否成功：

```sh
$ ykman list
YubiKey 4 [OTP+FIDO+CCID] Serial: 1234567
```

2. 生成密钥対，并导出公钥：

```sh
$ ykman piv generate-key -a RSA2048 9a pubkey.pem
```

3. 为该公钥生成自签名证书：

```sh
# 注意生成的证书是存储在Yubikey中的
$ ykman piv generate-certificate -s "SSH Key" 9a pubkey.pem
```

4. 安装 opensc 并修改~/.ssh/config 文件：

```sh
pacman -S opensc
```

在~/.ssh/config 文件中添加：

```sh
PKCS11Provider /usr/lib/opensc-pkcs11.so
```

5. 将之前导出的公钥转换成 SSH 协议可以识别的格式：

```sh
$ ssh-keygen -i -m PKCS8 -f pubkey.pem > pubkey.txt
```

6. 将生成的*pubkey.txt*配置到 server 端。

经过这一系列操作之后，在正常的通过 SSH 登录到 vps 的时候，就会弹出让你输入 Pin 的提示，当然这个 Pin 就是 Yubikey 的 Pin 啦，默认是 123456，也可以在[yubikey-manager-qt](https://www.archlinux.org/packages/?name=yubikey-manager-qt)中进行修改。

##### Mac 配置

1. 安装[opensc](https://github.com/OpenSC/OpenSC/wiki/Download-latest-OpenSC-stable-release)
2. 生成密钥对->导入 yubikey->配 vps(跟在 linux 的流程大致是一样的)
3. 登录

```sh
# '-I' 代表使用pkcs11
 ssh -I /usr/lib64/opensc-pkcs11.so <username>@<remote-host>
```

## 最后

以上就是我对 Yubikey 的初步认识，以后如果有更多的心得，会随时更新到本篇文章里的，感谢阅读！
