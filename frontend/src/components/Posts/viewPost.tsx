import "./Posts.css"
import { IconButton } from '@mui/material';
import { MdDelete } from "react-icons/md";
import { BsThreeDots } from "react-icons/bs";
import { useNavigate } from "react-router-dom";


export const ViewPost = (props: any) => {

    const navigate = useNavigate();

    const onDeleteClicked = async () => {
        await props.deletePost(props.post_id);
        console.log("Deleted!");
        navigate(`/`);
    }

    const onEditClicked = async () => {
        await props.editPost(props.post_id);
        console.log("Redirected!");
    }

    return (
        <div className="view-post">
            <div className = "post-info">
                Tag: {props.category}
            </div>
            <div className = "post-title"> 
                {props.title}
            </div>
            <div className = "post-content">
                {props.content}
            </div>
            <div className = "post-options">
                <span className = "post-delete">
                    <IconButton 
                        aria-label="delete-button" 
                        color = "inherit"
                        onClick = {onDeleteClicked}>
                        <MdDelete/>
                    </IconButton>
                </span>
                <span className = "post-edit">
                    <IconButton
                        onClick = {onEditClicked}>
                        <BsThreeDots/>
                    </IconButton>
                </span>
            </div>
        </div>
    )
}
