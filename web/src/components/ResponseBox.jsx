import "./ResponseBox.css";

const ResponseBox = ({ value }) => {
  return (
    <div className="response-container">
      <div className="response" dangerouslySetInnerHTML={{ __html: value }} />
    </div>
  );
};

export default ResponseBox;
