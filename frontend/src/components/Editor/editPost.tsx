import {useEffect, useState} from 'react'
import {useNavigate, useParams} from 'react-router-dom'
import axios from 'axios'
import Select from 'react-select/creatable';
import { get } from 'http'

const options = [
    { value: 'news', label: 'News' },
    { value: 'meme', label: 'Meme' },
    { value: 'study', label: 'Study' },
  ];

export default function EditPost(props: {
    [x: string]: any; editPost: (arg1: string, arg2: string, arg3: string) => void; 
}) {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const [category, setCategory] = useState('');
    let navigate = useNavigate();
    let {id} = useParams();

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        props.editPost(title, content, category);
        setTitle('');
        setContent('');
        setCategory('');
        setTimeout(() => {
        navigate(`/posts/${id}`);
        }, 1000);
    };    

    return (
        <div
        className = "edit-post">
        <div className = "edit-wrapper">
            <div className = "edit-intro">
                <h1> You are editing post </h1>
            </div>
            <form onSubmit ={handleSubmit}>
                <Select
                    options={options}
                    value={options.find((obj) => obj.value === category)}
                    onChange={(selectedOption) => {
                        if (selectedOption) {
                            setCategory(selectedOption.value);
                        } else {
                            setCategory('');
                        }}}
                    />
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