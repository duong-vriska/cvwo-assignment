import { AiFillLike } from "react-icons/ai";
import { FaCommentAlt } from "react-icons/fa";

export const SingleComment = () => {
    return (
    <div className = "comment-wrapper">
        <text className = "comment-creator">
            Created by XXX on YYYY-MM-DD
        </text>
        <div className = "comment-content">
            This is a comment
        </div>
        <div className = "comment-options">
            <div className = "comment-likes">
                <button className = "like-button"> 
                    <AiFillLike></AiFillLike> 
                    <span className = "like-count"> 0 </span>
                </button>
            </div>
            <div className = "comment-replies">
                <button className = "reply-button"> 
                    <FaCommentAlt></FaCommentAlt> Reply 
                </button>
                
            </div>
            <div className = "comment-edit">
                <button className = "edit-button"> 
                    Edit 
                </button>
            </div>
            <div className = "comment-delete">
                <button className = "delete-button"> 
                    Delete 
                </button>
            </div>
        </div>
    </div>
    )
}