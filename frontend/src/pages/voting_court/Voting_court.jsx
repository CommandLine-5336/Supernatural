import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Button from "@mui/material/Button";
import { useAuth } from "../../hooks/useAuth";
import { apiFetch } from "../../utils/apiFetch";
import promoIcon from "../../assets/images/promo.png";
import exIcon from "../../assets/images/ex.png";
import "./Voting_court.css";

export default function Votes() {
  const navigate = useNavigate();
  const { user, isAuthenticated } = useAuth();

  const [votes, setVotes] = useState([]);
  const [votedIds, setVotedIds] = useState(new Set());
  const [error, setError] = useState("");
  const [me, setMe] = useState(null);

  const load = async () => {
    setError("");
    try {
      const data = await apiFetch("/votes/", { method: "GET" });
      const v = Array.isArray(data?.votes) ? data.votes : [];
      const voted = Array.isArray(data?.user_voted) ? data.user_voted : [];
      setVotes(v);
      setVotedIds(new Set(voted.map((r) => r.vote)));
      setMe(data?.me ?? null);
    } catch (err) {
      setError(err?.message || "Couldn`t get amy votes");
    }
  };

  useEffect(() => {
    load();
  }, []);

  const onVote = async (voteId, res) => {
    if (!isAuthenticated) {
      navigate("/login");
      return;
    }
    if (votedIds.has(voteId)) return;

    try {
      const updated = await apiFetch(`/votes/${voteId}/set_vote/`, {
        method: "POST",
        body: { res },
      });
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
          <Button href="/">{"<-"}</Button>
          <Button href="/votes/new" style={{ display: session?.status === "copper" ? "none" : "inline-flex"}}>
            New vote
          </Button>
          <h1 className="vote-title">Voting Court</h1>
        </div>
      </header>

      {error && <p className="vote-error">{error}</p>}

      <section className="votes">
        {votes.map((vote) => {
          const { id: voteId, type, description, agree, disagree, user_alias } = vote;
          const alreadyVoted = votedIds.has(voteId);

          return (
            <article key={voteId} className="vote-card">
              <div className="vote-content">
                <img src={type === "promotion" ? promoIcon : exIcon} alt={type} className="vote-icon"/>

                <p className="vote-meta">{user_alias}</p>
                <p className="vote-meta">{description}</p>

                <div className="vote_res">
                  <div className="vote-btns">
                    <Button variant="contained" color="success" disabled={alreadyVoted} onClick={() => onVote(voteId, "+")}>
                      Yes
                    </Button>
                    <Button variant="contained" color="error" disabled={alreadyVoted} onClick={() => onVote(voteId, "-")}>
                      No
                    </Button>
                  </div>
                  <div className="vote-nums">
                    <p className="vote-num">{agree}%</p>
                    <p className="vote-num">{disagree}%</p>
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
