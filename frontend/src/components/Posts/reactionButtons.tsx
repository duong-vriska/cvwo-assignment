import { IconButton } from '@mui/material';
import { AiFillLike } from "react-icons/ai";
import { FaCommentAlt } from "react-icons/fa";

export const Reactions = () => {
    return (
    <div className = "post-reactions">
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
    </div> 
    )
}