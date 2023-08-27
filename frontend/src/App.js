import logo from './logo.svg';
import './App.css';

import React, { useState, useEffect } from "react";

function App() {

  const [message, setMessage] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/')
      .then(response => response.json())
      .then(body => setMessage(body.message))
      .catch(err => console.error('Error fetching data: ' + err));
  }, []);

  return (
    <h1 className='text-3xl font-bold underline'>{message}</h1>
  );
}

export default App;
