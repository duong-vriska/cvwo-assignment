import Select from 'react-select';
import { useState } from 'react';

const options = [
    { value: 'news', label: 'News' },
    { value: 'meme', label: 'Meme' },
    { value: 'study', label: 'Study' },
  ];

export const SelectCategory = () => {
    const [category, setCategory] = useState(null);
    const handleSelect = (selectedOption: any) => {
        setCategory(selectedOption);
    }

    return (
    <div className = "select-category">
        <form>
            <Select
            value={category} 
            options = {options}
            onChange = {handleSelect} 
            />
         </form>
    </div>
    )
}

