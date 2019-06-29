English | [中文](README_CN.md)  | [한국어](README_KR.md)

# The First Anniversary of Ontology MainNet Launch: Ontology Klein Bottle Contest

The Klein bottle was first described in 1882 by the German mathematician Felix Klein. It is an example of a non-orientable surface with no clear boundary. If likened to a bottle, it will be a bottle which you can never fill up. It is a two-dimensional manifold against which a system for determining a normal vector cannot be consistently defined. Its surface will never end.

On the occasion of the first anniversary of the Ontology MainNet launch, we would like to pay our respect to every blockchain developer by presenting “the Ontology Klein Bottle Contest”, a technical challenge with unlimited solutions. To encourage developers push back the frontiers of blockchain technology, we will choose the best solution in a week’s time and give away tens of thousands of ONG as rewards.

Just as you can never fill up a Klein Bottle, it is nevertheless a great achievement to explore the limits of the profound meaning within the limited space.

## Challenge One: Regular Expression

A regular expression is a sequence of characters that define a search pattern. Usually such patterns are used by string searching algorithms for "find" or "find and replace" operations on strings, or for input validation. Most programming languages today offer standard library of regular expression. However, smart contracts on the Ontology platform do not have a well-developed regular expression standard library yet.

The challenge is to develop a simple rule regular expression for smart contracts on the Ontology platform. The matching rules are excerpted from POSIX Basic Regular Expression as follows.

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
ASCII
For example, pattern <code>x.y</code> can match <code>xay</code> and <code>x2y</code>, but not <code>xy</code> or <code>xaby</code>. <code>^.$</code> can match any single-character string, while <code>^.*$</code> can match any string.

Smart contract template:

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

Test case examples:

```
"abc" ~ "abc"
"a.c" ~ "abc"
"a[123]c" ~ "a1c"
"a[^abc]c ~ "a1c"
"a?c" ~ "c"
"a*c" ~ "abbbbc"
"a+c" ~ "ac"
"^a" ~ "abc"
"^.*c" ~ "abbbbc"
"^$" ~ ""
```

### Judging Criteria:

1. Participants need to pass all the 200 test cases we have prepared. All characters are ASCII characters, the longest pattern has 16 characters and the longest string has 40 characters.
2. Participants who used the fewest Ontology NeoVM instructions to complete all the tests will be the winner.


## The Ontology Klein Bottle Contest will be divided into 2 rounds

### First Round

Reward:

* Basic Reward Pool 50,000 ONG
* Developer Boost Reward Pool (up to 50,000 ONG)

#### What is “Developer Boost Reward Pool”?
The Developer Boost Reward Pool is the extra incentives to encourage developer participation. 100 ONG will be added to the reward pool every time a new developer joins the Contest.

#### How to sign up?

Get GitHub-verified ONT ID

1. Download [ONTO](https://onto.app/) or install from appstore
2. Click “ONT ID” and sign up
3. Complete GitHub ID verification

Submit the following information via [Registration Form](http://bit.ly/2Jf02AE):

1. GitHub-verified ONT ID
2. GitHub account
3. Your email address

#### How to boost?
After submitting the above information, every time developers use their GitHub account to Fork and Star the Ontology MainNet code, the boost value will be added once, which means 100 ONG will be added to the Developer Boost Reward Pool. The amount of the Reward Pool will be updated every 24 hours, please refer to the [Ontology official Twitter account](https://twitter.com/OntologyNetwork) for the latest amount.
You can find the Ontology MainNet code [here](https://github.com/ontio/ontology).

#### How to submit code?
After completing development and testing, developers must call the registry contract to record their ONT ID and contract hash (SHA-256) on the blockchain. You'll have to pay 0.01 ONG for calling this contract as transaction fee. The tutorial is [here](register_tool/README.md).

After completing the above steps, developers need to submit their contract source code via [submission form](http://bit.ly/2XuH0Qb). The information you need to submit includes:

* Your ONT ID and email address
* Transaction hash called by the registry contract
* Contract code

#### Rules

1. Code submission time: 10:00 June 30th - 10:00 July 4th (UTC)
2. Participants who used the fewest instructions to pass all the tests will be the winner (for more details please refer to the challenge rules)
3. If there is more than one optimal solution, then participants who submitted first will be the winner. The submission time is determined by the time they called the registry contract
4. If developers submit their codes more than once, we will judge the result based on the codes in the final submission

#### Winner Announcement
The winner list will be announced on July 5th (UTC). 

For further details, please refer to Ontology’s official [Facebook](https://www.facebook.com/ONTnetwork/) or [Twitter](https://twitter.com/OntologyNetwork).


#### What are the rewards for winners?

If the winner confirms to claim the rewards for the first round, then they can get 60% of the entire reward pool (Basic Reward Pool + Developer Boost Reward Pool). After that, the winner’s journey in Ontology Klein Bottle Contest ends.

The winner needs to decide whether to claim the reward within 24 hours after the announcement by replying to our email, otherwise, the reward will be automatically given to the winner.

## Want more rewards?

All winners in the first round can accept challenges in the next round where all candidates have the chance to win the entire reward pool.

### Second round – Taking up the challenge

**Time**: Within 24hrs of the Twitter announcements of winners accepting challenges.

In this round, participants need to submit their codes to [ontio/Anniversary](https://github.com/ontio/Anniversary) via GitHub Pull Request.
The list of developers who ranked in the top 3 in the previous round will be announced at the beginning of this round.

**Rules**:

1. All developers can continue to sign up and submit their solutions, including the first-round winners.
2. If the optimal results in this round are better than the last round, then the developers will win the entire reward pool, instead of the 60% rewards in the last round.

If you have any questions, please ask us on [Ontology Discord](https://discord.gg/4TQujHj).
