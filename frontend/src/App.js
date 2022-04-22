import React from "react";
import './App.css'

import {Link} from "react-router-dom";
import {CssBaseline} from "@mui/material";

function App() {
    return (
        <div>
            <CssBaseline />
            <Link to="/">App</Link> |{" "}
            <Link to="/catalog">Catalog</Link> |{" "}
            <Link to="/cart">Cart</Link> |{" "}
            <Link to="/product/1">Product</Link> |{" "}
            <Link to="/sign-in">Sign-in</Link> |{" "}
            <Link to="/sign-up">Sign-up</Link> |{" "}
        </div>
    )
}

export default App;
