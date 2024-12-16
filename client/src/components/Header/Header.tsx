import React from 'react';
import {Link} from 'react-router-dom';
import './header.css';

const Header = () => {
    return (
        <header className="header">
            <Link id="web-title" to="/">moviewatchers</Link> {/* or put image for icon */}
            <div id="link-container">
                <Link id="link-item" to="/">Home</Link>
                <Link id="link-item" to="/library">Library</Link>
                <Link id="link-item" to="/about">About</Link>
            </div>
        </header>
    );
};

export default Header;