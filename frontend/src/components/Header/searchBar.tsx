import React, {useRef} from 'react'
import { FaMagnifyingGlass } from "react-icons/fa6";
import "./Header.css";

export const SearchBar = () => {
    return (
        <div className = "search-bar">
            <div className = "search-icon"> 
                <FaMagnifyingGlass></FaMagnifyingGlass>
            </div>
            <input 
                type = 'input'
                placeholder = "Search Mowou..." 
                className = "input-box">
            </input>

        </div>
    )
}

