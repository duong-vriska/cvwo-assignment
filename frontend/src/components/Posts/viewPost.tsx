import axios from 'axios';
import {useEffect, useState} from 'react';
import "./Posts.css"
import { AiFillLike } from 'react-icons/ai'
import { IconButton } from '@mui/material';
import { FaCommentAlt, FaEdit } from 'react-icons/fa'
import { MdDelete } from "react-icons/md";
import { BsThreeDots } from "react-icons/bs";
import { useNavigate, Link } from "react-router-dom";


export const ViewPost = (props: any) => {

    const navigate = useNavigate();

    const onDeleteClicked = () => {
        props.deletePost(props.id);
        navigate(`/`);
    }

    return (
        <div className="view-post">
            <div className = "post-info">
                Post {props.id}
            </div>
            <div className = "post-title"> 
                {props.title}
            </div>
            <div className = "post-content">
                {props.content}
            </div>
            <span className = "post-likes">
                <IconButton aria-label="like-button" color = "inherit">
                    <AiFillLike/>
                </IconButton>
                <span> 0 </span>
            </span>
            <span className = "post-comments">
                <IconButton aria-label="comment-button" color = "inherit"> 
                    <FaCommentAlt/>
                </IconButton>
                <span> 0 </span>
            </span>
            <span className = "post-options">
                <span className = "post-delete">
                    <IconButton 
                        aria-label="delete-button" 
                        color = "inherit"
                        onClick = {onDeleteClicked}>
                        <MdDelete/>
                    </IconButton>
                </span>
                <span className = "post-edit">
                    <Link to={`/posts/${props.id}/edit`}>
                    <IconButton>
                        <BsThreeDots/>
                    </IconButton>
                    </Link>
                </span>
            </span>
        </div>
    )
}