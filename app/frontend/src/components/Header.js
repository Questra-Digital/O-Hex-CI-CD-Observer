import React from "react";
import '../dist/output.css';
function Header(props)
{
    return (
        <div>
            <header className="App-header">
                <h4 className="Header-logo">SmartObserver  {props.test}</h4>
                <button className="Head-butt">Home</button>
                <button className="Head-butt">About us</button>
                <button className="Head-butt">Register</button>
                <button className="Head-butt">Login</button>
            </header>

        </div>
    );
}
export default Header; 