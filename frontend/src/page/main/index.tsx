import '../../style/App.css';

import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { Setting } from '../../components/Setting';
import logo from '../../logo.svg';
import { useLoginCheck } from '../../swr/useLoginCheck';

export type MainPageProps = {};

export const MainPage: React.FC<MainPageProps> = ({}) => {
    const [count, setCount] = useState(0);

    const { data, error } = useLoginCheck();

    const navigate = useNavigate();

    useEffect(() => {
        if (error) {
            console.error(error);
            navigate('/login');
            return;
        }
        if (!data) {
            navigate('/login');
        }
    }, [data, error, navigate]);

    return (
        <div className="App">
            <Setting />
            <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>Hello Vite + React!</p>
                <p>
                    <button type="button" onClick={() => setCount((count) => count + 1)}>
                        count is: {count}
                    </button>
                </p>
                <p>
                    Edit <code>App.tsx</code> and save to test HMR updates.
                </p>
                <p>
                    <a className="App-link" href="https://reactjs.org" target="_blank" rel="noopener noreferrer">
                        Learn React
                    </a>
                    {' | '}
                    <a
                        className="App-link"
                        href="https://vitejs.dev/guide/features.html"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        Vite Docs
                    </a>
                </p>
            </header>
        </div>
    );
};
