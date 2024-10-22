// src/App.js
import React from 'react';
import Chatbot from './components/Chatbot';
import FileExplorerWrapper from './components/FileExplorer';
import './App.css';

const App = () => {
  return (
    <div className="app-container">
      <div className="component-container">
        <FileExplorerWrapper />
      </div>
      <div className="component-container">
        <Chatbot />
      </div>
    </div>
  );
};

export default App;
