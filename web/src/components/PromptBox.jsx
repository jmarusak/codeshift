import { useState } from "react";
import "./PromptBox.css";

const PromptBox = ({ onSubmit, onPaste, onCopy }) => {
  const [prompt, setPrompt] = useState("Convert SAS code snippet in context to PySpark code (skip imports and creating spark session in PySpark code)");

  const handleSubmit = () => {
    onSubmit(prompt);
  };
  const handlePaste = () => {
    onPaste();
  };
  const handleCopy = () => {
    onCopy();
  };

  return (
    <div className="prompt-inputbox-container">
      <textarea
        id="prompt"
        rows="1"
        className="prompt-inputbox-textarea"
        value={prompt}
        onChange={(e) => setPrompt(e.target.value)}
      />
      <button className="prompt-button" onClick={handlePaste}>
        Paste SAS code 
      </button>
      <button className="prompt-button" onClick={handleSubmit}>
        Migrate 
      </button>
      <button className="prompt-button" onClick={handleCopy}>
        Copy Python code
      </button>
    </div>
  );
};

export default PromptBox;
