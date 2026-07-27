import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from "./pages/Home/Home";
import Mail from "./pages/Mail/Mail";
import VotingCourt from './pages/Voting_court/Voting_court';
import NewVote from './pages/New_vote/New_vote';
import Admin_page from './pages/Admin_page/Admin';

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/votes" element={<VotingCourt />} />
                <Route path="/votes/new" element={<NewVote />} />
                <Route path="/mail" element={<Mail />} />
                <Route path="/admin" element={<Admin_page />} />


            </Routes>
        </BrowserRouter>
    );
}

export default App
