import Select from 'react-select';

const options = [
    { value: 'chocolate', label: 'Chocolate' },
    { value: 'strawberry', label: 'Strawberry' },
    { value: 'vanilla', label: 'Vanilla' },
  ];

export const SelectCategory = () => {
    return (
    <div className = "select-category">
        <form>
            <Select 
            options = {options}
            />
         </form>
    </div>
    )
}

