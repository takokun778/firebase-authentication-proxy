import { useSetRecoilState } from 'recoil';

import { Button, useToast } from '@chakra-ui/react';

import { getKey, putChangePassword } from '../api/api';
import { encrypt } from '../api/enctypt';
import { isLockedAtom } from '../recoil/isLockedAtom';

export type SaveChangePasswordButtonProps = {
    oldPassword: string;
    newPassword: string;
};

export const SaveChangePasswordButton: React.FC<SaveChangePasswordButtonProps> = ({ oldPassword, newPassword }) => {
    const toast = useToast();
    const setLocked = useSetRecoilState(isLockedAtom);

    const save = async () => {
        try {
            setLocked(true);
            const key = await getKey();
            const oldPass = await encrypt(key, oldPassword);
            const newPass = await encrypt(key, newPassword);
            await putChangePassword(oldPass, newPass);
            toast({ title: 'change password success', status: 'success', isClosable: true });
        } catch (error) {
            toast({ title: 'change password failed', status: 'error', isClosable: true });
        } finally {
            setLocked(false);
        }
    };

    return (
        <Button colorScheme="orange" variant="outline" onClick={save} disabled={!oldPassword || !newPassword}>
            Save
        </Button>
    );
};
