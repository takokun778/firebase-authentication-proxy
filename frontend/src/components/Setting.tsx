import { useState } from "react";

import { SettingsIcon } from "@chakra-ui/icons";
import {
    Drawer,
    DrawerBody,
    DrawerCloseButton,
    DrawerContent,
    DrawerFooter,
    DrawerHeader,
    DrawerOverlay,
    HStack,
    IconButton,
    Stack,
    useDisclosure,
} from "@chakra-ui/react";

import { LogoutButton } from "./LogoutButton";
import { PasswordInputBox } from "./PasswordInputBox";
import { SaveChangePasswordButton } from "./SaveChangePasswordButton";
import { WithdrawButton } from "./WithdrawButton";

export const Setting: React.FC<{}> = ({}) => {
    const [oldPass, setOldPass] = useState("");
    const [newPass, setNewPass] = useState("");

    const { isOpen, onOpen, onClose } = useDisclosure();

    return (
        <>
            <IconButton
                aria-label="Setting Button"
                icon={<SettingsIcon />}
                onClick={onOpen}
                display={'block'}
                pos="absolute"
                zIndex="1000"
                right="0"
            />
            <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
                <DrawerOverlay />
                <DrawerContent>
                    <DrawerCloseButton />
                    <DrawerHeader borderBottomWidth="1px">Setting</DrawerHeader>

                    <DrawerBody>
                        <Stack spacing="10px">
                            <PasswordInputBox
                                setPass={(pass: string) => {
                                    setOldPass(pass);
                                }}
                                text={'OldPassword'}
                            />
                            <PasswordInputBox
                                setPass={(pass: string) => {
                                    setNewPass(pass);
                                }}
                                text={'NewPassword'}
                            />
                            <SaveChangePasswordButton oldPassword={oldPass} newPassword={newPass} />
                        </Stack>
                    </DrawerBody>

                    <DrawerFooter borderTopWidth="1px">
                        <HStack spacing="10px">
                            <WithdrawButton />
                            <LogoutButton />
                        </HStack>
                    </DrawerFooter>
                </DrawerContent>
            </Drawer>
        </>
    );
};
