import {useState, useEffect} from 'react';
import { Input, Button } from '@mui/material';

interface Username {
    username: string;
    password: string;
}

export const LoginBox = () => {
const [username, setUsername] = useState('');
const [password, setPassword] = useState(''); 
const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setUsername('');
    setPassword('');
};    

  return (
    <div className="login-box">
    <form onSubmit ={handleSubmit}>
    <div className="username-field">
    <Input
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}>
    </Input>
    </div> 
    <div className="password-field">
    <Input
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}>
    </Input>
    </div>
    <div className="login-button">
    <Button> Login</Button>
    </div>
    </form>
    </div>
  )
}

