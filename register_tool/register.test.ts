import RestClient from '../../src/network/rest/restClient';
import { Address} from './../../src/crypto';
import { PrivateKey } from './../../src/crypto/PrivateKey';
import { WebsocketClient } from './../../src/network/websocket/websocketClient';
import { signTransaction, makeInvokeTransaction } from './../../src/transaction/transactionBuilder';
import { reverseHex} from './../../src/utils';
import { Parameter, ParameterType} from './../../src/smartcontract/abi/parameter';

const registerContractAddr = 'fb5500b7ef5acb85f15adba6c7b799b55de95811'; // mainNet Contract Hash, DO NOT CHANGE
const registerContract = new Address(reverseHex(registerContractAddr));    
const gasPrice = '500';
const gasLimit = '20000';
const url = 'http://dappnode1.ont.io:'; // mainNet
// const url = 'http://polaris2.ont.io:'; // testNet
// const url = 'http://127.0.0.1:'; // privateNet
const restClient = new RestClient(url + '20334');
const socketClient = new WebsocketClient(url + '20335');


describe('register_tool', () => {
    // PLEASE FILL IN THE FOLLOWING PARAMETERS ------------- BEGIN 
    const base58Addr = new Address('AWf8NiLzXSDf1JB2Ae6YUKSHke4yLHMVCm');
    const privateKey = new PrivateKey('da213fb4cb1b12269c20307dadda35a7c89869c0c791b777fd8618d4159db99c');
    const ontId = 'did:ont:AWf8NiLzXSDf1JB2Ae6YUKSHke4yLHMVCm';
    const contractAddr = 'a34344eb3f7ea6ee8088ad6d13f5cfacc899dd88';
    // PLEASE FILL IN THE ABOVE PARAMETERS     ------------- END 

    test('test_bindContract',  async () => {
        const method = 'bindContract';
        let params = [
            new Parameter('account', ParameterType.ByteArray, base58Addr.serialize()),
            new Parameter('ontId', ParameterType.String, ontId),
            new Parameter('contractAddr', ParameterType.ByteArray, contractAddr)
        ];
        var tx = makeInvokeTransaction(method, params, registerContract, gasPrice, gasLimit, base58Addr)
        signTransaction(tx, privateKey);
        const result = await socketClient.sendRawTransaction(tx.serialize(), false, true);
        console.log('token bind contract ' + method + ' result is : \n' + JSON.stringify(result));
    });
    test('test_getBindContract', async()=>{
        let method = 'getBindContract';
        let params = [
            new Parameter('account', ParameterType.ByteArray, base58Addr.serialize()),
            new Parameter('ontId', ParameterType.String, ontId),
        ];
        var tx = makeInvokeTransaction(method, params, registerContract)
        const result = await restClient.sendRawTransaction(tx.serialize(), true);
        const resHex = result.Result.Result;
        console.log('token bind contract ' + method +' (' + base58Addr.toBase58() + '), ' + 'result is : \n' + resHex);
    });
});


