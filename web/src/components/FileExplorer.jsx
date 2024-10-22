import React, { useState } from 'react';
import './FileExplorer.css'; // Add a CSS file for styles

import fileStructure from './RepositoryContent';

// Recursive component for rendering folders and files
const FileExplorer = ({ structure }) => {
  const [isOpen, setIsOpen] = useState(false);

  const handleToggle = () => setIsOpen(!isOpen);

  if (structure.type === 'file') {
    return <div className="file-item">{structure.name}</div>;
  }

  return (
    <div>
      <div className={`folder-item ${isOpen ? 'open' : ''}`} onClick={handleToggle}>
        <span className="folder-icon">{isOpen ? '📂' : '📁'}</span>
        {structure.name}
      </div>
      {isOpen && (
        <div className="folder-children">
          {structure.children?.map((child, index) => (
            <FileExplorer key={index} structure={child} />
          ))}
        </div>
      )}
    </div>
  );
};

// Main component to pass file structure
const FileExplorerWrapper = () => {
  return (
    <div className="file-explorer-wrapper">
      <FileExplorer structure={fileStructure} />
    </div>
  );
};

export default FileExplorerWrapper;
