import React, { useState } from 'react';
import './App.css';

function App() {
  const [todos, setTodos] = useState([]);
  const [input, setInput] = useState('');
  const [editingId, setEditingId] = useState(null);
  const [editingText, setEditingText] = useState('');

  const handleAdd = () => {
    if (input) {
      const newTodo = {
        id: Date.now(),
        text: input
      };
      setTodos([...todos, newTodo]);
      setInput('');
    }
  };

  const handleDelete = (id) => {
    const updatedTodos = todos.filter(todo => todo.id !== id);
    setTodos(updatedTodos);
  };

  const handleEdit = (id) => {
    const todoToEdit = todos.find(todo => todo.id === id);
    if (todoToEdit) {
      setEditingId(id);
      setEditingText(todoToEdit.text);
    }
  };

  const handleUpdate = () => {
    const updatedTodos = todos.map(todo => 
      todo.id === editingId ? { ...todo, text: editingText } : todo
    );
    setTodos(updatedTodos);
    setEditingId(null);
    setEditingText('');
  };

  return (
    <div className="App bg-gray-100 min-h-screen flex items-center justify-center">
      <div className="bg-white p-8 rounded shadow-md w-96">
        <h1 className="text-2xl font-bold mb-4">Todo App</h1>
        {editingId ? (
          <div className="flex">
            <input
              className="flex-grow border p-2 rounded"
              value={editingText}
              onChange={e => setEditingText(e.target.value)}
            />
            <button className="ml-2 p-2 bg-blue-500 text-white rounded" onClick={handleUpdate}>Update</button>
          </div>
        ) : (
          <div className="flex">
            <input
              className="flex-grow border p-2 rounded"
              value={input}
              onChange={e => setInput(e.target.value)}
            />
            <button className="ml-2 p-2 bg-green-500 text-white rounded" onClick={handleAdd}>Add</button>
          </div>
        )}
        <ul className="mt-4 space-y-2">
          {todos.map((todo) => (
            <li key={todo.id} className="flex justify-between items-center border p-2 rounded">
              <span>{todo.text}</span>
              <div>
                <button className="mr-2 text-blue-500" onClick={() => handleEdit(todo.id)}>Edit</button>
                <button className="text-red-500" onClick={() => handleDelete(todo.id)}>Delete</button>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default App;
