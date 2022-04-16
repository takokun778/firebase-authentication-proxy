import { atom } from 'recoil';

export const isLockedAtom = atom<boolean>({
    key: 'isLockedKey',
    default: false,
});
