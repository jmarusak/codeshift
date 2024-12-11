import { useRef } from "react";
import "./TextEditor.css";

const TextEditor = ({ value, editable, onInput }) => {
  const editorRef = useRef(null);
  
  const handleInput = () => {
    const text = editorRef.current.innerText;
    onInput(text);
  };

  return (
    <div
      className="text-editor"
      contentEditable={editable}
      onInput={handleInput}
      ref={editorRef}
      suppressContentEditableWarning={true}
      spellCheck="false"
    >
      {value}
    </div>
  );
};

export default TextEditor;
