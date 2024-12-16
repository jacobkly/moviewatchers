import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Header from './components/Header/Header';
import Footer from './components/Footer/Footer';
import Home from './pages/Home/Home';
import Library from './pages/Library/Library';
import About from './pages/About/About';
import VideoPage from './pages/VideoPage/VideoPage';
import './assests/styles/app.css'

function App() {
    return (
        <Router>
            <div id="app-container">
                <Header/>
                <main>
                    <Routes>
                        <Route path="/" element={<Home/>}/>
                        <Route path="/library" element={<Library/>}/>
                        <Route path="/about" element={<About/>}/>
                        <Route path="/video/:videoName" element={<VideoPage/>}/>
                    </Routes>
                </main>
                <Footer/>
            </div>
        </Router>
    );
}

export default App;