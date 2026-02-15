import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import gatorLogo from './assets/gator-logo.png';
import './SignIn.css';

function SignIn({ onSignIn }) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e) => {
        e.preventDefault();
        onSignIn(email, password);
    };

    return (
        <div className="signin-page">
            <form className="signin-container" onSubmit={handleSubmit}>
                <div className="signin-header">
                    <img src={gatorLogo} alt="GatorDash" className="auth-logo" />
                    <h2>Welcome Back</h2>
                    <p className="signin-subtitle">Sign in to GatorDash</p>
                </div>

                <div className="input-group">
                    <label>Email</label>
                    <input
                        type="email"
                        placeholder="Enter your email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>

                <div className="input-group">
                    <label>Password</label>
                    <input
                        type="password"
                        placeholder="Enter your password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>

                <button className="signin-btn" type="submit">Sign In</button>

                <p className="signin-links">
                    Don't have an account? <Link to="/signup">Sign up here</Link>
                </p>
            </form>
        </div>
    );
}

export default SignIn;
