import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Header from './components/Header/Header';
import Footer from './components/Footer/Footer';
import Home from './pages/Home';
import Library from './pages/Library';
import About from './pages/About';
import VideoPage from './pages/VideoPage/VideoPage';
import {VideoProvider} from './contexts/VideoProvider';

import './assets/styles/app.css'

function App() {
    return (
        <VideoProvider>
            <Router>
                <div id="app-container">
                    <Header/>
                    <main>
                        <Routes>
                            <Route path="/" element={<Home/>}/>
                            <Route path="/library" element={<Library/>}/>
                            <Route path="/about" element={<About/>}/>
                            <Route path="/video/:videoId" element={<VideoPage/>}/>
                        </Routes>
                    </main>
                    <Footer/>
                </div>
            </Router>
        </VideoProvider>
    );
}

export default App;