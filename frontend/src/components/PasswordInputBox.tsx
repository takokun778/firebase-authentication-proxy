import { ViewIcon } from "@chakra-ui/icons";
import { Box, Input, InputGroup, InputRightElement, Text, useBoolean } from "@chakra-ui/react";

export type PasswordInputBoxProps = {
    setPass: (pass: string) => void,
    text?: string,
};

export const PasswordInputBox: React.FC<PasswordInputBoxProps> = (
    { setPass, text },
) => {
    const [show, setShow] = useBoolean();

    const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPass(event.target.value);
    };

    return (
        <Box my={'2vh'}>
            <Text>{text ?? 'Password'}</Text>
            <InputGroup size="lg">
                <Input onChange={onChange} type={show ? 'text' : 'password'} placeholder="Enter password" />
                <InputRightElement w="4.5rem">
                    <ViewIcon onMouseOver={setShow.on} onMouseOut={setShow.off} />
                </InputRightElement>
            </InputGroup>
        </Box>
    );
};
