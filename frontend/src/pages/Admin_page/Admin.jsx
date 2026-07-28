import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import BackButton from '../../shared/ui/BackButton/BackButton';
import "./Admin.css";
import EraseButton from '../../shared/ui/EraseButton/EraseButton';
import { erase } from '../../api/erase';

export default function Admin({ user }) {
  const navigate = useNavigate();
  const [eraseMessage, setEraseMessage] = useState("");

  const handleErase = async () => {
    const confirmed = window.confirm("This will permanently erase all data. Are you sure?");
    if (!confirmed) return;
    const response = await erase();
    const data = await response.json();
    setEraseMessage(data.message);
    navigate('/')
  };

  return (
<div>
        <header className="admin-header">
            <div className="admin-header-row">
              <BackButton onClick={() => window.location.href = '/'} style={{ margin: '60px' }}/>
              <h1 className="admin-title">Admin</h1>
            </div>
        </header>
      <div className="admin-container" style={{ display: ["gold", "silver"].includes(user?.status) ? "flex" : "none" }}>
      <button
          type="button"
          className="post-form__submit"
          onClick={() => navigate('/mail')}>
          Create mails
        </button>
        <EraseButton onClick={handleErase} style={{ display: user?.status === "gold" ? "flex" : "none" }}/>
        {eraseMessage && <p>{eraseMessage}</p>}
      </div>
  </div>
  );
}
