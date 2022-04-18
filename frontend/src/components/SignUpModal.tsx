import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useSetRecoilState } from "recoil";

import {
    Button,
    Modal,
    ModalBody,
    ModalCloseButton,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    useToast,
} from "@chakra-ui/react";

import { getKey, postRegister } from "../api/api";
import { encrypt } from "../api/enctypt";
import { isLockedAtom } from "../recoil/isLockedAtom";
import { EmailInputBox } from "./EmailInputBox";
import { PasswordInputBox } from "./PasswordInputBox";

export type SignUpModalProps = {
    isOpen: boolean,
    onOpen: () => void,
    onClose: () => void,
};

export const SignUpModal: React.FC<SignUpModalProps> = (
    { isOpen, onOpen, onClose },
) => {
    const [email, setEmail] = useState("");
    const [pass, setPass] = useState("");

    const navigate = useNavigate();
    const toast = useToast();
    const setLocked = useSetRecoilState(isLockedAtom);

    const register = async () => {
        if (!email || !pass) {
            return;
        }
        try {
            const key = await getKey();
            const password = await encrypt(key, pass);
            setLocked(true);
            await postRegister(email, password);
            navigate("/");
            toast({
                title: "sign up success",
                status: "success",
                isClosable: true,
            });
            onClose();
        } catch (e) {
            console.error(e);
            toast({ title: "sign up failed", status: "error", isClosable: true });
        } finally {
            setLocked(false);
        }
    };

    return (
        <>
            <Modal isOpen={isOpen} onClose={onClose} isCentered motionPreset="slideInBottom">
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>Create your account</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody pb={6}>
                        <EmailInputBox
                            setEmail={(email: string) => {
                                setEmail(email);
                            }}
                        />
                        <PasswordInputBox
                            setPass={(pass: string) => {
                                setPass(pass);
                            }}
                        />
                    </ModalBody>
                    <ModalFooter>
                        <Button colorScheme="orange" mr={3} onClick={register} disabled={!email || !pass}>
                            Submit
                        </Button>
                        <Button onClick={onClose}>Cancel</Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
};
