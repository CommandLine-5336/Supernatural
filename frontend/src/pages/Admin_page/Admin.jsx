import { useNavigate } from 'react-router-dom';
import BackButton from '../../shared/ui/BackButton/BackButton';
import "./Admin.css";

export default function Admin() {
  const navigate = useNavigate();
  return (
<div>
        <header className="admin-header">
            <div className="admin-header-row">
              <BackButton onClick={() => window.location.href = '/'} style={{ margin: '60px' }}/>
              <h1 className="admin-title">Admin pref</h1>
            </div>
        </header>
      <div className="mail-container">
      <button
          type="button"
          className="post-form__submit"
          onClick={() => navigate('/mail')}>
          Create mails
        </button>
        <button
          // style={{ display: user?.status === "gold" ? "inline-flex" : "none" }}
          type="button"
          className="post-form__submit"
          onClick={() => navigate('/invite')}>
          Create invitation
        </button>
      </div>
  </div>
  );
}
