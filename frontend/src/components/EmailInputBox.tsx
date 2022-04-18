import { Box, Input, Text } from "@chakra-ui/react";

export type EmailInputBoxProps = { setEmail: (email: string) => void };

export const EmailInputBox: React.FC<EmailInputBoxProps> = ({ setEmail }) => {
    const onChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(event.target.value);
    };

    return (
        <Box my={'2vh'}>
            <Text>Email</Text>
            <Input size="lg" type={'email'} onChange={onChange} placeholder="Enter email" />
        </Box>
    );
};
