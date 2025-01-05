import React, { useState } from 'react';
import './App.css';

function App() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        // Encode email and password as Base64
        const credentials = btoa(`${email}:${password}`);

        try {
            const response = await fetch('http://localhost:8080/v1/auth/login', {
                method: 'POST',
                headers: {
                    'Authorization': `Basic ${credentials}`,
                    'Content-Type': 'application/json'
                }
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            const data = await response.json();
            console.log('Login successful', data);
            // Handle successful login, like storing JWT token, redirecting, etc.

        } catch (error) {
            console.error('Error:', error);
            // Handle error (show error message, etc.)
        }
    };

    return (
        <div className="bg-white shadow-lg rounded-lg p-8 w-96">
            <h2 className="text-2xl font-bold text-center mb-6">Login to Your Account</h2>
            <form onSubmit={handleLogin}>
                <div className="mb-4">
                    <label htmlFor="email" className="block text-gray-700">Email</label>
                    <input
                        type="email"
                        id="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-blue-500"
                        placeholder="you@example.com"
                    />
                </div>
                <div className="mb-6">
                    <label htmlFor="password" className="block text-gray-700">Password</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                        className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-blue-500"
                        placeholder="********"
                    />
                </div>
                <button
                    type="submit"
                    className="w-full bg-blue-500 text-white py-2 rounded-md hover:bg-blue-600 transition duration-200"
                >
                    Login
                </button>
            </form>
            <div className="mt-4 text-center">
                <a href="#" className="text-blue-500 hover:underline">Forgot your password?</a>
            </div>
            <div className="mt-4 text-center">
                <p className="text-gray-600">
                    Don't have an account? <a href="#" className="text-blue-500 hover:underline">Sign up</a>
                </p>
            </div>
        </div>
    );
}

export default App;
