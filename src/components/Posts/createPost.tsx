import React from 'react'
import "./Posts.css"

export const CreatePost = () => {
    return (
        <div
            className = "create-post">
            <div className = "create-wrapper">
                <div className = "create-intro">
                    <h1> Create a new post </h1>
                </div>
                <div className = "title">
                    <input
                        type = "text"
                        placeholder = "Title"
                        className = "input-title">
                    </input>
                </div>
                <div className = "body">
                    <textarea
                        id = "post-body"
                        name = "post-body"
                        rows={5}
                        cols = {33}
                        className = "input-body">
                        Give your post content!
                    </textarea>
                </div>
                <div className = "submit">
                    <button
                        type = "submit"
                        className = "submit-button">
                        Submit
                    </button>
                </div>
            </div>
        </div>
    )
}