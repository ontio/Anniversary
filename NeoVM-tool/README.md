## NVM Test Tool

NVM-tool can help you testing how much NeoVM instructions your smart contract use to complete all your tests.

#### Requirements

The requirements to build NVM-tool are:

* Golang version 1.9 or later
* Glide (a third party package management tool for Golang)

#### Compile

1. Clone the ontio/anniversary

```
$ git clone https://github.com/ontio/Anniversary.git
```

2. Fetch the dependent third party packages with Glide:

```
$ export GOPATH=`pwd`/Anniversary/NeoVM-tool
$ cd Anniversary/NeoVM-tool/src/github.com/ontio/NeoVM-tool
$ glide install
```

3. Build the source code with make:

```
$ make
```

After building the source code successfully, you should see the executable program:

* nvm-tool

#### How to use

1. copy your nvm code to the folder, named with 'test.avm.str'
2. add your testcases to 'testcases.txt'
3. test with command 

```
$ ./nvm-tool
```

##### format of Testcases

CSV formatted as "\<patter\>, \<test string\>, \<result\>".

For \<result\>, "01" as matched, "00" as unmatched.
