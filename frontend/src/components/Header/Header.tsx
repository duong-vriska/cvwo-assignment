import React from 'react'
import "./Header.css";
import { Button } from '@mui/material'
import { Link } from 'react-router-dom';

export const Header = () => {
  return (
    <header className="header">
      <div className="text-wrapper"> 
      <div className = "logo">
        <a href = "/">
          <img src="logo-mowou.png" width="100" height="50"></img></a>
      </div>
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
          <Button
          variant='text'
          sx={{
            color: "black",
            font: "inherit",
          }}
          component = {Link} to = "/login">Login
          </Button>
      </nav>
    </header>
  );
};

