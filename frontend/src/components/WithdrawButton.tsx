import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSetRecoilState } from 'recoil';

import {
    Button,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    useDisclosure,
    useToast,
} from '@chakra-ui/react';

import { getKey, postLogout, postWithdraw } from '../api/api';
import { encrypt } from '../api/enctypt';
import { isLockedAtom } from '../recoil/isLockedAtom';
import { PasswordInputBox } from './PasswordInputBox';

export const WithdrawButton: React.FC<{}> = ({}) => {
    const [pass, setPass] = useState('');

    const navigate = useNavigate();
    const toast = useToast();
    const setLocked = useSetRecoilState(isLockedAtom);
    const { isOpen, onOpen, onClose } = useDisclosure();

    const withdraw = async () => {
        try {
            setLocked(true);
            const key = await getKey();
            const password = await encrypt(key, pass);
            await postWithdraw(password);
            await postLogout();
            navigate('/login');
            toast({ title: 'withdraw success', status: 'success', isClosable: true });
        } catch (error) {
            toast({ title: 'withdraw failed', status: 'error', isClosable: true });
        } finally {
            setLocked(false);
        }
    };

    return (
        <>
            <Button colorScheme="orange" variant="outline" onClick={onOpen}>
                Withdraw
            </Button>
            <Modal isOpen={isOpen} onClose={onClose} isCentered motionPreset="slideInBottom">
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>Withdraw</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody pb={6}>
                        <PasswordInputBox
                            setPass={(pass: string) => {
                                setPass(pass);
                            }}
                        />
                    </ModalBody>
                    <ModalFooter>
                        <Button colorScheme="orange" mr={3} onClick={withdraw} disabled={!pass}>
                            Withdraw
                        </Button>
                        <Button onClick={onClose}>Cancel</Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
};
