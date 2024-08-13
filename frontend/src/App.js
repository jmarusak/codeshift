
import React, { useState } from 'react';
import './App.css';

function App() {
  const [requestMessage, setRequestMessage] = useState('');
  const [responseMessage, setResponseMessage] = useState('');

  const handlePostRequest = async () => {
    try {
      const response = await fetch('https://jsonplaceholder.typicode.com/posts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ message: requestMessage }),
      });
      const data = await response.json();
      setResponseMessage(JSON.stringify(data, null, 2));
    } catch (error) {
      setResponseMessage('Error: ' + error.message);
    }
  };

  return (
    <div className="App">
      <div style={{ display: 'flex', gap: '10px' }}>
        <textarea
          value={requestMessage}
          onChange={(e) => setRequestMessage(e.target.value)}
          placeholder="Enter message for API POST"
          rows="10"
          cols="30"
        />
        <textarea
          value={responseMessage}
          readOnly
          placeholder="API response message"
          rows="10"
          cols="30"
        />
      </div>
      <button onClick={handlePostRequest}>Migrate</button>
    </div>
  );
}

export default App;
