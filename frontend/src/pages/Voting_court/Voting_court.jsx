import React, { useEffect, useState } from "react";
import { setVote } from "../../api/user";
import { getVotes } from "../../api/user";
import promoIcon from "../../assets/images/promo.png";
import exIcon from "../../assets/images/ex.png";
import "./Voting_court.css";

export default function VotingCourt() {

  const [votes, setVotes] = useState([]);
  const [votedIds, setVotedIds] = useState(new Set());
  const [error, setError] = useState("");
  const [me, setMe] = useState(null);

  const load = async () => {
    setError("");
    try {
      const data = await getVotes();
      const v = Array.isArray(data?.votes) ? data.votes : [];
      const voted = Array.isArray(data?.user_voted) ? data.user_voted : [];
      setVotes(v);
      setVotedIds(new Set(voted.map((r) => r.vote)));
      setMe(data?.me ?? null);
    } catch (err) {
      setError(err?.message || "Couldn't get amy votes");
    }
  };

  useEffect(() => {
    load();
  }, []);

  const onVote = async (voteId, res) => {
    if (votedIds.has(voteId)) return;

    try {
      const updated = await setVote(voteId, res);
      setVotes((prev) =>
        prev.map((v) => (v.id === voteId ? { ...v, ...updated } : v))
      );
      setVotedIds((prev) => new Set(prev).add(voteId));
    } catch (err) {
      window.alert(err?.message || "Err");
    }
  };

  return (
    <main className="vote-catalog">
      <header className="vote-header">
        <div className="vote-header-row">
          <a href="/">{"<-"}</a>
          <h1 className="vote-title">Voting Court</h1>
          <a href="/votes/new" style={{ display: me?.status === "copper" ? "none" : "inline-flex" }}>
            New vote
          </a>
          <div className="for-title" style={{ display: me?.status === "copper" ? "inline-flex" : "none" }}></div>
        </div>
      </header>

      {error && <div className="vote-error">
                        <p>{error}</p>
                    </div>}

      <section className="votes">
        {votes.map((vote) => {
          const { id: voteId, type, description, agree, disagree, user_alias } = vote;
          const alreadyVoted = votedIds.has(voteId);

          return (
            <article key={voteId} className="vote-card">
              <div className="vote-content">
                <img src={type === "promotion" ? promoIcon : exIcon} alt={type} className="vote-icon" />

                <p className="vote-meta-alias">{user_alias}</p>
                <p className="vote-meta-desc">{description}</p>

                <div className="vote_res">
                  <div className="vote-btns">
                    <div className="vote-yes">
                      <button type="button" className="vote-yes-btn" disabled={alreadyVoted} onClick={() => onVote(voteId, "+")}>
                        Yes
                      </button>
                      <p className="vote-num">{agree}</p>
                    </div>
                    <div className="vote-no">
                      <button type="button" className="vote-no-btn" disabled={alreadyVoted} onClick={() => onVote(voteId, "-")}>
                        No
                      </button>
                      <p className="vote-num">{disagree}</p>
                    </div>
                  </div>
                </div>
              </div>
            </article>
          );
        })}
      </section>

      <footer className="vote-footer">
        <h1 className="vote-descr">+ Promotion, - Excommunicado</h1>
      </footer>
    </main>
  );
}
