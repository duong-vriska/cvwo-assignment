import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { CardPost } from '../Home';
import "./Category.css";
import { Container } from '@mui/material';

interface Post {
    key_id: any;
    post_id: string;
    title: string;
    content: string;
    category: string;
}

export const CategoryButtons = () => {
    return (
        <div className="category-buttons">
            <div className="wrapper">
            <Link to={`/posts/category/news`}>
            <button className="tag-button">News</button>
            </Link>
            <Link to={`/posts/category/entertainment`}>
            <button className="tag-button">Entertainment</button>
            </Link>
            <Link to={`/posts/category/food`}>
            <button className="tag-button">Food</button>
            </Link>
            <Link to={`/posts/category/study`}>
            <button className="tag-button">Study</button>
            </Link>
            <Link to={`/posts/category/sports`}>
            <button className="tag-button">Sports</button>
            </Link>
            </div>
        </div>
    )
}

// a button for the "News" category
// a button for the "Sports" category 