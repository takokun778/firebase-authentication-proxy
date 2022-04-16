import Axios from 'axios';

import { client } from './client';

const errorHandle = (error: unknown) => {
    console.error(error);
    if (Axios.isAxiosError(error)) {
        switch (error.response?.status) {
            case 400:
                return new Error('BAD_REQUEST');
            case 401:
                return new Error('UNAUTHORIZE');
            default:
                return new Error('UNKNOWN');
        }
    }
    return new Error('UNKNOWN');
};

export const getKey = async () => {
    const result = await client.get('/api/key');
    return result.data;
};

export const postRegister = async (email: string, password: string) => {
    const result = await client.post(
        '/api/register',
        JSON.stringify({
            email,
            password,
        })
    );
    return result;
};

export const postLogin = async (email: string, password: string) => {
    const result = await client.post(
        '/api/login',
        JSON.stringify({
            email,
            password,
        })
    );
    return result;
};

export const putChangePassword = async (oldPassword: string, newPassword: string) => {
    try {
        const result = await client.put(
            '/api/change/password',
            JSON.stringify({
                oldPassword,
                newPassword,
            })
        );
        return result;
    } catch (error) {
        throw errorHandle(error);
    }
};

export const postLoginCheck = async () => {
    try {
        await client.post('/api/login/check');
        return 'OK';
    } catch (error) {
        throw errorHandle(error);
    }
};

export const postLogout = async () => {
    try {
        const result = await client.post('/api/logout');
        return result;
    } catch (error) {
        throw errorHandle(error);
    }
};

export const postWithdraw = async (password: string) => {
    try {
        await client.post(
            '/api/withdraw',
            JSON.stringify({
                password,
            })
        );
    } catch (error) {
        throw errorHandle(error);
    }
};
