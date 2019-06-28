1. Download [ontology-ts-sdk](https://github.com/ontio/ontology-ts-sdk).

2. Install ts-sdk based on the instruction.

3. Copy the ```register``` folder into  ```./ontology-ts-sdk/test/``` folder. 

4. Open the terminal, make sure your current is under ```./ontology-ts-sdk``` folder.

5. Run the invoking script within the ```register.test.ts``` file, before run the script, make sure you open the ```register.test.ts``` file, 
replace the parameters ```base58Addr```, ```privateKey```, ```ontId```, and ```hash``` with yours.

5.1 In order to bind your answer(contract hash) with your account and ontId, then
run the following command.

```
jest .\\test\\register\\register.test.ts -t 'test_bindContract'

Or 

jest .\test\register\register.test.ts -t 'test_bindContract'

```

5.2 In order to check that you have bind your answer(contract hash) with your account and ontId, then
run the following command.

```
jest .\\test\\register\\register.test.ts -t 'test_getBindContract'

Or 

jest .\test\register\register.test.ts -t 'test_getBindContract'

```

6. Congrats, you have bind your answer with your account and ontId, and nobody CAN modify that.

7. If you have to submit your answer (contract hash) more than one time, just do it.