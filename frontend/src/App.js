import logo from './logo.svg';
import './App.css';

import React, { useState, useEffect, useRef } from "react";

function App() {

  const [tasks, setTasks] = useState([]);
  const [newDescription, setNewDescription] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/tasks')
      .then(response => response.json())
      .then(tasks => setTasks(tasks))
      .catch(err => console.error(err));
  }, []);

  const addTask = () => {
    fetch('http://localhost:8080/tasks', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ description: newDescription })
    })
      .then(response => response.json())
      .then(task => {
        setTasks([...tasks, task]);
        setNewDescription('');
      })
      .catch(err => console.error('Error creating task: ' + err));
  };

  return (
    <div className='mt-32'>
      <div className='px-4 sm:px-8 max-w-5xl m-auto'>
        <h1 className='text-center font-semibold text-lg'>Task List</h1>
        <ul className='border border-gray-200 rounded overflow-hidden shadow-md'>
          {tasks.map(task => <li key={task.id} className='px-4 py-2 bg-white hover:bg-sky-100 hover:text-sky-900 border-b last:border-none border-gray-200 transition-all duration-300 ease-in-out'>{task.description} - {task.done ? 'Done' : 'Not done'}</li>)}
        </ul>
        <div className='text-center block mt-4'>
          <input type="text" value={newDescription} onChange={e => setNewDescription(e.target.value)} className='p-4 border rounded' placeholder='New Task Description' />
          <button onClick={addTask} className='p-4 bg-blue-500 text-white'>Add Task</button>
        </div>
      </div>
    </div>
  );
}

export default App;
