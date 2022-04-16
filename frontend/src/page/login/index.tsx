import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSetRecoilState } from 'recoil';

import {
    Button,
    Center,
    Container,
    Divider,
    Image,
    Text,
    useColorMode,
    useDisclosure,
    useToast,
} from '@chakra-ui/react';

import { getKey, postLogin } from '../../api/api';
import { encrypt } from '../../api/enctypt';
import { EmailInputBox } from '../../components/EmailInputBox';
import { LightDarkSwitchButton } from '../../components/LightDarkSwitchButton';
import { PasswordInputBox } from '../../components/PasswordInputBox';
import { isLockedAtom } from '../../recoil/isLockedAtom';
import { useLoginCheck } from '../../swr/useLoginCheck';
import { SignUpModal } from '../../components/SignUpModal';

const firebaseLogo = 'https://upload.wikimedia.org/wikipedia/commons/3/37/Firebase_Logo.svg';
const wide = '40vw';

export type LoginPageProps = {};

export const LoginPage: React.FC<LoginPageProps> = ({}) => {
    const [email, setEmail] = useState('');
    const [pass, setPass] = useState('');

    const navigate = useNavigate();
    const toast = useToast();
    const { data, error } = useLoginCheck();
    const setLocked = useSetRecoilState(isLockedAtom);
    const { colorMode } = useColorMode();
    const { isOpen, onOpen, onClose } = useDisclosure();

    const login = async () => {
        if (!email || !pass) {
            return;
        }
        try {
            const key = await getKey();
            const password = await encrypt(key, pass);
            setLocked(true);
            await postLogin(email, password);
            navigate('/');
            toast({ title: 'login', status: 'success', isClosable: true });
        } catch (e) {
            console.error(e);
            toast({ title: 'login error', status: 'error', isClosable: true });
        } finally {
            setLocked(false);
        }
    };

    useEffect(() => {
        if (error) {
            console.error(error);
            return;
        }
        if (data) {
            navigate('/');
        }
    }, [data, error, navigate]);

    return (
        <>
            <Center>
                <Image src={firebaseLogo} alt={'firebase-icon'} w="30vw" m={'2vw'} />
                <LightDarkSwitchButton />
            </Center>
            <Center>
                <Text>Authentication App</Text>
            </Center>
            <Container w={wide}>
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
            </Container>
            <Container w={wide}>
                <Center>
                    <Button w="100vw" colorScheme="orange" variant="outline" onClick={login} disabled={!email || !pass}>
                        Login
                    </Button>
                </Center>
            </Container>
            <Container w={wide} mt="2vh" mb="2vh">
                <Divider borderColor={colorMode === 'dark' ? 'whiteAlpha.500' : 'blackAlpha.500'} />
            </Container>
            <Container w={wide}>
                <Center>
                    <Button w="100vw" colorScheme="orange" onClick={onOpen}>
                        Sign Up
                    </Button>
                </Center>
            </Container>
            <SignUpModal isOpen={isOpen} onOpen={onOpen} onClose={onClose} />
        </>
    );
};
