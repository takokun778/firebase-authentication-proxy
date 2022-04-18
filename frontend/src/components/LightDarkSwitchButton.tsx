import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import { IconButton, useColorMode } from "@chakra-ui/react";

export const LightDarkSwitchButton: React.FC<{}> = ({}) => {
    const { colorMode, toggleColorMode } = useColorMode();

    return (
        <IconButton
            mb={10}
            aria-label="DarkMode Switch"
            icon={colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
            onClick={toggleColorMode}
        />
    );
};
