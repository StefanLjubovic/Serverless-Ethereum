import React,{useState} from 'react'
import "./SecondPage.css"
import DropdownCreate from './DropdownCreate/DropdownCreate'
function SecondPage() {

  const [isOpen, setIsOpen] = useState(false);
  const [dropdownItems,setDropdownItems] = useState(['Item 1', 'Item 2', 'Item 3']);
  const [newSectionName, setNewSectionName] = useState('');
  function toggleDropdown() {
      setIsOpen(!isOpen);
  }

  function addSection(){
    if (newSectionName.trim() !== '') {
      setDropdownItems([...dropdownItems, newSectionName]);
      setNewSectionName('');
    }
  }

  
  return (
    <div className='second'>
        <div>
        <input type="text" name="name" placeholder='Section name' className='name' value={newSectionName}
          onChange={(e) => setNewSectionName(e.target.value)}/>
        <div className='btn-div'><button className='add' onClick={addSection}>Add</button></div>
        </div>
        <div>
        {dropdownItems.map((item, index) => (
                <DropdownCreate key={index} items={dropdownItems} />
            ))}
        </div>
    </div>
  )
}

export default SecondPage
