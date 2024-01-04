import React from 'react'
import "./Header.css";
import { SearchBar } from './searchBar';

export const Header = () => {
  return (
    <header className="header">
      <div className="text-wrapper"> 
        MOWOU
      </div>
      <div className ="search-bar">
        <SearchBar></SearchBar>
      </div>
      <nav className="home-profile">
        New Post | Home | Profile
      </nav>
    </header>
  );
};

