import React from 'react'
import "./Header.css";
import { Button } from '@mui/material'
import { Link } from 'react-router-dom';
import { SearchBar } from './searchBar';

export const Header = () => {
  return (
    <header className="header">
      <div className="text-wrapper"> 
      <a href = "/"><img src="logo-mowou.png" width="40%" height="40%"></img></a>
      </div>
      <div className ="search-bar">
        <SearchBar></SearchBar>
      </div>
      <nav className="home-profile">
          <Button 
          variant="text"
          sx={{
            color: "black",
            font: "inherit",
          }
          }
          component = {Link} to = "/new">New Post
          </Button>
      </nav>
    </header>
  );
};

