// src/pages/Dashboard.js
import React from 'react';
import { Link } from 'react-router-dom';

const Dashboard = () => {
    return (
        <div>
            <h2>Dashboard</h2>
            <p>Welcome to the Dashboard!</p>
            <nav>
                <ul>
                    <li><Link to="/upload">File Upload</Link></li>
                    {/* 其他导航链接 */}
                </ul>
            </nav>
        </div>
    );
};

export default Dashboard;
