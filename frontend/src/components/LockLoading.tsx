import { useRecoilValue } from 'recoil';

import { Center, Spinner } from '@chakra-ui/react';

import { isLockedAtom } from '../recoil/isLockedAtom';

export const LockLoading: React.FC<{}> = ({}) => {
    const isLocked = useRecoilValue(isLockedAtom);

    return (
        <Center
            height="100vh"
            width="100wh"
            pos="fixed"
            top={0}
            right={0}
            bottom={0}
            left={0}
            backgroundColor="rgba(0, 0, 0, 0.5)"
            hidden={!isLocked}
            zIndex={9999}
        >
            <Spinner size="xl" thickness="4px" color="orange" speed="0.75s" />
        </Center>
    );
};
