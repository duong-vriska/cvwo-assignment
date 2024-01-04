import React from 'react'
import "./Posts.css"

export const CreateComment = () => {
   return (
        <div className="comments-section">
            <h2> Comments </h2>
            <div className="create-comment">
                <textarea
                id = "comment-body"
                name = "comment-body"
                placeholder= "Add a comment..." 
                className= "input-comment">
                </textarea>
                <button>
                    Submit
                </button>
            </div>
        </div>
    )
}