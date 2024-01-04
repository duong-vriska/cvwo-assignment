import React from 'react';
import "./Posts.css"
import { AiFillLike } from 'react-icons/ai'
import { FaCommentAlt } from 'react-icons/fa'

export const ViewPost = () => {
    return (
        <div className="post-wrapper">
            <div className = "post-info">
                <div className = "post-creator"> 
                Created by XXX on YYYY-MM-DD
                </div>
                <span className = "post-title"> 
                <h1> Post Title </h1>
                </span>
            </div>
            <div className = "post-body">
                <body>This is a post body</body>
            </div>
            <div className = "post-likes">
                <button className = "like-button"> 
                    <AiFillLike></AiFillLike> 
                    <span className = "like-count"> 0 </span>
                </button>
            </div>
            <div className = "post-comments">
                <button className = "comment-button"> 
                    <FaCommentAlt></FaCommentAlt> 
                    <span className = "comment-count"> 0 </span>
                </button>
            </div>
            <div className = "post-options">
                <div className = "post-edit">
                    <button className = "edit-button"> 
                        Edit 
                    </button>
                </div>
                <div className = "post-delete">
                    <button className = "delete-button"> 
                        Delete 
                    </button>
                </div>
            </div>
        </div>
    )
}