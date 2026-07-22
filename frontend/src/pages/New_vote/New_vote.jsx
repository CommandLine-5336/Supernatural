import React, { useEffect, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
import { apiFetch } from '../api/client';

export default function SpiceForm({ mode }) {
    const params = useParams();

    const [error, setError] = useState('');
    const [form, setForm] = useState({
        type: '',
        description: '',
        user: '',
        agree: 0
        disagree: 0
    });

    const isEdit = (mode === 'edit');

    useEffect(() => {
        if (!isEdit) return;

        (async () => {
            setError('');
            try
                const obj = await apiFetch(`/votes/`, { method: 'GET' });
                setForm({
                    type: obj.type || '',
                    description: obj.description || '',
                    agree: obj.agree || '',
                    disagree: obj.disagree || '',
                    user: obj.user || '',
                });
            } catch (err) {
                setError(err.message);
            }
        })();
    }, [isEdit, id]);

    const setField = (name, value) => {
        setForm((prev) => ({ ...prev, [name]: value }));
    };

    const onSubmit = async (e) => {
        e.preventDefault();
        setError('');

        try {
            await await apiFetch('/votes/', {
                method: 'POST',
                body: JSON.stringify({
                    type: form.type,
                    description: form.description,
                    user_alias: form.user_alias,
                    agree: 0,
                    disagree: 0,
                }),
            });
            navigate('/voting_court');
        } catch (err) {
            setError(err.message || 'Couldn`t create vote');
        }
    };


    return (
        <main className="vote-create-container">
            <header className="vote-header">
                <div className="vote-header-row">
                  <Button href="/voting_court">{"<-"}</Button>
                  <h1 className="vote-title">New vote</h1>
                </div>
            </header>


            <section className="form-section">
                {error ? (
                    <div className="error-message" role="alert">
                        <p>{error}</p>
                    </div>
                ) : null}

                <form className="create-form" onSubmit={onSubmit}>
                    <fieldset className="form-fieldset">

                        <div className="form-group">
                            <label htmlFor="id_type">Type</label>
                            <textarea
                                id="id_type"
                                name="type"
                                rows={2}
                                value={form.type}
                                onChange={(e) => setField('type', e.target.value)}
                            />
                        </div>

                        <div className="form-group">
                            <label htmlFor="id_user">User alias</label>
                            <input
                                type="text"
                                id="id_user"
                                name="user"
                                required
                                value={form.user}
                                onChange={(e) => setField('user', e.target.value)}
                            />
                        </div>

                        <div className="form-group">
                            <label htmlFor="id_description">Description</label>
                            <input
                                type="text"
                                id="id_description"
                                name="description"
                                required
                                value={form.description}
                                onChange={(e) => setField('description', e.target.value)}
                            />
                        </div>

                    </fieldset>

                    <footer className="form-actions">
                        <button type="submit" className="btn-main">Create vote</button>
                    </footer>
                </form>

            </section>
        </main>
    );
}
