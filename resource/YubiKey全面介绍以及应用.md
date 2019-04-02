# YubiKey全面介绍以及应用
[date] 2019-04-02 21:15:00
[tag] Yubikey 折腾

> 写在前面，本文基于[Yubikey](https://wiki.archlinux.org/index.php/YubiKey)润色整理而来。争取简单易懂的介绍一下Yubikey是个什么东东，以及如何正确使用，再以及如何有效~~吃灰~~利用。

## Yubikey是个什么鬼东东
Yubikey是一个**USB Security token**。它有以下能力：
1. 生成一次性密钥（OTP），基于**AES**或者**OATH-HOTP**
2. 生成一个63个字符长度的静态密码
3. Challenge/Response认证
4. **U2F**认证，也就是有些产品的二次验证，比如github、google的帐号登录都支持**U2F**
5. 存储**RSA**证书，从而可以实现ssh、邮件的签名/加密、git commit等能力

### 理解Yubikey
Yubikey利用一个小LED灯与一个按钮与用户进行交互。
Yubikey可以模拟一个usb键盘当你点击按钮的时候处处一个**OPT**或者静态密码。

### Yubikey的输入
* 利用USB接口的调用API
* 短按Yubikey按钮
* 长按Yubikey按钮

### Yubikey的输出
* 模拟键盘输出
  * 静态密码
  * **OTP**
* USB的API回调
  * **Challenge-Response**请求的response
  * **U2F**请求的response
  * **CCID**智能卡的请求的response
