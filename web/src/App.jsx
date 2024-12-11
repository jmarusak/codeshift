import { useState } from 'react';
import * as Diff from 'diff';

import PromptBox from './components/PromptBox';
import ResponseBox from './components/ResponseBox';
import TextEditor from './components/TextEditor';
import './App.css';

const App = () => {
  const [codePrompt, setCodePrompt] = useState("");
  const [textResponse, setTextResponse] = useState("");
  const [codeResponse, setCodeResponse] = useState("");

  const handleCodePromptInput = (codeText) => {
    setCodePrompt(codeText)
  };

  const handlePromptSubmit = async (promptText) => {
    const promptFinal = `${promptText}, output modified context as part of the response\n\nCONTEXT:\n${codePrompt}\n\nOUTPUT INSTRUCTION: Format output in JSON schema {"code":string, "explanation":string}\n`;

    try {
      const url = "http://localhost:8080/generate";
      const response = await fetch(url, {
        method: "POST",
        body: JSON.stringify({ message: promptFinal }),
        headers: { "Content-Type": "application/json" }
      });
      const data = await response.json();
      setTextResponse(markdownToHtml(data.explanation));
      setCodeResponse(data.code);
    }
    catch (error) {
      console.error('Error:', error);
    }
  };

  const handlePasteContext = async () => {
    const context = await navigator.clipboard.readText();
    setCodePrompt(context);
    setCodeResponse('');
  };

  const handleCopyResponse = () => {
    navigator.clipboard.writeText(codeResponse);
  };

  return (
    <div className="app-container" >
      <div>
        <PromptBox
          onSubmit={handlePromptSubmit}
          onPaste={handlePasteContext}
          onCopy={handleCopyResponse}
        />
        <ResponseBox value={textResponse}/>
      </div>
      <div className="editors-container">
        <div className="component-container">
          <TextEditor value={codePrompt} editable="true" onInput={handleCodePromptInput}/>
        </div>
        <div className="component-container">
          <TextEditor value={codeResponse} editable="false"/>
        </div>
      </div>
    </div>
  );
};

const markdownToHtml = (markdownText) => {
  const htmlText = markdownText
    .replace(/\*\*(.*?)\*\*/g, '<b>$1</b>')
    .replace(/`(.*?)`/g, '<span style="font-family: monospace; font-size: 13px; background-color: #e6e6e6;">$1</span>')
    .replace(/\n/g, '<br>');
  return '<p>' + htmlText + '</p>';
};

export default App;
