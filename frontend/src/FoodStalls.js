import React from 'react';
import Navbar from './Navbar';
import './FoodStalls.css';



const stalls = [
    { name: 'Restaurant 1', color: '#0D7377', desc: 'Coffee & Pastries' },
    { name: 'Restaurant 2', color: '#FFA500', desc: 'Burgers & Fries' },
    { name: 'Restaurant 3', color: '#FF0000', desc: 'Chinese Cuisine' },
    { name: 'Restaurant 4', color: '#5F8D4E', desc: 'Sandwiches & Salads' },
    { name: 'Restaurant 5', color: '#FEFFDE', desc: 'Street Food' },
    { name: 'Restaurant 6', color: '#FFD4D4', desc: 'Pizza & Italian' },
];

function FoodStalls({ onLogout }) {
    return (
        <div className="foodstalls-page">
            <Navbar onSignOut={onLogout} />
            <h2 className="stalls-heading">Explore Restaurants around you</h2>
            <div className="stalls-grid">
                {stalls.map((stall, index) => (
                    <div
                        className="stall-card"
                        key={index}
                        style={{ animationDelay: `${index * 0.1}s` }}
                    >
                        <div
                            className="stall-color-bar"
                            style={{ background: stall.color }}
                        />
                        <div className="stall-info">
                            <p className="stall-name">{stall.name}</p>
                            <p className="stall-desc">{stall.desc}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default FoodStalls;
