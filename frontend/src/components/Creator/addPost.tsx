import {useState} from 'react'
import { SelectCategory } from './selectCategory';
import { useNavigate} from 'react-router-dom';

export function AddPost(props: { addPost: (arg0: string, arg1: string) => void; }) {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const navigate = useNavigate();
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        props.addPost(title, content);
        setTitle('');
        setContent('');
        setTimeout(() => {
            navigate(`/posts`);
        }, 10);
    };    
    
    return (
        <div
        className = "create-post">
        <div className = "create-wrapper">
            <div className = "create-intro">
                <h1> Create a new post </h1>
            </div>
            <form onSubmit ={handleSubmit}>
                <SelectCategory></SelectCategory>
                <div className = "create-title">
                <input
                    type = "text"
                    required
                    placeholder="Title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                />
                </div>
                <div className = "create-content">
                <textarea
                    required
                    placeholder="Write something here..."
                    value={content} 
                    onChange={(e) => setContent(e.target.value)}
                />
                </div>
                <button 
                className = "create-button"
                type = "submit"> Submit
                </button>
            </form>
        </div>
        </div>
    )
}
