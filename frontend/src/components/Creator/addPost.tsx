import { useState } from 'react'
import { useNavigate } from 'react-router-dom';
import Select from 'react-select/creatable';

const options = [
    { value: 'News', label: 'News' },
    { value: 'Entertainment', label: 'Entertainment' },
    { value: 'Food', label: 'Food' },
    { value: 'Study', label: 'Study' },
    { value: 'Sports', label: 'Sports' },
  ];

export function AddPost(props: { addPost: (arg0: string, arg1: string, arg3: string) => void; }) {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const [category, setCategory] = useState('');
    const navigate = useNavigate();
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        props.addPost(title, content, category);
        setTitle('');
        setContent('');
        setCategory('');
        setTimeout(() => {
            navigate(`/posts`);
        }, 100);
    };    
    
    return (
        <div
        className = "create-post">
        <div className = "create-wrapper">
            <div className = "create-intro">
                <h1> Create a new post </h1>
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
