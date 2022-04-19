import React from "react";
import './App.css'

import Navbar from "./components/navbar/Navbar";
import CssBaseline from '@mui/material/CssBaseline';
import Catalog from "./components/catalog/Catalog";

function App() {
  return (
    <div className='App'>
        <CssBaseline />
        <Navbar/>
        <Catalog/>
    </div>
  );
}

export default App;
