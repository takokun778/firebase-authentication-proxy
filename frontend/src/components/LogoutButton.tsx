import { useNavigate } from "react-router-dom";
import { useSetRecoilState } from "recoil";

import { Button, useToast } from "@chakra-ui/react";

import { postLogout } from "../api/api";
import { isLockedAtom } from "../recoil/isLockedAtom";

export const LogoutButton: React.FC<{}> = ({}) => {
    const navigate = useNavigate();
    const toast = useToast();
    const setLocked = useSetRecoilState(isLockedAtom);

    const logout = async () => {
        try {
            setLocked(true);
            await postLogout();
            navigate("/login");
            toast({ title: "logout", status: "success", isClosable: true });
        } finally {
            setLocked(false);
        }
    };

    return (
        <Button colorScheme="orange" variant="outline" onClick={logout}>
            Logout
        </Button>
    );
};
