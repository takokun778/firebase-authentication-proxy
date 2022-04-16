import { encode } from 'base64-arraybuffer';

const str2ab = (str: string) => {
    const buf = new ArrayBuffer(str.length);
    const bufView = new Uint8Array(buf);
    for (let i = 0, strLen = str.length; i < strLen; i++) {
        bufView[i] = str.charCodeAt(i);
    }
    return buf;
};

const importRsaKey = async (pem: string) => {
    const pemHeader = '-----BEGIN PUBLIC KEY-----';
    const pemFooter = '-----END PUBLIC KEY-----';
    const pemContents = pem.trim().substring(pemHeader.length, pem.length - pemFooter.length - 1);
    const binaryDerString = window.atob(pemContents);
    const binaryDer = str2ab(binaryDerString);
    const result = await window.crypto.subtle.importKey(
        'spki',
        binaryDer,
        {
            name: 'RSA-OAEP',
            hash: 'SHA-256',
        },
        true,
        ['encrypt']
    );
    return result;
};

export const encrypt = async (pem: string, src: string) => {
    const enc = new TextEncoder();
    const encoded = enc.encode(src);
    const key = await importRsaKey(pem);
    const result = await window.crypto.subtle.encrypt(
        {
            name: 'RSA-OAEP',
        },
        key,
        encoded
    );
    if (result instanceof ArrayBuffer) {
        return encode(result);
    }
    throw new Error('');
};
