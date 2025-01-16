import React, { useState } from 'react';
import './Login.css';
import {useNavigate} from "react-router-dom";

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate(); // Initialize the navigate function

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        const credentials = btoa(`${email}:${password}`);

        try {
            const response = await fetch('http://localhost:8080/v1/auth/login', {
                method: 'POST',
                headers: {
                    'Authorization': `Basic ${credentials}`,
                    'Content-Type': 'application/json',
                },
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            const data = await response.json();
            console.log('Login successful', data);
            // Handle successful login
            navigate('/home/myfiles');
        } catch (error) {
            console.error('Error:', error);
            // Handle error
        }
    };

    return (
        <div className="login-container">
            <h2>Login to Your Account</h2>
            <form onSubmit={handleLogin}>
                <label htmlFor="email">Email</label>
                <input
                    type="email"
                    id="email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <label htmlFor="password">Password</label>
                <input
                    type="password"
                    id="password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                />
                <button type="submit">Login</button>
            </form>
            <div>
                <a href="#">Forgot your password?</a>
            </div>
            <div>
                <p>
                    Don't have an account? <a href="#">Sign up</a>
                </p>
            </div>
        </div>
    );
}

export default Login;
