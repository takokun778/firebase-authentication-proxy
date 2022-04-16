import { Route, Routes } from 'react-router-dom';

import { ChakraProvider } from '@chakra-ui/react';

import { LockLoading } from './components/LockLoading';
import { LoginPage } from './page/login';
import { MainPage } from './page/main';
import { theme } from './style/theme';

function App() {
    return (
        <ChakraProvider theme={theme}>
            <LockLoading />
            <Routes>
                <Route path="/" element={<MainPage />} />
                <Route path="/login" element={<LoginPage />} />
            </Routes>
        </ChakraProvider>
    );
}

export default App;
