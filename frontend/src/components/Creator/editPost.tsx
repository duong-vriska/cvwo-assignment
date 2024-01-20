import {useEffect, useState} from 'react'
import {useNavigate, useParams} from 'react-router-dom'
import axios from 'axios'
import { SelectCategory } from './selectCategory'
import { get } from 'http'

export default function EditPost(props: {
    [x: string]: any; editPost: (arg0: string, arg1: string) => void; 
}) {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    let navigate = useNavigate();
    let {id} = useParams();

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        props.editPost(title, content);
        setTitle('');
        setContent('');
        navigate(`/posts/${id}`);
    };    

    return (
        <div
        className = "edit-post">
        <div className = "edit-wrapper">
            <div className = "edit-intro">
                <h1> You are editing post </h1>
            </div>
            <form onSubmit ={handleSubmit}>
                <SelectCategory></SelectCategory>
                <div className = "edit-title">
                <input
                    type = "text"
                    required
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                </div>
                <div className = "edit-content">
                <textarea
                    required
                    placeholder="Write something here..."
                    value={content} 
                    onChange={(e) => setContent(e.target.value)}
                />
                </div>
                <button 
                className = "saved-button"
                type = "submit"> Submit
                </button>
            </form>
        </div>
        </div>
    )
}