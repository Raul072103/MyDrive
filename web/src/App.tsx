import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';
import Login from './components/pages/auth/Login';
import HomePage from "./components/pages/home/HomePage.tsx";

import "@styles/global.css";
import './App.css';

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Login />} />
                <Route path="/home/*" element={<HomePage />} />
            </Routes>
        </Router>
    );
}

export default App;
