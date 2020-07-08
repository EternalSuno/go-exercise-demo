package src

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func pwd() {
	//存储密码
	//普通方案
	//目前用的最多的密码存储方案是将明文密码做单向哈希后存储，单向哈希算法有一个特征：无法通过哈希后的摘要 (digest) 恢复原始数据，
	//这也是 “单向” 二字的来源。常用的单向哈希算法包括 SHA-256, SHA-1, MD5 等。
	//

	// import "crypto/sha256"
	h := sha256.New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))

	// import "crypto/sha1"
	hh := sha1.New()
	io.WriteString(hh, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", hh.Sum(nil))

	// import "crypto/md5"
	hhh := md5.New()
	io.WriteString(hhh, "需要加密的密码")
	fmt.Printf("%x", hhh.Sum(nil))

	//单向哈希有两个特性：
	//1）同一个密码进行单向哈希，得到的总是唯一确定的摘要。
	//2）计算速度快。随着技术进步，一秒钟能够完成数十亿次单向哈希计算。
	//结合上面两个特点，考虑到多数人所使用的密码为常见的组合，攻击者可以将所有密码的常见组合进行单向哈希，
	//得到一个摘要组合，然后与数据库中的摘要进行比对即可获得对应的密码。这个摘要组合也被称为 rainbow table。
	//因此通过单向加密之后存储的数据，和明文存储没有多大区别。因此，一旦网站的数据库泄露，所有用户的密码本身就大白于天下。
	//

	//进阶方案
	//通过上面介绍我们知道黑客可以用 rainbow table 来破解哈希后的密码，很大程度上是因为加密时使用的哈希算法是公开的。
	//如果黑客不知道加密的哈希算法是什么，那他也就无从下手了。

	//一个直接的解决办法是，自己设计一个哈希算法。然而，一个好的哈希算法是很难设计的 —— 既要避免碰撞，又不能有明显的规律，
	//做到这两点要比想象中的要困难很多。因此实际应用中更多的是利用已有的哈希算法进行多次哈希。

	//但是单纯的多次哈希，依然阻挡不住黑客。两次 MD5、三次 MD5 之类的方法，我们能想到，黑客自然也能想到。
	//特别是对于一些开源代码，这样哈希更是相当于直接把算法告诉了黑客。

	//没有攻不破的盾，但也没有折不断的矛。现在安全性比较好的网站，都会用一种叫做 “加盐” 的方式来存储密码，也就是常说的 “salt”。他们通常的做法是，先将用户输入的密码进行一次 MD5（或其它哈希算法）加密；将得到的 MD5 值前后加上一些只有管理员自己知道的随机串，再进行一次 MD5 加密。这个随机串中可以包括某些固定的串，也可以包括用户名（用来保证每个用户加密使用的密钥都不一样）。

	//
	// import "crypto/md5"
	// 假设用户名 abc，密码 123456
	h := md5.New()
	io.WriteString(h, "需要加密的密码")

	// pwmd5 等于 e10adc3949ba59abbe56e057f20f883e
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	// 指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	// salt1 + 用户名 + salt2 + MD5 拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	last := fmt.Sprintf("%x", h.Sum(nil))

	//在两个 salt 没有泄露的情况下，黑客如果拿到的是最后这个加密串，就几乎不可能推算出原始的密码是什么了。

	//专家方案
	//上面的进阶方案在几年前也许是足够安全的方案，因为攻击者没有足够的资源建立这么多的 rainbow table。
	//但是，时至今日，因为并行计算能力的提升，这种攻击已经完全可行。
	//
	//怎么解决这个问题呢？只要时间与资源允许，没有破译不了的密码，所以方案是：故意增加密码计算所需耗费的资源和时间，
	//使得任何人都不可获得足够的资源建立所需的 rainbow table。
	//
	//这类方案有一个特点，算法中都有个因子，用于指明计算密码摘要所需要的资源和时间，也就是计算强度。计算强度越大，
	//攻击者建立 rainbow table 越困难，以至于不可继续。
	//
	//这里推荐 scrypt 方案，scrypt 是由著名的 FreeBSD 黑客 Colin Percival 为他的备份服务 Tarsnap 开发的。
	//

	//
	//dk := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32) （scrypt无法使用）
	//通过上面的方法可以获取唯一的相应的密码值，这是目前为止最难破解的。

}
