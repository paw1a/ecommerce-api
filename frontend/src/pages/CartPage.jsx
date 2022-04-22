import React from 'react';
import CssBaseline from "@mui/material/CssBaseline";
import Navbar from "../components/navbar/Navbar";
import Catalog from "../components/catalog/Catalog";
import Footer from "../components/Footer";
import Cart from "../components/cart/Cart";

const CartPage = () => {
    return (
        <React.Fragment>
            <CssBaseline />
            <Navbar/>
            <Cart/>
            <Footer />
        </React.Fragment>
    );
};

export default CartPage;