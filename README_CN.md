[English](README.md) | 中文 | [한국어](README_KR.md)

# 本体主网上线一周年技术解密活动: 本体克莱因瓶竞赛

公元1882年，数学家菲利克斯·克莱因提出了一种自我封闭且没有明显边界的模型“克莱因瓶”。如果把克莱因瓶比做一个瓶子，那它则是一个永远无法装满的瓶子。瓶子的里面也是它的外面，没有里外之分，它的表面不会终结。

在本体主网上线一周年之际，我们致敬每一位区块链技术开发者，献出一道可有无限解决方案的技术题目--“克莱因瓶竞赛 ”，一周内寻求最优解决代码，并附赠数万ONG奖励，鼓励开发者探索区块链技术的极限。

正如在克莱因瓶里，瓶子无法装满，但在有限的空间内，能探索到极限的奥义，亦是一种伟大。

## 挑战: 正则表达式

正则表达式常用于程序设计，通常被用来检索、替换那些符合某个模式的字符串。当前大多程序设计语言，为了方便程序开发，都提供了正则表达式的标准库。但是本体平台的智能合约目前还没有较好的正则表达式的标准库。

本挑战为完成一个简单规则的正则表达式匹配的智能合约。需要实现的匹配规则节选自POSIX Basic Regular Expression，如下：

```
c    matches any literal character c
.    matches any single character
[]   matches any one character listed between the brackets
[^]  matches any one character not listed between the brackets
?    matches any character one time if it exists
*    matches zero or more occurrences of the previous character
+    matches declared element one or more times
^    matches the beginning of the input string
$    matches the end of the input string
```

比如模式 <code>x.y</code> 能匹配 <code>xay</code> 和 <code>x2y</code>等，但不能匹配 <code>xy</code> 或 <code>xaby</code>。<code>^.$</code>能够与任何单个字符的字符串匹配， 而 <code>^.*$</code>能够与任意字符串匹配。

#### 关于特殊字符
此简单规则的正则表达式需要支持在［］中的特殊字符。比如<code> [\\]] </code>将可以匹配 <code> abc\] </code>。
在此次比赛中，只需对\\后的字符做字面匹配，而不需要支持非字面意思的匹配，比如 \b 不需要作为词边界。

智能合约模版：

```
from ontology.interop.System.Runtime import Notify

def Main(operation, args):

    if operation == "match":
        pattern = args[0]
        text = args[1]
        return match(pattern, text)

    return False

def match(pattern, text):
    result = True      # Your Implementation Here
    Notify(["match", pattern, text, result])
    return result
```

示例测试用例：

```
"abc" ~ "abc"
"a.c" ~ "abc"
"a[123]c" ~ "a1c"
"a[^abc]c ~ "a1c"
"a?c" ~ "c"
"ab*c" ~ "abbbbc"
"a+c" ~ "ac"
"^a" ~ "abc"
"^.*c" ~ "abbbbc"
"^$" ~ ""
```

### 评测结果标准：

1. 通过所有测试用例。我们准备了200个测试用例，所有字符皆为ASCII字符，最长模式长度为16，最长字符串长度为40；
2. 完成所有测试的所需要的Ontology NeoVM指令数总和最少者获胜。


## “克莱因瓶竞赛”将分为2轮进行

### 第一轮:解密

奖励：
* 基础奖池 50,000 ONG
* 开发者助力奖池 (最多50,000 ONG) 

#### 什么是“开发者助力奖池”？
“开发者助力奖池”是针对本次活动开发者参与度的额外激励。“克莱因瓶竞赛”活动每多1位有效开发者报名并助力，开发者助力奖池将增长100 ONG，开发者助力奖池最高50,000 ONG。

#### 如何报名？

获取经GitHub认证的ONT ID
1. 下载[ONTO](https://onto.app/)或者从应用商店中取得并安装；
2. 进入点击“ONT ID”并注册；
3. 完成Github身份认证。

请使用表单[Registration Form](http://bit.ly/2Jf02AE)提交以下信息
1. 经GitHub认证的ONT ID；正则表达式常用于程序设计；
2. GitHub账号；
3. 联络邮箱。

#### 如何助力？
开发者在提交完以上信息后，使用对应的GitHub账号在Github上对[Ontology主网代码](https://github.com/ontio/ontology)进行Fork和Star，即可积累有效“助力”1次，对应“开发者助力奖池”将增加100ONG奖励。该奖池数字每24小时更新一次，请以Ontology官方[Twitter](https://twitter.com/OntologyNetwork)公告为准。

#### 如何提交参赛代码？
开发者完成开发、测试后，请务必调用存证合约将开发者的ONT ID，合约hash（SHA-256）进行存证。注意：合约调用需支付0.01 ONG。存证教程可以在[这里](register_tool/README.md)找到。

完成以上后，请开发者将合约源码通过 [Submission Form](http://bit.ly/2XuH0Qb)提交。需要提交的信息包括：

* ONT ID及联络邮箱
* 存证合约调用的交易Hash
* 合约代码

#### 比赛规则

1. 参赛时间: 新加坡时间2019年6月30日18:00-7月4日18:00；
2. 最终取参赛代码中“功能完备”、 “完成所有测试用例且总步数最短”者获胜（详细要求请见题目）；
3. 如有多人提交最优答案，最早提交者获胜。提交时间以上链存证时间为准；
4. 如一人提交多次参赛代码，将以最后一次提交内容为准。

#### 公布获奖
此轮获奖者将在新加坡时间7月5日内宣布，具体时间请关注Ontology官方[Facebook](https://www.facebook.com/ONTnetwork/)和[Twitter](https://twitter.com/OntologyNetwork)。


#### 获奖者能获得什么？

如确认领奖：直接领取总奖池奖金（基础奖池+开发者助力奖池）的60%，“克莱因瓶竞赛”也就此结束。

获奖者需要在名单公布后24小时内决定是否领奖，回复获奖通知邮件。如不回复，将自动领取此轮奖项。
	
## 还不过瘾？

获奖者可以接受打擂，进入下一轮，所有参赛者将有机会竞争100%奖池奖励。

### 第二轮：打擂

**时间**：Ontology官方Twitter公布上轮获奖者接受打擂时间后的24小时内。

此轮提交参赛代码请以GitHub Pull Request方式提交至[ontio/Anniversary](https://github.com/ontio/Anniversary)。

此轮开始时将公布上轮成绩排名前3的参赛代码。

**比赛规则**:

1. 此轮竞赛中所有开发者可以继续报名并提交自己的方案，包含上轮获奖者；
2. 在此轮期间的最优结果如能优于上一轮解密轮结果，提交者将获得总奖池100%的奖励。上一轮获奖者60%奖励将不予发放。

## 请注意
* 可以从[这里](https://dev-docs.ont.io/#/docs-en/smartcontract/01-started)获取更多本体智能合约开发和测试的相关文档；
* 请使用Ontology TestNet进行测试，可以从[这里](https://developer.ont.io/applyOng)获取测试网的测试币；
* 智能合约开发利器[SmartX](https://smartx.ont.io/)，可以在[这里](https://dev-docs.ont.io/#/docs-en/SmartX/00-overview)找到相关教程；
* 如有问题，请加入[Ontology Discord](https://discord.gg/4TQujHj)。
