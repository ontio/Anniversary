import RestClient from '../../src/network/rest/restClient';
import { Address} from './../../src/crypto';
import { PrivateKey } from './../../src/crypto/PrivateKey';
import { WebsocketClient } from './../../src/network/websocket/websocketClient';
import { addSign, signTransaction, makeInvokeTransaction, makeDeployCodeTransaction} from './../../src/transaction/transactionBuilder';
import { hexstr2str, reverseHex, str2hexstr, ab2hexstring, hexstring2ab, num2hexstring, bigIntToBytes, sha256, StringReader} from './../../src/utils';
import { Parameter, ParameterType} from './../../src/smartcontract/abi/parameter';



describe('register', () => {
    const registerContractHash = '2df19b2562838e2234a435ac9ce341d374485db4';
    const registerContractAddr = new Address(reverseHex(registerContractHash));    
    const gasPrice = '500';
    const gasLimit = '20000';
    
    // const url = 'http://dappnode1.ont.io:'; // mainNet
    const url = 'http://polaris2.ont.io:'; // testNet
    // const url = 'http://127.0.0.1:'; // privateNet
    const restClient = new RestClient(url + '20334');
    const socketClient = new WebsocketClient(url + '20335');

    const base58Addr = new Address('AWf8NiLzXSDf1JB2Ae6YUKSHke4yLHMVCm');
    const privateKey = new PrivateKey('da213fb4cb1b12269c20307dadda35a7c89869c0c791b777fd8618d4159db99c');
    const ontId = 'did:ont:AMks3CoKPrgo2YvPWezVw2ckP8MkuDkfFf';
    const hash = 'a34344eb3f7ea6ee8088ad6d13f5cfacc899dd88';

    test('test_bindContract',  async () => {
        const method = 'bindContract';
        let params = [
            new Parameter('account', ParameterType.ByteArray, base58Addr.serialize()),
            new Parameter('ontid', ParameterType.String, ontId),
            new Parameter('hash', ParameterType.ByteArray, hash)
        ];
        var tx = makeInvokeTransaction(method, params, registerContractAddr, gasPrice, gasLimit, base58Addr)
        signTransaction(tx, privateKey);
        const result = await socketClient.sendRawTransaction(tx.serialize(), false, true);
        console.log('token bind contract ' + method + ' result is : \n' + JSON.stringify(result));
    });
    test('test_getBindContract', async()=>{
        let method = 'getBindContract';
        const base58Addr = new Address('AQf4Mzu1YJrhz9f3aRkkwSm9n3qhXGSh4p');
        let params = [
            new Parameter('account', ParameterType.ByteArray, base58Addr.serialize()),
            new Parameter('ontid', ParameterType.String, ontId),
        ];
        var tx = makeInvokeTransaction(method, params, registerContractAddr)
        const result = await restClient.sendRawTransaction(tx.serialize(), true);
        const resHex = result.Result.Result;
        console.log('token bind contract ' + method +' (' + base58Addr.toBase58() + '), ' + 'result is : \n' + resHex);
    });
});


