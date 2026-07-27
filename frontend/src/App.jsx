import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from "./pages/Home/Home";
import VotingCourt from './pages/Voting_court/Voting_court';
import NewVote from './pages/New_vote/New_vote';

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/votes" element={<VotingCourt />} />
                <Route path="/votes/new" element={<NewVote />} />
            </Routes>
        </BrowserRouter>
    );
}

export default App
