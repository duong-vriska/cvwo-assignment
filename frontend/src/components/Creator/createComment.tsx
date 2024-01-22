import "./Creator.css"
import { Divider, TextareaAutosize } from "@mui/material"

export const CommentCreator = () => {
   return (
            <div className="create-comment">
                <h2> Add a comment </h2>
                <TextareaAutosize
                id = "comment-body"
                name = "comment-body"
                placeholder= "Add a comment..." 
                className= "input-comment"
                minRows = {2}
                maxRows = {7}>
                </TextareaAutosize>
                <div>
                <button>
                    Submit
                </button>
                </div>
            </div>
    )
}

