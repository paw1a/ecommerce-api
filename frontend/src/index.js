import React from 'react';
import * as ReactDOM from 'react-dom';
import {
    BrowserRouter,
    Routes,
    Route,
} from "react-router-dom";

import CatalogPage from "./pages/CatalogPage";
import ProductPage from "./pages/ProductPage";
import App from "./App";
import {CssBaseline} from "@mui/material";

import './App.css'
import SignInPage from "./pages/SignInPage";
import SignUpPage from "./pages/SignUpPage";
import CartPage from "./pages/CartPage";

const container = document.getElementById('root');
const root = ReactDOM.createRoot(container);

root.render(
    <BrowserRouter>
        <CssBaseline />
        <div className='App'>
            <Routes>
                <Route path="/" element={<App />} />
                <Route path="catalog" element={<CatalogPage />} />
                <Route path="cart" element={<CartPage />} />
                <Route path="product/:productID" element={<ProductPage/>} />
                <Route path="sign-in" element={<SignInPage/>} />
                <Route path="sign-up" element={<SignUpPage/>} />
                <Route path="*" element={<div>Page not found</div>} />
            </Routes>
        </div>
    </BrowserRouter>,
    document.getElementById("root")
);
