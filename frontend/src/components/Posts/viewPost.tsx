import "./Posts.css"
import { IconButton } from '@mui/material';
import { Reactions } from './reactionButtons';
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
            <div className = "post-options">
            <Reactions></Reactions>
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
            </div>
        </div>
    )
}
