[English](README.md) | [中文](README_CN.md) | 한국어

# 온톨로지 메인넷 출시 1주년 기술 비밀 풀이 이벤트: Ontology Klein Bottle Contest

서기 1882년 수학자 펠릭스 클라인은 바깥쪽과 안쪽을 구별할 수 없는 단측곡면의 모형“Klein Bottle “을 제안하였습니다. “Klein Bottle “을 병에 비한다면, 이 병은 영원히 채워질 수 없는 병입니다.

온톨로지 메인넷 출시 일주년을 맞이하며, 모는 블록체인 기술 개발자들에게 감사의 인사를 전달하고 싶습니다. 또한 우리는 더 많은 개발자들의 블록체인 기술 탐색을 지원하기 위해--“Ontology Klein Bottle Contest”솔루션 도전을 개최하며, 수만개의 ONG바운티를 배포할 예정입니다.

Klein Bottle과 같이 채워질 수 없지만, 한정된 3차원 공간에서 무한의 가능성을 탐색할 수 있습니다.

## 도전: 정규표현식(Regular Expression)

정규표현식(Regular Expression)은 프로그래밍에서 흔히 사용됩니다. 통상 시 특별한 규칙을 가진 문자열의 집합 혹은 교환을 표현하기 위해서 만들어졌습니다.

대부분 프로그램 개발 시 편리를 위해 프로그래밍 언어(Programming Language)는 정규표현식의 표준 라이브러리를 제공합니다. 하지만 온톨로지 플랫폼의 스마트 컨트랙트는 아직 적절한 정규표현식 표준 라이브러리를 제공하지 못했습니다.

본 도전은 간단한 규칙의 정규표현식을 매칭한 스마트 컨트랙트를 작성하는 것입니다. 구현해야 할 매칭 규칙은 다음과 같은 POSIX Basic Regular Expression에서 일부 뽑은 정규표현식입니다：

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

예를 들어 모드 <code>x.y</code>는 <code>xay</code>, <code>x2y</code> 등을 매칭할 수 있지만, <code>xy</code> 또는 <code>xaby</code>는 매칭할 수 없습니다.  <code>^.$</code>는 모든 단일 문자열을 매칭할 수 있으며, <code>^.*$</code>는 모든 문자열을 매칭할 수 있습니다.

스마트 컨트랙트 템플레이트:

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

Test Case 예시:

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

### 평가 기준:

1. 모든 Test Case를 통과합니다. 우리는 200개의 Test Case를 준비하였습니다. 모두 ASCII 문자열이며, 제일 긴 패턴은 길이는 16이며, 제일 긴 문자열 길이는 40입니다.
2. 모든 테스트를 완료하는 데 총합해서 사용된Ontology NeoVM 명령 수가 제일 적은 자가 승리합니다.


## “Ontology Klein Bottle Contest”은 두 차례로 나눠서 진행됩니다

### 제1차    비밀 풀이

바운티:

* 바운티: 50000 ONG
* 개발자 리워드 (최고 50,000 ONG)

#### 개발자 리워드:
“개발자 리워드”란— “개발자 리워드”는 이번 이벤트 참여도에 따라 추가되는 바운티를 말합니다. “Ontology Klein Bottle Contest”이벤트 유효 참가자가 한 명이 늘 시 개발자 리워드 인센티브 또한 100 ONG 씩 증가합니다. 또한 리워드 최고 수량은 50000ONG입니다.

#### 신청 방법：

GitHub인증된 ONT ID받는 방법

1. [ONTO](https://onto.app/) 다운로드
2. "ONT ID" 클릭 후 가입
3. GitHub 신분 인증을 완료

[Registration Form](http://bit.ly/2Jf02AE) 에 아래와 같은 정보를 제출하세요.

1. Github에서 인정받은ONT ID
2. Github아이디
3. 이메일 주소

#### 개발자 리워드 참가 방법:

개발자는 이와 같은 정보를 제출한 뒤, GitHub에서 온토로지 메인넷 코드에 Fork와Star를 진행하세요. 유효 참가자로 인정될 시 인센티브 또한 100 ONG 가 증가합니다. 리워드 수량은 24시간마다 업데이트합니다. 온톨로지 공식계정[twitter](https://twitter.com/OntologyNetwork)을 기반으로 확인하세요.

Ontology메인넷 코드 싸이트：https://github.com/ontio/ontology

####  참가 코드 제출 방법:
개발자는 개발과 테스트를 완성 후, 개발자의 ONT ID, hash（SHA-256）컨트랙트를 꼭 증거 보관(저장)하세요. 사용 방법은[here](register_tool/README.md)에서 확인하세요.

완료 후, 개발자는 컨트랙트 소스 코드를 [Submission Form](http://bit.ly/2XuH0Qb)에 제출해 주세요. 

* ONT ID및 이메일
* 증거 보관한 컨트랙트 사용한 거래 Hash
* 컨트랙트 코드

#### 경기 규칙:

1. 참가 코드 제출 시기: 싱가포르 시간 2019년 6월 30일 18:00-7월 10일 18:00
2. 최종 참가 코드 중 "기능 완전", "모든 테스트 통과 또는 최소 시간 소비" 를 달성한 참가자 승리
3. 파일크기가 같을 시, 먼저 제출한 자가 승리합니다
4. 한 참가자가 참가 코드를 여러 번 제출할 시 마지막 제출한 내용을 기준으로 평가합니다.

#### 수상자 발표
이번 수상자는 싱가포르 시간으로 7월5일에 발표할 예정입니다. 상세 정보는 Ontology 공식 계정 [Facebook](https://www.facebook.com/ONTnetwork/)&[Twitter](https://twitter.com/OntologyNetwork)를 주목해 주세요. 

For further details, please refer to Ontology’s official [Facebook](https://www.facebook.com/ONTnetwork/) or [Twitter](https://twitter.com/OntologyNetwork).


#### 수상자 바운티

수상 받는 방법: 개발자 리워드 총 수량(기본 수량+추가 수량)의 60%，“Ontology Klein Bottle Game”종료.

수상자는 명단 발표후 24시간 안에 수상 여부를 결정할 수 있으며, 통지 메일에 회신을 하시면 됩니다. 회신이 없을 시 자동으로 바운티를 지불합니다.

## 아쉽다면?

참가자는 수상자에게 경진을 신청할 수 있으며, 참가자는 100%바운티를 받을 수 있는 기회가 주어집니다.

### 제2차   경진

**시간**: Ontology Twitter 공식 계정이 발표한 수상자가 경진 신청을 승인한 뒤 24시간 내.
이번 참가 코드는 GitHub Pull Request방식으로 [ontio/Anniversary](https://github.com/ontio/Anniversary)에 제출해 주세요.

이번 경진이 시작 시, 전 경진의 시합 TOP3 참가자 순위를 발표할 것입니다

**시합 규칙**:

1. 이번 경쟁은 수상자 또는 1차 참가자 등 모든 개발자가 신청할 수 있습니다. 
2. 여러 차 방안을 제출할 시 마지막으로 제출한 방안을 기준으로 평가할
   것입니다. 
3. 이번 라운드 결과가 전 라운드 결과보다 보다 우수할 경우, 제출자는 리워드의 바운티를 100%받아갈 수 있으며, 또한 지난 라운드 수장자는 60% 바운티를 지급받지 못합니다.

문제 있을 시, [Ontology Discord](https://discord.gg/4TQujHj)로 들어와 주세요.
